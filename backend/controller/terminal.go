package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"web-terminal/api"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(*http.Request) bool {
			return true
		},
	}

	Server *api.WebsocketServer
)

// TerminalController implement beego Controller
type TerminalController struct {
	beego.Controller
}

// Get handle get method for terminal controller
func (c *TerminalController) Get() {
	conn, err := upgrader.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		if _, ok := err.(websocket.HandshakeError); ok {
			c.Ctx.WriteString(fmt.Sprintf("websocket handshake error: %s", err.Error()))
		} else {
			c.Ctx.WriteString(fmt.Sprintf("websocket connection meet error: %s", err.Error()))
		}
		c.StopRun()
	}

	Server = &api.WebsocketServer{
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	go Server.WriteMessage()
	go Server.ReadMessage()
}
