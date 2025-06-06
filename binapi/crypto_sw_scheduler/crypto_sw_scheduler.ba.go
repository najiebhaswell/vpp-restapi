// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.12.0
//  VPP:              24.10.0-1~gb3a21b9b0
// source: plugins/crypto_sw_scheduler.api.json

// Package crypto_sw_scheduler contains generated bindings for API file crypto_sw_scheduler.api.
//
// Contents:
// -  2 messages
package crypto_sw_scheduler

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
	APIFile    = "crypto_sw_scheduler"
	APIVersion = "1.1.0"
	VersionCrc = 0xf4b02951
)

// crypto sw scheduler: Enable or disable workers
//   - worker_index - Worker index to enable / disable
//   - crypto_enable - On/Off
//
// CryptoSwSchedulerSetWorker defines message 'crypto_sw_scheduler_set_worker'.
type CryptoSwSchedulerSetWorker struct {
	WorkerIndex  uint32 `binapi:"u32,name=worker_index" json:"worker_index,omitempty"`
	CryptoEnable bool   `binapi:"bool,name=crypto_enable" json:"crypto_enable,omitempty"`
}

func (m *CryptoSwSchedulerSetWorker) Reset()               { *m = CryptoSwSchedulerSetWorker{} }
func (*CryptoSwSchedulerSetWorker) GetMessageName() string { return "crypto_sw_scheduler_set_worker" }
func (*CryptoSwSchedulerSetWorker) GetCrcString() string   { return "b4274502" }
func (*CryptoSwSchedulerSetWorker) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CryptoSwSchedulerSetWorker) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.WorkerIndex
	size += 1 // m.CryptoEnable
	return size
}
func (m *CryptoSwSchedulerSetWorker) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.WorkerIndex)
	buf.EncodeBool(m.CryptoEnable)
	return buf.Bytes(), nil
}
func (m *CryptoSwSchedulerSetWorker) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.WorkerIndex = buf.DecodeUint32()
	m.CryptoEnable = buf.DecodeBool()
	return nil
}

// CryptoSwSchedulerSetWorkerReply defines message 'crypto_sw_scheduler_set_worker_reply'.
type CryptoSwSchedulerSetWorkerReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *CryptoSwSchedulerSetWorkerReply) Reset() { *m = CryptoSwSchedulerSetWorkerReply{} }
func (*CryptoSwSchedulerSetWorkerReply) GetMessageName() string {
	return "crypto_sw_scheduler_set_worker_reply"
}
func (*CryptoSwSchedulerSetWorkerReply) GetCrcString() string { return "e8d4e804" }
func (*CryptoSwSchedulerSetWorkerReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CryptoSwSchedulerSetWorkerReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *CryptoSwSchedulerSetWorkerReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *CryptoSwSchedulerSetWorkerReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_crypto_sw_scheduler_binapi_init() }
func file_crypto_sw_scheduler_binapi_init() {
	api.RegisterMessage((*CryptoSwSchedulerSetWorker)(nil), "crypto_sw_scheduler_set_worker_b4274502")
	api.RegisterMessage((*CryptoSwSchedulerSetWorkerReply)(nil), "crypto_sw_scheduler_set_worker_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*CryptoSwSchedulerSetWorker)(nil),
		(*CryptoSwSchedulerSetWorkerReply)(nil),
	}
}
