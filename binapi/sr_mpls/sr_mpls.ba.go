// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.12.0
//  VPP:              24.10.0-1~gb3a21b9b0
// source: plugins/sr_mpls.api.json

// Package sr_mpls contains generated bindings for API file sr_mpls.api.
//
// Contents:
// - 10 messages
package sr_mpls

import (
	_ "vpp-restapi/binapi/interface_types"
	"vpp-restapi/binapi/ip_types"
	"vpp-restapi/binapi/sr_types"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "sr_mpls"
	APIVersion = "3.0.0"
	VersionCrc = 0x156edb17
)

// MPLS SR policy add
//   - bsid - is the bindingSID of the SR Policy. MPLS label (20bit)
//   - weight - is the weight of the sid list. optional.
//   - is_spray - is the type of the SR policy. (0.Default // 1.Spray)
//   - segments - vector of labels (20bit) composing the segment list
//
// SrMplsPolicyAdd defines message 'sr_mpls_policy_add'.
type SrMplsPolicyAdd struct {
	Bsid      uint32   `binapi:"u32,name=bsid" json:"bsid,omitempty"`
	Weight    uint32   `binapi:"u32,name=weight" json:"weight,omitempty"`
	IsSpray   bool     `binapi:"bool,name=is_spray" json:"is_spray,omitempty"`
	NSegments uint8    `binapi:"u8,name=n_segments" json:"-"`
	Segments  []uint32 `binapi:"u32[n_segments],name=segments" json:"segments,omitempty"`
}

func (m *SrMplsPolicyAdd) Reset()               { *m = SrMplsPolicyAdd{} }
func (*SrMplsPolicyAdd) GetMessageName() string { return "sr_mpls_policy_add" }
func (*SrMplsPolicyAdd) GetCrcString() string   { return "a1a70c70" }
func (*SrMplsPolicyAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrMplsPolicyAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                   // m.Bsid
	size += 4                   // m.Weight
	size += 1                   // m.IsSpray
	size += 1                   // m.NSegments
	size += 4 * len(m.Segments) // m.Segments
	return size
}
func (m *SrMplsPolicyAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Bsid)
	buf.EncodeUint32(m.Weight)
	buf.EncodeBool(m.IsSpray)
	buf.EncodeUint8(uint8(len(m.Segments)))
	for i := 0; i < len(m.Segments); i++ {
		var x uint32
		if i < len(m.Segments) {
			x = uint32(m.Segments[i])
		}
		buf.EncodeUint32(x)
	}
	return buf.Bytes(), nil
}
func (m *SrMplsPolicyAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Bsid = buf.DecodeUint32()
	m.Weight = buf.DecodeUint32()
	m.IsSpray = buf.DecodeBool()
	m.NSegments = buf.DecodeUint8()
	m.Segments = make([]uint32, m.NSegments)
	for i := 0; i < len(m.Segments); i++ {
		m.Segments[i] = buf.DecodeUint32()
	}
	return nil
}

// SrMplsPolicyAddReply defines message 'sr_mpls_policy_add_reply'.
type SrMplsPolicyAddReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrMplsPolicyAddReply) Reset()               { *m = SrMplsPolicyAddReply{} }
func (*SrMplsPolicyAddReply) GetMessageName() string { return "sr_mpls_policy_add_reply" }
func (*SrMplsPolicyAddReply) GetCrcString() string   { return "e8d4e804" }
func (*SrMplsPolicyAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrMplsPolicyAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrMplsPolicyAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrMplsPolicyAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// MPLS SR steering add/del
//   - bsid is the bindingSID of the SR Policy
//   - endpoint is the endpoint of the SR policy
//   - color is the color of the sr policy
//
// SrMplsPolicyAssignEndpointColor defines message 'sr_mpls_policy_assign_endpoint_color'.
type SrMplsPolicyAssignEndpointColor struct {
	Bsid     uint32           `binapi:"u32,name=bsid" json:"bsid,omitempty"`
	Endpoint ip_types.Address `binapi:"address,name=endpoint" json:"endpoint,omitempty"`
	Color    uint32           `binapi:"u32,name=color" json:"color,omitempty"`
}

func (m *SrMplsPolicyAssignEndpointColor) Reset() { *m = SrMplsPolicyAssignEndpointColor{} }
func (*SrMplsPolicyAssignEndpointColor) GetMessageName() string {
	return "sr_mpls_policy_assign_endpoint_color"
}
func (*SrMplsPolicyAssignEndpointColor) GetCrcString() string { return "0e7eb978" }
func (*SrMplsPolicyAssignEndpointColor) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrMplsPolicyAssignEndpointColor) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.Bsid
	size += 1      // m.Endpoint.Af
	size += 1 * 16 // m.Endpoint.Un
	size += 4      // m.Color
	return size
}
func (m *SrMplsPolicyAssignEndpointColor) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Bsid)
	buf.EncodeUint8(uint8(m.Endpoint.Af))
	buf.EncodeBytes(m.Endpoint.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.Color)
	return buf.Bytes(), nil
}
func (m *SrMplsPolicyAssignEndpointColor) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Bsid = buf.DecodeUint32()
	m.Endpoint.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Endpoint.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Color = buf.DecodeUint32()
	return nil
}

