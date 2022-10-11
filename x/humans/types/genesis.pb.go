// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: humans/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the humans module's genesis state.
type GenesisState struct {
	Params              Params            `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	FeeBalanceList      []FeeBalance      `protobuf:"bytes,2,rep,name=feeBalanceList,proto3" json:"feeBalanceList"`
	KeysignVoteDataList []KeysignVoteData `protobuf:"bytes,3,rep,name=keysignVoteDataList,proto3" json:"keysignVoteDataList"`
	ObserveVoteList     []ObserveVote     `protobuf:"bytes,4,rep,name=observeVoteList,proto3" json:"observeVoteList"`
	PoolBalanceList     []PoolBalance     `protobuf:"bytes,5,rep,name=poolBalanceList,proto3" json:"poolBalanceList"`
	SuperadminList      []Superadmin      `protobuf:"bytes,6,rep,name=superadminList,proto3" json:"superadminList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a4ac69d020c093, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetFeeBalanceList() []FeeBalance {
	if m != nil {
		return m.FeeBalanceList
	}
	return nil
}

func (m *GenesisState) GetKeysignVoteDataList() []KeysignVoteData {
	if m != nil {
		return m.KeysignVoteDataList
	}
	return nil
}

func (m *GenesisState) GetObserveVoteList() []ObserveVote {
	if m != nil {
		return m.ObserveVoteList
	}
	return nil
}

func (m *GenesisState) GetPoolBalanceList() []PoolBalance {
	if m != nil {
		return m.PoolBalanceList
	}
	return nil
}

func (m *GenesisState) GetSuperadminList() []Superadmin {
	if m != nil {
		return m.SuperadminList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "humansdotai.humans.humans.GenesisState")
}

func init() { proto.RegisterFile("humans/genesis.proto", fileDescriptor_f1a4ac69d020c093) }

