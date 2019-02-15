// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protocol.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Vector struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vector) Reset()         { *m = Vector{} }
func (m *Vector) String() string { return proto.CompactTextString(m) }
func (*Vector) ProtoMessage()    {}
func (*Vector) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_a162741ecb4fa1ff, []int{0}
}
func (m *Vector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vector.Unmarshal(m, b)
}
func (m *Vector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vector.Marshal(b, m, deterministic)
}
func (dst *Vector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vector.Merge(dst, src)
}
func (m *Vector) XXX_Size() int {
	return xxx_messageInfo_Vector.Size(m)
}
func (m *Vector) XXX_DiscardUnknown() {
	xxx_messageInfo_Vector.DiscardUnknown(m)
}

var xxx_messageInfo_Vector proto.InternalMessageInfo

func (m *Vector) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vector) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Vector) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

type Connection struct {
	ConnectedTo          uint64   `protobuf:"varint,1,opt,name=connected_to,json=connectedTo,proto3" json:"connected_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_a162741ecb4fa1ff, []int{1}
}
func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (dst *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(dst, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetConnectedTo() uint64 {
	if m != nil {
		return m.ConnectedTo
	}
	return 0
}

type Cell struct {
	Id                   uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	EnergyLevel          uint64        `protobuf:"varint,2,opt,name=energy_level,json=energyLevel,proto3" json:"energy_level,omitempty"`
	Pos                  *Vector       `protobuf:"bytes,3,opt,name=pos,proto3" json:"pos,omitempty"`
	Vel                  *Vector       `protobuf:"bytes,4,opt,name=vel,proto3" json:"vel,omitempty"`
	Dna                  []byte        `protobuf:"bytes,5,opt,name=dna,proto3" json:"dna,omitempty"`
	Connections          []*Connection `protobuf:"bytes,6,rep,name=connections,proto3" json:"connections,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Cell) Reset()         { *m = Cell{} }
func (m *Cell) String() string { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()    {}
func (*Cell) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_a162741ecb4fa1ff, []int{2}
}
func (m *Cell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cell.Unmarshal(m, b)
}
func (m *Cell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cell.Marshal(b, m, deterministic)
}
func (dst *Cell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cell.Merge(dst, src)
}
func (m *Cell) XXX_Size() int {
	return xxx_messageInfo_Cell.Size(m)
}
func (m *Cell) XXX_DiscardUnknown() {
	xxx_messageInfo_Cell.DiscardUnknown(m)
}

var xxx_messageInfo_Cell proto.InternalMessageInfo

func (m *Cell) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Cell) GetEnergyLevel() uint64 {
	if m != nil {
		return m.EnergyLevel
	}
	return 0
}

func (m *Cell) GetPos() *Vector {
	if m != nil {
		return m.Pos
	}
	return nil
}

func (m *Cell) GetVel() *Vector {
	if m != nil {
		return m.Vel
	}
	return nil
}

func (m *Cell) GetDna() []byte {
	if m != nil {
		return m.Dna
	}
	return nil
}

func (m *Cell) GetConnections() []*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

