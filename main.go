package main

import (
	// "io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var updgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/", test)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Printf("Error: %v", err.Error())
	}
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error read: %v", err.Error())
			return
		}

		log.Println(string(p))

		if err := conn.WriteMessage(messageType, []byte("Hi there from server")); err != nil {
			log.Printf("Error write: %v", err.Error())
			return
		}
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	updgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := updgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrade: %v", err.Error())
	}

	log.Println("CLient Succes Connected....")

	reader(ws)
}
