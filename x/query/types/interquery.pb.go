// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query/interquery.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/tendermint/tendermint/proto/tendermint/crypto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Interquery struct {
	Storeid       string `protobuf:"bytes,1,opt,name=storeid,proto3" json:"storeid,omitempty"`
	Chainid       string `protobuf:"bytes,2,opt,name=chainid,proto3" json:"chainid,omitempty"`
	Path          string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Key           []byte `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	TimeoutHeight uint64 `protobuf:"varint,5,opt,name=timeoutHeight,proto3" json:"timeoutHeight,omitempty"`
	ConnectionId  string `protobuf:"bytes,6,opt,name=connectionId,proto3" json:"connectionId,omitempty"`
	ClientId      string `protobuf:"bytes,7,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (m *Interquery) Reset()         { *m = Interquery{} }
func (m *Interquery) String() string { return proto.CompactTextString(m) }
func (*Interquery) ProtoMessage()    {}
func (*Interquery) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfa9beb887ed5f4, []int{0}
}
func (m *Interquery) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Interquery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Interquery.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Interquery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interquery.Merge(m, src)
}
func (m *Interquery) XXX_Size() int {
	return m.Size()
}
func (m *Interquery) XXX_DiscardUnknown() {
	xxx_messageInfo_Interquery.DiscardUnknown(m)
}

var xxx_messageInfo_Interquery proto.InternalMessageInfo

func (m *Interquery) GetStoreid() string {
	if m != nil {
		return m.Storeid
	}
	return ""
}

func (m *Interquery) GetChainid() string {
	if m != nil {
		return m.Chainid
	}
	return ""
}

func (m *Interquery) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Interquery) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Interquery) GetTimeoutHeight() uint64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func (m *Interquery) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *Interquery) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type InterqueryResult struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Storeid string `protobuf:"bytes,2,opt,name=storeid,proto3" json:"storeid,omitempty"`
	Chainid string `protobuf:"bytes,3,opt,name=chainid,proto3" json:"chainid,omitempty"`
	Data    []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// queried chain height on submission
	Height *types.Height `protobuf:"bytes,5,opt,name=height,proto3" json:"height,omitempty"`
	// querying chain height on submission
	LocalHeight uint64 `protobuf:"varint,6,opt,name=localHeight,proto3" json:"localHeight,omitempty"`
	Success     bool   `protobuf:"varint,7,opt,name=success,proto3" json:"success,omitempty"`
	Proved      bool   `protobuf:"varint,8,opt,name=proved,proto3" json:"proved,omitempty"`
}

func (m *InterqueryResult) Reset()         { *m = InterqueryResult{} }
func (m *InterqueryResult) String() string { return proto.CompactTextString(m) }
func (*InterqueryResult) ProtoMessage()    {}
func (*InterqueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfa9beb887ed5f4, []int{1}
}
func (m *InterqueryResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterqueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterqueryResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterqueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterqueryResult.Merge(m, src)
}
func (m *InterqueryResult) XXX_Size() int {
	return m.Size()
}
func (m *InterqueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_InterqueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_InterqueryResult proto.InternalMessageInfo

func (m *InterqueryResult) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *InterqueryResult) GetStoreid() string {
	if m != nil {
		return m.Storeid
	}
	return ""
}

func (m *InterqueryResult) GetChainid() string {
	if m != nil {
		return m.Chainid
	}
	return ""
}

func (m *InterqueryResult) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *InterqueryResult) GetHeight() *types.Height {
	if m != nil {
		return m.Height
	}
	return nil
}

func (m *InterqueryResult) GetLocalHeight() uint64 {
	if m != nil {
		return m.LocalHeight
	}
	return 0
}

func (m *InterqueryResult) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *InterqueryResult) GetProved() bool {
	if m != nil {
		return m.Proved
	}
	return false
}

type InterqueryTimeoutResult struct {
	Storeid       string `protobuf:"bytes,1,opt,name=storeid,proto3" json:"storeid,omitempty"`
	TimeoutHeight uint64 `protobuf:"varint,2,opt,name=timeoutHeight,proto3" json:"timeoutHeight,omitempty"`
}

func (m *InterqueryTimeoutResult) Reset()         { *m = InterqueryTimeoutResult{} }
func (m *InterqueryTimeoutResult) String() string { return proto.CompactTextString(m) }
func (*InterqueryTimeoutResult) ProtoMessage()    {}
func (*InterqueryTimeoutResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_edfa9beb887ed5f4, []int{2}
}
func (m *InterqueryTimeoutResult) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InterqueryTimeoutResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InterqueryTimeoutResult.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InterqueryTimeoutResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterqueryTimeoutResult.Merge(m, src)
}
func (m *InterqueryTimeoutResult) XXX_Size() int {
	return m.Size()
}
func (m *InterqueryTimeoutResult) XXX_DiscardUnknown() {
	xxx_messageInfo_InterqueryTimeoutResult.DiscardUnknown(m)
}

