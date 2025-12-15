package login_user

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	jwttoken "backend/pkg/jwt_token"
	"context"
)

type Postgres interface {
	GetUserByTagAndPassword(ctx context.Context, tag string, password string) (*domain.User, error)
}

type JWTManager interface {
	CreateToken(userID int, usertag, username string) (string, error)
	GetExpiresInSeconds() int
}

type Usecase struct {
	postgres   Postgres
	jwtManager JWTManager
}

func New(postgres *postgres.Pool, jwtManager *jwttoken.JWTManager) *Usecase {
	uc := &Usecase{
		postgres:   postgres,
		jwtManager: jwtManager,
	}
	usecase = uc
	return uc
}

func (u *Usecase) LoginUser(ctx context.Context, input Input) (Output, error) {
	user, err := u.postgres.GetUserByTagAndPassword(context.Background(), input.Usertag, input.Password)
	if err != nil {
		return Output{}, err
	}

	token, err := u.jwtManager.CreateToken(user.ID, user.Tag, user.Username)
	if err != nil {
		return Output{}, err
	}

	return Output{
		User: User{
			ID:   user.ID,
			Tag:  user.Tag,
			Name: user.Username,
		},
		Token: Token{
			Refresh:   token,
			ExpiresIn: u.jwtManager.GetExpiresInSeconds(),
			Type:      "Bearer",
		},
	}, nil

}
