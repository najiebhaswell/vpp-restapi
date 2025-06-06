// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package sr

import (
	"context"
	"fmt"
	"io"
	"vpp-restapi/binapi/memclnt"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service sr.
type RPCService interface {
	SrLocalsidAddDel(ctx context.Context, in *SrLocalsidAddDel) (*SrLocalsidAddDelReply, error)
	SrLocalsidsDump(ctx context.Context, in *SrLocalsidsDump) (RPCService_SrLocalsidsDumpClient, error)
	SrLocalsidsWithPacketStatsDump(ctx context.Context, in *SrLocalsidsWithPacketStatsDump) (RPCService_SrLocalsidsWithPacketStatsDumpClient, error)
	SrPoliciesDump(ctx context.Context, in *SrPoliciesDump) (RPCService_SrPoliciesDumpClient, error)
	SrPoliciesV2Dump(ctx context.Context, in *SrPoliciesV2Dump) (RPCService_SrPoliciesV2DumpClient, error)
	SrPoliciesWithSlIndexDump(ctx context.Context, in *SrPoliciesWithSlIndexDump) (RPCService_SrPoliciesWithSlIndexDumpClient, error)
	SrPolicyAdd(ctx context.Context, in *SrPolicyAdd) (*SrPolicyAddReply, error)
	SrPolicyAddV2(ctx context.Context, in *SrPolicyAddV2) (*SrPolicyAddV2Reply, error)
	SrPolicyDel(ctx context.Context, in *SrPolicyDel) (*SrPolicyDelReply, error)
	SrPolicyMod(ctx context.Context, in *SrPolicyMod) (*SrPolicyModReply, error)
	SrPolicyModV2(ctx context.Context, in *SrPolicyModV2) (*SrPolicyModV2Reply, error)
	SrSetEncapHopLimit(ctx context.Context, in *SrSetEncapHopLimit) (*SrSetEncapHopLimitReply, error)
	SrSetEncapSource(ctx context.Context, in *SrSetEncapSource) (*SrSetEncapSourceReply, error)
	SrSteeringAddDel(ctx context.Context, in *SrSteeringAddDel) (*SrSteeringAddDelReply, error)
	SrSteeringPolDump(ctx context.Context, in *SrSteeringPolDump) (RPCService_SrSteeringPolDumpClient, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) SrLocalsidAddDel(ctx context.Context, in *SrLocalsidAddDel) (*SrLocalsidAddDelReply, error) {
	out := new(SrLocalsidAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SrLocalsidsDump(ctx context.Context, in *SrLocalsidsDump) (RPCService_SrLocalsidsDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SrLocalsidsDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SrLocalsidsDumpClient interface {
	Recv() (*SrLocalsidsDetails, error)
	api.Stream
}

type serviceClient_SrLocalsidsDumpClient struct {
	api.Stream
}

func (c *serviceClient_SrLocalsidsDumpClient) Recv() (*SrLocalsidsDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SrLocalsidsDetails:
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

func (c *serviceClient) SrLocalsidsWithPacketStatsDump(ctx context.Context, in *SrLocalsidsWithPacketStatsDump) (RPCService_SrLocalsidsWithPacketStatsDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SrLocalsidsWithPacketStatsDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SrLocalsidsWithPacketStatsDumpClient interface {
	Recv() (*SrLocalsidsWithPacketStatsDetails, error)
	api.Stream
}

type serviceClient_SrLocalsidsWithPacketStatsDumpClient struct {
	api.Stream
}

func (c *serviceClient_SrLocalsidsWithPacketStatsDumpClient) Recv() (*SrLocalsidsWithPacketStatsDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SrLocalsidsWithPacketStatsDetails:
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

func (c *serviceClient) SrPoliciesDump(ctx context.Context, in *SrPoliciesDump) (RPCService_SrPoliciesDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SrPoliciesDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SrPoliciesDumpClient interface {
	Recv() (*SrPoliciesDetails, error)
	api.Stream
}

type serviceClient_SrPoliciesDumpClient struct {
	api.Stream
}

func (c *serviceClient_SrPoliciesDumpClient) Recv() (*SrPoliciesDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SrPoliciesDetails:
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

func (c *serviceClient) SrPoliciesV2Dump(ctx context.Context, in *SrPoliciesV2Dump) (RPCService_SrPoliciesV2DumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SrPoliciesV2DumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SrPoliciesV2DumpClient interface {
	Recv() (*SrPoliciesV2Details, error)
	api.Stream
}

type serviceClient_SrPoliciesV2DumpClient struct {
	api.Stream
}

func (c *serviceClient_SrPoliciesV2DumpClient) Recv() (*SrPoliciesV2Details, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SrPoliciesV2Details:
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

func (c *serviceClient) SrPoliciesWithSlIndexDump(ctx context.Context, in *SrPoliciesWithSlIndexDump) (RPCService_SrPoliciesWithSlIndexDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SrPoliciesWithSlIndexDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SrPoliciesWithSlIndexDumpClient interface {
	Recv() (*SrPoliciesWithSlIndexDetails, error)
	api.Stream
}

type serviceClient_SrPoliciesWithSlIndexDumpClient struct {
	api.Stream
}

func (c *serviceClient_SrPoliciesWithSlIndexDumpClient) Recv() (*SrPoliciesWithSlIndexDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SrPoliciesWithSlIndexDetails:
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

func (c *serviceClient) SrPolicyAdd(ctx context.Context, in *SrPolicyAdd) (*SrPolicyAddReply, error) {
	out := new(SrPolicyAddReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SrPolicyAddV2(ctx context.Context, in *SrPolicyAddV2) (*SrPolicyAddV2Reply, error) {
	out := new(SrPolicyAddV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SrPolicyDel(ctx context.Context, in *SrPolicyDel) (*SrPolicyDelReply, error) {
	out := new(SrPolicyDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SrPolicyMod(ctx context.Context, in *SrPolicyMod) (*SrPolicyModReply, error) {
	out := new(SrPolicyModReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SrPolicyModV2(ctx context.Context, in *SrPolicyModV2) (*SrPolicyModV2Reply, error) {
	out := new(SrPolicyModV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SrSetEncapHopLimit(ctx context.Context, in *SrSetEncapHopLimit) (*SrSetEncapHopLimitReply, error) {
	out := new(SrSetEncapHopLimitReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SrSetEncapSource(ctx context.Context, in *SrSetEncapSource) (*SrSetEncapSourceReply, error) {
	out := new(SrSetEncapSourceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SrSteeringAddDel(ctx context.Context, in *SrSteeringAddDel) (*SrSteeringAddDelReply, error) {
	out := new(SrSteeringAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SrSteeringPolDump(ctx context.Context, in *SrSteeringPolDump) (RPCService_SrSteeringPolDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_SrSteeringPolDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_SrSteeringPolDumpClient interface {
	Recv() (*SrSteeringPolDetails, error)
	api.Stream
}

type serviceClient_SrSteeringPolDumpClient struct {
	api.Stream
}

func (c *serviceClient_SrSteeringPolDumpClient) Recv() (*SrSteeringPolDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *SrSteeringPolDetails:
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
