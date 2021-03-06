// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tunnel.proto

package kube_tunnel

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LogLevel int32

const (
	LogLevel_INFO    LogLevel = 0
	LogLevel_VERBOSE LogLevel = 1
	LogLevel_DEBUG   LogLevel = 2
	LogLevel_WARNING LogLevel = 3
	LogLevel_ERROR   LogLevel = 4
)

var LogLevel_name = map[int32]string{
	0: "INFO",
	1: "VERBOSE",
	2: "DEBUG",
	3: "WARNING",
	4: "ERROR",
}

var LogLevel_value = map[string]int32{
	"INFO":    0,
	"VERBOSE": 1,
	"DEBUG":   2,
	"WARNING": 3,
	"ERROR":   4,
}

func (x LogLevel) String() string {
	return proto.EnumName(LogLevel_name, int32(x))
}

func (LogLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{0}
}

type TunnelScheme int32

const (
	TunnelScheme_TCP TunnelScheme = 0
	TunnelScheme_UDP TunnelScheme = 1
)

var TunnelScheme_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}

var TunnelScheme_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x TunnelScheme) String() string {
	return proto.EnumName(TunnelScheme_name, int32(x))
}

func (TunnelScheme) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{1}
}

type LogMessage struct {
	LogLevel             LogLevel `protobuf:"varint,1,opt,name=logLevel,proto3,enum=grpc_tunnel.LogLevel" json:"logLevel,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogMessage) Reset()         { *m = LogMessage{} }
func (m *LogMessage) String() string { return proto.CompactTextString(m) }
func (*LogMessage) ProtoMessage()    {}
func (*LogMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{0}
}

func (m *LogMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogMessage.Unmarshal(m, b)
}
func (m *LogMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogMessage.Marshal(b, m, deterministic)
}
func (m *LogMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogMessage.Merge(m, src)
}
func (m *LogMessage) XXX_Size() int {
	return xxx_messageInfo_LogMessage.Size(m)
}
func (m *LogMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LogMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LogMessage proto.InternalMessageInfo

func (m *LogMessage) GetLogLevel() LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return LogLevel_INFO
}

func (m *LogMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SocketDataRequest struct {
	Port                 int32        `protobuf:"varint,1,opt,name=port,proto3" json:"port,omitempty"`
	RequestId            string       `protobuf:"bytes,2,opt,name=requestId,proto3" json:"requestId,omitempty"`
	LogLevel             LogLevel     `protobuf:"varint,3,opt,name=logLevel,proto3,enum=grpc_tunnel.LogLevel" json:"logLevel,omitempty"`
	Scheme               TunnelScheme `protobuf:"varint,4,opt,name=scheme,proto3,enum=grpc_tunnel.TunnelScheme" json:"scheme,omitempty"`
	Data                 []byte       `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	ShouldClose          bool         `protobuf:"varint,6,opt,name=shouldClose,proto3" json:"shouldClose,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SocketDataRequest) Reset()         { *m = SocketDataRequest{} }
func (m *SocketDataRequest) String() string { return proto.CompactTextString(m) }
func (*SocketDataRequest) ProtoMessage()    {}
func (*SocketDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{1}
}

func (m *SocketDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketDataRequest.Unmarshal(m, b)
}
func (m *SocketDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketDataRequest.Marshal(b, m, deterministic)
}
func (m *SocketDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketDataRequest.Merge(m, src)
}
func (m *SocketDataRequest) XXX_Size() int {
	return xxx_messageInfo_SocketDataRequest.Size(m)
}
func (m *SocketDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SocketDataRequest proto.InternalMessageInfo

func (m *SocketDataRequest) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *SocketDataRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *SocketDataRequest) GetLogLevel() LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return LogLevel_INFO
}

func (m *SocketDataRequest) GetScheme() TunnelScheme {
	if m != nil {
		return m.Scheme
	}
	return TunnelScheme_TCP
}

func (m *SocketDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SocketDataRequest) GetShouldClose() bool {
	if m != nil {
		return m.ShouldClose
	}
	return false
}

type SocketDataResponse struct {
	HasErr               bool        `protobuf:"varint,1,opt,name=hasErr,proto3" json:"hasErr,omitempty"`
	LogMessage           *LogMessage `protobuf:"bytes,2,opt,name=logMessage,proto3" json:"logMessage,omitempty"`
	RequestId            string      `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Data                 []byte      `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ShouldClose          bool        `protobuf:"varint,5,opt,name=shouldClose,proto3" json:"shouldClose,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SocketDataResponse) Reset()         { *m = SocketDataResponse{} }
func (m *SocketDataResponse) String() string { return proto.CompactTextString(m) }
func (*SocketDataResponse) ProtoMessage()    {}
func (*SocketDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f51ddaa7891a711, []int{2}
}

func (m *SocketDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SocketDataResponse.Unmarshal(m, b)
}
func (m *SocketDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SocketDataResponse.Marshal(b, m, deterministic)
}
func (m *SocketDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SocketDataResponse.Merge(m, src)
}
func (m *SocketDataResponse) XXX_Size() int {
	return xxx_messageInfo_SocketDataResponse.Size(m)
}
func (m *SocketDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SocketDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SocketDataResponse proto.InternalMessageInfo

func (m *SocketDataResponse) GetHasErr() bool {
	if m != nil {
		return m.HasErr
	}
	return false
}

func (m *SocketDataResponse) GetLogMessage() *LogMessage {
	if m != nil {
		return m.LogMessage
	}
	return nil
}

func (m *SocketDataResponse) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *SocketDataResponse) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *SocketDataResponse) GetShouldClose() bool {
	if m != nil {
		return m.ShouldClose
	}
	return false
}