var xxx_messageInfo_InterqueryTimeoutResult proto.InternalMessageInfo

func (m *InterqueryTimeoutResult) GetStoreid() string {
	if m != nil {
		return m.Storeid
	}
	return ""
}

func (m *InterqueryTimeoutResult) GetTimeoutHeight() uint64 {
	if m != nil {
		return m.TimeoutHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*Interquery)(nil), "defundlabs.defund.query.Interquery")
	proto.RegisterType((*InterqueryResult)(nil), "defundlabs.defund.query.InterqueryResult")
	proto.RegisterType((*InterqueryTimeoutResult)(nil), "defundlabs.defund.query.InterqueryTimeoutResult")
}

func init() { proto.RegisterFile("query/interquery.proto", fileDescriptor_edfa9beb887ed5f4) }

var fileDescriptor_edfa9beb887ed5f4 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0xb6, 0x64, 0xc5, 0x1b, 0xd2, 0xe4, 0xc3, 0x66, 0x55, 0x22, 0x54, 0x15, 0x87,
	0x4a, 0x88, 0x58, 0x1b, 0xdf, 0x00, 0x09, 0x89, 0x5e, 0x23, 0x2e, 0x70, 0x4b, 0xec, 0xb7, 0xc5,
	0x22, 0xb5, 0x83, 0xf3, 0x52, 0xd1, 0x6f, 0xc1, 0xc7, 0xe2, 0xc0, 0x61, 0x47, 0x8e, 0xa8, 0xfd,
	0x1a, 0x1c, 0x50, 0x6d, 0x77, 0xd9, 0x34, 0x7a, 0xfb, 0xff, 0xdf, 0x73, 0x62, 0xff, 0xfe, 0xef,
	0xd1, 0x8b, 0x6f, 0x1d, 0xb8, 0x8d, 0xd0, 0x06, 0xc1, 0x79, 0x99, 0x35, 0xce, 0xa2, 0x65, 0x97,
	0x0a, 0x6e, 0x3a, 0xa3, 0xea, 0xa2, 0x6c, 0xb3, 0x20, 0x33, 0xdf, 0x9e, 0xbe, 0x44, 0x30, 0x0a,
	0xdc, 0x4a, 0x1b, 0x14, 0xd2, 0x6d, 0x1a, 0xb4, 0xa2, 0x71, 0xd6, 0xde, 0x84, 0xef, 0xa6, 0xaf,
	0x74, 0x29, 0x85, 0xb4, 0x0e, 0x84, 0xac, 0x35, 0x18, 0x14, 0xeb, 0xab, 0xa8, 0xc2, 0x81, 0xf9,
	0x2f, 0x42, 0xe9, 0xf2, 0xfe, 0x36, 0xc6, 0xe9, 0x49, 0x8b, 0xd6, 0x81, 0x56, 0x9c, 0xcc, 0xc8,
	0xe2, 0x79, 0x7e, 0xb0, 0xfb, 0x8e, 0xac, 0x0a, 0x6d, 0xb4, 0xe2, 0xc3, 0xd0, 0x89, 0x96, 0x31,
	0x3a, 0x6e, 0x0a, 0xac, 0xf8, 0xc8, 0x97, 0xbd, 0x66, 0xe7, 0x74, 0xf4, 0x15, 0x36, 0x7c, 0x3c,
	0x23, 0x8b, 0xb3, 0x7c, 0x2f, 0xd9, 0x6b, 0xfa, 0x02, 0xf5, 0x0a, 0x6c, 0x87, 0x1f, 0x41, 0xdf,
	0x56, 0xc8, 0x9f, 0xcd, 0xc8, 0x62, 0x9c, 0x3f, 0x2e, 0xb2, 0x39, 0x3d, 0x93, 0xd6, 0x18, 0x90,
	0xa8, 0xad, 0x59, 0x2a, 0x9e, 0xf8, 0x7f, 0x3e, 0xaa, 0xb1, 0x29, 0x9d, 0x04, 0x84, 0xa5, 0xe2,
	0x27, 0xbe, 0x7f, 0xef, 0xe7, 0x7f, 0x09, 0x3d, 0xef, 0x71, 0x72, 0x68, 0xbb, 0x1a, 0xfd, 0xd3,
	0x1d, 0x14, 0x68, 0xdd, 0x01, 0x2a, 0xda, 0x87, 0xb8, 0xc3, 0xa3, 0xb8, 0xa3, 0x27, 0xb8, 0xaa,
	0xc0, 0x22, 0xb2, 0x79, 0xcd, 0xae, 0x69, 0x52, 0xf5, 0x54, 0xa7, 0xd7, 0xd3, 0x4c, 0x97, 0x32,
	0xdb, 0xe7, 0x9e, 0xc5, 0xb4, 0xd7, 0x57, 0x59, 0x40, 0xcc, 0xe3, 0x49, 0x36, 0xa3, 0xa7, 0xb5,
	0x95, 0x45, 0x1d, 0xe3, 0x48, 0x7c, 0x1c, 0x0f, 0x4b, 0xfe, 0x75, 0x9d, 0x94, 0xd0, 0xb6, 0x9e,
	0x73, 0x92, 0x1f, 0x2c, 0xbb, 0xa0, 0x49, 0xe3, 0xec, 0x1a, 0x14, 0x9f, 0xf8, 0x46, 0x74, 0xf3,
	0xcf, 0xf4, 0xb2, 0xa7, 0xff, 0x14, 0x92, 0xed, 0x43, 0x38, 0x32, 0xd9, 0x27, 0x93, 0x19, 0xfe,
	0x67, 0x32, 0xef, 0x3f, 0xfc, 0xdc, 0xa6, 0xe4, 0x6e, 0x9b, 0x92, 0x3f, 0xdb, 0x94, 0xfc, 0xd8,
	0xa5, 0x83, 0xbb, 0x5d, 0x3a, 0xf8, 0xbd, 0x4b, 0x07, 0x5f, 0xde, 0xdc, 0x6a, 0xac, 0xba, 0x32,
	0x93, 0x76, 0x25, 0xc2, 0x6e, 0xbe, 0xdd, 0xef, 0x69, 0xd4, 0xe2, 0xbb, 0x08, 0x3b, 0x8d, 0x9b,
	0x06, 0xda, 0x32, 0xf1, 0x6b, 0xf7, 0xee, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x31, 0xc8,
	0xb2, 0xe9, 0x02, 0x00, 0x00,
}

