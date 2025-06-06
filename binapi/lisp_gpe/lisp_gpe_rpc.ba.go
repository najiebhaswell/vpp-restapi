// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package lisp_gpe

import (
	"context"
	"fmt"
	"io"
	"vpp-restapi/binapi/memclnt"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service lisp_gpe.
type RPCService interface {
	GpeAddDelFwdEntry(ctx context.Context, in *GpeAddDelFwdEntry) (*GpeAddDelFwdEntryReply, error)
	GpeAddDelIface(ctx context.Context, in *GpeAddDelIface) (*GpeAddDelIfaceReply, error)
	GpeAddDelNativeFwdRpath(ctx context.Context, in *GpeAddDelNativeFwdRpath) (*GpeAddDelNativeFwdRpathReply, error)
	GpeEnableDisable(ctx context.Context, in *GpeEnableDisable) (*GpeEnableDisableReply, error)
	GpeFwdEntriesGet(ctx context.Context, in *GpeFwdEntriesGet) (*GpeFwdEntriesGetReply, error)
	GpeFwdEntryPathDump(ctx context.Context, in *GpeFwdEntryPathDump) (RPCService_GpeFwdEntryPathDumpClient, error)
	GpeFwdEntryVnisGet(ctx context.Context, in *GpeFwdEntryVnisGet) (*GpeFwdEntryVnisGetReply, error)
	GpeGetEncapMode(ctx context.Context, in *GpeGetEncapMode) (*GpeGetEncapModeReply, error)
	GpeNativeFwdRpathsGet(ctx context.Context, in *GpeNativeFwdRpathsGet) (*GpeNativeFwdRpathsGetReply, error)
	GpeSetEncapMode(ctx context.Context, in *GpeSetEncapMode) (*GpeSetEncapModeReply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) GpeAddDelFwdEntry(ctx context.Context, in *GpeAddDelFwdEntry) (*GpeAddDelFwdEntryReply, error) {
	out := new(GpeAddDelFwdEntryReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GpeAddDelIface(ctx context.Context, in *GpeAddDelIface) (*GpeAddDelIfaceReply, error) {
	out := new(GpeAddDelIfaceReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GpeAddDelNativeFwdRpath(ctx context.Context, in *GpeAddDelNativeFwdRpath) (*GpeAddDelNativeFwdRpathReply, error) {
	out := new(GpeAddDelNativeFwdRpathReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GpeEnableDisable(ctx context.Context, in *GpeEnableDisable) (*GpeEnableDisableReply, error) {
	out := new(GpeEnableDisableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GpeFwdEntriesGet(ctx context.Context, in *GpeFwdEntriesGet) (*GpeFwdEntriesGetReply, error) {
	out := new(GpeFwdEntriesGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GpeFwdEntryPathDump(ctx context.Context, in *GpeFwdEntryPathDump) (RPCService_GpeFwdEntryPathDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_GpeFwdEntryPathDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_GpeFwdEntryPathDumpClient interface {
	Recv() (*GpeFwdEntryPathDetails, error)
	api.Stream
}

type serviceClient_GpeFwdEntryPathDumpClient struct {
	api.Stream
}

func (c *serviceClient_GpeFwdEntryPathDumpClient) Recv() (*GpeFwdEntryPathDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *GpeFwdEntryPathDetails:
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

func (c *serviceClient) GpeFwdEntryVnisGet(ctx context.Context, in *GpeFwdEntryVnisGet) (*GpeFwdEntryVnisGetReply, error) {
	out := new(GpeFwdEntryVnisGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GpeGetEncapMode(ctx context.Context, in *GpeGetEncapMode) (*GpeGetEncapModeReply, error) {
	out := new(GpeGetEncapModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GpeNativeFwdRpathsGet(ctx context.Context, in *GpeNativeFwdRpathsGet) (*GpeNativeFwdRpathsGetReply, error) {
	out := new(GpeNativeFwdRpathsGetReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) GpeSetEncapMode(ctx context.Context, in *GpeSetEncapMode) (*GpeSetEncapModeReply, error) {
	out := new(GpeSetEncapModeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
