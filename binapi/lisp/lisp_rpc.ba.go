// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package lisp

import (
	"context"
	"fmt"
	"io"
	"vpp-restapi/binapi/memclnt"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service lisp.
type RPCService interface {
	LispAddDelAdjacency(ctx context.Context, in *LispAddDelAdjacency) (*LispAddDelAdjacencyReply, error)
	LispAddDelLocalEid(ctx context.Context, in *LispAddDelLocalEid) (*LispAddDelLocalEidReply, error)
	LispAddDelLocator(ctx context.Context, in *LispAddDelLocator) (*LispAddDelLocatorReply, error)
	LispAddDelLocatorSet(ctx context.Context, in *LispAddDelLocatorSet) (*LispAddDelLocatorSetReply, error)
	LispAddDelMapRequestItrRlocs(ctx context.Context, in *LispAddDelMapRequestItrRlocs) (*LispAddDelMapRequestItrRlocsReply, error)
	LispAddDelMapResolver(ctx context.Context, in *LispAddDelMapResolver) (*LispAddDelMapResolverReply, error)
	LispAddDelMapServer(ctx context.Context, in *LispAddDelMapServer) (*LispAddDelMapServerReply, error)
	LispAddDelRemoteMapping(ctx context.Context, in *LispAddDelRemoteMapping) (*LispAddDelRemoteMappingReply, error)
	LispAdjacenciesGet(ctx context.Context, in *LispAdjacenciesGet) (*LispAdjacenciesGetReply, error)
	LispEidTableAddDelMap(ctx context.Context, in *LispEidTableAddDelMap) (*LispEidTableAddDelMapReply, error)
	LispEidTableDump(ctx context.Context, in *LispEidTableDump) (RPCService_LispEidTableDumpClient, error)
	LispEidTableMapDump(ctx context.Context, in *LispEidTableMapDump) (RPCService_LispEidTableMapDumpClient, error)
	LispEidTableVniDump(ctx context.Context, in *LispEidTableVniDump) (RPCService_LispEidTableVniDumpClient, error)
	LispEnableDisable(ctx context.Context, in *LispEnableDisable) (*LispEnableDisableReply, error)
	LispGetMapRequestItrRlocs(ctx context.Context, in *LispGetMapRequestItrRlocs) (*LispGetMapRequestItrRlocsReply, error)
	LispLocatorDump(ctx context.Context, in *LispLocatorDump) (RPCService_LispLocatorDumpClient, error)
	LispLocatorSetDump(ctx context.Context, in *LispLocatorSetDump) (RPCService_LispLocatorSetDumpClient, error)
	LispMapRegisterEnableDisable(ctx context.Context, in *LispMapRegisterEnableDisable) (*LispMapRegisterEnableDisableReply, error)
	LispMapRequestMode(ctx context.Context, in *LispMapRequestMode) (*LispMapRequestModeReply, error)
	LispMapResolverDump(ctx context.Context, in *LispMapResolverDump) (RPCService_LispMapResolverDumpClient, error)
	LispMapServerDump(ctx context.Context, in *LispMapServerDump) (RPCService_LispMapServerDumpClient, error)
	LispPitrSetLocatorSet(ctx context.Context, in *LispPitrSetLocatorSet) (*LispPitrSetLocatorSetReply, error)
	LispRlocProbeEnableDisable(ctx context.Context, in *LispRlocProbeEnableDisable) (*LispRlocProbeEnableDisableReply, error)
	LispUsePetr(ctx context.Context, in *LispUsePetr) (*LispUsePetrReply, error)
	ShowLispMapRegisterState(ctx context.Context, in *ShowLispMapRegisterState) (*ShowLispMapRegisterStateReply, error)
	ShowLispMapRequestMode(ctx context.Context, in *ShowLispMapRequestMode) (*ShowLispMapRequestModeReply, error)
	ShowLispPitr(ctx context.Context, in *ShowLispPitr) (*ShowLispPitrReply, error)
	ShowLispRlocProbeState(ctx context.Context, in *ShowLispRlocProbeState) (*ShowLispRlocProbeStateReply, error)
	ShowLispStatus(ctx context.Context, in *ShowLispStatus) (*ShowLispStatusReply, error)
	ShowLispUsePetr(ctx context.Context, in *ShowLispUsePetr) (*ShowLispUsePetrReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) LispAddDelAdjacency(ctx context.Context, in *LispAddDelAdjacency) (*LispAddDelAdjacencyReply, error) {
	out := new(LispAddDelAdjacencyReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispAddDelLocalEid(ctx context.Context, in *LispAddDelLocalEid) (*LispAddDelLocalEidReply, error) {
	out := new(LispAddDelLocalEidReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispAddDelLocator(ctx context.Context, in *LispAddDelLocator) (*LispAddDelLocatorReply, error) {
	out := new(LispAddDelLocatorReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispAddDelLocatorSet(ctx context.Context, in *LispAddDelLocatorSet) (*LispAddDelLocatorSetReply, error) {
	out := new(LispAddDelLocatorSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispAddDelMapRequestItrRlocs(ctx context.Context, in *LispAddDelMapRequestItrRlocs) (*LispAddDelMapRequestItrRlocsReply, error) {
	out := new(LispAddDelMapRequestItrRlocsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispAddDelMapResolver(ctx context.Context, in *LispAddDelMapResolver) (*LispAddDelMapResolverReply, error) {
	out := new(LispAddDelMapResolverReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispAddDelMapServer(ctx context.Context, in *LispAddDelMapServer) (*LispAddDelMapServerReply, error) {
	out := new(LispAddDelMapServerReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispAddDelRemoteMapping(ctx context.Context, in *LispAddDelRemoteMapping) (*LispAddDelRemoteMappingReply, error) {
	out := new(LispAddDelRemoteMappingReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispAdjacenciesGet(ctx context.Context, in *LispAdjacenciesGet) (*LispAdjacenciesGetReply, error) {
	out := new(LispAdjacenciesGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispEidTableAddDelMap(ctx context.Context, in *LispEidTableAddDelMap) (*LispEidTableAddDelMapReply, error) {
	out := new(LispEidTableAddDelMapReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispEidTableDump(ctx context.Context, in *LispEidTableDump) (RPCService_LispEidTableDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LispEidTableDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LispEidTableDumpClient interface {
	Recv() (*LispEidTableDetails, error)
	api.Stream
}

type serviceClient_LispEidTableDumpClient struct {
	api.Stream
}

func (c *serviceClient_LispEidTableDumpClient) Recv() (*LispEidTableDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *LispEidTableDetails:
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

func (c *serviceClient) LispEidTableMapDump(ctx context.Context, in *LispEidTableMapDump) (RPCService_LispEidTableMapDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LispEidTableMapDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LispEidTableMapDumpClient interface {
	Recv() (*LispEidTableMapDetails, error)
	api.Stream
}

type serviceClient_LispEidTableMapDumpClient struct {
	api.Stream
}

func (c *serviceClient_LispEidTableMapDumpClient) Recv() (*LispEidTableMapDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *LispEidTableMapDetails:
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

func (c *serviceClient) LispEidTableVniDump(ctx context.Context, in *LispEidTableVniDump) (RPCService_LispEidTableVniDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LispEidTableVniDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LispEidTableVniDumpClient interface {
	Recv() (*LispEidTableVniDetails, error)
	api.Stream
}

type serviceClient_LispEidTableVniDumpClient struct {
	api.Stream
}

func (c *serviceClient_LispEidTableVniDumpClient) Recv() (*LispEidTableVniDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *LispEidTableVniDetails:
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

func (c *serviceClient) LispEnableDisable(ctx context.Context, in *LispEnableDisable) (*LispEnableDisableReply, error) {
	out := new(LispEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispGetMapRequestItrRlocs(ctx context.Context, in *LispGetMapRequestItrRlocs) (*LispGetMapRequestItrRlocsReply, error) {
	out := new(LispGetMapRequestItrRlocsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispLocatorDump(ctx context.Context, in *LispLocatorDump) (RPCService_LispLocatorDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LispLocatorDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LispLocatorDumpClient interface {
	Recv() (*LispLocatorDetails, error)
	api.Stream
}

type serviceClient_LispLocatorDumpClient struct {
	api.Stream
}

func (c *serviceClient_LispLocatorDumpClient) Recv() (*LispLocatorDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *LispLocatorDetails:
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

func (c *serviceClient) LispLocatorSetDump(ctx context.Context, in *LispLocatorSetDump) (RPCService_LispLocatorSetDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LispLocatorSetDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LispLocatorSetDumpClient interface {
	Recv() (*LispLocatorSetDetails, error)
	api.Stream
}

type serviceClient_LispLocatorSetDumpClient struct {
	api.Stream
}

func (c *serviceClient_LispLocatorSetDumpClient) Recv() (*LispLocatorSetDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *LispLocatorSetDetails:
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

func (c *serviceClient) LispMapRegisterEnableDisable(ctx context.Context, in *LispMapRegisterEnableDisable) (*LispMapRegisterEnableDisableReply, error) {
	out := new(LispMapRegisterEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispMapRequestMode(ctx context.Context, in *LispMapRequestMode) (*LispMapRequestModeReply, error) {
	out := new(LispMapRequestModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispMapResolverDump(ctx context.Context, in *LispMapResolverDump) (RPCService_LispMapResolverDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LispMapResolverDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LispMapResolverDumpClient interface {
	Recv() (*LispMapResolverDetails, error)
	api.Stream
}

type serviceClient_LispMapResolverDumpClient struct {
	api.Stream
}

func (c *serviceClient_LispMapResolverDumpClient) Recv() (*LispMapResolverDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *LispMapResolverDetails:
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

func (c *serviceClient) LispMapServerDump(ctx context.Context, in *LispMapServerDump) (RPCService_LispMapServerDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_LispMapServerDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_LispMapServerDumpClient interface {
	Recv() (*LispMapServerDetails, error)
	api.Stream
}

type serviceClient_LispMapServerDumpClient struct {
	api.Stream
}

func (c *serviceClient_LispMapServerDumpClient) Recv() (*LispMapServerDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *LispMapServerDetails:
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

func (c *serviceClient) LispPitrSetLocatorSet(ctx context.Context, in *LispPitrSetLocatorSet) (*LispPitrSetLocatorSetReply, error) {
	out := new(LispPitrSetLocatorSetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispRlocProbeEnableDisable(ctx context.Context, in *LispRlocProbeEnableDisable) (*LispRlocProbeEnableDisableReply, error) {
	out := new(LispRlocProbeEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) LispUsePetr(ctx context.Context, in *LispUsePetr) (*LispUsePetrReply, error) {
	out := new(LispUsePetrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowLispMapRegisterState(ctx context.Context, in *ShowLispMapRegisterState) (*ShowLispMapRegisterStateReply, error) {
	out := new(ShowLispMapRegisterStateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowLispMapRequestMode(ctx context.Context, in *ShowLispMapRequestMode) (*ShowLispMapRequestModeReply, error) {
	out := new(ShowLispMapRequestModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowLispPitr(ctx context.Context, in *ShowLispPitr) (*ShowLispPitrReply, error) {
	out := new(ShowLispPitrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowLispRlocProbeState(ctx context.Context, in *ShowLispRlocProbeState) (*ShowLispRlocProbeStateReply, error) {
	out := new(ShowLispRlocProbeStateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowLispStatus(ctx context.Context, in *ShowLispStatus) (*ShowLispStatusReply, error) {
	out := new(ShowLispStatusReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) ShowLispUsePetr(ctx context.Context, in *ShowLispUsePetr) (*ShowLispUsePetrReply, error) {
	out := new(ShowLispUsePetrReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
