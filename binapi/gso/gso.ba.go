// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.12.0
//  VPP:              unknown
// source: gso.api.json

// Package gso contains generated bindings for API file gso.api.
//
// Contents:
// -  2 messages
package gso

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
	APIFile    = "gso"
	APIVersion = "1.0.0"
	VersionCrc = 0x81a73026
)

// Enable or disable interface feature gso arc
//   - sw_if_index - The interface to enable/disable gso feature arc.
//   - enable_disable - set to 1 to enable, 0 to disable gso feature arc
//
// FeatureGsoEnableDisable defines message 'feature_gso_enable_disable'.
type FeatureGsoEnableDisable struct {
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	EnableDisable bool                           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
}

func (m *FeatureGsoEnableDisable) Reset()               { *m = FeatureGsoEnableDisable{} }
func (*FeatureGsoEnableDisable) GetMessageName() string { return "feature_gso_enable_disable" }
func (*FeatureGsoEnableDisable) GetCrcString() string   { return "5501adee" }
func (*FeatureGsoEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *FeatureGsoEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.EnableDisable
	return size
}
func (m *FeatureGsoEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.EnableDisable)
	return buf.Bytes(), nil
}
func (m *FeatureGsoEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.EnableDisable = buf.DecodeBool()
	return nil
}

// FeatureGsoEnableDisableReply defines message 'feature_gso_enable_disable_reply'.
type FeatureGsoEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *FeatureGsoEnableDisableReply) Reset() { *m = FeatureGsoEnableDisableReply{} }
func (*FeatureGsoEnableDisableReply) GetMessageName() string {
	return "feature_gso_enable_disable_reply"
}
func (*FeatureGsoEnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*FeatureGsoEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *FeatureGsoEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *FeatureGsoEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *FeatureGsoEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_gso_binapi_init() }
func file_gso_binapi_init() {
	api.RegisterMessage((*FeatureGsoEnableDisable)(nil), "feature_gso_enable_disable_5501adee")
	api.RegisterMessage((*FeatureGsoEnableDisableReply)(nil), "feature_gso_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*FeatureGsoEnableDisable)(nil),
		(*FeatureGsoEnableDisableReply)(nil),
	}
}
