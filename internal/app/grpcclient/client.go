package grpcclient

import (
	"context"
	"fmt"
	"log"

	"github.com/BillyBones007/pwdm_client/internal/storage"
	"google.golang.org/grpc/metadata"
)

// ClientGRPC - client structure.
type ClientGRPC struct {
	Config   *Config
	Storage  *storage.Storage
	AuthFlag bool
	Token    string
}

// InitClient - client initialization.
func InitClient() *ClientGRPC {
	config, err := setFileConfig()
	if err != nil {
		log.Fatal(err)
	}
	storage := storage.NewStorage()
	fmt.Printf("INFO: server address: %s\n", config.ServerAddr)
	return &ClientGRPC{Config: config, Storage: storage}
}

// getContext - returns a context with a token.
func (c *ClientGRPC) getContext() context.Context {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "token", c.Token)
	return ctx
}
