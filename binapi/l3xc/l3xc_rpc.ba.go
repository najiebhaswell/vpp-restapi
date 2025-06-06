// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package l3xc

import (
	"context"
	"fmt"
	"io"
	"vpp-restapi/binapi/memclnt"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service l3xc.
type RPCService interface {
	L3xcDel(ctx context.Context, in *L3xcDel) (*L3xcDelReply, error)
	L3xcDump(ctx context.Context, in *L3xcDump) (RPCService_L3xcDumpClient, error)
	L3xcPluginGetVersion(ctx context.Context, in *L3xcPluginGetVersion) (*L3xcPluginGetVersionReply, error)
	L3xcUpdate(ctx context.Context, in *L3xcUpdate) (*L3xcUpdateReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) L3xcDel(ctx context.Context, in *L3xcDel) (*L3xcDelReply, error) {
	out := new(L3xcDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L3xcDump(ctx context.Context, in *L3xcDump) (RPCService_L3xcDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_L3xcDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_L3xcDumpClient interface {
	Recv() (*L3xcDetails, error)
	api.Stream
}

type serviceClient_L3xcDumpClient struct {
	api.Stream
}

func (c *serviceClient_L3xcDumpClient) Recv() (*L3xcDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *L3xcDetails:
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

func (c *serviceClient) L3xcPluginGetVersion(ctx context.Context, in *L3xcPluginGetVersion) (*L3xcPluginGetVersionReply, error) {
	out := new(L3xcPluginGetVersionReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) L3xcUpdate(ctx context.Context, in *L3xcUpdate) (*L3xcUpdateReply, error) {
	out := new(L3xcUpdateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
