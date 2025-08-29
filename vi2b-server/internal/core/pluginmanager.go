package core

import (
	"log"
	"github.com/gorilla/websocket"
)

type Plugin interface {
	Init() error
	OnMessage(conn *websocket.Conn, dataType string, data map[string]interface{})
}

var plugins map[string]Plugin

func init() {
	plugins = make(map[string]Plugin)
}

func PluginRegister(name string, p Plugin) {
	plugins[name] = p
	log.Printf("Plugin %s registered\n", name)
}