func init() {
	proto.RegisterEnum("grpc_tunnel.LogLevel", LogLevel_name, LogLevel_value)
	proto.RegisterEnum("grpc_tunnel.TunnelScheme", TunnelScheme_name, TunnelScheme_value)
	proto.RegisterType((*LogMessage)(nil), "grpc_tunnel.LogMessage")
	proto.RegisterType((*SocketDataRequest)(nil), "grpc_tunnel.SocketDataRequest")
	proto.RegisterType((*SocketDataResponse)(nil), "grpc_tunnel.SocketDataResponse")
}

func init() { proto.RegisterFile("tunnel.proto", fileDescriptor_6f51ddaa7891a711) }

var fileDescriptor_6f51ddaa7891a711 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xed, 0x6c, 0x3e, 0x9a, 0xde, 0x54, 0x89, 0x03, 0x6a, 0x14, 0xd1, 0x90, 0xa7, 0xb0, 0x0f,
	0xc5, 0xad, 0x0f, 0x3e, 0xdb, 0x4d, 0x5c, 0x02, 0xb5, 0x5d, 0x26, 0xad, 0xa2, 0x20, 0x92, 0xb6,
	0x43, 0x2a, 0x4d, 0x33, 0x31, 0x33, 0xf1, 0xbf, 0xf9, 0x73, 0xfc, 0x27, 0x92, 0x49, 0x5a, 0x93,
	0x96, 0xc2, 0x3e, 0xcd, 0xbd, 0xe7, 0x9e, 0x3b, 0x9c, 0x73, 0xb8, 0x30, 0x14, 0x65, 0x96, 0xd1,
	0x74, 0x94, 0x17, 0x4c, 0x30, 0x6c, 0x26, 0x45, 0xbe, 0xfe, 0x51, 0x43, 0xee, 0x57, 0x80, 0x29,
	0x4b, 0x3e, 0x51, 0xce, 0xe3, 0x84, 0xe2, 0x1b, 0x30, 0x52, 0x96, 0x4c, 0xe9, 0x6f, 0x9a, 0xda,
	0xc8, 0x41, 0xde, 0xe3, 0xf1, 0xd3, 0x51, 0x8b, 0x3d, 0x9a, 0x36, 0x43, 0x72, 0xa4, 0x61, 0x1b,
	0xfa, 0xfb, 0x7a, 0xdb, 0xbe, 0x72, 0x90, 0x37, 0x20, 0x87, 0xd6, 0xfd, 0x8b, 0xe0, 0x49, 0xc4,
	0xd6, 0x3b, 0x2a, 0xfc, 0x58, 0xc4, 0x84, 0xfe, 0x2a, 0x29, 0x17, 0x18, 0x83, 0x9a, 0xb3, 0x42,
	0xc8, 0xef, 0x35, 0x22, 0x6b, 0xfc, 0x0a, 0x06, 0x45, 0x3d, 0x0e, 0x37, 0xcd, 0x2f, 0xff, 0x81,
	0x8e, 0x28, 0xe5, 0x61, 0xa2, 0x6e, 0x40, 0xe7, 0xeb, 0x2d, 0xdd, 0x53, 0x5b, 0x95, 0x0b, 0x2f,
	0x3a, 0x0b, 0x0b, 0xf9, 0x44, 0x92, 0x40, 0x1a, 0x62, 0xa5, 0x6b, 0x13, 0x8b, 0xd8, 0xd6, 0x1c,
	0xe4, 0x0d, 0x89, 0xac, 0xb1, 0x03, 0x26, 0xdf, 0xb2, 0x32, 0xdd, 0xdc, 0xa6, 0x8c, 0x53, 0x5b,
	0x77, 0x90, 0x67, 0x90, 0x36, 0xe4, 0xfe, 0x41, 0x80, 0xdb, 0x1e, 0x79, 0xce, 0x32, 0x4e, 0xf1,
	0x33, 0xd0, 0xb7, 0x31, 0x0f, 0x8a, 0x42, 0xda, 0x34, 0x48, 0xd3, 0xe1, 0xf7, 0x00, 0xe9, 0x31,
	0x6d, 0xe9, 0xd4, 0x1c, 0x3f, 0x3f, 0x35, 0xd3, 0x8c, 0x49, 0x8b, 0xda, 0x4d, 0x48, 0x39, 0x4d,
	0xe8, 0xa0, 0x5d, 0xbd, 0xac, 0x5d, 0x3b, 0xd3, 0x7e, 0xed, 0x83, 0x71, 0x88, 0x0e, 0x1b, 0xa0,
	0x86, 0xb3, 0x8f, 0x73, 0xab, 0x87, 0x4d, 0xe8, 0x7f, 0x0e, 0xc8, 0x64, 0x1e, 0x05, 0x16, 0xc2,
	0x03, 0xd0, 0xfc, 0x60, 0xb2, 0xbc, 0xb3, 0xae, 0x2a, 0xfc, 0xcb, 0x07, 0x32, 0x0b, 0x67, 0x77,
	0x96, 0x52, 0xe1, 0x01, 0x21, 0x73, 0x62, 0xa9, 0xd7, 0x0e, 0x0c, 0xdb, 0x79, 0xe2, 0x3e, 0x28,
	0x8b, 0xdb, 0x7b, 0xab, 0x57, 0x15, 0x4b, 0xff, 0xde, 0x42, 0xe3, 0xef, 0xa0, 0xd7, 0x0c, 0x1c,
	0x01, 0x84, 0xd9, 0x4f, 0xd1, 0x74, 0xaf, 0x3b, 0xc6, 0xcf, 0x2e, 0xe5, 0xe5, 0x9b, 0x8b, 0xf3,
	0x3a, 0x65, 0xb7, 0xe7, 0xa1, 0xb7, 0x68, 0xf2, 0xe8, 0x9b, 0xb9, 0x2b, 0x57, 0xb4, 0xe1, 0xad,
	0x74, 0x79, 0xe4, 0xef, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xca, 0x11, 0x4c, 0xe9, 0xf4, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TunnelClient is the client API for Tunnel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TunnelClient interface {
	InitTunnel(ctx context.Context, opts ...grpc.CallOption) (Tunnel_InitTunnelClient, error)
}

