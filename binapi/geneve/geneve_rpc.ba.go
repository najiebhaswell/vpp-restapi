// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package geneve

import (
	"context"
	"fmt"
	"io"
	"vpp-restapi/binapi/memclnt"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service geneve.
type RPCService interface {
	GeneveAddDelTunnel(ctx context.Context, in *GeneveAddDelTunnel) (*GeneveAddDelTunnelReply, error)
	GeneveAddDelTunnel2(ctx context.Context, in *GeneveAddDelTunnel2) (*GeneveAddDelTunnel2Reply, error)
	GeneveTunnelDump(ctx context.Context, in *GeneveTunnelDump) (RPCService_GeneveTunnelDumpClient, error)
	SwInterfaceSetGeneveBypass(ctx context.Context, in *SwInterfaceSetGeneveBypass) (*SwInterfaceSetGeneveBypassReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) GeneveAddDelTunnel(ctx context.Context, in *GeneveAddDelTunnel) (*GeneveAddDelTunnelReply, error) {
	out := new(GeneveAddDelTunnelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GeneveAddDelTunnel2(ctx context.Context, in *GeneveAddDelTunnel2) (*GeneveAddDelTunnel2Reply, error) {
	out := new(GeneveAddDelTunnel2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GeneveTunnelDump(ctx context.Context, in *GeneveTunnelDump) (RPCService_GeneveTunnelDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GeneveTunnelDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GeneveTunnelDumpClient interface {
	Recv() (*GeneveTunnelDetails, error)
	api.Stream
}

type serviceClient_GeneveTunnelDumpClient struct {
	api.Stream
}

func (c *serviceClient_GeneveTunnelDumpClient) Recv() (*GeneveTunnelDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GeneveTunnelDetails:
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

func (c *serviceClient) SwInterfaceSetGeneveBypass(ctx context.Context, in *SwInterfaceSetGeneveBypass) (*SwInterfaceSetGeneveBypassReply, error) {
	out := new(SwInterfaceSetGeneveBypassReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
