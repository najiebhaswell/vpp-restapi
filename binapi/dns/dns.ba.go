// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.12.0
//  VPP:              24.10.0-1~gb3a21b9b0
// source: plugins/dns.api.json

// Package dns contains generated bindings for API file dns.api.
//
// Contents:
// -  8 messages
package dns

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "dns"
	APIVersion = "1.0.0"
	VersionCrc = 0x269575cd
)

// enable/disable name resolution
//   - is_enable - 1 = enable, 0 = disable
//
// DNSEnableDisable defines message 'dns_enable_disable'.
type DNSEnableDisable struct {
	Enable uint8 `binapi:"u8,name=enable" json:"enable,omitempty"`
}

func (m *DNSEnableDisable) Reset()               { *m = DNSEnableDisable{} }
func (*DNSEnableDisable) GetMessageName() string { return "dns_enable_disable" }
func (*DNSEnableDisable) GetCrcString() string   { return "8050327d" }
func (*DNSEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DNSEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Enable
	return size
}
func (m *DNSEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.Enable)
	return buf.Bytes(), nil
}
func (m *DNSEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Enable = buf.DecodeUint8()
	return nil
}

// DNSEnableDisableReply defines message 'dns_enable_disable_reply'.
type DNSEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *DNSEnableDisableReply) Reset()               { *m = DNSEnableDisableReply{} }
func (*DNSEnableDisableReply) GetMessageName() string { return "dns_enable_disable_reply" }
func (*DNSEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*DNSEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DNSEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *DNSEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *DNSEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// add or delete an upstream name server
//   - is_ip6 - an ip6 name server
//   - is_add - add = 1, delete = 0
//   - server_address - server ip address
//
// DNSNameServerAddDel defines message 'dns_name_server_add_del'.
type DNSNameServerAddDel struct {
	IsIP6         uint8  `binapi:"u8,name=is_ip6" json:"is_ip6,omitempty"`
	IsAdd         uint8  `binapi:"u8,name=is_add" json:"is_add,omitempty"`
	ServerAddress []byte `binapi:"u8[16],name=server_address" json:"server_address,omitempty"`
}

func (m *DNSNameServerAddDel) Reset()               { *m = DNSNameServerAddDel{} }
func (*DNSNameServerAddDel) GetMessageName() string { return "dns_name_server_add_del" }
func (*DNSNameServerAddDel) GetCrcString() string   { return "3bb05d8c" }
func (*DNSNameServerAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DNSNameServerAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsIP6
	size += 1      // m.IsAdd
	size += 1 * 16 // m.ServerAddress
	return size
}
func (m *DNSNameServerAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.IsIP6)
	buf.EncodeUint8(m.IsAdd)
	buf.EncodeBytes(m.ServerAddress, 16)
	return buf.Bytes(), nil
}
func (m *DNSNameServerAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsIP6 = buf.DecodeUint8()
	m.IsAdd = buf.DecodeUint8()
	m.ServerAddress = make([]byte, 16)
	copy(m.ServerAddress, buf.DecodeBytes(len(m.ServerAddress)))
	return nil
}

// DNSNameServerAddDelReply defines message 'dns_name_server_add_del_reply'.
type DNSNameServerAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *DNSNameServerAddDelReply) Reset()               { *m = DNSNameServerAddDelReply{} }
func (*DNSNameServerAddDelReply) GetMessageName() string { return "dns_name_server_add_del_reply" }
func (*DNSNameServerAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*DNSNameServerAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DNSNameServerAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *DNSNameServerAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *DNSNameServerAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// DNS IP -> name resolution request
//   - is_ip6 - set if the reverse-DNS request is an ip6 address
//   - address - the address to map to a name
//
// DNSResolveIP defines message 'dns_resolve_ip'.
type DNSResolveIP struct {
	IsIP6   uint8  `binapi:"u8,name=is_ip6" json:"is_ip6,omitempty"`
	Address []byte `binapi:"u8[16],name=address" json:"address,omitempty"`
}

func (m *DNSResolveIP) Reset()               { *m = DNSResolveIP{} }
func (*DNSResolveIP) GetMessageName() string { return "dns_resolve_ip" }
func (*DNSResolveIP) GetCrcString() string   { return "ae96a1a3" }
func (*DNSResolveIP) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DNSResolveIP) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsIP6
	size += 1 * 16 // m.Address
	return size
}
func (m *DNSResolveIP) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.IsIP6)
	buf.EncodeBytes(m.Address, 16)
	return buf.Bytes(), nil
}
func (m *DNSResolveIP) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsIP6 = buf.DecodeUint8()
	m.Address = make([]byte, 16)
	copy(m.Address, buf.DecodeBytes(len(m.Address)))
	return nil
}

// DNS ip->name resolution reply
//   - retval - return value, 0 => success
//   - name - canonical name for the indicated IP address
//
// DNSResolveIPReply defines message 'dns_resolve_ip_reply'.
type DNSResolveIPReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Name   []byte `binapi:"u8[256],name=name" json:"name,omitempty"`
}