// SrMplsPolicyAssignEndpointColorReply defines message 'sr_mpls_policy_assign_endpoint_color_reply'.
type SrMplsPolicyAssignEndpointColorReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrMplsPolicyAssignEndpointColorReply) Reset() { *m = SrMplsPolicyAssignEndpointColorReply{} }
func (*SrMplsPolicyAssignEndpointColorReply) GetMessageName() string {
	return "sr_mpls_policy_assign_endpoint_color_reply"
}
func (*SrMplsPolicyAssignEndpointColorReply) GetCrcString() string { return "e8d4e804" }
func (*SrMplsPolicyAssignEndpointColorReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrMplsPolicyAssignEndpointColorReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrMplsPolicyAssignEndpointColorReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrMplsPolicyAssignEndpointColorReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// MPLS SR policy deletion
//   - bsid is the bindingSID of the SR Policy. MPLS label (20bit)
//
// SrMplsPolicyDel defines message 'sr_mpls_policy_del'.
type SrMplsPolicyDel struct {
	Bsid uint32 `binapi:"u32,name=bsid" json:"bsid,omitempty"`
}

func (m *SrMplsPolicyDel) Reset()               { *m = SrMplsPolicyDel{} }
func (*SrMplsPolicyDel) GetMessageName() string { return "sr_mpls_policy_del" }
func (*SrMplsPolicyDel) GetCrcString() string   { return "e29d34fa" }
func (*SrMplsPolicyDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrMplsPolicyDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Bsid
	return size
}
func (m *SrMplsPolicyDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Bsid)
	return buf.Bytes(), nil
}
func (m *SrMplsPolicyDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Bsid = buf.DecodeUint32()
	return nil
}

