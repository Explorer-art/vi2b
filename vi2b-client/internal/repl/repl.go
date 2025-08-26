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
	setCommandCallback("echo", EchoCommand)
	setCommandCallback("exit", ExitCommand)
	
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to vi2b! Type 'help' for commands.")
	
	for {
		fmt.Print("> ")
		
		if !scanner.Scan() {
			break
		}
		
		input := scanner.Text()
		args := strings.Fields(input)

		commands_callback[args[0]](args)
	}
}
