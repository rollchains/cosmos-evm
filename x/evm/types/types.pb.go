// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: seiprotocol/seichain/evm/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type Whitelist struct {
	Hashes []string `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty" yaml:"hashes"`
}

func (m *Whitelist) Reset()         { *m = Whitelist{} }
func (m *Whitelist) String() string { return proto.CompactTextString(m) }
func (*Whitelist) ProtoMessage()    {}
func (*Whitelist) Descriptor() ([]byte, []int) {
	return fileDescriptor_c80e42d822facef3, []int{0}
}
func (m *Whitelist) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Whitelist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Whitelist.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Whitelist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Whitelist.Merge(m, src)
}
func (m *Whitelist) XXX_Size() int {
	return m.Size()
}
func (m *Whitelist) XXX_DiscardUnknown() {
	xxx_messageInfo_Whitelist.DiscardUnknown(m)
}

var xxx_messageInfo_Whitelist proto.InternalMessageInfo

func (m *Whitelist) GetHashes() []string {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func init() {
	proto.RegisterType((*Whitelist)(nil), "seiprotocol.seichain.evm.Whitelist")
}

func init() {
	proto.RegisterFile("seiprotocol/seichain/evm/types.proto", fileDescriptor_c80e42d822facef3)
}

var fileDescriptor_c80e42d822facef3 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x4e, 0xcd, 0x2c,
	0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4e, 0xcd, 0x4c, 0xce, 0x48, 0xcc, 0xcc,
	0xd3, 0x4f, 0x2d, 0xcb, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x03, 0xcb, 0x09, 0x49, 0x20,
	0xa9, 0xd2, 0x83, 0xa9, 0xd2, 0x4b, 0x2d, 0xcb, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x4b,
	0xe9, 0x83, 0x58, 0x10, 0xf5, 0x4a, 0x66, 0x5c, 0x9c, 0xe1, 0x19, 0x99, 0x25, 0xa9, 0x39, 0x99,
	0xc5, 0x25, 0x42, 0x9a, 0x5c, 0x6c, 0x19, 0x89, 0xc5, 0x19, 0xa9, 0xc5, 0x12, 0x8c, 0x0a, 0xcc,
	0x1a, 0x9c, 0x4e, 0x82, 0x9f, 0xee, 0xc9, 0xf3, 0x56, 0x26, 0xe6, 0xe6, 0x58, 0x29, 0x41, 0xc4,
	0x95, 0x82, 0xa0, 0x0a, 0x9c, 0xdc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1,
	0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21,
	0x4a, 0x37, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x17, 0xe4, 0x4c, 0x5d, 0x64,
	0x37, 0xeb, 0x42, 0x1c, 0x5d, 0x81, 0x70, 0x76, 0x12, 0x1b, 0x58, 0xde, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x7a, 0x6a, 0xf6, 0x42, 0xdf, 0x00, 0x00, 0x00,
}

func (m *Whitelist) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Whitelist) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Whitelist) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for iNdEx := len(m.Hashes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Hashes[iNdEx])
			copy(dAtA[i:], m.Hashes[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.Hashes[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Whitelist) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Hashes) > 0 {
		for _, s := range m.Hashes {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Whitelist) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Whitelist: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Whitelist: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hashes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hashes = append(m.Hashes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