// SrMplsPolicyDelReply defines message 'sr_mpls_policy_del_reply'.
type SrMplsPolicyDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrMplsPolicyDelReply) Reset()               { *m = SrMplsPolicyDelReply{} }
func (*SrMplsPolicyDelReply) GetMessageName() string { return "sr_mpls_policy_del_reply" }
func (*SrMplsPolicyDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SrMplsPolicyDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrMplsPolicyDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrMplsPolicyDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrMplsPolicyDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// MPLS SR policy modification
//   - bsid is the bindingSID of the SR Policy. MPLS label (20bit)
//   - sr_policy_index is the index of the SR policy
//   - fib_table is the VRF where to install the FIB entry for the BSID
//   - operation is the operation to perform (among the top ones)
//   - segments is a vector of MPLS labels composing the segment list
//   - sl_index is the index of the Segment List to modify/delete
//   - weight is the weight of the sid list. optional.
//   - is_encap Mode. Encapsulation or SRH insertion.
//
// SrMplsPolicyMod defines message 'sr_mpls_policy_mod'.
type SrMplsPolicyMod struct {
	Bsid      uint32              `binapi:"u32,name=bsid" json:"bsid,omitempty"`
	Operation sr_types.SrPolicyOp `binapi:"sr_policy_op,name=operation" json:"operation,omitempty"`
	SlIndex   uint32              `binapi:"u32,name=sl_index" json:"sl_index,omitempty"`
	Weight    uint32              `binapi:"u32,name=weight" json:"weight,omitempty"`
	NSegments uint8               `binapi:"u8,name=n_segments" json:"-"`
	Segments  []uint32            `binapi:"u32[n_segments],name=segments" json:"segments,omitempty"`
}

func (m *SrMplsPolicyMod) Reset()               { *m = SrMplsPolicyMod{} }
func (*SrMplsPolicyMod) GetMessageName() string { return "sr_mpls_policy_mod" }
func (*SrMplsPolicyMod) GetCrcString() string   { return "88482c17" }
func (*SrMplsPolicyMod) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrMplsPolicyMod) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                   // m.Bsid
	size += 1                   // m.Operation
	size += 4                   // m.SlIndex
	size += 4                   // m.Weight
	size += 1                   // m.NSegments
	size += 4 * len(m.Segments) // m.Segments
	return size
}
func (m *SrMplsPolicyMod) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Bsid)
	buf.EncodeUint8(uint8(m.Operation))
	buf.EncodeUint32(m.SlIndex)
	buf.EncodeUint32(m.Weight)
	buf.EncodeUint8(uint8(len(m.Segments)))
	for i := 0; i < len(m.Segments); i++ {
		var x uint32
		if i < len(m.Segments) {
			x = uint32(m.Segments[i])
		}
		buf.EncodeUint32(x)
	}
	return buf.Bytes(), nil
}
func (m *SrMplsPolicyMod) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Bsid = buf.DecodeUint32()
	m.Operation = sr_types.SrPolicyOp(buf.DecodeUint8())
	m.SlIndex = buf.DecodeUint32()
	m.Weight = buf.DecodeUint32()
	m.NSegments = buf.DecodeUint8()
	m.Segments = make([]uint32, m.NSegments)
	for i := 0; i < len(m.Segments); i++ {
		m.Segments[i] = buf.DecodeUint32()
	}
	return nil
}

