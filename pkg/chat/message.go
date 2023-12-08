package chat

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Message struct {
	ID      uuid.UUID `json:"id"`
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
