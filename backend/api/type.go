package api

import (
	"bytes"
	"log"

	"github.com/gorilla/websocket"
)

var (
	// WebsocketServers will contain all online websocket
	WebsocketServers map[string]*WebsocketServer
)

// WebsocketServer extends gorilla websocket
type WebsocketServer struct {
	Conn *websocket.Conn
	Send chan []byte
}

// WriteMessage write message through websocket
func (ws *WebsocketServer) WriteMessage() {
	defer ws.Conn.Close()
	for {
		select {
		case message, ok := <-ws.Send:
			log.Printf("receive message: %v", message)
			if !ok {
				_ = ws.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			_ = ws.Conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

// ReadMessage read message from websocket client
func (ws *WebsocketServer) ReadMessage() {
	defer ws.Conn.Close()
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			log.Printf("error: %v", err)
			break
		}
		if bytes.Equal(message, []byte{uint8(13)}){
			message = []byte("\r\n")
		}
		log.Printf("message: %v", message)
		ws.Send <- message
	}
}