func (m *DNSResolveIPReply) Reset()               { *m = DNSResolveIPReply{} }
func (*DNSResolveIPReply) GetMessageName() string { return "dns_resolve_ip_reply" }
func (*DNSResolveIPReply) GetCrcString() string   { return "49ed78d6" }
func (*DNSResolveIPReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DNSResolveIPReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4       // m.Retval
	size += 1 * 256 // m.Name
	return size
}
func (m *DNSResolveIPReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeBytes(m.Name, 256)
	return buf.Bytes(), nil
}
func (m *DNSResolveIPReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Name = make([]byte, 256)
	copy(m.Name, buf.DecodeBytes(len(m.Name)))
	return nil
}

// DNS name resolution request
//   - name - the name to resolve
//
// DNSResolveName defines message 'dns_resolve_name'.
type DNSResolveName struct {
	Name []byte `binapi:"u8[256],name=name" json:"name,omitempty"`
}

func (m *DNSResolveName) Reset()               { *m = DNSResolveName{} }
func (*DNSResolveName) GetMessageName() string { return "dns_resolve_name" }
func (*DNSResolveName) GetCrcString() string   { return "c6566676" }
func (*DNSResolveName) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *DNSResolveName) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 * 256 // m.Name
	return size
}
func (m *DNSResolveName) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBytes(m.Name, 256)
	return buf.Bytes(), nil
}
func (m *DNSResolveName) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Name = make([]byte, 256)
	copy(m.Name, buf.DecodeBytes(len(m.Name)))
	return nil
}

// DNS name resolution reply
//   - retval - return value, 0 => success
//   - ip4_set - indicates that the ip4 address is valid
//   - ip6_set - indicates that the ip6 address is valid
//   - ip4_address - the ip4 name resolution reply
//   - ip6_address - the ip6 name resolution reply
//
// DNSResolveNameReply defines message 'dns_resolve_name_reply'.
type DNSResolveNameReply struct {
	Retval     int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	IP4Set     uint8  `binapi:"u8,name=ip4_set" json:"ip4_set,omitempty"`
	IP6Set     uint8  `binapi:"u8,name=ip6_set" json:"ip6_set,omitempty"`
	IP4Address []byte `binapi:"u8[4],name=ip4_address" json:"ip4_address,omitempty"`
	IP6Address []byte `binapi:"u8[16],name=ip6_address" json:"ip6_address,omitempty"`
}

func (m *DNSResolveNameReply) Reset()               { *m = DNSResolveNameReply{} }
func (*DNSResolveNameReply) GetMessageName() string { return "dns_resolve_name_reply" }
func (*DNSResolveNameReply) GetCrcString() string   { return "c2d758c3" }
func (*DNSResolveNameReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *DNSResolveNameReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Retval
	size += 1      // m.IP4Set
	size += 1      // m.IP6Set
	size += 1 * 4  // m.IP4Address
	size += 1 * 16 // m.IP6Address
	return size
}
func (m *DNSResolveNameReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint8(m.IP4Set)
	buf.EncodeUint8(m.IP6Set)
	buf.EncodeBytes(m.IP4Address, 4)
	buf.EncodeBytes(m.IP6Address, 16)
	return buf.Bytes(), nil
}
func (m *DNSResolveNameReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.IP4Set = buf.DecodeUint8()
	m.IP6Set = buf.DecodeUint8()
	m.IP4Address = make([]byte, 4)
	copy(m.IP4Address, buf.DecodeBytes(len(m.IP4Address)))
	m.IP6Address = make([]byte, 16)
	copy(m.IP6Address, buf.DecodeBytes(len(m.IP6Address)))
	return nil
}

func init() { file_dns_binapi_init() }
func file_dns_binapi_init() {
	api.RegisterMessage((*DNSEnableDisable)(nil), "dns_enable_disable_8050327d")
	api.RegisterMessage((*DNSEnableDisableReply)(nil), "dns_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*DNSNameServerAddDel)(nil), "dns_name_server_add_del_3bb05d8c")
	api.RegisterMessage((*DNSNameServerAddDelReply)(nil), "dns_name_server_add_del_reply_e8d4e804")
	api.RegisterMessage((*DNSResolveIP)(nil), "dns_resolve_ip_ae96a1a3")
	api.RegisterMessage((*DNSResolveIPReply)(nil), "dns_resolve_ip_reply_49ed78d6")
	api.RegisterMessage((*DNSResolveName)(nil), "dns_resolve_name_c6566676")
	api.RegisterMessage((*DNSResolveNameReply)(nil), "dns_resolve_name_reply_c2d758c3")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*DNSEnableDisable)(nil),
		(*DNSEnableDisableReply)(nil),
		(*DNSNameServerAddDel)(nil),
		(*DNSNameServerAddDelReply)(nil),
		(*DNSResolveIP)(nil),
		(*DNSResolveIPReply)(nil),
		(*DNSResolveName)(nil),
		(*DNSResolveNameReply)(nil),
	}
}
