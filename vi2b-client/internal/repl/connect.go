package repl

import (
	"fmt"
	"github.com/Explorer-art/vi2b/services"
)

func ConnectCommand(args []string) {
	if len(args) < 2 {
		fmt.Println("Syntax: connect <ip:port> <password>\n")
		return
	}

	var server *services.TCPServer

	if len(args) < 3 {
		server = services.NewTCPServer(args[1], args[2])
	} else {
		server = services.NewTCPServer(args[1], "")
	}

	server.Connect()
}