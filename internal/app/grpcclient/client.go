package grpcclient

import (
	"context"
	"log"

	"github.com/BillyBones007/pwdm_client/internal/storage"
	"github.com/BillyBones007/pwdm_client/internal/tools/encrypttools"
	"google.golang.org/grpc/metadata"
)

// ClientGRPC - client structure.
type ClientGRPC struct {
	Config    *Config
	Storage   *storage.Storage
	AuthFlag  bool
	Token     string
	Encrypter *encrypttools.Encrypter
}

// InitClient - client initialization.
func InitClient() *ClientGRPC {
	config, err := setFileConfig()
	if err != nil {
		log.Fatal(err)
	}
	storage := storage.NewStorage()
	return &ClientGRPC{Config: config, Storage: storage}
}

// getContext - returns a context with a token.
func (c *ClientGRPC) getContext() context.Context {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "token", c.Token)
	return ctx
}
