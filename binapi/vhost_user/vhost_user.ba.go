// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.12.0
//  VPP:              24.10.0-1~gb3a21b9b0
// source: plugins/vhost_user.api.json

// Package vhost_user contains generated bindings for API file vhost_user.api.
//
// Contents:
// - 12 messages
package vhost_user

import (
	"vpp-restapi/binapi/ethernet_types"
	"vpp-restapi/binapi/interface_types"
	"vpp-restapi/binapi/virtio_types"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "vhost_user"
	APIVersion = "4.1.1"
	VersionCrc = 0xd49ae8cd
)

// vhost-user interface create request
//   - is_server - our side is socket server
//   - sock_filename - unix socket filename, used to speak with frontend
//   - use_custom_mac - enable or disable the use of the provided hardware address
//   - disable_mrg_rxbuf - disable the use of merge receive buffers
//   - disable_indirect_desc - disable the use of indirect descriptors which driver can use
//   - enable_gso - enable gso support (default 0)
//   - enable_packed - enable packed ring support (default 0)
//   - mac_address - hardware address to use if 'use_custom_mac' is set
//
// CreateVhostUserIf defines message 'create_vhost_user_if'.
// Deprecated: the message will be removed in the future versions
type CreateVhostUserIf struct {
	IsServer            bool                      `binapi:"bool,name=is_server" json:"is_server,omitempty"`
	SockFilename        string                    `binapi:"string[256],name=sock_filename" json:"sock_filename,omitempty"`
	Renumber            bool                      `binapi:"bool,name=renumber" json:"renumber,omitempty"`
	DisableMrgRxbuf     bool                      `binapi:"bool,name=disable_mrg_rxbuf" json:"disable_mrg_rxbuf,omitempty"`
	DisableIndirectDesc bool                      `binapi:"bool,name=disable_indirect_desc" json:"disable_indirect_desc,omitempty"`
	EnableGso           bool                      `binapi:"bool,name=enable_gso" json:"enable_gso,omitempty"`
	EnablePacked        bool                      `binapi:"bool,name=enable_packed" json:"enable_packed,omitempty"`
	CustomDevInstance   uint32                    `binapi:"u32,name=custom_dev_instance" json:"custom_dev_instance,omitempty"`
	UseCustomMac        bool                      `binapi:"bool,name=use_custom_mac" json:"use_custom_mac,omitempty"`
	MacAddress          ethernet_types.MacAddress `binapi:"mac_address,name=mac_address" json:"mac_address,omitempty"`
	Tag                 string                    `binapi:"string[64],name=tag" json:"tag,omitempty"`
}

func (m *CreateVhostUserIf) Reset()               { *m = CreateVhostUserIf{} }
func (*CreateVhostUserIf) GetMessageName() string { return "create_vhost_user_if" }
func (*CreateVhostUserIf) GetCrcString() string   { return "c785c6fc" }
func (*CreateVhostUserIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CreateVhostUserIf) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1     // m.IsServer
	size += 256   // m.SockFilename
	size += 1     // m.Renumber
	size += 1     // m.DisableMrgRxbuf
	size += 1     // m.DisableIndirectDesc
	size += 1     // m.EnableGso
	size += 1     // m.EnablePacked
	size += 4     // m.CustomDevInstance
	size += 1     // m.UseCustomMac
	size += 1 * 6 // m.MacAddress
	size += 64    // m.Tag
	return size
}
func (m *CreateVhostUserIf) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsServer)
	buf.EncodeString(m.SockFilename, 256)
	buf.EncodeBool(m.Renumber)
	buf.EncodeBool(m.DisableMrgRxbuf)
	buf.EncodeBool(m.DisableIndirectDesc)
	buf.EncodeBool(m.EnableGso)
	buf.EncodeBool(m.EnablePacked)
	buf.EncodeUint32(m.CustomDevInstance)
	buf.EncodeBool(m.UseCustomMac)
	buf.EncodeBytes(m.MacAddress[:], 6)
	buf.EncodeString(m.Tag, 64)
	return buf.Bytes(), nil
}
func (m *CreateVhostUserIf) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsServer = buf.DecodeBool()
	m.SockFilename = buf.DecodeString(256)
	m.Renumber = buf.DecodeBool()
	m.DisableMrgRxbuf = buf.DecodeBool()
	m.DisableIndirectDesc = buf.DecodeBool()
	m.EnableGso = buf.DecodeBool()
	m.EnablePacked = buf.DecodeBool()
	m.CustomDevInstance = buf.DecodeUint32()
	m.UseCustomMac = buf.DecodeBool()
	copy(m.MacAddress[:], buf.DecodeBytes(6))
	m.Tag = buf.DecodeString(64)
	return nil
}

