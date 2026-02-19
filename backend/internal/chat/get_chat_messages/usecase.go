package get_chat_messages

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	"context"
)

type Postgres interface {
	GetChatMessages(ctx context.Context, chatID, limit int, beforeMessageID *int) ([]domain.ChatMessage, error)
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

//TODO: написать метод для проверки принадлежности пользователя к чату

func (u *Usecase) GetChatMessages(ctx context.Context, input Input) (Output, error) {
	messages, err := u.postgres.GetChatMessages(ctx, input.ChatID, input.Limit, input.BeforeMessageID)
	if err != nil {
		return Output{}, err
	}

	outputMessages := make([]OutputMessage, 0)
	for _, message := range messages {
		outputMessages = append(outputMessages, OutputMessage{
			UserAuthor: OutputUser{
				ID:   message.UserID,
				Tag:  message.Usertag,
				Name: message.Username,
			},
			Message: OutputMessageContent{
				ID:   message.MessageID,
				Text: message.Text,
				Time: message.Time,
			},
		})
	}

	return Output{
		Messages: outputMessages,
	}, nil
}
