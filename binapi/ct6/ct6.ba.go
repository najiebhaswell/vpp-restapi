// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.12.0
//  VPP:              24.10.0-1~gb3a21b9b0
// source: plugins/ct6.api.json

// Package ct6 contains generated bindings for API file ct6.api.
//
// Contents:
// -  2 messages
package ct6

import (
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
	APIFile    = "ct6"
	APIVersion = "1.0.0"
	VersionCrc = 0x5c824a95
)

// /* Define a simple enable-disable binary API to control the feature
// Ct6EnableDisable defines message 'ct6_enable_disable'.
type Ct6EnableDisable struct {
	EnableDisable bool                           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
	IsInside      bool                           `binapi:"bool,name=is_inside" json:"is_inside,omitempty"`
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *Ct6EnableDisable) Reset()               { *m = Ct6EnableDisable{} }
func (*Ct6EnableDisable) GetMessageName() string { return "ct6_enable_disable" }
func (*Ct6EnableDisable) GetCrcString() string   { return "5d02ac02" }
func (*Ct6EnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *Ct6EnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.EnableDisable
	size += 1 // m.IsInside
	size += 4 // m.SwIfIndex
	return size
}
func (m *Ct6EnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.EnableDisable)
	buf.EncodeBool(m.IsInside)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *Ct6EnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	m.IsInside = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// Ct6EnableDisableReply defines message 'ct6_enable_disable_reply'.
type Ct6EnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *Ct6EnableDisableReply) Reset()               { *m = Ct6EnableDisableReply{} }
func (*Ct6EnableDisableReply) GetMessageName() string { return "ct6_enable_disable_reply" }
func (*Ct6EnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*Ct6EnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *Ct6EnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *Ct6EnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *Ct6EnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_ct6_binapi_init() }
func file_ct6_binapi_init() {
	api.RegisterMessage((*Ct6EnableDisable)(nil), "ct6_enable_disable_5d02ac02")
	api.RegisterMessage((*Ct6EnableDisableReply)(nil), "ct6_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*Ct6EnableDisable)(nil),
		(*Ct6EnableDisableReply)(nil),
	}
}
