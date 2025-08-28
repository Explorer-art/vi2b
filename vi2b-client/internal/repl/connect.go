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

	var client *core.Client

	if len(args) == 3 {
		client = core.NewClient(args[1], args[2])
	} else {
		client = core.NewClient(args[1], "")
	}

	client.Connect()
}