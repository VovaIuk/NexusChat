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
