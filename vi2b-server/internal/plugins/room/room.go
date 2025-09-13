package room

import (
	"log"

	"github.com/Explorer-art/vi2b-server/core"
	"github.com/gorilla/websocket"
)

type Plugin struct{}

func helloCallback(client *core.Client, args []string) {
	client.SendMessage("chat", SendChatMessage{Message: "Hello!"})
}

func init() {
	core.RegisterPlugin("room", &Plugin{})
}

func (p *Plugin) OnMessage(conn *websocket.Conn, dataType string, data map[string]interface{}) {
	log.Printf("Room Example on message: %s", data["message"])
}

func (p *Plugin) Init() error {
	RegisterCommand("hello", helloCallback, "hello", true)
	return nil
}
