package example

import (
	"github.com/gorilla/websocket"
	"github.com/Explorer-art/vi2b-server/internal/core"
)

type Plugin struct {}

func helloCallback(client *core.Client, args []string) {
	client.SendMessage("chat", core.SendChatMessage{Message: "Hello!"})
}

func init() {
	core.RegisterPlugin("example", &Plugin{})
}

func (p *Plugin) OnMessage(conn *websocket.Conn, dataType string, data map[string]interface{}) {
}

func (p *Plugin) Init() error {
	core.RegisterCommand("hello", helloCallback, "hello", true)
	return nil
}
