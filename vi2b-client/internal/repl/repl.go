package repl

import (
	"fmt"
	"log"
	"strings"
	"github.com/chzyer/readline"
	"github.com/Explorer-art/vi2b-client/internal/core"
)

var rl *readline.Instance
var commands_callback = make(map[string]func([]string))

func setCommandCallback(command_name string, command_callback func([]string)) {
	commands_callback[command_name] = command_callback
}

func Init(l *readline.Instance) {
	rl = l

	setCommandCallback("connect", ConnectCommand)
	setCommandCallback("disconnect", DisconnectCommand)
	setCommandCallback("echo", EchoCommand)
	setCommandCallback("exit", ExitCommand)
	setCommandCallback("help", HelpCommand)
}

func Start() {
	fmt.Fprintf(rl.Stdout(), "Welcome to vi2b! Type 'help' for commands.")

	for {
		line, err := rl.Readline()
		if err != nil {
			log.Printf("Error read line: %s\n", err)
			continue
		}

		if line == "exit" {
			break
		}
		
		args := strings.Fields(line)

		_, ok := commands_callback[args[0]]

		if !ok {
			core.GetServer().SendMessage("cmd", core.SendCommand{Command: strings.Replace(line, "\n", "", -1)})
			continue
		}

		commands_callback[args[0]](args)
	}
}
