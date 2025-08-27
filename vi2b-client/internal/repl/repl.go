package repl

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

var commands_callback = make(map[string]func([]string))

func setCommandCallback(command_name string, command_callback func([]string)) {
	commands_callback[command_name] = command_callback
}

func Start() {
	setCommandCallback("connect", ConnectCommand)
	setCommandCallback("disconnect", DisconnectCommand)
	setCommandCallback("echo", EchoCommand)
	setCommandCallback("exit", ExitCommand)
	setCommandCallback("help", HelpCommand)
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to vi2b! Type 'help' for commands.")
	
	for {
		fmt.Print("> ")
		
		if !scanner.Scan() {
			break
		}
		
		input := scanner.Text()
		args := strings.Fields(input)

		_, ok := commands_callback[args[0]]

		if !ok {
			fmt.Println("Unknown command! Type 'help' for commands.")
			continue
		}

		commands_callback[args[0]](args)
	}
}
