package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func homeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func registerRoutes() {
	http.HandleFunc("/", homeRoute)
}

func main() {
	registerRoutes()
	http.ListenAndServe(":80", nil)
}
