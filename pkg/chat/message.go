package chat

import (
	"encoding/json"
	"github.com/google/uuid"
	"webchat/pkg/utils"
)

const (
	MessageTypeNone   = "none"
	MessageTypeUser   = "user"
	MessageTypeSystem = "system"
)

type Message struct {
	ID      uuid.UUID `json:"id"`
	Type    string    `json:"type"`
	RoomID  uuid.UUID `json:"room_id"`
	User    User      `json:"user"`
	Content string    `json:"msg"`
}

func ParseMessage(data []byte) (Message, error) {
	var message Message
	if err := json.Unmarshal(data, &message); err != nil {
		return Message{}, err
	}

	return message, nil
}

func DefaultUserMessage(roomID uuid.UUID, currentUser User, content string) utils.Response {
	return defaultMessage(roomID, currentUser, MessageTypeUser, content)
}

func DefaultSystemMessage(roomID uuid.UUID, currentUser User, content string) utils.Response {
	return defaultMessage(roomID, currentUser, MessageTypeSystem, content)
}

func defaultMessage(roomID uuid.UUID, currentUser User, messageType string, content string) utils.Response {
	return utils.Response{
		Error:   false,
		Code:    200,
		Message: "",
		Data: Message{
			ID:      uuid.New(),
			Type:    MessageTypeSystem,
			RoomID:  roomID,
			User:    currentUser,
			Content: content,
		},
	}
}
