package main

import (
	"log"
	"os"
	"os/signal"
	"github.com/chzyer/readline"
	"github.com/Explorer-art/vi2b-client/internal/core"
	"github.com/Explorer-art/vi2b-client/internal/repl"
)

func main() {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:				"> ",
		InterruptPrompt:	"^C",
		EOFPrompt:			"exit",
	})
	if err != nil {
		log.Fatal(err)
	}
	defer rl.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	core.Init(rl)
	repl.Init(rl)
	repl.Start()
}