// vhost-user interface create response
//   - retval - return code for the request
//   - sw_if_index - interface the operation is applied to
//
// CreateVhostUserIfReply defines message 'create_vhost_user_if_reply'.
// Deprecated: the message will be removed in the future versions
type CreateVhostUserIfReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *CreateVhostUserIfReply) Reset()               { *m = CreateVhostUserIfReply{} }
func (*CreateVhostUserIfReply) GetMessageName() string { return "create_vhost_user_if_reply" }
func (*CreateVhostUserIfReply) GetCrcString() string   { return "5383d31f" }
func (*CreateVhostUserIfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CreateVhostUserIfReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *CreateVhostUserIfReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *CreateVhostUserIfReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// vhost-user interface create request
//   - is_server - our side is socket server
//   - sock_filename - unix socket filename, used to speak with frontend
//   - use_custom_mac - enable or disable the use of the provided hardware address
//   - disable_mrg_rxbuf - disable the use of merge receive buffers
//   - disable_indirect_desc - disable the use of indirect descriptors which driver can use
//   - enable_gso - enable gso support (default 0)
//   - enable_packed - enable packed ring support (default 0)
//   - enable_event_idx - enable event_idx support (default 0)
//   - mac_address - hardware address to use if 'use_custom_mac' is set
//   - renumber - if true, use custom_dev_instance is valid
//   - custom_dev_instance - custom device instance number
//
// CreateVhostUserIfV2 defines message 'create_vhost_user_if_v2'.
type CreateVhostUserIfV2 struct {
	IsServer            bool                      `binapi:"bool,name=is_server" json:"is_server,omitempty"`
	SockFilename        string                    `binapi:"string[256],name=sock_filename" json:"sock_filename,omitempty"`
	Renumber            bool                      `binapi:"bool,name=renumber" json:"renumber,omitempty"`
	DisableMrgRxbuf     bool                      `binapi:"bool,name=disable_mrg_rxbuf" json:"disable_mrg_rxbuf,omitempty"`
	DisableIndirectDesc bool                      `binapi:"bool,name=disable_indirect_desc" json:"disable_indirect_desc,omitempty"`
	EnableGso           bool                      `binapi:"bool,name=enable_gso" json:"enable_gso,omitempty"`
	EnablePacked        bool                      `binapi:"bool,name=enable_packed" json:"enable_packed,omitempty"`
	EnableEventIdx      bool                      `binapi:"bool,name=enable_event_idx" json:"enable_event_idx,omitempty"`
	CustomDevInstance   uint32                    `binapi:"u32,name=custom_dev_instance" json:"custom_dev_instance,omitempty"`
	UseCustomMac        bool                      `binapi:"bool,name=use_custom_mac" json:"use_custom_mac,omitempty"`
	MacAddress          ethernet_types.MacAddress `binapi:"mac_address,name=mac_address" json:"mac_address,omitempty"`
	Tag                 string                    `binapi:"string[64],name=tag" json:"tag,omitempty"`
}

func (m *CreateVhostUserIfV2) Reset()               { *m = CreateVhostUserIfV2{} }
func (*CreateVhostUserIfV2) GetMessageName() string { return "create_vhost_user_if_v2" }
func (*CreateVhostUserIfV2) GetCrcString() string   { return "dba1cc1d" }
func (*CreateVhostUserIfV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CreateVhostUserIfV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1     // m.IsServer
	size += 256   // m.SockFilename
	size += 1     // m.Renumber
	size += 1     // m.DisableMrgRxbuf
	size += 1     // m.DisableIndirectDesc
	size += 1     // m.EnableGso
	size += 1     // m.EnablePacked
	size += 1     // m.EnableEventIdx
	size += 4     // m.CustomDevInstance
	size += 1     // m.UseCustomMac
	size += 1 * 6 // m.MacAddress
	size += 64    // m.Tag
	return size
}
func (m *CreateVhostUserIfV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsServer)
	buf.EncodeString(m.SockFilename, 256)
	buf.EncodeBool(m.Renumber)
	buf.EncodeBool(m.DisableMrgRxbuf)
	buf.EncodeBool(m.DisableIndirectDesc)
	buf.EncodeBool(m.EnableGso)
	buf.EncodeBool(m.EnablePacked)
	buf.EncodeBool(m.EnableEventIdx)
	buf.EncodeUint32(m.CustomDevInstance)
	buf.EncodeBool(m.UseCustomMac)
	buf.EncodeBytes(m.MacAddress[:], 6)
	buf.EncodeString(m.Tag, 64)
	return buf.Bytes(), nil
}
func (m *CreateVhostUserIfV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsServer = buf.DecodeBool()
	m.SockFilename = buf.DecodeString(256)
	m.Renumber = buf.DecodeBool()
	m.DisableMrgRxbuf = buf.DecodeBool()
	m.DisableIndirectDesc = buf.DecodeBool()
	m.EnableGso = buf.DecodeBool()
	m.EnablePacked = buf.DecodeBool()
	m.EnableEventIdx = buf.DecodeBool()
	m.CustomDevInstance = buf.DecodeUint32()
	m.UseCustomMac = buf.DecodeBool()
	copy(m.MacAddress[:], buf.DecodeBytes(6))
	m.Tag = buf.DecodeString(64)
	return nil
}

