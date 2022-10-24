// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/intertx/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

// MsgRegisterAccount defines the payload for Msg/RegisterAccount
type MsgRegisterAccount struct {
	Owner        string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	ConnectionId string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty" yaml:"connection_id"`
	Version      string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (m *MsgRegisterAccount) Reset()         { *m = MsgRegisterAccount{} }
func (m *MsgRegisterAccount) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterAccount) ProtoMessage()    {}
func (*MsgRegisterAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f79e1c9e86a55bd, []int{0}
}
func (m *MsgRegisterAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterAccount.Merge(m, src)
}
func (m *MsgRegisterAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterAccount proto.InternalMessageInfo

// MsgRegisterAccountResponse defines the response for Msg/RegisterAccount
type MsgRegisterAccountResponse struct {
}

func (m *MsgRegisterAccountResponse) Reset()         { *m = MsgRegisterAccountResponse{} }
func (m *MsgRegisterAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRegisterAccountResponse) ProtoMessage()    {}
func (*MsgRegisterAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f79e1c9e86a55bd, []int{1}
}
func (m *MsgRegisterAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRegisterAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRegisterAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRegisterAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRegisterAccountResponse.Merge(m, src)
}
func (m *MsgRegisterAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRegisterAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRegisterAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRegisterAccountResponse proto.InternalMessageInfo

// MsgSubmitTx defines the payload for Msg/SubmitTx
type MsgSubmitTx struct {
	Owner        string     `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	ConnectionId string     `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty" yaml:"connection_id"`
	Msg          *types.Any `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *MsgSubmitTx) Reset()         { *m = MsgSubmitTx{} }
func (m *MsgSubmitTx) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitTx) ProtoMessage()    {}
func (*MsgSubmitTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f79e1c9e86a55bd, []int{2}
}
func (m *MsgSubmitTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitTx.Merge(m, src)
}
func (m *MsgSubmitTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitTx proto.InternalMessageInfo

// MsgSubmitTxResponse defines the response for Msg/SubmitTx
type MsgSubmitTxResponse struct {
}

func (m *MsgSubmitTxResponse) Reset()         { *m = MsgSubmitTxResponse{} }
func (m *MsgSubmitTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitTxResponse) ProtoMessage()    {}
func (*MsgSubmitTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f79e1c9e86a55bd, []int{3}
}
func (m *MsgSubmitTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitTxResponse.Merge(m, src)
}
func (m *MsgSubmitTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitTxResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRegisterAccount)(nil), "intertx.MsgRegisterAccount")
	proto.RegisterType((*MsgRegisterAccountResponse)(nil), "intertx.MsgRegisterAccountResponse")
	proto.RegisterType((*MsgSubmitTx)(nil), "intertx.MsgSubmitTx")
	proto.RegisterType((*MsgSubmitTxResponse)(nil), "intertx.MsgSubmitTxResponse")
}

func init() { proto.RegisterFile("provenance/intertx/tx.proto", fileDescriptor_6f79e1c9e86a55bd) }

var fileDescriptor_6f79e1c9e86a55bd = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0x31, 0x4b, 0xfb, 0x40,
	0x1c, 0xcd, 0xfd, 0xcb, 0xdf, 0xd6, 0xab, 0x22, 0xc4, 0x08, 0x31, 0x2d, 0xa9, 0x44, 0x10, 0x97,
	0x26, 0x50, 0x9d, 0x0a, 0x0a, 0xed, 0xe6, 0x50, 0x87, 0xd4, 0xc9, 0x45, 0x9a, 0xf4, 0x3c, 0x0f,
	0x9a, 0xbb, 0x90, 0xbb, 0xd4, 0xe4, 0x1b, 0xb8, 0x08, 0x0e, 0x4e, 0x4e, 0xfd, 0x38, 0x8e, 0x1d,
	0x9d, 0x44, 0xda, 0xc5, 0xd9, 0x4f, 0x20, 0x49, 0x1a, 0x1a, 0xad, 0xb8, 0xb9, 0xdd, 0x7b, 0xef,
	0x7e, 0x3f, 0xde, 0xbd, 0x7b, 0xb0, 0xe6, 0x07, 0x6c, 0x8c, 0xe8, 0x80, 0xba, 0xc8, 0x22, 0x54,
	0xa0, 0x40, 0x44, 0x96, 0x88, 0x4c, 0x3f, 0x60, 0x82, 0xc9, 0xe5, 0x05, 0xa3, 0x29, 0x98, 0x61,
	0x96, 0x72, 0x56, 0x72, 0xca, 0x64, 0x6d, 0x17, 0x33, 0x86, 0x47, 0xc8, 0x4a, 0x91, 0x13, 0x5e,
	0x5b, 0x03, 0x1a, 0x67, 0x92, 0x71, 0x0f, 0xa0, 0xdc, 0xe3, 0xd8, 0x46, 0x98, 0x70, 0x81, 0x82,
	0x8e, 0xeb, 0xb2, 0x90, 0x0a, 0x59, 0x81, 0xff, 0xd9, 0x2d, 0x45, 0x81, 0x0a, 0xf6, 0xc0, 0xe1,
	0xba, 0x9d, 0x01, 0xf9, 0x04, 0x6e, 0xba, 0x8c, 0x52, 0xe4, 0x0a, 0xc2, 0xe8, 0x15, 0x19, 0xaa,
	0xff, 0x12, 0xb5, 0xab, 0x7e, 0xbc, 0x36, 0x94, 0x78, 0xe0, 0x8d, 0xda, 0xc6, 0x17, 0xd9, 0xb0,
	0x37, 0x96, 0xf8, 0x6c, 0x28, 0xab, 0xb0, 0x3c, 0x46, 0x01, 0x27, 0x8c, 0xaa, 0xa5, 0x74, 0x6d,
	0x0e, 0xdb, 0x95, 0xbb, 0x49, 0x43, 0x7a, 0x9f, 0x34, 0x24, 0xa3, 0x0e, 0xb5, 0x55, 0x3b, 0x36,
	0xe2, 0x3e, 0xa3, 0x1c, 0x19, 0x8f, 0x00, 0x56, 0x7b, 0x1c, 0xf7, 0x43, 0xc7, 0x23, 0xe2, 0x22,
	0xfa, 0x1b, 0x9b, 0x07, 0xb0, 0xe4, 0x71, 0x9c, 0x5a, 0xac, 0xb6, 0x14, 0x33, 0xcb, 0xce, 0xcc,
	0xb3, 0x33, 0x3b, 0x34, 0xb6, 0x93, 0x0b, 0x05, 0xd3, 0x3b, 0x70, 0xbb, 0xe0, 0x2a, 0x77, 0xdb,
	0x7a, 0x02, 0xb0, 0xd4, 0xe3, 0x58, 0xee, 0xc3, 0xad, 0xef, 0xf9, 0xd6, 0xcc, 0xc5, 0x8f, 0x99,
	0xab, 0xaf, 0xd5, 0xf6, 0x7f, 0x11, 0xf3, 0xe5, 0xf2, 0x29, 0xac, 0x2c, 0x63, 0x28, 0x0e, 0xe4,
	0xac, 0x56, 0xff, 0x89, 0xcd, 0xe7, 0xbb, 0xe7, 0xcf, 0x33, 0x1d, 0x4c, 0x67, 0x3a, 0x78, 0x9b,
	0xe9, 0xe0, 0x61, 0xae, 0x4b, 0xd3, 0xb9, 0x2e, 0xbd, 0xcc, 0x75, 0xe9, 0xf2, 0x18, 0x13, 0x71,
	0x13, 0x3a, 0xa6, 0xcb, 0x3c, 0x6b, 0x59, 0xba, 0x26, 0x61, 0x05, 0x64, 0x45, 0x59, 0x09, 0x9b,
	0x49, 0x0b, 0x63, 0x1f, 0x71, 0x67, 0x2d, 0x0d, 0xe8, 0xe8, 0x33, 0x00, 0x00, 0xff, 0xff, 0xf9,
	0xcb, 0xbd, 0xdc, 0xa8, 0x02, 0x00, 0x00,
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
	// Register defines a rpc handler for MsgRegisterAccount
	RegisterAccount(ctx context.Context, in *MsgRegisterAccount, opts ...grpc.CallOption) (*MsgRegisterAccountResponse, error)
	// SubmitTx defines a rpc handler for MsgSubmitTx
	SubmitTx(ctx context.Context, in *MsgSubmitTx, opts ...grpc.CallOption) (*MsgSubmitTxResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RegisterAccount(ctx context.Context, in *MsgRegisterAccount, opts ...grpc.CallOption) (*MsgRegisterAccountResponse, error) {
	out := new(MsgRegisterAccountResponse)
	err := c.cc.Invoke(ctx, "/intertx.Msg/RegisterAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitTx(ctx context.Context, in *MsgSubmitTx, opts ...grpc.CallOption) (*MsgSubmitTxResponse, error) {
	out := new(MsgSubmitTxResponse)
	err := c.cc.Invoke(ctx, "/intertx.Msg/SubmitTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Register defines a rpc handler for MsgRegisterAccount
	RegisterAccount(context.Context, *MsgRegisterAccount) (*MsgRegisterAccountResponse, error)
	// SubmitTx defines a rpc handler for MsgSubmitTx
	SubmitTx(context.Context, *MsgSubmitTx) (*MsgSubmitTxResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RegisterAccount(ctx context.Context, req *MsgRegisterAccount) (*MsgRegisterAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterAccount not implemented")
}
func (*UnimplementedMsgServer) SubmitTx(ctx context.Context, req *MsgSubmitTx) (*MsgSubmitTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTx not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RegisterAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/intertx.Msg/RegisterAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterAccount(ctx, req.(*MsgRegisterAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/intertx.Msg/SubmitTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitTx(ctx, req.(*MsgSubmitTx))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "intertx.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterAccount",
			Handler:    _Msg_RegisterAccount_Handler,
		},
		{
			MethodName: "SubmitTx",
			Handler:    _Msg_SubmitTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provenance/intertx/tx.proto",
}

func (m *MsgRegisterAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConnectionId)))
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

func (m *MsgRegisterAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRegisterAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRegisterAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConnectionId)))
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

func (m *MsgSubmitTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgRegisterAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRegisterAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitTxResponse) Size() (n int) {
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
func (m *MsgRegisterAccount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterAccount: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
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
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
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
			m.Version = string(dAtA[iNdEx:postIndex])
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
func (m *MsgRegisterAccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRegisterAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRegisterAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgSubmitTx) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitTx: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
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
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &types.Any{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgSubmitTxResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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