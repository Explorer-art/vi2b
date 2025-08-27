package repl

import (
	"github.com/Explorer-art/vi2b/services"
)

func DisconnectCommand(args []string) {
	services.GetServer().Disconnect()
}