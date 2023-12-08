package chat

import (
	"errors"
	"github.com/google/uuid"
)

type Chat struct {
	Rooms       map[string]*Room
	HandleRooms chan *Room
}

var CHAT *Chat

func init() {
	CHAT = New()
}

func New() *Chat {
	return &Chat{
		Rooms:       make(map[string]*Room),
		HandleRooms: make(chan *Room),
	}
}

func (c *Chat) CreateRoom(owner User) *Room {
	room := &Room{
		ID:        uuid.New(),
		Owner:     owner,
		Listeners: make([]*Listener, 0),
		Queue:     make(chan Message),
	}
	c.Rooms[room.ID.String()] = room
	return room
}

func (c *Chat) GetRoom(id string) (*Room, error) {
	if room, ok := c.Rooms[id]; !ok {
		return nil, errors.New("room doesn't exists")
	} else {
		return room, nil
	}
}
