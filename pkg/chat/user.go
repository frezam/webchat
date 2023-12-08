package chat

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func NewUser(username string) User {
	return User{
		ID:   uuid.New(),
		Name: username,
	}
}
