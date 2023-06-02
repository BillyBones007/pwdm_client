package uishell

import (
	"strings"

	"github.com/BillyBones007/pwdm_client/internal/app/grpcclient"
	"github.com/c-bata/go-prompt"
)

// Completer - main complete structure.
type Completer struct {
	client *grpcclient.ClientGRPC
}

// NewCompleter - constructor.
func NewCompleter(client *grpcclient.ClientGRPC) *Completer {
	return &Completer{client: client}
}

// Complete - main completer.
func (c *Completer) Complete(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")

	return c.argumentsCompleter(args)
}
