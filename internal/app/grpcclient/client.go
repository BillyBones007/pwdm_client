package grpcclient

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"

	"github.com/BillyBones007/pwdm_client/internal/tools/encrypttools"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// ClientGRPC - client structure.
type ClientGRPC struct {
	Config    *Config
	Encrypter *encrypttools.Encrypter
	creds     credentials.TransportCredentials
	Token     string
	AuthFlag  bool
}

// InitClient - client initialization.
func InitClient() *ClientGRPC {
	config, err := setConfig()
	if err != nil {
		log.Fatal(err)
	}
	caCert, err := os.ReadFile("cert/ca.crt")
	if err != nil {
		log.Fatalf("failed to load CA certificate: %s", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		RootCAs: caCertPool,
	}
	creds := credentials.NewTLS(tlsConfig)
	return &ClientGRPC{Config: config, creds: creds}
}

// getContext - returns a context with a token.
func (c *ClientGRPC) getContext() context.Context {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "token", c.Token)
	return ctx
}
