package repl

import (
	"github.com/Explorer-art/vi2b-client/internal/core"
)

func DisconnectCommand(args []string) {
	core.GetServer().Disconnect()
}