// vhost-user interface create response
//   - retval - return code for the request
//   - sw_if_index - interface the operation is applied to
//
// CreateVhostUserIfV2Reply defines message 'create_vhost_user_if_v2_reply'.
type CreateVhostUserIfV2Reply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *CreateVhostUserIfV2Reply) Reset()               { *m = CreateVhostUserIfV2Reply{} }
func (*CreateVhostUserIfV2Reply) GetMessageName() string { return "create_vhost_user_if_v2_reply" }
func (*CreateVhostUserIfV2Reply) GetCrcString() string   { return "5383d31f" }
func (*CreateVhostUserIfV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CreateVhostUserIfV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *CreateVhostUserIfV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *CreateVhostUserIfV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// vhost-user interface delete request
// DeleteVhostUserIf defines message 'delete_vhost_user_if'.
type DeleteVhostUserIf struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *DeleteVhostUserIf) Reset()               { *m = DeleteVhostUserIf{} }
func (*DeleteVhostUserIf) GetMessageName() string { return "delete_vhost_user_if" }
func (*DeleteVhostUserIf) GetCrcString() string   { return "f9e6675e" }
func (*DeleteVhostUserIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DeleteVhostUserIf) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *DeleteVhostUserIf) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *DeleteVhostUserIf) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// DeleteVhostUserIfReply defines message 'delete_vhost_user_if_reply'.
type DeleteVhostUserIfReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *DeleteVhostUserIfReply) Reset()               { *m = DeleteVhostUserIfReply{} }
func (*DeleteVhostUserIfReply) GetMessageName() string { return "delete_vhost_user_if_reply" }
func (*DeleteVhostUserIfReply) GetCrcString() string   { return "e8d4e804" }
func (*DeleteVhostUserIfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DeleteVhostUserIfReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *DeleteVhostUserIfReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *DeleteVhostUserIfReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// vhost-user interface modify request
//   - is_server - our side is socket server
//   - sock_filename - unix socket filename, used to speak with frontend
//   - enable_gso - enable gso support (default 0)
//   - enable_packed - enable packed ring support (default 0)
//
// ModifyVhostUserIf defines message 'modify_vhost_user_if'.
// Deprecated: the message will be removed in the future versions
type ModifyVhostUserIf struct {
	SwIfIndex         interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsServer          bool                           `binapi:"bool,name=is_server" json:"is_server,omitempty"`
	SockFilename      string                         `binapi:"string[256],name=sock_filename" json:"sock_filename,omitempty"`
	Renumber          bool                           `binapi:"bool,name=renumber" json:"renumber,omitempty"`
	EnableGso         bool                           `binapi:"bool,name=enable_gso" json:"enable_gso,omitempty"`
	EnablePacked      bool                           `binapi:"bool,name=enable_packed" json:"enable_packed,omitempty"`
	CustomDevInstance uint32                         `binapi:"u32,name=custom_dev_instance" json:"custom_dev_instance,omitempty"`
}

func (m *ModifyVhostUserIf) Reset()               { *m = ModifyVhostUserIf{} }
func (*ModifyVhostUserIf) GetMessageName() string { return "modify_vhost_user_if" }
func (*ModifyVhostUserIf) GetCrcString() string   { return "0e71d40b" }
func (*ModifyVhostUserIf) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ModifyVhostUserIf) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4   // m.SwIfIndex
	size += 1   // m.IsServer
	size += 256 // m.SockFilename
	size += 1   // m.Renumber
	size += 1   // m.EnableGso
	size += 1   // m.EnablePacked
	size += 4   // m.CustomDevInstance
	return size
}
func (m *ModifyVhostUserIf) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsServer)
	buf.EncodeString(m.SockFilename, 256)
	buf.EncodeBool(m.Renumber)
	buf.EncodeBool(m.EnableGso)
	buf.EncodeBool(m.EnablePacked)
	buf.EncodeUint32(m.CustomDevInstance)
	return buf.Bytes(), nil
}
func (m *ModifyVhostUserIf) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsServer = buf.DecodeBool()
	m.SockFilename = buf.DecodeString(256)
	m.Renumber = buf.DecodeBool()
	m.EnableGso = buf.DecodeBool()
	m.EnablePacked = buf.DecodeBool()
	m.CustomDevInstance = buf.DecodeUint32()
	return nil
}

