package search_user

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	"context"
)

type Postgres interface {
	SearchUserByTag(ctx context.Context, tag string, limit int) ([]domain.PublicUser, error)
}

type Usecase struct {
	postgres Postgres
}

func New(postgres *postgres.Pool) *Usecase {
	uc := &Usecase{
		postgres: postgres,
	}
	usecase = uc
	return uc
}

func (u *Usecase) SearchUser(ctx context.Context, input Input) (Output, error) {
	users, err := u.postgres.SearchUserByTag(ctx, input.Tag, input.Limit)
	if err != nil {
		return Output{}, err
	}

	outputUsers := make([]OutputUser, 0)
	for _, user := range users {
		outputUsers = append(outputUsers, OutputUser{
			ID:   user.ID,
			Tag:  user.Tag,
			Name: user.Username,
		})
	}

	return Output{Users: outputUsers}, nil
}
