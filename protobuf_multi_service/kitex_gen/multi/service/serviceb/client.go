// Code generated by Kitex v0.9.0. DO NOT EDIT.

package serviceb

import (
	"context"
	service "github.com/cloudwego/kitex-examples/protobuf_multi_service/kitex_gen/multi/service"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ChatB(ctx context.Context, Req *service.RequestB, callOptions ...callopt.Option) (r *service.Reply, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kServiceBClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kServiceBClient struct {
	*kClient
}

func (p *kServiceBClient) ChatB(ctx context.Context, Req *service.RequestB, callOptions ...callopt.Option) (r *service.Reply, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ChatB(ctx, Req)
}
