package main

import (
	"log"
	"webchat/pkg/webserver"
)

func main() {
	server := webserver.New()
	log.Fatal(server.ListenAndServe())
}
