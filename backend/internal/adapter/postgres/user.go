package postgres

import (
	"backend/internal/domain"
	"context"
)

func (p *Pool) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	var id int
	err := p.pool.QueryRow(ctx,
		`INSERT INTO users (tag, username, password) 
		VALUES ($1, $2, $3) 
		RETURNING id`,
		user.Tag, user.Username, user.Password,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	user.ID = id
	return &user, nil
}
