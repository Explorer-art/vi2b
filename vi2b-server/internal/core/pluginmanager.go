package core

import (
	"log"
	"strings"
	"github.com/gorilla/websocket"
)

type Plugin interface {
	Init() error
	OnMessage(conn *websocket.Conn, dataType string, data map[string]interface{})
}

type PermissionsGroup struct {
	Name string
	Permissions []string
}

type Command struct {
	Name 			string
	FuncCallback	func (client *Client, args []string)
	Permission		string
	IsDefault		bool
}

var plugins map[string]Plugin
var permissionsGroups map[string]PermissionsGroup
var commands map[string]Command

// func checkPermission(permissionsGroup string, permission string) bool {
// 	for _, v := range permissionsGroups {
// 		if v == permission {
// 			return true
// 		}
// 	}

// 	return false
// }

func RegisterPlugin(name string, p Plugin) {
	plugins[name] = p
	log.Printf("Plugin %s registered\n", name)
}

func RegisterCommand(commandName string, funcCallback func (client *Client, args []string), permission string, isDefault bool) {
	commands[commandName] = Command{
		Name: commandName,
		FuncCallback: funcCallback,
		Permission: permission,
		IsDefault: isDefault,
	}

	log.Printf("Command %s registered\n", commandName)
}

func ParseCommand(client *Client, command string) {
	if len(command) < 1 {
		client.Close()
		return
	}

	commandSplited := strings.Split(command, " ")

	commandName := commandSplited[0]
	args := commandSplited[1:]

	_, ok := commands[commandName]
	if !ok {
		client.SendMessage("chat", ChatMessage{Message: "Unknown command"})
		return
	}

	if commands[commandName].IsDefault {
		commands[commandName].FuncCallback(client, args)
		return
	}

	// if checkPermission(client.PermissionsGroup, commands[commandName].Permission) {
	// 	log.Println("Success")
	// 	commands[commandName].FuncCallback(client, args)
	// } else {
	// 	client.SendMessage("chat", ChatMessage{Message: "You don't have permission!"})
	// }
}

func init() {
	plugins = make(map[string]Plugin)
	permissionsGroups = make(map[string]PermissionsGroup)
	commands = make(map[string]Command)

	permissionsGroups["default"] = PermissionsGroup{
		Name: "default",
		Permissions: []string{
			"",
		},
	}

	permissionsGroups["admin"] = PermissionsGroup{
		Name: "admin",
		Permissions: []string{
			"kick",
			"mute",
			"ban",
		},
	}
}