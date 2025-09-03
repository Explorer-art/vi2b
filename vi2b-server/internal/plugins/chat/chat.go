package chat

import (
	"strings"
	"github.com/gorilla/websocket"
	"github.com/Explorer-art/vi2b-server/internal/core"
)

type Plugin struct {}

func sayCallback(client *core.Client, args []string) {
	for _, client := range core.ClientsData {
		client.SendMessage("chat", core.ChatMessage{Message: strings.Join(args, " ")})
	}
}

func init() {
	core.RegisterPlugin("chat", &Plugin{})
}

func (p *Plugin) OnMessage(conn *websocket.Conn, dataType string, data map[string]interface{}) {

}

func (p *Plugin) Init() error {
	core.RegisterCommand("say", sayCallback, "say", true)
	return nil
}