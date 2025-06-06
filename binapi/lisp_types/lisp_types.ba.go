// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.12.0
//  VPP:              24.10.0-1~gb3a21b9b0
// source: plugins/lisp_types.api.json

// Package lisp_types contains generated bindings for API file lisp_types.api.
//
// Contents:
// -  2 enums
// -  5 structs
// -  1 union
package lisp_types

import (
	"strconv"
	"vpp-restapi/binapi/ethernet_types"
	"vpp-restapi/binapi/interface_types"
	"vpp-restapi/binapi/ip_types"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "lisp_types"
	APIVersion = "1.0.0"
	VersionCrc = 0xf05d92a6
)

// EidType defines enum 'eid_type'.
type EidType uint8

const (
	EID_TYPE_API_PREFIX EidType = 0
	EID_TYPE_API_MAC    EidType = 1
	EID_TYPE_API_NSH    EidType = 2
)

var (
	EidType_name = map[uint8]string{
		0: "EID_TYPE_API_PREFIX",
		1: "EID_TYPE_API_MAC",
		2: "EID_TYPE_API_NSH",
	}
	EidType_value = map[string]uint8{
		"EID_TYPE_API_PREFIX": 0,
		"EID_TYPE_API_MAC":    1,
		"EID_TYPE_API_NSH":    2,
	}
)

func (x EidType) String() string {
	s, ok := EidType_name[uint8(x)]
	if ok {
		return s
	}
	return "EidType(" + strconv.Itoa(int(x)) + ")"
}

// HmacKeyID defines enum 'hmac_key_id'.
type HmacKeyID uint8

const (
	KEY_ID_API_HMAC_NO_KEY      HmacKeyID = 0
	KEY_ID_API_HMAC_SHA_1_96    HmacKeyID = 1
	KEY_ID_API_HMAC_SHA_256_128 HmacKeyID = 2
)

var (
	HmacKeyID_name = map[uint8]string{
		0: "KEY_ID_API_HMAC_NO_KEY",
		1: "KEY_ID_API_HMAC_SHA_1_96",
		2: "KEY_ID_API_HMAC_SHA_256_128",
	}
	HmacKeyID_value = map[string]uint8{
		"KEY_ID_API_HMAC_NO_KEY":      0,
		"KEY_ID_API_HMAC_SHA_1_96":    1,
		"KEY_ID_API_HMAC_SHA_256_128": 2,
	}
)

func (x HmacKeyID) String() string {
	s, ok := HmacKeyID_name[uint8(x)]
	if ok {
		return s
	}
	return "HmacKeyID(" + strconv.Itoa(int(x)) + ")"
}

// Eid defines type 'eid'.
type Eid struct {
	Type    EidType         `binapi:"eid_type,name=type" json:"type,omitempty"`
	Address EidAddressUnion `binapi:"eid_address,name=address" json:"address,omitempty"`
}

// HmacKey defines type 'hmac_key'.
type HmacKey struct {
	ID  HmacKeyID `binapi:"hmac_key_id,name=id" json:"id,omitempty"`
	Key []byte    `binapi:"u8[64],name=key" json:"key,omitempty"`
}

// LocalLocator defines type 'local_locator'.
type LocalLocator struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Priority  uint8                          `binapi:"u8,name=priority" json:"priority,omitempty"`
	Weight    uint8                          `binapi:"u8,name=weight" json:"weight,omitempty"`
}

// Nsh defines type 'nsh'.
type Nsh struct {
	Spi uint32 `binapi:"u32,name=spi" json:"spi,omitempty"`
	Si  uint8  `binapi:"u8,name=si" json:"si,omitempty"`
}

// RemoteLocator defines type 'remote_locator'.
type RemoteLocator struct {
	Priority  uint8            `binapi:"u8,name=priority" json:"priority,omitempty"`
	Weight    uint8            `binapi:"u8,name=weight" json:"weight,omitempty"`
	IPAddress ip_types.Address `binapi:"address,name=ip_address" json:"ip_address,omitempty"`
}

// EidAddressUnion defines union 'eid_address'.
type EidAddressUnion struct {
	// EidAddressUnion can be one of:
	// - Prefix *ip_types.Prefix
	// - Mac *ethernet_types.MacAddress
	// - Nsh *Nsh
	XXX_UnionData [18]byte
}

func EidAddressUnionPrefix(a ip_types.Prefix) (u EidAddressUnion) {
	u.SetPrefix(a)
	return
}
func (u *EidAddressUnion) SetPrefix(a ip_types.Prefix) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint8(uint8(a.Address.Af))
	buf.EncodeBytes(a.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(a.Len)
}
func (u *EidAddressUnion) GetPrefix() (a ip_types.Prefix) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(a.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	a.Len = buf.DecodeUint8()
	return
}

func EidAddressUnionMac(a ethernet_types.MacAddress) (u EidAddressUnion) {
	u.SetMac(a)
	return
}
func (u *EidAddressUnion) SetMac(a ethernet_types.MacAddress) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeBytes(a[:], 6)
}
func (u *EidAddressUnion) GetMac() (a ethernet_types.MacAddress) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	copy(a[:], buf.DecodeBytes(6))
	return
}

func EidAddressUnionNsh(a Nsh) (u EidAddressUnion) {
	u.SetNsh(a)
	return
}
func (u *EidAddressUnion) SetNsh(a Nsh) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	buf.EncodeUint32(a.Spi)
	buf.EncodeUint8(a.Si)
}
func (u *EidAddressUnion) GetNsh() (a Nsh) {
	buf := codec.NewBuffer(u.XXX_UnionData[:])
	a.Spi = buf.DecodeUint32()
	a.Si = buf.DecodeUint8()
	return
}
