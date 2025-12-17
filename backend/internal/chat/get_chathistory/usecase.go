package get_chathistory

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	jwttoken "backend/pkg/jwt_token"
	"context"
	"fmt"
)

type Postgres interface {
	GetMessagesByChatID(ctx context.Context, chatId int) ([]domain.Message, error)
	GetUsersByChatID(ctx context.Context, chatID int) (map[int]domain.User, error)
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

func (u *Usecase) GetChatHistory(ctx context.Context, input Input) (Output, error) {
	claims, err := u.jwtManager.ParseToken(input.JWTToken)
	if err != nil {
		return Output{}, err
	}

	usersInChat, err := u.postgres.GetUsersByChatID(ctx, input.ChatID)
	if err != nil {
		return Output{}, err
	}
	if _, exists := usersInChat[claims.UserID]; !exists {
		return Output{}, fmt.Errorf("user %d is not a member of chat %d",
			claims.UserID,
			input.ChatID)
	}

	messages, err := u.postgres.GetMessagesByChatID(ctx, input.ChatID)
	if err != nil {
		return Output{}, err
	}

	var output Output
	for _, message := range messages {
		output.Messages = append(output.Messages, OutputMessage{
			ID:       message.ID,
			UserID:   message.UserID,
			Username: usersInChat[message.UserID].Username,
			Message:  message.Text,
			Time:     message.Time.Format("15.04"),
		})
	}
	return output, nil
}
