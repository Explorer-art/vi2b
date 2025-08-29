package repl

import (
	"fmt"
	"strings"
	"github.com/Explorer-art/vi2b-client/internal/core"
)

type Message struct {
	Message 	string `json:"message"`
}

func SayCommand(args []string) {
	if len(args) < 2 {
		fmt.Println("Syntax: say <message>")
		return
	}

	core.GetServer().SendMessage("chat", Message{Message: strings.Join(args[1:], " ")})
}