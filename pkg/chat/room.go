package chat

import (
	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID
	Owner     User
	Listeners []*Listener
	Queue     chan Message
}

type Listener struct {
	User      User
	onMessage MessageCallback
}

type MessageCallback func(message Message)

func (r *Room) AddUser(user User) *Listener {
	listener := &Listener{
		User: user,
	}
	r.Listeners = append(r.Listeners, listener)
	return listener
}

func (r *Room) RemoveUser(user User) {
	for i, u := range r.Listeners {
		if u.User == user {
			r.Listeners = append(r.Listeners[:i], r.Listeners[i+1:]...)
		}
	}
}

func (r *Room) Listen() {
	for {
		message := <-r.Queue
		for _, listener := range r.Listeners {
			go listener.onMessage(message)
		}
	}
}

func (l *Listener) OnMessage(onMessage MessageCallback) {
	l.onMessage = onMessage
}
