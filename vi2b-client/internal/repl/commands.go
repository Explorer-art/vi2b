package repl

import (
	"fmt"
	"os"
	"strings"
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

func DisconnectCommand(args []string) {
	core.GetServer().Disconnect()
}

func EchoCommand(args []string) {
	args = append(args[:0], args[1:]...)
	fmt.Println(strings.Join(args, " "))
}

func ExitCommand(args []string) {
	os.Exit(0)
}

func HelpCommand(args []string) {
	fmt.Println("Commands:\nserver\necho\nexit\nhelp\n")
}