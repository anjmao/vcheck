package client

import (
	"context"
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewGRPC(target string, method string) Client {
	return &grpcClient{target: target, method: method}
}

type grpcClient struct {
	target string
	method string
}

func (c grpcClient) GetVersion(ctx context.Context) (*GetVersionReply, error) {
	creds := credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: true,
	})

	conn, err := grpc.Dial(c.target, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	in := new(GetVersionRequest)
	out := new(GetVersionReply)
	err = conn.Invoke(ctx, c.method, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
