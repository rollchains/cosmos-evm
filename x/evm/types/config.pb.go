// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: seiprotocol/seichain/evm/config.proto

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

// XXTime fields indicate upgrade timestamps. For example, a ShanghaiTime
// of 42198537129 means the chain upgraded to the Shanghai version at timestamp 42198537129.
// A value of 0 means the upgrade is included in the genesis of the EVM on Sei.
// -1 means upgrade not reached yet.
type ChainConfig struct {
	CancunTime int64 `protobuf:"varint,1,opt,name=cancun_time,json=cancunTime,proto3" json:"cancun_time,omitempty" yaml:"cancun_time"`
	PragueTime int64 `protobuf:"varint,2,opt,name=prague_time,json=pragueTime,proto3" json:"prague_time,omitempty" yaml:"prague_time"`
	VerkleTime int64 `protobuf:"varint,3,opt,name=verkle_time,json=verkleTime,proto3" json:"verkle_time,omitempty" yaml:"verkle_time"`
}

func (m *ChainConfig) Reset()         { *m = ChainConfig{} }
func (m *ChainConfig) String() string { return proto.CompactTextString(m) }
func (*ChainConfig) ProtoMessage()    {}
func (*ChainConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_88ac6789f3577184, []int{0}
}
func (m *ChainConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfig.Merge(m, src)
}
func (m *ChainConfig) XXX_Size() int {
	return m.Size()
}
func (m *ChainConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfig proto.InternalMessageInfo

func (m *ChainConfig) GetCancunTime() int64 {
	if m != nil {
		return m.CancunTime
	}
	return 0
}

func (m *ChainConfig) GetPragueTime() int64 {
	if m != nil {
		return m.PragueTime
	}
	return 0
}

func (m *ChainConfig) GetVerkleTime() int64 {
	if m != nil {
		return m.VerkleTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ChainConfig)(nil), "seiprotocol.seichain.evm.ChainConfig")
}

func init() {
	proto.RegisterFile("seiprotocol/seichain/evm/config.proto", fileDescriptor_88ac6789f3577184)
}

var fileDescriptor_88ac6789f3577184 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x4e, 0xcd, 0x2c,
	0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x4e, 0xcd, 0x4c, 0xce, 0x48, 0xcc, 0xcc,
	0xd3, 0x4f, 0x2d, 0xcb, 0xd5, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x03, 0x4b, 0x0a, 0x49,
	0x20, 0x29, 0xd3, 0x83, 0x29, 0xd3, 0x4b, 0x2d, 0xcb, 0x95, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07,
	0x4b, 0xe9, 0x83, 0x58, 0x10, 0xf5, 0x4a, 0x3b, 0x18, 0xb9, 0xb8, 0x9d, 0x41, 0x6a, 0x9c, 0xc1,
	0xa6, 0x08, 0x99, 0x73, 0x71, 0x27, 0x27, 0xe6, 0x25, 0x97, 0xe6, 0xc5, 0x97, 0x64, 0xe6, 0xa6,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x3b, 0x89, 0x7d, 0xba, 0x27, 0x2f, 0x54, 0x99, 0x98, 0x9b,
	0x63, 0xa5, 0x84, 0x24, 0xa9, 0x14, 0xc4, 0x05, 0xe1, 0x85, 0x64, 0xe6, 0xa6, 0x82, 0x34, 0x16,
	0x14, 0x25, 0xa6, 0x97, 0xa6, 0x42, 0x34, 0x32, 0xa1, 0x6b, 0x44, 0x92, 0x54, 0x0a, 0xe2, 0x82,
	0xf0, 0x60, 0x1a, 0xcb, 0x52, 0x8b, 0xb2, 0x73, 0xa0, 0x1a, 0x99, 0xd1, 0x35, 0x22, 0x49, 0x2a,
	0x05, 0x71, 0x41, 0x78, 0x20, 0x8d, 0x4e, 0xee, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c,
	0xc7, 0x10, 0xa5, 0x9b, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0x0b, 0x0a, 0x2a,
	0x5d, 0xe4, 0x70, 0xd3, 0x85, 0x04, 0x5c, 0x05, 0x38, 0xe8, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93,
	0xd8, 0xc0, 0xf2, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x6c, 0x06, 0xfd, 0x63, 0x01,
	0x00, 0x00,
}

func (m *ChainConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VerkleTime != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.VerkleTime))
		i--
		dAtA[i] = 0x18
	}
	if m.PragueTime != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.PragueTime))
		i--
		dAtA[i] = 0x10
	}
	if m.CancunTime != 0 {
		i = encodeVarintConfig(dAtA, i, uint64(m.CancunTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	offset -= sovConfig(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChainConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CancunTime != 0 {
		n += 1 + sovConfig(uint64(m.CancunTime))
	}
	if m.PragueTime != 0 {
		n += 1 + sovConfig(uint64(m.PragueTime))
	}
	if m.VerkleTime != 0 {
		n += 1 + sovConfig(uint64(m.VerkleTime))
	}
	return n
}

func sovConfig(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: ChainConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CancunTime", wireType)
			}
			m.CancunTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CancunTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PragueTime", wireType)
			}
			m.PragueTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PragueTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerkleTime", wireType)
			}
			m.VerkleTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VerkleTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConfig
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
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConfig
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConfig
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConfig        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConfig = fmt.Errorf("proto: unexpected end of group")
)
