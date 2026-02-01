package postgres

import (
	"backend/internal/domain"
	"context"
	"database/sql"
	"fmt"
)

func (p *Pool) GetChatHeaders(ctx context.Context, userID int) ([]domain.ChatHeader, error) {
	var headers []domain.ChatHeader = make([]domain.ChatHeader, 0)

	query := `
	SELECT 
		c.id AS chat_id,
		CASE 
			WHEN participant_count = 2 THEN other_user.username
			ELSE 'Групповой чат'
		END AS chat_name,
		last_message.text AS last_message,
		last_message.time AS last_time
	FROM (
		SELECT 
			uc.chat_id,
			COUNT(uc2.user_id) AS participant_count,
			MAX(uc2.user_id) FILTER (WHERE uc2.user_id <> $1) AS other_user_id
		FROM user_chat uc
		JOIN user_chat uc2 ON uc2.chat_id = uc.chat_id
		WHERE uc.user_id = $1
		GROUP BY uc.chat_id
	) chat_info
	JOIN chats c ON c.id = chat_info.chat_id
	LEFT JOIN users other_user ON other_user.id = chat_info.other_user_id
	LEFT JOIN LATERAL (
		SELECT text, time
		FROM messages m
		WHERE m.chat_id = chat_info.chat_id
		ORDER BY time DESC
		LIMIT 1
	) last_message ON true
	ORDER BY last_message.time DESC NULLS LAST;`

	rows, err := p.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query chat headers: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var header domain.ChatHeader
		var lastMessage sql.NullString
		var lastTime sql.NullTime

		err := rows.Scan(
			&header.ChatID,
			&header.ChatName,
			&lastMessage,
			&lastTime,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan chat header row: %w", err)
		}

		header.LastMessage = lastMessage.String
		if lastTime.Valid {
			header.LastTime = lastTime.Time
		}

		headers = append(headers, header)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return headers, nil
}

func (p *Pool) GetChatsMembersByUserID(ctx context.Context, userID int) (map[int][]domain.ChatMember, error) {
	query := `
	SELECT uc1.chat_id, u.id as user_id, u.tag as usertag, u.username FROM user_chat uc1
	JOIN user_chat uc2 on uc1.chat_id = uc2.chat_id and uc1.user_id = $1
	JOIN users u on uc2.user_id = u.id 
	`

	rows, err := p.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to query chat members: %w", err)
	}
	defer rows.Close()

	members := make(map[int][]domain.ChatMember, 0)
	for rows.Next() {
		var member domain.ChatMember
		err := rows.Scan(&member.ChatID, &member.UserID, &member.Usertag, &member.Username)
		if err != nil {
			return nil, fmt.Errorf("failed to scan chat member row: %w", err)
		}
		members[member.ChatID] = append(members[member.ChatID], member)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return members, nil
}

func (p *Pool) GetChatsMessagesByUserID(ctx context.Context, userID int, limit int) (map[int][]domain.ChatMessage, error) {
	query := `
		SELECT chat_id, user_id, usertag, username, message_id, text, time
		FROM (
			SELECT
				m.chat_id,
				m.user_id,
				u.tag AS usertag,
				u.username,
				m.id AS message_id,
				m.text,
				m.time,
				ROW_NUMBER() OVER (PARTITION BY m.chat_id ORDER BY m.time DESC) AS rn
			FROM messages m
			JOIN user_chat uc ON uc.chat_id = m.chat_id AND uc.user_id = $1
			JOIN users u ON u.id = m.user_id
		) ranked
		WHERE rn <= $2
		ORDER BY time ASC`

	rows, err := p.pool.Query(ctx, query, userID, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query chat messages: %w", err)
	}
	defer rows.Close()

	messages := make(map[int][]domain.ChatMessage, 0)
	for rows.Next() {
		var msg domain.ChatMessage
		err := rows.Scan(
			&msg.ChatID,
			&msg.UserID,
			&msg.Usertag,
			&msg.Username,
			&msg.MessageID,
			&msg.Text,
			&msg.Time,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan chat message row: %w", err)
		}
		messages[msg.ChatID] = append(messages[msg.ChatID], msg)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return messages, nil
}