type CellComputeBatch struct {
	TimeStep             uint64   `protobuf:"varint,1,opt,name=time_step,json=timeStep,proto3" json:"time_step,omitempty"`
	CellsToCompute       []*Cell  `protobuf:"bytes,2,rep,name=cells_to_compute,json=cellsToCompute,proto3" json:"cells_to_compute,omitempty"`
	CellsInProximity     []*Cell  `protobuf:"bytes,3,rep,name=cells_in_proximity,json=cellsInProximity,proto3" json:"cells_in_proximity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CellComputeBatch) Reset()         { *m = CellComputeBatch{} }
func (m *CellComputeBatch) String() string { return proto.CompactTextString(m) }
func (*CellComputeBatch) ProtoMessage()    {}
func (*CellComputeBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_a162741ecb4fa1ff, []int{3}
}
func (m *CellComputeBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CellComputeBatch.Unmarshal(m, b)
}
func (m *CellComputeBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CellComputeBatch.Marshal(b, m, deterministic)
}
func (dst *CellComputeBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CellComputeBatch.Merge(dst, src)
}
func (m *CellComputeBatch) XXX_Size() int {
	return xxx_messageInfo_CellComputeBatch.Size(m)
}
func (m *CellComputeBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_CellComputeBatch.DiscardUnknown(m)
}

var xxx_messageInfo_CellComputeBatch proto.InternalMessageInfo

func (m *CellComputeBatch) GetTimeStep() uint64 {
	if m != nil {
		return m.TimeStep
	}
	return 0
}

func (m *CellComputeBatch) GetCellsToCompute() []*Cell {
	if m != nil {
		return m.CellsToCompute
	}
	return nil
}

func (m *CellComputeBatch) GetCellsInProximity() []*Cell {
	if m != nil {
		return m.CellsInProximity
	}
	return nil
}

type BigBangRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BigBangRequest) Reset()         { *m = BigBangRequest{} }
func (m *BigBangRequest) String() string { return proto.CompactTextString(m) }
func (*BigBangRequest) ProtoMessage()    {}
func (*BigBangRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_a162741ecb4fa1ff, []int{4}
}
func (m *BigBangRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigBangRequest.Unmarshal(m, b)
}
func (m *BigBangRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigBangRequest.Marshal(b, m, deterministic)
}
func (dst *BigBangRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigBangRequest.Merge(dst, src)
}
func (m *BigBangRequest) XXX_Size() int {
	return xxx_messageInfo_BigBangRequest.Size(m)
}
func (m *BigBangRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BigBangRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BigBangRequest proto.InternalMessageInfo

type SlaveRegistration struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Threads              uint32   `protobuf:"varint,2,opt,name=threads,proto3" json:"threads,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlaveRegistration) Reset()         { *m = SlaveRegistration{} }
func (m *SlaveRegistration) String() string { return proto.CompactTextString(m) }
func (*SlaveRegistration) ProtoMessage()    {}
func (*SlaveRegistration) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_a162741ecb4fa1ff, []int{5}
}
func (m *SlaveRegistration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlaveRegistration.Unmarshal(m, b)
}
func (m *SlaveRegistration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlaveRegistration.Marshal(b, m, deterministic)
}
func (dst *SlaveRegistration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlaveRegistration.Merge(dst, src)
}
func (m *SlaveRegistration) XXX_Size() int {
	return xxx_messageInfo_SlaveRegistration.Size(m)
}
func (m *SlaveRegistration) XXX_DiscardUnknown() {
	xxx_messageInfo_SlaveRegistration.DiscardUnknown(m)
}

var xxx_messageInfo_SlaveRegistration proto.InternalMessageInfo

func (m *SlaveRegistration) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SlaveRegistration) GetThreads() uint32 {
	if m != nil {
		return m.Threads
	}
	return 0
}

type SlaveRegistrationResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SlaveRegistrationResponse) Reset()         { *m = SlaveRegistrationResponse{} }
func (m *SlaveRegistrationResponse) String() string { return proto.CompactTextString(m) }
func (*SlaveRegistrationResponse) ProtoMessage()    {}
func (*SlaveRegistrationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_protocol_a162741ecb4fa1ff, []int{6}
}
func (m *SlaveRegistrationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlaveRegistrationResponse.Unmarshal(m, b)
}
func (m *SlaveRegistrationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlaveRegistrationResponse.Marshal(b, m, deterministic)
}
func (dst *SlaveRegistrationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlaveRegistrationResponse.Merge(dst, src)
}
func (m *SlaveRegistrationResponse) XXX_Size() int {
	return xxx_messageInfo_SlaveRegistrationResponse.Size(m)
}
func (m *SlaveRegistrationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SlaveRegistrationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SlaveRegistrationResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Vector)(nil), "proto.Vector")
	proto.RegisterType((*Connection)(nil), "proto.Connection")
	proto.RegisterType((*Cell)(nil), "proto.Cell")
	proto.RegisterType((*CellComputeBatch)(nil), "proto.CellComputeBatch")
	proto.RegisterType((*BigBangRequest)(nil), "proto.BigBangRequest")
	proto.RegisterType((*SlaveRegistration)(nil), "proto.SlaveRegistration")
	proto.RegisterType((*SlaveRegistrationResponse)(nil), "proto.SlaveRegistrationResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CellInteractionServiceClient is the client API for CellInteractionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CellInteractionServiceClient interface {
	ComputeCellInteractions(ctx context.Context, in *CellComputeBatch, opts ...grpc.CallOption) (*CellComputeBatch, error)
	BigBang(ctx context.Context, in *BigBangRequest, opts ...grpc.CallOption) (CellInteractionService_BigBangClient, error)
}

type cellInteractionServiceClient struct {
	cc *grpc.ClientConn
}

func NewCellInteractionServiceClient(cc *grpc.ClientConn) CellInteractionServiceClient {
	return &cellInteractionServiceClient{cc}
}

func (c *cellInteractionServiceClient) ComputeCellInteractions(ctx context.Context, in *CellComputeBatch, opts ...grpc.CallOption) (*CellComputeBatch, error) {
	out := new(CellComputeBatch)
	err := c.cc.Invoke(ctx, "/proto.CellInteractionService/ComputeCellInteractions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cellInteractionServiceClient) BigBang(ctx context.Context, in *BigBangRequest, opts ...grpc.CallOption) (CellInteractionService_BigBangClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CellInteractionService_serviceDesc.Streams[0], "/proto.CellInteractionService/BigBang", opts...)
	if err != nil {
		return nil, err
	}
	x := &cellInteractionServiceBigBangClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CellInteractionService_BigBangClient interface {
	Recv() (*Cell, error)
	grpc.ClientStream
}

type cellInteractionServiceBigBangClient struct {
	grpc.ClientStream
}

func (x *cellInteractionServiceBigBangClient) Recv() (*Cell, error) {
	m := new(Cell)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CellInteractionServiceServer is the server API for CellInteractionService service.
type CellInteractionServiceServer interface {
	ComputeCellInteractions(context.Context, *CellComputeBatch) (*CellComputeBatch, error)
	BigBang(*BigBangRequest, CellInteractionService_BigBangServer) error
}

func RegisterCellInteractionServiceServer(s *grpc.Server, srv CellInteractionServiceServer) {
	s.RegisterService(&_CellInteractionService_serviceDesc, srv)
}

func _CellInteractionService_ComputeCellInteractions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CellComputeBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CellInteractionServiceServer).ComputeCellInteractions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CellInteractionService/ComputeCellInteractions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CellInteractionServiceServer).ComputeCellInteractions(ctx, req.(*CellComputeBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _CellInteractionService_BigBang_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BigBangRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CellInteractionServiceServer).BigBang(m, &cellInteractionServiceBigBangServer{stream})
}

type CellInteractionService_BigBangServer interface {
	Send(*Cell) error
	grpc.ServerStream
}

type cellInteractionServiceBigBangServer struct {
	grpc.ServerStream
}

func (x *cellInteractionServiceBigBangServer) Send(m *Cell) error {
	return x.ServerStream.SendMsg(m)
}

var _CellInteractionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CellInteractionService",
	HandlerType: (*CellInteractionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ComputeCellInteractions",
			Handler:    _CellInteractionService_ComputeCellInteractions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BigBang",
			Handler:       _CellInteractionService_BigBang_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protocol.proto",
}

// SlaveRegistrationServiceClient is the client API for SlaveRegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SlaveRegistrationServiceClient interface {
	Register(ctx context.Context, in *SlaveRegistration, opts ...grpc.CallOption) (*SlaveRegistrationResponse, error)
}

