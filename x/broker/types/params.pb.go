// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: broker/params.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type BaseDenoms struct {
	AtomTrace *types.DenomTrace `protobuf:"bytes,1,opt,name=AtomTrace,proto3" json:"AtomTrace,omitempty"`
	OsmoTrace *types.DenomTrace `protobuf:"bytes,2,opt,name=OsmoTrace,proto3" json:"OsmoTrace,omitempty"`
}

func (m *BaseDenoms) Reset()         { *m = BaseDenoms{} }
func (m *BaseDenoms) String() string { return proto.CompactTextString(m) }
func (*BaseDenoms) ProtoMessage()    {}
func (*BaseDenoms) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e45f8c3e5dbb96c, []int{0}
}
func (m *BaseDenoms) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseDenoms) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseDenoms.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseDenoms) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseDenoms.Merge(m, src)
}
func (m *BaseDenoms) XXX_Size() int {
	return m.Size()
}
func (m *BaseDenoms) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseDenoms.DiscardUnknown(m)
}

var xxx_messageInfo_BaseDenoms proto.InternalMessageInfo

func (m *BaseDenoms) GetAtomTrace() *types.DenomTrace {
	if m != nil {
		return m.AtomTrace
	}
	return nil
}

func (m *BaseDenoms) GetOsmoTrace() *types.DenomTrace {
	if m != nil {
		return m.OsmoTrace
	}
	return nil
}

// Params defines the parameters for the broker module.
type Params struct {
	// set the base denoms
	BaseDenoms *BaseDenoms `protobuf:"bytes,1,opt,name=base_denoms,json=baseDenoms,proto3" json:"base_denoms,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e45f8c3e5dbb96c, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetBaseDenoms() *BaseDenoms {
	if m != nil {
		return m.BaseDenoms
	}
	return nil
}

func init() {
	proto.RegisterType((*BaseDenoms)(nil), "defundlabs.defund.broker.BaseDenoms")
	proto.RegisterType((*Params)(nil), "defundlabs.defund.broker.Params")
}

func init() { proto.RegisterFile("broker/params.proto", fileDescriptor_2e45f8c3e5dbb96c) }

var fileDescriptor_2e45f8c3e5dbb96c = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0x7f, 0xa8, 0xf4, 0xbb, 0x5b, 0x60, 0xa8, 0x3a, 0x58, 0xa8, 0x62, 0xa8, 0x04,
	0x5c, 0xab, 0xf0, 0x04, 0x54, 0xd0, 0xb5, 0xa8, 0x62, 0x62, 0x41, 0x76, 0xe2, 0x86, 0x88, 0x3a,
	0xd7, 0xb2, 0xdd, 0x0a, 0xde, 0x82, 0x07, 0xe0, 0x81, 0x18, 0x3b, 0x32, 0xa2, 0xe4, 0x45, 0x50,
	0xec, 0x28, 0x65, 0x61, 0x60, 0x3b, 0xc3, 0x77, 0x3e, 0x5d, 0x1f, 0xd3, 0x23, 0x69, 0xf1, 0x59,
	0x59, 0x6e, 0x84, 0x15, 0xda, 0x81, 0xb1, 0xe8, 0x31, 0x1d, 0xe5, 0x6a, 0xbd, 0xad, 0xf2, 0x8d,
	0x90, 0x0e, 0x62, 0x84, 0x88, 0x8d, 0x8f, 0x0b, 0x2c, 0x30, 0x40, 0xbc, 0x4d, 0x91, 0x1f, 0x9f,
	0x95, 0x32, 0xe3, 0xc2, 0x98, 0x4d, 0x99, 0x09, 0x5f, 0x62, 0xe5, 0xb8, 0xb7, 0xa2, 0x72, 0x6b,
	0x65, 0xf9, 0x6e, 0xd6, 0xe7, 0x08, 0x4f, 0xde, 0x09, 0xa5, 0x73, 0xe1, 0xd4, 0x8d, 0xaa, 0x50,
	0xbb, 0x74, 0x41, 0xff, 0x5f, 0x7b, 0xd4, 0xf7, 0x56, 0x64, 0x6a, 0x44, 0x4e, 0xc8, 0x74, 0x78,
	0x39, 0x85, 0x52, 0x66, 0xf0, 0xd3, 0x07, 0xbd, 0x63, 0x37, 0x83, 0x50, 0x0c, 0xfc, 0xea, 0x50,
	0x6d, 0x3d, 0x4b, 0xa7, 0x31, 0x7a, 0xfe, 0xfd, 0xd5, 0xd3, 0x57, 0x27, 0x4b, 0x3a, 0xb8, 0x0b,
	0x5b, 0xa4, 0xb7, 0x74, 0x28, 0x85, 0x53, 0x8f, 0x79, 0x38, 0xb4, 0xbb, 0xed, 0x14, 0x7e, 0xdb,
	0x06, 0x0e, 0x8f, 0x5a, 0x51, 0xd9, 0xe7, 0xf9, 0xe2, 0xa3, 0x66, 0x64, 0x5f, 0x33, 0xf2, 0x55,
	0x33, 0xf2, 0xd6, 0xb0, 0x64, 0xdf, 0xb0, 0xe4, 0xb3, 0x61, 0xc9, 0xc3, 0x79, 0x51, 0xfa, 0xa7,
	0xad, 0x84, 0x0c, 0x35, 0x8f, 0xaa, 0x8b, 0x56, 0xdb, 0x65, 0xfe, 0xc2, 0xbb, 0xbf, 0xf1, 0xaf,
	0x46, 0x39, 0x39, 0x08, 0xf3, 0x5d, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x38, 0x09, 0x84, 0xf7,
	0xb2, 0x01, 0x00, 0x00,
}

func (m *BaseDenoms) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseDenoms) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseDenoms) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OsmoTrace != nil {
		{
			size, err := m.OsmoTrace.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.AtomTrace != nil {
		{
			size, err := m.AtomTrace.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseDenoms != nil {
		{
			size, err := m.BaseDenoms.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintParams(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseDenoms) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AtomTrace != nil {
		l = m.AtomTrace.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	if m.OsmoTrace != nil {
		l = m.OsmoTrace.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseDenoms != nil {
		l = m.BaseDenoms.Size()
		n += 1 + l + sovParams(uint64(l))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseDenoms) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: BaseDenoms: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseDenoms: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AtomTrace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AtomTrace == nil {
				m.AtomTrace = &types.DenomTrace{}
			}
			if err := m.AtomTrace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OsmoTrace", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.OsmoTrace == nil {
				m.OsmoTrace = &types.DenomTrace{}
			}
			if err := m.OsmoTrace.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenoms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseDenoms == nil {
				m.BaseDenoms = &BaseDenoms{}
			}
			if err := m.BaseDenoms.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
