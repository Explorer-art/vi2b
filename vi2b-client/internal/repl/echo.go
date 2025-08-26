package repl

import (
	"fmt"
	"strings"
)

func EchoCommand(args []string) {
	args = append(args[:0], args[1:]...)
	fmt.Println(strings.Join(args, " "))
}