type slaveRegistrationServiceClient struct {
	cc *grpc.ClientConn
}

func NewSlaveRegistrationServiceClient(cc *grpc.ClientConn) SlaveRegistrationServiceClient {
	return &slaveRegistrationServiceClient{cc}
}

func (c *slaveRegistrationServiceClient) Register(ctx context.Context, in *SlaveRegistration, opts ...grpc.CallOption) (*SlaveRegistrationResponse, error) {
	out := new(SlaveRegistrationResponse)
	err := c.cc.Invoke(ctx, "/proto.SlaveRegistrationService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SlaveRegistrationServiceServer is the server API for SlaveRegistrationService service.
type SlaveRegistrationServiceServer interface {
	Register(context.Context, *SlaveRegistration) (*SlaveRegistrationResponse, error)
}

func RegisterSlaveRegistrationServiceServer(s *grpc.Server, srv SlaveRegistrationServiceServer) {
	s.RegisterService(&_SlaveRegistrationService_serviceDesc, srv)
}

func _SlaveRegistrationService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveRegistration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlaveRegistrationServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SlaveRegistrationService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlaveRegistrationServiceServer).Register(ctx, req.(*SlaveRegistration))
	}
	return interceptor(ctx, in, info, handler)
}

var _SlaveRegistrationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SlaveRegistrationService",
	HandlerType: (*SlaveRegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _SlaveRegistrationService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol.proto",
}

func init() { proto.RegisterFile("protocol.proto", fileDescriptor_protocol_a162741ecb4fa1ff) }

