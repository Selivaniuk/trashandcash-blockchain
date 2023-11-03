// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trashandcashblockchain/trashandcashblockchain/spin.proto

package types

import (
	fmt "fmt"
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

type Spin struct {
	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bet     string `protobuf:"bytes,2,opt,name=bet,proto3" json:"bet,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Result  string `protobuf:"bytes,4,opt,name=result,proto3" json:"result,omitempty"`
	Winning string `protobuf:"bytes,5,opt,name=winning,proto3" json:"winning,omitempty"`
}

func (m *Spin) Reset()         { *m = Spin{} }
func (m *Spin) String() string { return proto.CompactTextString(m) }
func (*Spin) ProtoMessage()    {}
func (*Spin) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f45174c6b2df570, []int{0}
}
func (m *Spin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Spin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Spin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Spin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Spin.Merge(m, src)
}
func (m *Spin) XXX_Size() int {
	return m.Size()
}
func (m *Spin) XXX_DiscardUnknown() {
	xxx_messageInfo_Spin.DiscardUnknown(m)
}

var xxx_messageInfo_Spin proto.InternalMessageInfo

func (m *Spin) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Spin) GetBet() string {
	if m != nil {
		return m.Bet
	}
	return ""
}

func (m *Spin) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Spin) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func (m *Spin) GetWinning() string {
	if m != nil {
		return m.Winning
	}
	return ""
}

func init() {
	proto.RegisterType((*Spin)(nil), "trashandcashblockchain.trashandcashblockchain.Spin")
}

func init() {
	proto.RegisterFile("trashandcashblockchain/trashandcashblockchain/spin.proto", fileDescriptor_5f45174c6b2df570)
}

var fileDescriptor_5f45174c6b2df570 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x28, 0x29, 0x4a, 0x2c,
	0xce, 0x48, 0xcc, 0x4b, 0x49, 0x4e, 0x2c, 0xce, 0x48, 0xca, 0xc9, 0x4f, 0xce, 0x4e, 0xce, 0x48,
	0xcc, 0xcc, 0xd3, 0xc7, 0x21, 0x5c, 0x5c, 0x90, 0x99, 0xa7, 0x57, 0x50, 0x94, 0x5f, 0x92, 0x2f,
	0xa4, 0x8b, 0x5d, 0x89, 0x1e, 0x76, 0x61, 0xa5, 0x12, 0x2e, 0x96, 0xe0, 0x82, 0xcc, 0x3c, 0x21,
	0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0xa6, 0xcc, 0x14, 0x21,
	0x01, 0x2e, 0xe6, 0xa4, 0xd4, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x10, 0x53, 0x48,
	0x82, 0x8b, 0x3d, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x19, 0x2c, 0x0a, 0xe3, 0x0a,
	0x89, 0x71, 0xb1, 0x15, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x48, 0xb0, 0x80, 0x25, 0xa0, 0x3c, 0x90,
	0x8e, 0xf2, 0xcc, 0xbc, 0xbc, 0xcc, 0xbc, 0x74, 0x09, 0x56, 0x88, 0x0e, 0x28, 0xd7, 0x29, 0xe0,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xcc, 0x90, 0xdd, 0xa9, 0x8b, 0xe4,
	0xc5, 0x0a, 0x5c, 0x7e, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xde, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0x57, 0x26, 0x09, 0x90, 0x39, 0x01, 0x00, 0x00,
}

func (m *Spin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Spin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Spin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Winning) > 0 {
		i -= len(m.Winning)
		copy(dAtA[i:], m.Winning)
		i = encodeVarintSpin(dAtA, i, uint64(len(m.Winning)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintSpin(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintSpin(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Bet) > 0 {
		i -= len(m.Bet)
		copy(dAtA[i:], m.Bet)
		i = encodeVarintSpin(dAtA, i, uint64(len(m.Bet)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintSpin(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSpin(dAtA []byte, offset int, v uint64) int {
	offset -= sovSpin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Spin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovSpin(uint64(m.Id))
	}
	l = len(m.Bet)
	if l > 0 {
		n += 1 + l + sovSpin(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovSpin(uint64(l))
	}
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovSpin(uint64(l))
	}
	l = len(m.Winning)
	if l > 0 {
		n += 1 + l + sovSpin(uint64(l))
	}
	return n
}

func sovSpin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSpin(x uint64) (n int) {
	return sovSpin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Spin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSpin
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
			return fmt.Errorf("proto: Spin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Spin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpin
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
				return ErrInvalidLengthSpin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bet = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpin
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
				return ErrInvalidLengthSpin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpin
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
				return ErrInvalidLengthSpin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winning", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSpin
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
				return ErrInvalidLengthSpin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSpin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winning = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSpin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSpin
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
func skipSpin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSpin
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
					return 0, ErrIntOverflowSpin
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
					return 0, ErrIntOverflowSpin
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
				return 0, ErrInvalidLengthSpin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSpin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSpin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSpin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSpin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSpin = fmt.Errorf("proto: unexpected end of group")
)
