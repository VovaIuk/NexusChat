package login_user

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	"context"
)

type Postgres interface {
	CreateUser(ctx context.Context, user domain.User) (*domain.User, error)
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

func (u *Usecase) RegisterUser(ctx context.Context, input Input) (Output, error) {
	user, err := domain.NewUser(input.Usertag, input.Username, input.Password)
	if err != nil {
		return Output{}, err
	}

	createUser, err := u.postgres.CreateUser(ctx, user)
	if err != nil {
		return Output{}, err
	}

	return Output{
		Id:       createUser.ID,
		Usertag:  createUser.Tag,
		Username: createUser.Username,
		Password: createUser.Password,
	}, nil

}
