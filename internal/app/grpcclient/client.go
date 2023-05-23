package grpcclient

import (
	"context"
	"log"

	"github.com/BillyBones007/pwdm_client/internal/tools/encrypttools"
	"google.golang.org/grpc/metadata"
)

// ClientGRPC - client structure.
type ClientGRPC struct {
	Config    *Config
	Encrypter *encrypttools.Encrypter
	Token     string
	AuthFlag  bool
}

// InitClient - client initialization.
func InitClient() *ClientGRPC {
	config, err := setConfig()
	if err != nil {
		log.Fatal(err)
	}
	return &ClientGRPC{Config: config}
}

// getContext - returns a context with a token.
func (c *ClientGRPC) getContext() context.Context {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "token", c.Token)
	return ctx
}
