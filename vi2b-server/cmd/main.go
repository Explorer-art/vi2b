package main

import (
	"github.com/Explorer-art/vi2b-server/internal/core"
)

func main() {
	server := core.NewServer(":8000", "")
	server.Start()
}
