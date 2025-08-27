package repl

import (
	"fmt"
)

func HelpCommand(args []string) {
	fmt.Println("Commands:\nserver\necho\nexit\nhelp\n")
}