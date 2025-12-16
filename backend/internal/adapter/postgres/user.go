package postgres

import (
	"backend/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
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

func (p *Pool) GetUserByTagAndPassword(ctx context.Context, tag string, password string) (*domain.User, error) {
	var user domain.User
	err := p.pool.QueryRow(ctx,
		`SELECT id, tag, username, password FROM users WHERE tag = $1`,
		tag,
	).Scan(&user.ID, &user.Tag, &user.Username, &user.Password)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("invalid credentials")
		}
		return nil, fmt.Errorf("database query failed: %w", err)
	}
	if !user.CheckPassword(password) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}
