package repl

import (
	"fmt"
	"github.com/Explorer-art/vi2b-client/internal/core"
)

func ConnectCommand(args []string) {
	if len(args) < 2 {
		fmt.Println("Syntax: connect <ip:port> <password>\n")
		return
	}

	var server *core.Server

	if len(args) == 3 {
		server = core.NewServer(args[1], args[2])
	} else {
		server = core.NewServer(args[1], "")
	}

	server.Connect()
}
