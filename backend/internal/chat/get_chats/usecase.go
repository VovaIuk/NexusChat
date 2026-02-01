package getchats

import (
	"backend/internal/adapter/postgres"
	"backend/internal/domain"
	"context"
)

type Postgres interface {
	GetChatsMembersByUserID(ctx context.Context, userID int) (map[int][]domain.ChatMember, error)
	GetChatsMessagesByUserID(ctx context.Context, userID int, limit int) (map[int][]domain.ChatMessage, error)
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

func (u *Usecase) GetChats(ctx context.Context, input Input) (Output, error) {

	members, err := u.postgres.GetChatsMembersByUserID(ctx, input.UserID)
	if err != nil {
		return Output{}, err
	}
	messages, err := u.postgres.GetChatsMessagesByUserID(ctx, input.UserID, input.LimitMessages)
	if err != nil {
		return Output{}, err
	}

	chatIDs := make(map[int]struct{}, 0)
	for chatID := range members {
		chatIDs[chatID] = struct{}{}
	}
	for chatID := range messages {
		chatIDs[chatID] = struct{}{}
	}

	outputChats := make([]OutputChat, 0, len(chatIDs))
	for chatID := range chatIDs {
		oc := OutputChat{
			ID:       chatID,
			Name:     getChatName(members[chatID], input.UserID),
			Users:    mapMembersToOutputUsers(members[chatID]),
			Messages: mapMessagesToOutputMessages(messages[chatID]),
		}
		outputChats = append(outputChats, oc)
	}

	return Output{
		Chats: outputChats,
	}, nil
}

func mapMembersToOutputUsers(members []domain.ChatMember) []OutputUser {
	users := make([]OutputUser, 0, len(members))
	for _, m := range members {
		users = append(users, OutputUser{
			ID:   m.UserID,
			Tag:  m.Usertag,
			Name: m.Username,
		})
	}
	return users
}

func mapMessagesToOutputMessages(msgs []domain.ChatMessage) []OutputMessage {
	result := make([]OutputMessage, 0, len(msgs))
	for _, m := range msgs {
		result = append(result, OutputMessage{
			UserAuthor: OutputUser{
				ID:   m.UserID,
				Tag:  m.Usertag,
				Name: m.Username,
			},
			Message: OutputMessageContent{
				ID:   m.MessageID,
				Text: m.Text,
				Time: m.Time,
			},
		})
	}
	return result
}

func getChatName(members []domain.ChatMember, currentUserID int) string {
	if len(members) == 2 {
		for _, m := range members {
			if m.UserID != currentUserID {
				return m.Username
			}
		}
	}
	return "Групповой чат"
}
