package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

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
		return nil, fmt.Errorf("failed to query chats messages: %w", err)
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

func (p *Pool) GetChatMessages(ctx context.Context, chatID, limit, offset int) ([]domain.ChatMessage, error) {
	query := `
		SELECT *
		FROM (
			SELECT m.user_id, u.tag AS usertag, u.username, m.id AS message_id, m.text, m.time
			FROM messages m
			JOIN users u ON u.id = m.user_id
			WHERE m.chat_id = $1
			ORDER BY m.time DESC
			LIMIT $2
			OFFSET $3
		) AS last_messages
		ORDER BY time ASC
	`

	rows, err := p.pool.Query(ctx, query, chatID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query chat messages: %w", err)
	}
	messages := make([]domain.ChatMessage, 0)
	for rows.Next() {
		var msg domain.ChatMessage
		msg.ChatID = chatID
		err := rows.Scan(
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
		messages = append(messages, msg)
	}
	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}
	return messages, nil
}