type tunnelClient struct {
	cc *grpc.ClientConn
}

func NewTunnelClient(cc *grpc.ClientConn) TunnelClient {
	return &tunnelClient{cc}
}

func (c *tunnelClient) InitTunnel(ctx context.Context, opts ...grpc.CallOption) (Tunnel_InitTunnelClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Tunnel_serviceDesc.Streams[0], "/grpc_tunnel.Tunnel/InitTunnel", opts...)
	if err != nil {
		return nil, err
	}
	x := &tunnelInitTunnelClient{stream}
	return x, nil
}

type Tunnel_InitTunnelClient interface {
	Send(*SocketDataRequest) error
	Recv() (*SocketDataResponse, error)
	grpc.ClientStream
}

type tunnelInitTunnelClient struct {
	grpc.ClientStream
}

func (x *tunnelInitTunnelClient) Send(m *SocketDataRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tunnelInitTunnelClient) Recv() (*SocketDataResponse, error) {
	m := new(SocketDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TunnelServer is the server API for Tunnel service.
type TunnelServer interface {
	InitTunnel(Tunnel_InitTunnelServer) error
}

// UnimplementedTunnelServer can be embedded to have forward compatible implementations.
type UnimplementedTunnelServer struct {
}

func (*UnimplementedTunnelServer) InitTunnel(srv Tunnel_InitTunnelServer) error {
	return status.Errorf(codes.Unimplemented, "method InitTunnel not implemented")
}

func RegisterTunnelServer(s *grpc.Server, srv TunnelServer) {
	s.RegisterService(&_Tunnel_serviceDesc, srv)
}

func _Tunnel_InitTunnel_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TunnelServer).InitTunnel(&tunnelInitTunnelServer{stream})
}

type Tunnel_InitTunnelServer interface {
	Send(*SocketDataResponse) error
	Recv() (*SocketDataRequest, error)
	grpc.ServerStream
}

type tunnelInitTunnelServer struct {
	grpc.ServerStream
}

func (x *tunnelInitTunnelServer) Send(m *SocketDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tunnelInitTunnelServer) Recv() (*SocketDataRequest, error) {
	m := new(SocketDataRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Tunnel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_tunnel.Tunnel",
	HandlerType: (*TunnelServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "InitTunnel",
			Handler:       _Tunnel_InitTunnel_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "tunnel.proto",
}