// ModifyVhostUserIfReply defines message 'modify_vhost_user_if_reply'.
// Deprecated: the message will be removed in the future versions
type ModifyVhostUserIfReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *ModifyVhostUserIfReply) Reset()               { *m = ModifyVhostUserIfReply{} }
func (*ModifyVhostUserIfReply) GetMessageName() string { return "modify_vhost_user_if_reply" }
func (*ModifyVhostUserIfReply) GetCrcString() string   { return "e8d4e804" }
func (*ModifyVhostUserIfReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ModifyVhostUserIfReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *ModifyVhostUserIfReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *ModifyVhostUserIfReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// vhost-user interface modify request
//   - is_server - our side is socket server
//   - sock_filename - unix socket filename, used to speak with frontend
//   - enable_gso - enable gso support (default 0)
//   - enable_packed - enable packed ring support (default 0)
//   - enable_event_idx - enable event idx support (default 0)
//   - renumber - if true, use custom_dev_instance is valid
//   - custom_dev_instance - custom device instance number
//
// ModifyVhostUserIfV2 defines message 'modify_vhost_user_if_v2'.
type ModifyVhostUserIfV2 struct {
	SwIfIndex         interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsServer          bool                           `binapi:"bool,name=is_server" json:"is_server,omitempty"`
	SockFilename      string                         `binapi:"string[256],name=sock_filename" json:"sock_filename,omitempty"`
	Renumber          bool                           `binapi:"bool,name=renumber" json:"renumber,omitempty"`
	EnableGso         bool                           `binapi:"bool,name=enable_gso" json:"enable_gso,omitempty"`
	EnablePacked      bool                           `binapi:"bool,name=enable_packed" json:"enable_packed,omitempty"`
	EnableEventIdx    bool                           `binapi:"bool,name=enable_event_idx" json:"enable_event_idx,omitempty"`
	CustomDevInstance uint32                         `binapi:"u32,name=custom_dev_instance" json:"custom_dev_instance,omitempty"`
}

func (m *ModifyVhostUserIfV2) Reset()               { *m = ModifyVhostUserIfV2{} }
func (*ModifyVhostUserIfV2) GetMessageName() string { return "modify_vhost_user_if_v2" }
func (*ModifyVhostUserIfV2) GetCrcString() string   { return "b2483771" }
func (*ModifyVhostUserIfV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ModifyVhostUserIfV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4   // m.SwIfIndex
	size += 1   // m.IsServer
	size += 256 // m.SockFilename
	size += 1   // m.Renumber
	size += 1   // m.EnableGso
	size += 1   // m.EnablePacked
	size += 1   // m.EnableEventIdx
	size += 4   // m.CustomDevInstance
	return size
}
func (m *ModifyVhostUserIfV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsServer)
	buf.EncodeString(m.SockFilename, 256)
	buf.EncodeBool(m.Renumber)
	buf.EncodeBool(m.EnableGso)
	buf.EncodeBool(m.EnablePacked)
	buf.EncodeBool(m.EnableEventIdx)
	buf.EncodeUint32(m.CustomDevInstance)
	return buf.Bytes(), nil
}
func (m *ModifyVhostUserIfV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsServer = buf.DecodeBool()
	m.SockFilename = buf.DecodeString(256)
	m.Renumber = buf.DecodeBool()
	m.EnableGso = buf.DecodeBool()
	m.EnablePacked = buf.DecodeBool()
	m.EnableEventIdx = buf.DecodeBool()
	m.CustomDevInstance = buf.DecodeUint32()
	return nil
}

