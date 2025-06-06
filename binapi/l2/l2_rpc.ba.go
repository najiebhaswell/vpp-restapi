// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package l2

import (
	"context"
	"fmt"
	"io"
	"vpp-restapi/binapi/memclnt"

	api "go.fd.io/govpp/api"
)

// RPCService defines RPC service l2.
type RPCService interface {
	BdIPMacAddDel(ctx context.Context, in *BdIPMacAddDel) (*BdIPMacAddDelReply, error)
	BdIPMacDump(ctx context.Context, in *BdIPMacDump) (RPCService_BdIPMacDumpClient, error)
	BdIPMacFlush(ctx context.Context, in *BdIPMacFlush) (*BdIPMacFlushReply, error)
	BridgeDomainAddDel(ctx context.Context, in *BridgeDomainAddDel) (*BridgeDomainAddDelReply, error)
	BridgeDomainAddDelV2(ctx context.Context, in *BridgeDomainAddDelV2) (*BridgeDomainAddDelV2Reply, error)
	BridgeDomainDump(ctx context.Context, in *BridgeDomainDump) (RPCService_BridgeDomainDumpClient, error)
	BridgeDomainSetDefaultLearnLimit(ctx context.Context, in *BridgeDomainSetDefaultLearnLimit) (*BridgeDomainSetDefaultLearnLimitReply, error)
	BridgeDomainSetLearnLimit(ctx context.Context, in *BridgeDomainSetLearnLimit) (*BridgeDomainSetLearnLimitReply, error)
	BridgeDomainSetMacAge(ctx context.Context, in *BridgeDomainSetMacAge) (*BridgeDomainSetMacAgeReply, error)
	BridgeFlags(ctx context.Context, in *BridgeFlags) (*BridgeFlagsReply, error)
	BviCreate(ctx context.Context, in *BviCreate) (*BviCreateReply, error)
	BviDelete(ctx context.Context, in *BviDelete) (*BviDeleteReply, error)
	L2FibClearTable(ctx context.Context, in *L2FibClearTable) (*L2FibClearTableReply, error)
	L2FibTableDump(ctx context.Context, in *L2FibTableDump) (RPCService_L2FibTableDumpClient, error)
	L2Flags(ctx context.Context, in *L2Flags) (*L2FlagsReply, error)
	L2InterfaceEfpFilter(ctx context.Context, in *L2InterfaceEfpFilter) (*L2InterfaceEfpFilterReply, error)
	L2InterfacePbbTagRewrite(ctx context.Context, in *L2InterfacePbbTagRewrite) (*L2InterfacePbbTagRewriteReply, error)
	L2InterfaceVlanTagRewrite(ctx context.Context, in *L2InterfaceVlanTagRewrite) (*L2InterfaceVlanTagRewriteReply, error)
	L2PatchAddDel(ctx context.Context, in *L2PatchAddDel) (*L2PatchAddDelReply, error)
	L2XconnectDump(ctx context.Context, in *L2XconnectDump) (RPCService_L2XconnectDumpClient, error)
	L2fibAddDel(ctx context.Context, in *L2fibAddDel) (*L2fibAddDelReply, error)
	L2fibFlushAll(ctx context.Context, in *L2fibFlushAll) (*L2fibFlushAllReply, error)
	L2fibFlushBd(ctx context.Context, in *L2fibFlushBd) (*L2fibFlushBdReply, error)
	L2fibFlushInt(ctx context.Context, in *L2fibFlushInt) (*L2fibFlushIntReply, error)
	L2fibSetScanDelay(ctx context.Context, in *L2fibSetScanDelay) (*L2fibSetScanDelayReply, error)
	SwInterfaceSetL2Bridge(ctx context.Context, in *SwInterfaceSetL2Bridge) (*SwInterfaceSetL2BridgeReply, error)
	SwInterfaceSetL2Xconnect(ctx context.Context, in *SwInterfaceSetL2Xconnect) (*SwInterfaceSetL2XconnectReply, error)
	SwInterfaceSetVpath(ctx context.Context, in *SwInterfaceSetVpath) (*SwInterfaceSetVpathReply, error)
	WantL2ArpTermEvents(ctx context.Context, in *WantL2ArpTermEvents) (*WantL2ArpTermEventsReply, error)
	WantL2MacsEvents(ctx context.Context, in *WantL2MacsEvents) (*WantL2MacsEventsReply, error)
	WantL2MacsEvents2(ctx context.Context, in *WantL2MacsEvents2) (*WantL2MacsEvents2Reply, error)
}

