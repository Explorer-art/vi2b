package main

import (
	"os"
	"github.com/Explorer-art/vi2b-server/internal/core"
)

func main() {
	address := ":8000"

	if len(os.Args) > 1 {
		address = os.Args[1]
	}
	
	server := core.Init(address, "", "db.sqlite3")
	server.Start()
}
