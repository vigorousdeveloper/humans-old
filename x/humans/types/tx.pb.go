// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: humans/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MsgRequestTransaction struct {
	Creator       string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	OriginChain   string `protobuf:"bytes,2,opt,name=originChain,proto3" json:"originChain,omitempty"`
	OriginAddress string `protobuf:"bytes,3,opt,name=originAddress,proto3" json:"originAddress,omitempty"`
	TargetChain   string `protobuf:"bytes,4,opt,name=targetChain,proto3" json:"targetChain,omitempty"`
	TargetAddress string `protobuf:"bytes,5,opt,name=targetAddress,proto3" json:"targetAddress,omitempty"`
	Amount        string `protobuf:"bytes,6,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee           string `protobuf:"bytes,7,opt,name=fee,proto3" json:"fee,omitempty"`
}

func (m *MsgRequestTransaction) Reset()         { *m = MsgRequestTransaction{} }
func (m *MsgRequestTransaction) String() string { return proto.CompactTextString(m) }
func (*MsgRequestTransaction) ProtoMessage()    {}
func (*MsgRequestTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e886a756dbee256, []int{0}
}
func (m *MsgRequestTransaction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestTransaction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestTransaction.Merge(m, src)
}
func (m *MsgRequestTransaction) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestTransaction proto.InternalMessageInfo

func (m *MsgRequestTransaction) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRequestTransaction) GetOriginChain() string {
	if m != nil {
		return m.OriginChain
	}
	return ""
}

func (m *MsgRequestTransaction) GetOriginAddress() string {
	if m != nil {
		return m.OriginAddress
	}
	return ""
}

func (m *MsgRequestTransaction) GetTargetChain() string {
	if m != nil {
		return m.TargetChain
	}
	return ""
}

func (m *MsgRequestTransaction) GetTargetAddress() string {
	if m != nil {
		return m.TargetAddress
	}
	return ""
}

func (m *MsgRequestTransaction) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *MsgRequestTransaction) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

type MsgRequestTransactionResponse struct {
}

func (m *MsgRequestTransactionResponse) Reset()         { *m = MsgRequestTransactionResponse{} }
func (m *MsgRequestTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestTransactionResponse) ProtoMessage()    {}
func (*MsgRequestTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e886a756dbee256, []int{1}
}
func (m *MsgRequestTransactionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestTransactionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestTransactionResponse.Merge(m, src)
}
func (m *MsgRequestTransactionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestTransactionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRequestTransaction)(nil), "humansdotai.humans.humans.MsgRequestTransaction")
	proto.RegisterType((*MsgRequestTransactionResponse)(nil), "humansdotai.humans.humans.MsgRequestTransactionResponse")
}

func init() { proto.RegisterFile("humans/tx.proto", fileDescriptor_8e886a756dbee256) }

var fileDescriptor_8e886a756dbee256 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x31, 0x4b, 0xc4, 0x30,
	0x18, 0x86, 0x2f, 0x9e, 0xf6, 0x30, 0x22, 0x4a, 0x40, 0x89, 0x82, 0xf1, 0x38, 0x1c, 0x1c, 0xa4,
	0x15, 0x5d, 0x5c, 0x55, 0x70, 0xbb, 0xa5, 0x38, 0xb9, 0xe5, 0xda, 0xd8, 0x66, 0x68, 0x52, 0x93,
	0xaf, 0x70, 0x6e, 0x8e, 0x8e, 0xfe, 0x2c, 0xc7, 0x1b, 0x1d, 0xa5, 0x1d, 0xfd, 0x13, 0xd2, 0xa6,
	0x85, 0x1e, 0xd4, 0xc1, 0xa9, 0xef, 0xfb, 0x7c, 0xed, 0x43, 0xf3, 0x11, 0xbc, 0x97, 0x16, 0x19,
	0x57, 0x36, 0x80, 0xa5, 0x9f, 0x1b, 0x0d, 0x9a, 0x1c, 0x39, 0x10, 0x6b, 0xe0, 0xd2, 0x77, 0xb9,
	0x7d, 0xcc, 0x7e, 0x10, 0x3e, 0x98, 0xdb, 0x24, 0x14, 0x2f, 0x85, 0xb0, 0xf0, 0x68, 0xb8, 0xb2,
	0x3c, 0x02, 0xa9, 0x15, 0xa1, 0x78, 0x12, 0x19, 0xc1, 0x41, 0x1b, 0x8a, 0xa6, 0xe8, 0x7c, 0x3b,
	0xec, 0x2a, 0x99, 0xe2, 0x1d, 0x6d, 0x64, 0x22, 0xd5, 0x7d, 0xca, 0xa5, 0xa2, 0x1b, 0xcd, 0xb4,
	0x8f, 0xc8, 0x19, 0xde, 0x75, 0xf5, 0x36, 0x8e, 0x8d, 0xb0, 0x96, 0x8e, 0x9b, 0x77, 0xd6, 0x61,
	0xed, 0x01, 0x6e, 0x12, 0x01, 0xce, 0xb3, 0xe9, 0x3c, 0x3d, 0x54, 0x7b, 0x5c, 0xed, 0x3c, 0x5b,
	0xce, 0xb3, 0x06, 0xc9, 0x21, 0xf6, 0x78, 0xa6, 0x0b, 0x05, 0xd4, 0x6b, 0xc6, 0x6d, 0x23, 0xfb,
	0x78, 0xfc, 0x2c, 0x04, 0x9d, 0x34, 0xb0, 0x8e, 0xb3, 0x53, 0x7c, 0x32, 0x78, 0xd8, 0x50, 0xd8,
	0x5c, 0x2b, 0x2b, 0xae, 0xde, 0x11, 0x1e, 0xcf, 0x6d, 0x42, 0xde, 0x10, 0x26, 0x03, 0x3b, 0xb9,
	0xf4, 0xff, 0xdc, 0xa4, 0x3f, 0x28, 0x3e, 0xbe, 0xf9, 0xef, 0x17, 0xdd, 0xaf, 0xdc, 0x3d, 0x7c,
	0x96, 0x0c, 0xad, 0x4a, 0x86, 0xbe, 0x4b, 0x86, 0x3e, 0x2a, 0x36, 0x5a, 0x55, 0x6c, 0xf4, 0x55,
	0xb1, 0xd1, 0xd3, 0x45, 0x22, 0x21, 0x2d, 0x16, 0x7e, 0xa4, 0xb3, 0xa0, 0x67, 0x6f, 0x73, 0xb0,
	0xec, 0x02, 0xbc, 0xe6, 0xc2, 0x2e, 0xbc, 0xe6, 0x0e, 0x5c, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xea, 0x62, 0xd4, 0x3f, 0x16, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RequestTransaction(ctx context.Context, in *MsgRequestTransaction, opts ...grpc.CallOption) (*MsgRequestTransactionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RequestTransaction(ctx context.Context, in *MsgRequestTransaction, opts ...grpc.CallOption) (*MsgRequestTransactionResponse, error) {
	out := new(MsgRequestTransactionResponse)
	err := c.cc.Invoke(ctx, "/humansdotai.humans.humans.Msg/RequestTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RequestTransaction(context.Context, *MsgRequestTransaction) (*MsgRequestTransactionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RequestTransaction(ctx context.Context, req *MsgRequestTransaction) (*MsgRequestTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestTransaction not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RequestTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/humansdotai.humans.humans.Msg/RequestTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestTransaction(ctx, req.(*MsgRequestTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "humansdotai.humans.humans.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestTransaction",
			Handler:    _Msg_RequestTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "humans/tx.proto",
}

func (m *MsgRequestTransaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestTransaction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestTransaction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TargetAddress) > 0 {
		i -= len(m.TargetAddress)
		copy(dAtA[i:], m.TargetAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TargetAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.TargetChain) > 0 {
		i -= len(m.TargetChain)
		copy(dAtA[i:], m.TargetChain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.TargetChain)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OriginAddress) > 0 {
		i -= len(m.OriginAddress)
		copy(dAtA[i:], m.OriginAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OriginAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.OriginChain) > 0 {
		i -= len(m.OriginChain)
		copy(dAtA[i:], m.OriginChain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OriginChain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestTransactionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestTransactionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestTransactionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgRequestTransaction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.OriginChain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.OriginAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TargetChain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.TargetAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRequestTransactionResponse) Size() (n int) {
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
func (m *MsgRequestTransaction) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRequestTransaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestTransaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginChain", wireType)
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
			m.OriginChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginAddress", wireType)
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
			m.OriginAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetChain", wireType)
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
			m.TargetChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetAddress", wireType)
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
			m.TargetAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
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
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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
func (m *MsgRequestTransactionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRequestTransactionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestTransactionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