type serviceClient struct {
	conn api.Connection
}

func NewServiceClient(conn api.Connection) RPCService {
	return &serviceClient{conn}
}

func (c *serviceClient) BdIPMacAddDel(ctx context.Context, in *BdIPMacAddDel) (*BdIPMacAddDelReply, error) {
	out := new(BdIPMacAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BdIPMacDump(ctx context.Context, in *BdIPMacDump) (RPCService_BdIPMacDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_BdIPMacDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_BdIPMacDumpClient interface {
	Recv() (*BdIPMacDetails, error)
	api.Stream
}

type serviceClient_BdIPMacDumpClient struct {
	api.Stream
}

func (c *serviceClient_BdIPMacDumpClient) Recv() (*BdIPMacDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *BdIPMacDetails:
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

func (c *serviceClient) BdIPMacFlush(ctx context.Context, in *BdIPMacFlush) (*BdIPMacFlushReply, error) {
	out := new(BdIPMacFlushReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BridgeDomainAddDel(ctx context.Context, in *BridgeDomainAddDel) (*BridgeDomainAddDelReply, error) {
	out := new(BridgeDomainAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BridgeDomainAddDelV2(ctx context.Context, in *BridgeDomainAddDelV2) (*BridgeDomainAddDelV2Reply, error) {
	out := new(BridgeDomainAddDelV2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BridgeDomainDump(ctx context.Context, in *BridgeDomainDump) (RPCService_BridgeDomainDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_BridgeDomainDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_BridgeDomainDumpClient interface {
	Recv() (*BridgeDomainDetails, error)
	api.Stream
}

type serviceClient_BridgeDomainDumpClient struct {
	api.Stream
}

func (c *serviceClient_BridgeDomainDumpClient) Recv() (*BridgeDomainDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *BridgeDomainDetails:
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

func (c *serviceClient) BridgeDomainSetDefaultLearnLimit(ctx context.Context, in *BridgeDomainSetDefaultLearnLimit) (*BridgeDomainSetDefaultLearnLimitReply, error) {
	out := new(BridgeDomainSetDefaultLearnLimitReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BridgeDomainSetLearnLimit(ctx context.Context, in *BridgeDomainSetLearnLimit) (*BridgeDomainSetLearnLimitReply, error) {
	out := new(BridgeDomainSetLearnLimitReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BridgeDomainSetMacAge(ctx context.Context, in *BridgeDomainSetMacAge) (*BridgeDomainSetMacAgeReply, error) {
	out := new(BridgeDomainSetMacAgeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BridgeFlags(ctx context.Context, in *BridgeFlags) (*BridgeFlagsReply, error) {
	out := new(BridgeFlagsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BviCreate(ctx context.Context, in *BviCreate) (*BviCreateReply, error) {
	out := new(BviCreateReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) BviDelete(ctx context.Context, in *BviDelete) (*BviDeleteReply, error) {
	out := new(BviDeleteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2FibClearTable(ctx context.Context, in *L2FibClearTable) (*L2FibClearTableReply, error) {
	out := new(L2FibClearTableReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2FibTableDump(ctx context.Context, in *L2FibTableDump) (RPCService_L2FibTableDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_L2FibTableDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_L2FibTableDumpClient interface {
	Recv() (*L2FibTableDetails, error)
	api.Stream
}

type serviceClient_L2FibTableDumpClient struct {
	api.Stream
}

func (c *serviceClient_L2FibTableDumpClient) Recv() (*L2FibTableDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *L2FibTableDetails:
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

func (c *serviceClient) L2Flags(ctx context.Context, in *L2Flags) (*L2FlagsReply, error) {
	out := new(L2FlagsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2InterfaceEfpFilter(ctx context.Context, in *L2InterfaceEfpFilter) (*L2InterfaceEfpFilterReply, error) {
	out := new(L2InterfaceEfpFilterReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2InterfacePbbTagRewrite(ctx context.Context, in *L2InterfacePbbTagRewrite) (*L2InterfacePbbTagRewriteReply, error) {
	out := new(L2InterfacePbbTagRewriteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2InterfaceVlanTagRewrite(ctx context.Context, in *L2InterfaceVlanTagRewrite) (*L2InterfaceVlanTagRewriteReply, error) {
	out := new(L2InterfaceVlanTagRewriteReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2PatchAddDel(ctx context.Context, in *L2PatchAddDel) (*L2PatchAddDelReply, error) {
	out := new(L2PatchAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2XconnectDump(ctx context.Context, in *L2XconnectDump) (RPCService_L2XconnectDumpClient, error) {
	stream, err := c.conn.NewStream(ctx)
	if err != nil {
		return nil, err
	}
	x := &serviceClient_L2XconnectDumpClient{stream}
	if err := x.Stream.SendMsg(in); err != nil {
		return nil, err
	}
	if err = x.Stream.SendMsg(&memclnt.ControlPing{}); err != nil {
		return nil, err
	}
	return x, nil
}

type RPCService_L2XconnectDumpClient interface {
	Recv() (*L2XconnectDetails, error)
	api.Stream
}

type serviceClient_L2XconnectDumpClient struct {
	api.Stream
}

func (c *serviceClient_L2XconnectDumpClient) Recv() (*L2XconnectDetails, error) {
	msg, err := c.Stream.RecvMsg()
	if err != nil {
		return nil, err
	}
	switch m := msg.(type) {
	case *L2XconnectDetails:
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

func (c *serviceClient) L2fibAddDel(ctx context.Context, in *L2fibAddDel) (*L2fibAddDelReply, error) {
	out := new(L2fibAddDelReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2fibFlushAll(ctx context.Context, in *L2fibFlushAll) (*L2fibFlushAllReply, error) {
	out := new(L2fibFlushAllReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2fibFlushBd(ctx context.Context, in *L2fibFlushBd) (*L2fibFlushBdReply, error) {
	out := new(L2fibFlushBdReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2fibFlushInt(ctx context.Context, in *L2fibFlushInt) (*L2fibFlushIntReply, error) {
	out := new(L2fibFlushIntReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) L2fibSetScanDelay(ctx context.Context, in *L2fibSetScanDelay) (*L2fibSetScanDelayReply, error) {
	out := new(L2fibSetScanDelayReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetL2Bridge(ctx context.Context, in *SwInterfaceSetL2Bridge) (*SwInterfaceSetL2BridgeReply, error) {
	out := new(SwInterfaceSetL2BridgeReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetL2Xconnect(ctx context.Context, in *SwInterfaceSetL2Xconnect) (*SwInterfaceSetL2XconnectReply, error) {
	out := new(SwInterfaceSetL2XconnectReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) SwInterfaceSetVpath(ctx context.Context, in *SwInterfaceSetVpath) (*SwInterfaceSetVpathReply, error) {
	out := new(SwInterfaceSetVpathReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WantL2ArpTermEvents(ctx context.Context, in *WantL2ArpTermEvents) (*WantL2ArpTermEventsReply, error) {
	out := new(WantL2ArpTermEventsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WantL2MacsEvents(ctx context.Context, in *WantL2MacsEvents) (*WantL2MacsEventsReply, error) {
	out := new(WantL2MacsEventsReply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}

func (c *serviceClient) WantL2MacsEvents2(ctx context.Context, in *WantL2MacsEvents2) (*WantL2MacsEvents2Reply, error) {
	out := new(WantL2MacsEvents2Reply)
	err := c.conn.Invoke(ctx, in, out)
	if err != nil {
		return nil, err
	}
	return out, api.RetvalToVPPApiError(out.Retval)
}
