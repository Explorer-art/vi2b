package main

import (
	"os"
	"github.com/Explorer-art/vi2b-server/internal/core"
)

func main() {
	address := ":8080"

	if len(os.Args) > 1 {
		address = os.Args[1]
	}
	
	server := core.NewServer(address, "")
	server.Start()
}
