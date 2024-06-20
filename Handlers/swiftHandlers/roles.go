package swifthandlers

import (
	"log"

	"github.com/gorilla/websocket"
)

func FetchData(conn *websocket.Conn, message []byte) {
	log.Println("Received message for /fetch/data")
}
