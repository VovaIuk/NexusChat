package getme

import (
	"backend/internal/adapter/postgres"
	jwttoken "backend/pkg/jwt_token"
	"context"
	"strings"

	"github.com/sirupsen/logrus"
)

type Postgres interface {
}

type JWTManager interface {
	ParseToken(tokenStr string) (*jwttoken.Claims, error)
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

func (u *Usecase) GetUserByToken(ctx context.Context, input Input) (Output, error) {
	//TODO: потом добавить проверку пользователя
	tokenParts := strings.Split(input.Token, " ")

	claims, err := u.jwtManager.ParseToken(tokenParts[1])
	if err != nil {
		logrus.Debugf("debug usecase me: %s", err)
		return Output{}, err
	}
	output := Output{
		Id:       claims.UserID,
		Usertag:  claims.Usertag,
		Username: claims.Username,
	}
	return output, nil
}
