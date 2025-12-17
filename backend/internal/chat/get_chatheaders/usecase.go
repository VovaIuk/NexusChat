package get_chatheaders

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	jwttoken "backend/pkg/jwt_token"
	"context"
)

type Postgres interface {
	GetChatHeaders(ctx context.Context, userID int) ([]domain.ChatHeader, error)
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

func (u *Usecase) GetChatHeaders(ctx context.Context, input Input) (Output, error) {
	claims, err := u.jwtManager.ParseToken(input.JWTToken)
	if err != nil {
		return Output{}, err
	}

	headers, err := u.postgres.GetChatHeaders(ctx, claims.UserID)
	if err != nil {
		return Output{}, err
	}

	var output Output
	for _, header := range headers {
		output.Chats = append(output.Chats, OutputChat{
			ID:      header.ChatID,
			Name:    header.ChatName,
			Message: header.LastMessage,
			Time:    header.LastTime.Format("15.04"),
		})
	}
	return output, nil
}
