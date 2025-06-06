// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.12.0
//  VPP:              unknown
// source: p2p_ethernet.api.json

// Package p2p_ethernet contains generated bindings for API file p2p_ethernet.api.
//
// Contents:
// -  4 messages
package p2p_ethernet

import (
	"vpp-restapi/binapi/ethernet_types"
	"vpp-restapi/binapi/interface_types"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "p2p_ethernet"
	APIVersion = "1.0.0"
	VersionCrc = 0x3cfe55da
)

// Create a point-to-point (p2p) Ethernet sub-interface
//   - parent_if_index - index of the parent interface
//   - subif_id - subinterface index identifier
//   - remote_mac - client MAC address
//     @retval VNET_API_ERROR_INVALID_SW_IF_INDEX on invalid parent_if_index
//     @retval VNET_API_ERROR_INVALID_SW_IF_INDEX_2 on invalid subif_id
//     @retval VNET_API_ERROR_BOND_SLAVE_NOT_ALLOWED
//     @retval VNET_API_ERROR_SUBIF_ALREADY_EXISTS
//     @retval VNET_API_ERROR_SUBIF_CREATE_FAILED
//
// P2pEthernetAdd defines message 'p2p_ethernet_add'.
type P2pEthernetAdd struct {
	ParentIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=parent_if_index" json:"parent_if_index,omitempty"`
	SubifID       uint32                         `binapi:"u32,name=subif_id" json:"subif_id,omitempty"`
	RemoteMac     ethernet_types.MacAddress      `binapi:"mac_address,name=remote_mac" json:"remote_mac,omitempty"`
}

func (m *P2pEthernetAdd) Reset()               { *m = P2pEthernetAdd{} }
func (*P2pEthernetAdd) GetMessageName() string { return "p2p_ethernet_add" }
func (*P2pEthernetAdd) GetCrcString() string   { return "36a1a6dc" }
func (*P2pEthernetAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *P2pEthernetAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.ParentIfIndex
	size += 4     // m.SubifID
	size += 1 * 6 // m.RemoteMac
	return size
}
func (m *P2pEthernetAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.ParentIfIndex))
	buf.EncodeUint32(m.SubifID)
	buf.EncodeBytes(m.RemoteMac[:], 6)
	return buf.Bytes(), nil
}
func (m *P2pEthernetAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ParentIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.SubifID = buf.DecodeUint32()
	copy(m.RemoteMac[:], buf.DecodeBytes(6))
	return nil
}

// P2pEthernetAddReply defines message 'p2p_ethernet_add_reply'.
type P2pEthernetAddReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *P2pEthernetAddReply) Reset()               { *m = P2pEthernetAddReply{} }
func (*P2pEthernetAddReply) GetMessageName() string { return "p2p_ethernet_add_reply" }
func (*P2pEthernetAddReply) GetCrcString() string   { return "5383d31f" }
func (*P2pEthernetAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *P2pEthernetAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *P2pEthernetAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *P2pEthernetAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// Delete a point-to-point (p2p) Ethernet sub-interface
//   - parent_if_index - index of the parent interface
//   - remote_mac - client MAC address
//     @retval VNET_API_ERROR_SUBIF_DOESNT_EXIST
//
// P2pEthernetDel defines message 'p2p_ethernet_del'.
type P2pEthernetDel struct {
	ParentIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=parent_if_index" json:"parent_if_index,omitempty"`
	RemoteMac     ethernet_types.MacAddress      `binapi:"mac_address,name=remote_mac" json:"remote_mac,omitempty"`
}

func (m *P2pEthernetDel) Reset()               { *m = P2pEthernetDel{} }
func (*P2pEthernetDel) GetMessageName() string { return "p2p_ethernet_del" }
func (*P2pEthernetDel) GetCrcString() string   { return "62f81c8c" }
func (*P2pEthernetDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *P2pEthernetDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.ParentIfIndex
	size += 1 * 6 // m.RemoteMac
	return size
}
func (m *P2pEthernetDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.ParentIfIndex))
	buf.EncodeBytes(m.RemoteMac[:], 6)
	return buf.Bytes(), nil
}
func (m *P2pEthernetDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ParentIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	copy(m.RemoteMac[:], buf.DecodeBytes(6))
	return nil
}

// P2pEthernetDelReply defines message 'p2p_ethernet_del_reply'.
type P2pEthernetDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *P2pEthernetDelReply) Reset()               { *m = P2pEthernetDelReply{} }
func (*P2pEthernetDelReply) GetMessageName() string { return "p2p_ethernet_del_reply" }
func (*P2pEthernetDelReply) GetCrcString() string   { return "e8d4e804" }
func (*P2pEthernetDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *P2pEthernetDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *P2pEthernetDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *P2pEthernetDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_p2p_ethernet_binapi_init() }
func file_p2p_ethernet_binapi_init() {
	api.RegisterMessage((*P2pEthernetAdd)(nil), "p2p_ethernet_add_36a1a6dc")
	api.RegisterMessage((*P2pEthernetAddReply)(nil), "p2p_ethernet_add_reply_5383d31f")
	api.RegisterMessage((*P2pEthernetDel)(nil), "p2p_ethernet_del_62f81c8c")
	api.RegisterMessage((*P2pEthernetDelReply)(nil), "p2p_ethernet_del_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*P2pEthernetAdd)(nil),
		(*P2pEthernetAddReply)(nil),
		(*P2pEthernetDel)(nil),
		(*P2pEthernetDelReply)(nil),
	}
}
