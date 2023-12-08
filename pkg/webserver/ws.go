package webserver

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"webchat/pkg/webserver/routes"
)

func New() http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/api/room", routes.HandleRoom)
	r.HandleFunc("/api/new_room", routes.HandleNewRoom)
	r.PathPrefix("/").Handler(routes.StaticPublicDir("/"))

	return http.Server{
		Addr:              "127.0.0.1:8000",
		Handler:           r,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}
}