var fileDescriptor_f1a4ac69d020c093 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x6a, 0xf2, 0x40,
	0x14, 0xc5, 0x93, 0x4f, 0x3f, 0x17, 0xb1, 0xb4, 0x10, 0x85, 0x5a, 0x17, 0x53, 0x5b, 0x68, 0x91,
	0x52, 0x12, 0xb0, 0x0f, 0x50, 0x90, 0x62, 0x17, 0x2d, 0x54, 0x2a, 0xb8, 0xe8, 0x46, 0x26, 0x7a,
	0x8d, 0x43, 0x4d, 0x26, 0x64, 0x46, 0xa9, 0x0f, 0x51, 0xe8, 0x63, 0xb9, 0x74, 0xd9, 0x55, 0x29,
	0xfa, 0x22, 0x25, 0x33, 0xd7, 0x7f, 0xa1, 0xd5, 0x55, 0x2e, 0xf7, 0xcc, 0xf9, 0xdd, 0xcc, 0xb9,
	0x63, 0x15, 0x07, 0xa3, 0x80, 0x86, 0xc2, 0xf5, 0x21, 0x04, 0xc1, 0x84, 0x13, 0xc5, 0x5c, 0x72,
	0xfb, 0x44, 0x77, 0x7b, 0x5c, 0x52, 0xe6, 0xe8, 0x1a, 0x3f, 0xe5, 0xa2, 0xcf, 0x7d, 0xae, 0x4e,
	0xb9, 0x49, 0xa5, 0x0d, 0xe5, 0x02, 0x62, 0x22, 0x1a, 0xd3, 0x00, 0x29, 0xe5, 0x12, 0x36, 0xfb,
	0x00, 0x1d, 0x8f, 0x0e, 0x69, 0xd8, 0x05, 0x54, 0x08, 0x2a, 0xaf, 0x30, 0x11, 0xcc, 0x0f, 0x3b,
	0x63, 0x2e, 0xa1, 0xd3, 0xa3, 0x92, 0xa2, 0x8e, 0xf3, 0x5d, 0xee, 0x09, 0x88, 0xc7, 0xa0, 0xf4,
	0x94, 0x14, 0x71, 0x3e, 0x4c, 0x51, 0x8f, 0x51, 0x12, 0xa3, 0x08, 0x62, 0xda, 0x0b, 0x58, 0xa8,
	0x85, 0xf3, 0xf7, 0xac, 0x75, 0x70, 0xaf, 0x2f, 0xd8, 0x92, 0x54, 0x82, 0x7d, 0x6b, 0xe5, 0xf4,
	0x9f, 0x96, 0xcc, 0x8a, 0x59, 0xcd, 0xd7, 0xce, 0x9c, 0x3f, 0x2f, 0xec, 0x34, 0xd5, 0xc1, 0x7a,
	0x76, 0xfa, 0x75, 0x6a, 0x3c, 0xa3, 0xcd, 0x6e, 0x59, 0x87, 0x7d, 0x80, 0xba, 0x1e, 0xff, 0xc8,
	0x84, 0x2c, 0xfd, 0xab, 0x64, 0xaa, 0xf9, 0xda, 0xc5, 0x0e, 0x50, 0x63, 0x65, 0x40, 0x58, 0x0a,
	0x61, 0x7b, 0x56, 0x01, 0x03, 0x69, 0x73, 0x09, 0x77, 0x54, 0x52, 0x45, 0xce, 0x28, 0xf2, 0xd5,
	0x0e, 0xf2, 0xc3, 0xb6, 0x0b, 0xf1, 0xbf, 0xc1, 0xec, 0xb6, 0x75, 0x84, 0xa1, 0x26, 0x6d, 0xc5,
	0xcf, 0x2a, 0xfe, 0xe5, 0x0e, 0xfe, 0xd3, 0xda, 0x81, 0xec, 0x34, 0x24, 0xe1, 0x26, 0x1b, 0xd9,
	0x4c, 0xe4, 0xff, 0x5e, 0x6e, 0x73, 0xed, 0x58, 0x72, 0x53, 0x90, 0x24, 0xe8, 0xf5, 0x3a, 0x15,
	0x36, 0xb7, 0x37, 0xe8, 0xd6, 0xca, 0xb0, 0x0c, 0x7a, 0x1b, 0x51, 0x6f, 0x4c, 0xe7, 0xc4, 0x9c,
	0xcd, 0x89, 0xf9, 0x3d, 0x27, 0xe6, 0xc7, 0x82, 0x18, 0xb3, 0x05, 0x31, 0x3e, 0x17, 0xc4, 0x78,
	0xb9, 0xf6, 0x99, 0x1c, 0x8c, 0x3c, 0xa7, 0xcb, 0x03, 0x77, 0x63, 0x00, 0xd6, 0xee, 0xdb, 0xb2,
	0x90, 0x93, 0x08, 0x84, 0x97, 0x53, 0xcf, 0xeb, 0xe6, 0x27, 0x00, 0x00, 0xff, 0xff, 0x10, 0x03,
	0x38, 0x35, 0x45, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SuperadminList) > 0 {
		for iNdEx := len(m.SuperadminList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SuperadminList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.PoolBalanceList) > 0 {
		for iNdEx := len(m.PoolBalanceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolBalanceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.ObserveVoteList) > 0 {
		for iNdEx := len(m.ObserveVoteList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ObserveVoteList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.KeysignVoteDataList) > 0 {
		for iNdEx := len(m.KeysignVoteDataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.KeysignVoteDataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.FeeBalanceList) > 0 {
		for iNdEx := len(m.FeeBalanceList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FeeBalanceList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.FeeBalanceList) > 0 {
		for _, e := range m.FeeBalanceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.KeysignVoteDataList) > 0 {
		for _, e := range m.KeysignVoteDataList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ObserveVoteList) > 0 {
		for _, e := range m.ObserveVoteList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PoolBalanceList) > 0 {
		for _, e := range m.PoolBalanceList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SuperadminList) > 0 {
		for _, e := range m.SuperadminList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeBalanceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeBalanceList = append(m.FeeBalanceList, FeeBalance{})
			if err := m.FeeBalanceList[len(m.FeeBalanceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeysignVoteDataList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeysignVoteDataList = append(m.KeysignVoteDataList, KeysignVoteData{})
			if err := m.KeysignVoteDataList[len(m.KeysignVoteDataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObserveVoteList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ObserveVoteList = append(m.ObserveVoteList, ObserveVote{})
			if err := m.ObserveVoteList[len(m.ObserveVoteList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolBalanceList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PoolBalanceList = append(m.PoolBalanceList, PoolBalance{})
			if err := m.PoolBalanceList[len(m.PoolBalanceList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuperadminList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SuperadminList = append(m.SuperadminList, Superadmin{})
			if err := m.SuperadminList[len(m.SuperadminList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