// SrMplsPolicyModReply defines message 'sr_mpls_policy_mod_reply'.
type SrMplsPolicyModReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrMplsPolicyModReply) Reset()               { *m = SrMplsPolicyModReply{} }
func (*SrMplsPolicyModReply) GetMessageName() string { return "sr_mpls_policy_mod_reply" }
func (*SrMplsPolicyModReply) GetCrcString() string   { return "e8d4e804" }
func (*SrMplsPolicyModReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrMplsPolicyModReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrMplsPolicyModReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrMplsPolicyModReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// MPLS SR steering add/del
//   - is_del
//   - bsid - is the bindingSID of the SR Policy (~0 is no bsid)
//   - table_id - is the VRF where to install the FIB entry for the BSID
//   - prefix - is the IPv4/v6 address for L3 traffic type.
//   - mask_width - is the mask for L3 traffic type
//   - next_hop - describes the next_hop (in case no BSID)
//   - color - describes the color
//   - co_bits - are the CO_bits of the steering policy
//   - vpn_label - is an additonal last VPN label. (~0 is no label)
//
// SrMplsSteeringAddDel defines message 'sr_mpls_steering_add_del'.
type SrMplsSteeringAddDel struct {
	IsDel     bool             `binapi:"bool,name=is_del,default=false" json:"is_del,omitempty"`
	Bsid      uint32           `binapi:"u32,name=bsid" json:"bsid,omitempty"`
	TableID   uint32           `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	Prefix    ip_types.Prefix  `binapi:"prefix,name=prefix" json:"prefix,omitempty"`
	MaskWidth uint32           `binapi:"u32,name=mask_width" json:"mask_width,omitempty"`
	NextHop   ip_types.Address `binapi:"address,name=next_hop" json:"next_hop,omitempty"`
	Color     uint32           `binapi:"u32,name=color" json:"color,omitempty"`
	CoBits    uint8            `binapi:"u8,name=co_bits" json:"co_bits,omitempty"`
	VPNLabel  uint32           `binapi:"u32,name=vpn_label" json:"vpn_label,omitempty"`
}

func (m *SrMplsSteeringAddDel) Reset()               { *m = SrMplsSteeringAddDel{} }
func (*SrMplsSteeringAddDel) GetMessageName() string { return "sr_mpls_steering_add_del" }
func (*SrMplsSteeringAddDel) GetCrcString() string   { return "64acff63" }
func (*SrMplsSteeringAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *SrMplsSteeringAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsDel
	size += 4      // m.Bsid
	size += 4      // m.TableID
	size += 1      // m.Prefix.Address.Af
	size += 1 * 16 // m.Prefix.Address.Un
	size += 1      // m.Prefix.Len
	size += 4      // m.MaskWidth
	size += 1      // m.NextHop.Af
	size += 1 * 16 // m.NextHop.Un
	size += 4      // m.Color
	size += 1      // m.CoBits
	size += 4      // m.VPNLabel
	return size
}
func (m *SrMplsSteeringAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsDel)
	buf.EncodeUint32(m.Bsid)
	buf.EncodeUint32(m.TableID)
	buf.EncodeUint8(uint8(m.Prefix.Address.Af))
	buf.EncodeBytes(m.Prefix.Address.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.Prefix.Len)
	buf.EncodeUint32(m.MaskWidth)
	buf.EncodeUint8(uint8(m.NextHop.Af))
	buf.EncodeBytes(m.NextHop.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(m.Color)
	buf.EncodeUint8(m.CoBits)
	buf.EncodeUint32(m.VPNLabel)
	return buf.Bytes(), nil
}
func (m *SrMplsSteeringAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsDel = buf.DecodeBool()
	m.Bsid = buf.DecodeUint32()
	m.TableID = buf.DecodeUint32()
	m.Prefix.Address.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Prefix.Address.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Prefix.Len = buf.DecodeUint8()
	m.MaskWidth = buf.DecodeUint32()
	m.NextHop.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.NextHop.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.Color = buf.DecodeUint32()
	m.CoBits = buf.DecodeUint8()
	m.VPNLabel = buf.DecodeUint32()
	return nil
}

// SrMplsSteeringAddDelReply defines message 'sr_mpls_steering_add_del_reply'.
type SrMplsSteeringAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SrMplsSteeringAddDelReply) Reset()               { *m = SrMplsSteeringAddDelReply{} }
func (*SrMplsSteeringAddDelReply) GetMessageName() string { return "sr_mpls_steering_add_del_reply" }
func (*SrMplsSteeringAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*SrMplsSteeringAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *SrMplsSteeringAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *SrMplsSteeringAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *SrMplsSteeringAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_sr_mpls_binapi_init() }
func file_sr_mpls_binapi_init() {
	api.RegisterMessage((*SrMplsPolicyAdd)(nil), "sr_mpls_policy_add_a1a70c70")
	api.RegisterMessage((*SrMplsPolicyAddReply)(nil), "sr_mpls_policy_add_reply_e8d4e804")
	api.RegisterMessage((*SrMplsPolicyAssignEndpointColor)(nil), "sr_mpls_policy_assign_endpoint_color_0e7eb978")
	api.RegisterMessage((*SrMplsPolicyAssignEndpointColorReply)(nil), "sr_mpls_policy_assign_endpoint_color_reply_e8d4e804")
	api.RegisterMessage((*SrMplsPolicyDel)(nil), "sr_mpls_policy_del_e29d34fa")
	api.RegisterMessage((*SrMplsPolicyDelReply)(nil), "sr_mpls_policy_del_reply_e8d4e804")
	api.RegisterMessage((*SrMplsPolicyMod)(nil), "sr_mpls_policy_mod_88482c17")
	api.RegisterMessage((*SrMplsPolicyModReply)(nil), "sr_mpls_policy_mod_reply_e8d4e804")
	api.RegisterMessage((*SrMplsSteeringAddDel)(nil), "sr_mpls_steering_add_del_64acff63")
	api.RegisterMessage((*SrMplsSteeringAddDelReply)(nil), "sr_mpls_steering_add_del_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*SrMplsPolicyAdd)(nil),
		(*SrMplsPolicyAddReply)(nil),
		(*SrMplsPolicyAssignEndpointColor)(nil),
		(*SrMplsPolicyAssignEndpointColorReply)(nil),
		(*SrMplsPolicyDel)(nil),
		(*SrMplsPolicyDelReply)(nil),
		(*SrMplsPolicyMod)(nil),
		(*SrMplsPolicyModReply)(nil),
		(*SrMplsSteeringAddDel)(nil),
		(*SrMplsSteeringAddDelReply)(nil),
	}
}
