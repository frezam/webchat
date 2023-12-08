package routes

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"webchat/pkg/chat"
	"webchat/pkg/utils"
)

func HandleRoom(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")
	if username == "" {
		_, _ = w.Write(utils.GetInternalError())
		return
	}

	roomID, errorParseId := uuid.Parse(r.URL.Query().Get("id"))
	if errorParseId != nil {
		_, _ = w.Write(utils.GetInternalError())
		return
	}

	currentUser := chat.NewUser(username)

	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer ws.Close()

	room, errGetRoom := chat.CHAT.GetRoom(roomID.String())
	if errGetRoom != nil {
		response := utils.Response{
			Error:   true,
			Code:    404,
			Message: errGetRoom.Error(),
			Data:    nil,
		}
		_ = ws.WriteJSON(response)
		return
	}

	listener := room.AddUser(currentUser)

	closed := make(chan bool)
	go func() {
		for {
			_, message, errRead := ws.ReadMessage()
			if errRead != nil {
				closed <- true
				continue
			}
			room.Queue <- chat.Message{
				ID:      uuid.New(),
				RoomID:  room.ID,
				User:    currentUser,
				Content: string(message),
			}
		}
	}()

	go listener.OnMessage(func(message chat.Message) {
		_ = ws.WriteJSON(message)
	})

	<-closed

	room.RemoveUser(currentUser)
}

func HandleNewRoom(w http.ResponseWriter, r *http.Request) {
	user := chat.NewUser("anonymous")
	room := chat.CHAT.CreateRoom(user)
	go room.Listen()
	http.Redirect(w, r, fmt.Sprintf("/room.html?id=%s&username=%s", room.ID.String(), user.Name), 307)
}
