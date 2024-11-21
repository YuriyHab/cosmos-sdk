// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/accounts/defaults/asset/v1/tx.proto

package v1

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/cosmos/gogoproto/types/any"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// MsgInitAssetAccount defines a message that enables creating a asset account.
type MsgInitAssetAccount struct {
	// owner of the asset account
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	// denom
	Denom    string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Symbol   string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Decimals uint64 `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (m *MsgInitAssetAccount) Reset()         { *m = MsgInitAssetAccount{} }
func (m *MsgInitAssetAccount) String() string { return proto.CompactTextString(m) }
func (*MsgInitAssetAccount) ProtoMessage()    {}
func (*MsgInitAssetAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4834eb883c54202, []int{0}
}
func (m *MsgInitAssetAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitAssetAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitAssetAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitAssetAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitAssetAccount.Merge(m, src)
}
func (m *MsgInitAssetAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitAssetAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitAssetAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitAssetAccount proto.InternalMessageInfo

func (m *MsgInitAssetAccount) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *MsgInitAssetAccount) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *MsgInitAssetAccount) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MsgInitAssetAccount) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

// MsgInitAssetAccountResponse defines the Msg/InitLockupAccount response type.
type MsgInitAssetAccountResponse struct {
}

func (m *MsgInitAssetAccountResponse) Reset()         { *m = MsgInitAssetAccountResponse{} }
func (m *MsgInitAssetAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgInitAssetAccountResponse) ProtoMessage()    {}
func (*MsgInitAssetAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4834eb883c54202, []int{1}
}
func (m *MsgInitAssetAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInitAssetAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInitAssetAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInitAssetAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInitAssetAccountResponse.Merge(m, src)
}
func (m *MsgInitAssetAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgInitAssetAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInitAssetAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInitAssetAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgInitAssetAccount)(nil), "cosmos.accounts.defaults.asset.v1.MsgInitAssetAccount")
	proto.RegisterType((*MsgInitAssetAccountResponse)(nil), "cosmos.accounts.defaults.asset.v1.MsgInitAssetAccountResponse")
}

func init() {
	proto.RegisterFile("cosmos/accounts/defaults/asset/v1/tx.proto", fileDescriptor_b4834eb883c54202)
}

var fileDescriptor_b4834eb883c54202 = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x6e, 0xdb, 0x30,
	0x14, 0xc6, 0xcd, 0xd6, 0x36, 0x5a, 0x6d, 0x55, 0x8d, 0x56, 0x75, 0x51, 0xd9, 0xf5, 0x52, 0xd7,
	0x80, 0x45, 0x08, 0xd9, 0xb2, 0xd9, 0x99, 0x32, 0x64, 0x51, 0xb6, 0x2c, 0x01, 0x25, 0xd1, 0x84,
	0x60, 0x91, 0xcf, 0xd0, 0xa3, 0x1d, 0xfb, 0x0a, 0x99, 0x72, 0x84, 0x1c, 0x21, 0x43, 0x86, 0x1c,
	0x21, 0xa3, 0x91, 0x29, 0x63, 0x60, 0x0f, 0xc9, 0x31, 0x02, 0x89, 0x74, 0x86, 0xc0, 0xc8, 0x42,
	0xf0, 0x7b, 0xdf, 0xef, 0xf1, 0xfd, 0x01, 0x9d, 0x41, 0x02, 0x28, 0x01, 0x29, 0x4b, 0x12, 0x98,
	0x2b, 0x8d, 0x34, 0xe5, 0x13, 0x36, 0xcf, 0x35, 0x52, 0x86, 0xc8, 0x35, 0x5d, 0x84, 0x54, 0x2f,
	0x83, 0x59, 0x01, 0x1a, 0xdc, 0xbf, 0x86, 0x0d, 0x76, 0x6c, 0xb0, 0x63, 0x83, 0x8a, 0x0d, 0x16,
	0x61, 0xfb, 0x1b, 0x93, 0x99, 0x02, 0x5a, 0x9d, 0x26, 0xab, 0xed, 0xdb, 0x0a, 0x31, 0x43, 0x4e,
	0x17, 0x61, 0xcc, 0x35, 0x0b, 0x69, 0x02, 0x99, 0xb2, 0xfe, 0x4f, 0xeb, 0x4b, 0x14, 0x65, 0x35,
	0x89, 0xc2, 0x1a, 0xbf, 0x8c, 0x71, 0x5e, 0x29, 0x6a, 0x6b, 0x1b, 0xab, 0x25, 0x40, 0x80, 0x89,
	0x97, 0xb7, 0x5d, 0x82, 0x00, 0x10, 0x39, 0xa7, 0x95, 0x8a, 0xe7, 0x13, 0xca, 0xd4, 0xca, 0x5a,
	0x9d, 0xf7, 0x96, 0xce, 0x24, 0x47, 0xcd, 0xe4, 0xcc, 0x00, 0xbd, 0x3b, 0xe2, 0x7c, 0x3f, 0x41,
	0x71, 0xac, 0x32, 0x3d, 0x2a, 0x87, 0x19, 0x99, 0x19, 0xdd, 0xc0, 0x69, 0xc0, 0x85, 0xe2, 0x85,
	0x47, 0xba, 0xa4, 0xff, 0x75, 0xec, 0x3d, 0xdc, 0x0e, 0x5b, 0xb6, 0x95, 0x51, 0x9a, 0x16, 0x1c,
	0xf1, 0x54, 0x17, 0x99, 0x12, 0x91, 0xc1, 0xdc, 0x96, 0xd3, 0x48, 0xb9, 0x02, 0xe9, 0x7d, 0x2a,
	0xf9, 0xc8, 0x08, 0xf7, 0x87, 0xd3, 0xc4, 0x95, 0x8c, 0x21, 0xf7, 0x3e, 0x57, 0x61, 0xab, 0xdc,
	0xb6, 0xf3, 0x25, 0xe5, 0x49, 0x26, 0x59, 0x8e, 0x5e, 0xbd, 0x4b, 0xfa, 0xf5, 0xe8, 0x4d, 0x1f,
	0xfe, 0x7b, 0xb9, 0xee, 0x90, 0xcb, 0xe7, 0x9b, 0x81, 0x5d, 0xe0, 0x10, 0xd3, 0x29, 0xdd, 0xd3,
	0x62, 0xef, 0x8f, 0xf3, 0x7b, 0x4f, 0x38, 0xe2, 0x38, 0x03, 0x85, 0x7c, 0x7c, 0x74, 0xbf, 0xf1,
	0xc9, 0x7a, 0xe3, 0x93, 0xa7, 0x8d, 0x4f, 0xae, 0xb6, 0x7e, 0x6d, 0xbd, 0xf5, 0x6b, 0x8f, 0x5b,
	0xbf, 0x76, 0xf6, 0xdf, 0x3c, 0x8c, 0xe9, 0x34, 0xc8, 0x80, 0x2e, 0x3f, 0xf8, 0x03, 0x71, 0xb3,
	0xda, 0xd2, 0xc1, 0x6b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x62, 0x81, 0x2f, 0x2f, 0x02, 0x00,
	0x00,
}

func (this *MsgInitAssetAccount) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgInitAssetAccount)
	if !ok {
		that2, ok := that.(MsgInitAssetAccount)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if this.Symbol != that1.Symbol {
		return false
	}
	if this.Decimals != that1.Decimals {
		return false
	}
	return true
}
func (m *MsgInitAssetAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitAssetAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitAssetAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimals != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgInitAssetAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInitAssetAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInitAssetAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInitAssetAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovTx(uint64(m.Decimals))
	}
	return n
}

func (m *MsgInitAssetAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInitAssetAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInitAssetAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitAssetAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgInitAssetAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgInitAssetAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInitAssetAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