func (m *Interquery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Interquery) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Interquery) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0x32
	}
	if m.TimeoutHeight != 0 {
		i = encodeVarintInterquery(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Chainid) > 0 {
		i -= len(m.Chainid)
		copy(dAtA[i:], m.Chainid)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Chainid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Storeid) > 0 {
		i -= len(m.Storeid)
		copy(dAtA[i:], m.Storeid)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Storeid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InterqueryResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterqueryResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterqueryResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proved {
		i--
		if m.Proved {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x40
	}
	if m.Success {
		i--
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.LocalHeight != 0 {
		i = encodeVarintInterquery(dAtA, i, uint64(m.LocalHeight))
		i--
		dAtA[i] = 0x30
	}
	if m.Height != nil {
		{
			size, err := m.Height.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintInterquery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Chainid) > 0 {
		i -= len(m.Chainid)
		copy(dAtA[i:], m.Chainid)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Chainid)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Storeid) > 0 {
		i -= len(m.Storeid)
		copy(dAtA[i:], m.Storeid)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Storeid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InterqueryTimeoutResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InterqueryTimeoutResult) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InterqueryTimeoutResult) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TimeoutHeight != 0 {
		i = encodeVarintInterquery(dAtA, i, uint64(m.TimeoutHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Storeid) > 0 {
		i -= len(m.Storeid)
		copy(dAtA[i:], m.Storeid)
		i = encodeVarintInterquery(dAtA, i, uint64(len(m.Storeid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintInterquery(dAtA []byte, offset int, v uint64) int {
	offset -= sovInterquery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Interquery) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Storeid)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Chainid)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovInterquery(uint64(m.TimeoutHeight))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	return n
}

func (m *InterqueryResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Storeid)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Chainid)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	if m.Height != nil {
		l = m.Height.Size()
		n += 1 + l + sovInterquery(uint64(l))
	}
	if m.LocalHeight != 0 {
		n += 1 + sovInterquery(uint64(m.LocalHeight))
	}
	if m.Success {
		n += 2
	}
	if m.Proved {
		n += 2
	}
	return n
}

func (m *InterqueryTimeoutResult) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Storeid)
	if l > 0 {
		n += 1 + l + sovInterquery(uint64(l))
	}
	if m.TimeoutHeight != 0 {
		n += 1 + sovInterquery(uint64(m.TimeoutHeight))
	}
	return n
}

func sovInterquery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInterquery(x uint64) (n int) {
	return sovInterquery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Interquery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterquery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Interquery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Interquery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storeid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storeid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chainid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chainid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInterquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterquery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InterqueryResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterquery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InterqueryResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterqueryResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storeid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storeid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chainid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chainid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Height == nil {
				m.Height = &types.Height{}
			}
			if err := m.Height.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LocalHeight", wireType)
			}
			m.LocalHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LocalHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proved", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Proved = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipInterquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterquery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *InterqueryTimeoutResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInterquery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: InterqueryTimeoutResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InterqueryTimeoutResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storeid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthInterquery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInterquery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storeid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
			}
			m.TimeoutHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipInterquery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInterquery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipInterquery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInterquery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowInterquery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthInterquery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInterquery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInterquery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInterquery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInterquery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInterquery = fmt.Errorf("proto: unexpected end of group")
)