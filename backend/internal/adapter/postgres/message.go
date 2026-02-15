package postgres

import (
	"backend/internal/domain"
	"context"
	"fmt"
)

func Start() {

}

func (p *Pool) GetMessagesByChatID(ctx context.Context, chatID int) ([]domain.Message, error) {
	messages := make([]domain.Message, 0)
	rows, err := p.pool.Query(ctx, `SELECT id, user_id, chat_id, text, time FROM messages WHERE chat_id = $1 ORDER BY time ASC`, chatID)
	if err != nil {
		return nil, fmt.Errorf("failed to query messages: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var msg domain.Message
		err = rows.Scan(&msg.ID, &msg.UserID, &msg.ChatID, &msg.Text, &msg.Time)
		if err != nil {
			return nil, fmt.Errorf("failed to scan message row: %w", err)
		}
		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return messages, nil
}