var fileDescriptor_protocol_a162741ecb4fa1ff = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x95, 0xb3, 0x69, 0xda, 0xce, 0xa6, 0x51, 0x6a, 0x09, 0x6a, 0xd2, 0x03, 0x61, 0x4f, 0x39,
	0xb5, 0x68, 0x2b, 0x0e, 0x5c, 0x13, 0x09, 0x54, 0xc1, 0x01, 0x39, 0x15, 0xd7, 0xd5, 0x76, 0x77,
	0x94, 0x5a, 0x72, 0xec, 0xc5, 0x76, 0xa3, 0xa4, 0x1f, 0xc2, 0x9d, 0x5f, 0xe1, 0xcb, 0x90, 0xed,
	0x5d, 0x08, 0x84, 0x9e, 0x76, 0xdf, 0xcc, 0x7b, 0x33, 0xef, 0x8d, 0x61, 0xd4, 0x18, 0xed, 0x74,
	0xa5, 0xe5, 0x55, 0xf8, 0xa1, 0x47, 0xe1, 0x93, 0xe5, 0x30, 0xf8, 0x8a, 0x95, 0xd3, 0x86, 0x0e,
	0x81, 0x6c, 0x19, 0x99, 0x92, 0x59, 0x8f, 0x93, 0xad, 0x47, 0x3b, 0xd6, 0x8b, 0x68, 0xe7, 0xd1,
	0x13, 0x4b, 0x22, 0x7a, 0xca, 0xae, 0x01, 0x16, 0x5a, 0x29, 0xac, 0x9c, 0xd0, 0x8a, 0xbe, 0x81,
	0x61, 0x15, 0x11, 0xd6, 0x85, 0xd3, 0x61, 0x44, 0x9f, 0xa7, 0xbf, 0x6b, 0x77, 0x3a, 0xfb, 0x49,
	0xa0, 0xbf, 0x40, 0x29, 0xe9, 0x08, 0x7a, 0xa2, 0x6e, 0x19, 0x3d, 0x51, 0x7b, 0x2d, 0x2a, 0x34,
	0xab, 0x5d, 0x21, 0x71, 0x83, 0x32, 0x2c, 0xec, 0xf3, 0x34, 0xd6, 0x3e, 0xfb, 0x12, 0x7d, 0x0d,
	0x49, 0xa3, 0x6d, 0x58, 0x9e, 0xe6, 0x67, 0xd1, 0xfc, 0x55, 0xb4, 0xcc, 0x7d, 0xc7, 0x13, 0xbc,
	0xb4, 0xff, 0x5f, 0x82, 0x9f, 0x30, 0x86, 0xa4, 0x56, 0x25, 0x3b, 0x9a, 0x92, 0xd9, 0x90, 0xfb,
	0x5f, 0x7a, 0x03, 0x9d, 0x3d, 0xa1, 0x95, 0x65, 0x83, 0x69, 0x32, 0x4b, 0xf3, 0xf3, 0x56, 0xfa,
	0x27, 0x1a, 0xdf, 0x67, 0x65, 0x3f, 0x08, 0x8c, 0x7d, 0x88, 0x85, 0x5e, 0x37, 0x8f, 0x0e, 0xe7,
	0xa5, 0xab, 0x1e, 0xe8, 0x25, 0x9c, 0x3a, 0xb1, 0xc6, 0xc2, 0x3a, 0x6c, 0xda, 0x5c, 0x27, 0xbe,
	0xb0, 0x74, 0xd8, 0xd0, 0x77, 0x30, 0xae, 0x50, 0x4a, 0x5b, 0x38, 0x5d, 0x54, 0x51, 0xc5, 0x7a,
	0x61, 0x57, 0xda, 0xed, 0x42, 0x29, 0xf9, 0x28, 0x90, 0xee, 0x74, 0x3b, 0x98, 0xbe, 0x07, 0x1a,
	0x65, 0x42, 0x15, 0x8d, 0xd1, 0x5b, 0xb1, 0x16, 0x6e, 0xc7, 0x92, 0x43, 0x61, 0x9c, 0x7e, 0xab,
	0xbe, 0x74, 0xa4, 0x6c, 0x0c, 0xa3, 0xb9, 0x58, 0xcd, 0x4b, 0xb5, 0xe2, 0xf8, 0xed, 0x11, 0xad,
	0xcb, 0x3e, 0xc2, 0xf9, 0x52, 0x96, 0x1b, 0xe4, 0xb8, 0x12, 0xd6, 0x99, 0x32, 0x3c, 0x19, 0x83,
	0xe3, 0xb2, 0xae, 0x0d, 0x5a, 0x1b, 0x3c, 0x9f, 0xf2, 0x0e, 0xfa, 0x8e, 0x7b, 0x30, 0x58, 0xd6,
	0x36, 0xbc, 0xc5, 0x19, 0xef, 0x60, 0x76, 0x09, 0xaf, 0x0e, 0x06, 0x71, 0xb4, 0x8d, 0x56, 0x16,
	0xf3, 0xef, 0x04, 0x5e, 0x7a, 0x4b, 0xb7, 0xca, 0xa1, 0x29, 0xc3, 0xc1, 0x96, 0x68, 0x36, 0xa2,
	0x42, 0xfa, 0x09, 0x2e, 0xda, 0x60, 0xff, 0x10, 0x2c, 0xbd, 0xd8, 0x0b, 0xb3, 0x7f, 0xd5, 0xc9,
	0x73, 0x0d, 0x7a, 0x0d, 0xc7, 0x6d, 0x3e, 0xfa, 0xa2, 0xe5, 0xfc, 0x9d, 0x77, 0xb2, 0x7f, 0xa0,
	0xb7, 0x24, 0xbf, 0x07, 0x76, 0xe0, 0xba, 0x73, 0xf6, 0x01, 0x4e, 0x62, 0x19, 0x0d, 0x65, 0xad,
	0xec, 0x80, 0x3c, 0x99, 0x3e, 0xd7, 0xe9, 0xc2, 0xdf, 0x0f, 0x02, 0xe1, 0xe6, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf0, 0x6e, 0x0c, 0xd3, 0x62, 0x03, 0x00, 0x00,
}
