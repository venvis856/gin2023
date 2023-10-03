package jump_proxy_rpc_sdk

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"gin/internal/app/library/jump_proxy_rpc_sdk/auth"
	"gin/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"os"
	"sync"
	"time"
)

var (
	Conn       *grpc.ClientConn
	once       sync.Once
	ConnMutex  sync.Mutex
	ConnErr    error
	ConnErrMux sync.Mutex
)

func RPCConnection() (*grpc.ClientConn, error) {
	once.Do(func() {
		creds, err := DoubleCreds()
		if err != nil {
			ConnErrMux.Lock()
			ConnErr = err
			ConnErrMux.Unlock()
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()
		conn, err := grpc.DialContext(ctx, config.Cfg.Rpc.Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(new(auth.ClientToken)))
		if err != nil {
			ConnErrMux.Lock()
			ConnErr = err
			ConnErrMux.Unlock()
			return
		}

		ConnMutex.Lock()
		Conn = conn
		ConnMutex.Unlock()
	})

	ConnErrMux.Lock()
	defer ConnErrMux.Unlock()
	return Conn, ConnErr
}

// DoubleCreds loads the client certificate and constructs the transport credentials
func DoubleCreds() (credentials.TransportCredentials, error) {
	// Load client certificate
	fmt.Println(config.Cfg.Rpc.ClientPem, "==config.Cfg.Rpc.ClientPem")
	certificate, err := tls.LoadX509KeyPair(config.Cfg.Rpc.ClientPem, config.Cfg.Rpc.ClientKey)
	if err != nil {
		return nil, err
	}
	// Build CertPool to validate server certificate
	certPool := x509.NewCertPool()
	ca, err := os.ReadFile(config.Cfg.Rpc.CaPem)
	if err != nil {
		return nil, err
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   config.Cfg.Rpc.ServerName,
		RootCAs:      certPool,
	})

	return creds, nil
}

func GetCtx() context.Context {
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("X-GRPC", "grpc"))
	return ctx
}
