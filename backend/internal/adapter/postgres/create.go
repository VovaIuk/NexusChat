package postgres

import (
	"backend/internal/domain"
	"context"

	"github.com/sirupsen/logrus"
)

func (p *Pool) CreateMessage(ctx context.Context, msg domain.MessageCreate) (int, error) {
	query := `
		INSERT INTO Messages (user_id, chat_id, text, time)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var id int
	err := p.pool.QueryRow(ctx, query, msg.UserID, msg.ChatID, msg.Text, msg.Time).Scan(&id)
	if err != nil {
		logrus.Errorf("create message: %v", err)
		return 0, err
	}
	return id, nil
}
