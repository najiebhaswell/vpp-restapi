// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package urpf

import (
	"context"
	"fmt"
	"io"
	"vpp-restapi/binapi/memclnt"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service urpf.
type RPCService interface {
	UrpfInterfaceDump(ctx context.Context, in *UrpfInterfaceDump) (RPCService_UrpfInterfaceDumpClient, error)
	UrpfUpdate(ctx context.Context, in *UrpfUpdate) (*UrpfUpdateReply, error)
	UrpfUpdateV2(ctx context.Context, in *UrpfUpdateV2) (*UrpfUpdateV2Reply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) UrpfInterfaceDump(ctx context.Context, in *UrpfInterfaceDump) (RPCService_UrpfInterfaceDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_UrpfInterfaceDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_UrpfInterfaceDumpClient interface {
	Recv() (*UrpfInterfaceDetails, error)
	api.Stream
}

type serviceClient_UrpfInterfaceDumpClient struct {
	api.Stream
}

func (c *serviceClient_UrpfInterfaceDumpClient) Recv() (*UrpfInterfaceDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *UrpfInterfaceDetails:
		return m, nil
	case *memclnt.ControlPingReply:
		err = c.Stream.Close()
		if err != nil {
			return nil, err
		}
		return nil, io.EOF
	default:
		return nil, fmt.Errorf("unexpected message: %T %v", m, m)
	}
}

func (c *serviceClient) UrpfUpdate(ctx context.Context, in *UrpfUpdate) (*UrpfUpdateReply, error) {
	out := new(UrpfUpdateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) UrpfUpdateV2(ctx context.Context, in *UrpfUpdateV2) (*UrpfUpdateV2Reply, error) {
	out := new(UrpfUpdateV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
