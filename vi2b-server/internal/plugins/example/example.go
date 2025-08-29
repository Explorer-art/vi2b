package example

import (
	"log"
	"github.com/gorilla/websocket"
	"github.com/Explorer-art/vi2b-server/internal/core"
)

type Plugin struct {}

func init() {
	core.PluginRegister("example", &Plugin{})
}

func (p *Plugin) OnMessage(conn *websocket.Conn, dataType string, data map[string]interface{}) {
	log.Printf("Room Example on message: %s", data["message"])
}

func (p *Plugin) Init() error {
	log.Println("Example plugin init!")
	return nil
}