// ModifyVhostUserIfV2Reply defines message 'modify_vhost_user_if_v2_reply'.
type ModifyVhostUserIfV2Reply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *ModifyVhostUserIfV2Reply) Reset()               { *m = ModifyVhostUserIfV2Reply{} }
func (*ModifyVhostUserIfV2Reply) GetMessageName() string { return "modify_vhost_user_if_v2_reply" }
func (*ModifyVhostUserIfV2Reply) GetCrcString() string   { return "e8d4e804" }
func (*ModifyVhostUserIfV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ModifyVhostUserIfV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *ModifyVhostUserIfV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *ModifyVhostUserIfV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Vhost-user interface details structure (fix this)
//   - sw_if_index - index of the interface
//   - interface_name - name of interface
//   - virtio_net_hdr_sz - net header size
//   - features_first_32 - interface features, first 32 bits
//   - features_last_32 - interface features, last 32 bits
//   - is_server - vhost-user server socket
//   - sock_filename - socket filename
//   - num_regions - number of used memory regions
//   - sock_errno - socket errno
//
// SwInterfaceVhostUserDetails defines message 'sw_interface_vhost_user_details'.
type SwInterfaceVhostUserDetails struct {
	SwIfIndex       interface_types.InterfaceIndex        `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	InterfaceName   string                                `binapi:"string[64],name=interface_name" json:"interface_name,omitempty"`
	VirtioNetHdrSz  uint32                                `binapi:"u32,name=virtio_net_hdr_sz" json:"virtio_net_hdr_sz,omitempty"`
	FeaturesFirst32 virtio_types.VirtioNetFeaturesFirst32 `binapi:"virtio_net_features_first_32,name=features_first_32" json:"features_first_32,omitempty"`
	FeaturesLast32  virtio_types.VirtioNetFeaturesLast32  `binapi:"virtio_net_features_last_32,name=features_last_32" json:"features_last_32,omitempty"`
	IsServer        bool                                  `binapi:"bool,name=is_server" json:"is_server,omitempty"`
	SockFilename    string                                `binapi:"string[256],name=sock_filename" json:"sock_filename,omitempty"`
	NumRegions      uint32                                `binapi:"u32,name=num_regions" json:"num_regions,omitempty"`
	SockErrno       int32                                 `binapi:"i32,name=sock_errno" json:"sock_errno,omitempty"`
}

func (m *SwInterfaceVhostUserDetails) Reset()               { *m = SwInterfaceVhostUserDetails{} }
func (*SwInterfaceVhostUserDetails) GetMessageName() string { return "sw_interface_vhost_user_details" }
func (*SwInterfaceVhostUserDetails) GetCrcString() string   { return "0cee1e53" }
func (*SwInterfaceVhostUserDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SwInterfaceVhostUserDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4   // m.SwIfIndex
	size += 64  // m.InterfaceName
	size += 4   // m.VirtioNetHdrSz
	size += 4   // m.FeaturesFirst32
	size += 4   // m.FeaturesLast32
	size += 1   // m.IsServer
	size += 256 // m.SockFilename
	size += 4   // m.NumRegions
	size += 4   // m.SockErrno
	return size
}
func (m *SwInterfaceVhostUserDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.InterfaceName, 64)
	buf.EncodeUint32(m.VirtioNetHdrSz)
	buf.EncodeUint32(uint32(m.FeaturesFirst32))
	buf.EncodeUint32(uint32(m.FeaturesLast32))
	buf.EncodeBool(m.IsServer)
	buf.EncodeString(m.SockFilename, 256)
	buf.EncodeUint32(m.NumRegions)
	buf.EncodeInt32(m.SockErrno)
	return buf.Bytes(), nil
}
func (m *SwInterfaceVhostUserDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.InterfaceName = buf.DecodeString(64)
	m.VirtioNetHdrSz = buf.DecodeUint32()
	m.FeaturesFirst32 = virtio_types.VirtioNetFeaturesFirst32(buf.DecodeUint32())
	m.FeaturesLast32 = virtio_types.VirtioNetFeaturesLast32(buf.DecodeUint32())
	m.IsServer = buf.DecodeBool()
	m.SockFilename = buf.DecodeString(256)
	m.NumRegions = buf.DecodeUint32()
	m.SockErrno = buf.DecodeInt32()
	return nil
}

// Vhost-user interface dump request
//   - sw_if_index - filter by sw_if_index
//
// SwInterfaceVhostUserDump defines message 'sw_interface_vhost_user_dump'.
type SwInterfaceVhostUserDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
}

func (m *SwInterfaceVhostUserDump) Reset()               { *m = SwInterfaceVhostUserDump{} }
func (*SwInterfaceVhostUserDump) GetMessageName() string { return "sw_interface_vhost_user_dump" }
func (*SwInterfaceVhostUserDump) GetCrcString() string   { return "f9e6675e" }
func (*SwInterfaceVhostUserDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SwInterfaceVhostUserDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *SwInterfaceVhostUserDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *SwInterfaceVhostUserDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

func init() { file_vhost_user_binapi_init() }
func file_vhost_user_binapi_init() {
	api.RegisterMessage((*CreateVhostUserIf)(nil), "create_vhost_user_if_c785c6fc")
	api.RegisterMessage((*CreateVhostUserIfReply)(nil), "create_vhost_user_if_reply_5383d31f")
	api.RegisterMessage((*CreateVhostUserIfV2)(nil), "create_vhost_user_if_v2_dba1cc1d")
	api.RegisterMessage((*CreateVhostUserIfV2Reply)(nil), "create_vhost_user_if_v2_reply_5383d31f")
	api.RegisterMessage((*DeleteVhostUserIf)(nil), "delete_vhost_user_if_f9e6675e")
	api.RegisterMessage((*DeleteVhostUserIfReply)(nil), "delete_vhost_user_if_reply_e8d4e804")
	api.RegisterMessage((*ModifyVhostUserIf)(nil), "modify_vhost_user_if_0e71d40b")
	api.RegisterMessage((*ModifyVhostUserIfReply)(nil), "modify_vhost_user_if_reply_e8d4e804")
	api.RegisterMessage((*ModifyVhostUserIfV2)(nil), "modify_vhost_user_if_v2_b2483771")
	api.RegisterMessage((*ModifyVhostUserIfV2Reply)(nil), "modify_vhost_user_if_v2_reply_e8d4e804")
	api.RegisterMessage((*SwInterfaceVhostUserDetails)(nil), "sw_interface_vhost_user_details_0cee1e53")
	api.RegisterMessage((*SwInterfaceVhostUserDump)(nil), "sw_interface_vhost_user_dump_f9e6675e")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*CreateVhostUserIf)(nil),
		(*CreateVhostUserIfReply)(nil),
		(*CreateVhostUserIfV2)(nil),
		(*CreateVhostUserIfV2Reply)(nil),
		(*DeleteVhostUserIf)(nil),
		(*DeleteVhostUserIfReply)(nil),
		(*ModifyVhostUserIf)(nil),
		(*ModifyVhostUserIfReply)(nil),
		(*ModifyVhostUserIfV2)(nil),
		(*ModifyVhostUserIfV2Reply)(nil),
		(*SwInterfaceVhostUserDetails)(nil),
		(*SwInterfaceVhostUserDump)(nil),
	}
}
