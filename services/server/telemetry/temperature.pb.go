// Code generated by protoc-gen-go. DO NOT EDIT.
// source: temperature.proto

package telemetry

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Temp struct {
	Temp                 float32  `protobuf:"fixed32,1,opt,name=temp,proto3" json:"temp,omitempty"`
	LocationId           int32    `protobuf:"varint,2,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Temp) Reset()         { *m = Temp{} }
func (m *Temp) String() string { return proto.CompactTextString(m) }
func (*Temp) ProtoMessage()    {}
func (*Temp) Descriptor() ([]byte, []int) {
	return fileDescriptor_b42fa9d0f0973a72, []int{0}
}

func (m *Temp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Temp.Unmarshal(m, b)
}
func (m *Temp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Temp.Marshal(b, m, deterministic)
}
func (m *Temp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Temp.Merge(m, src)
}
func (m *Temp) XXX_Size() int {
	return xxx_messageInfo_Temp.Size(m)
}
func (m *Temp) XXX_DiscardUnknown() {
	xxx_messageInfo_Temp.DiscardUnknown(m)
}

var xxx_messageInfo_Temp proto.InternalMessageInfo

func (m *Temp) GetTemp() float32 {
	if m != nil {
		return m.Temp
	}
	return 0
}

func (m *Temp) GetLocationId() int32 {
	if m != nil {
		return m.LocationId
	}
	return 0
}

func init() {
	proto.RegisterType((*Temp)(nil), "telemetry.Temp")
}

func init() { proto.RegisterFile("temperature.proto", fileDescriptor_b42fa9d0f0973a72) }

var fileDescriptor_b42fa9d0f0973a72 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x49, 0xcd, 0x2d,
	0x48, 0x2d, 0x4a, 0x2c, 0x29, 0x2d, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2c,
	0x49, 0xcd, 0x49, 0xcd, 0x4d, 0x2d, 0x29, 0xaa, 0x94, 0x92, 0x4e, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x07, 0x4b, 0x24, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0xd4, 0x29, 0x59,
	0x73, 0xb1, 0x84, 0xa4, 0xe6, 0x16, 0x08, 0x09, 0x71, 0xb1, 0x80, 0x0c, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0a, 0x02, 0xb3, 0x85, 0xe4, 0xb9, 0xb8, 0x73, 0xf2, 0x93, 0x13, 0x4b, 0x32, 0xf3,
	0xf3, 0xe2, 0x33, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x83, 0xb8, 0x60, 0x42, 0x9e, 0x29,
	0x46, 0x0e, 0x5c, 0xac, 0x20, 0xcd, 0xc5, 0x42, 0xe6, 0x5c, 0x5c, 0xce, 0x45, 0xa9, 0x89, 0x25,
	0xa9, 0x60, 0xb3, 0xf8, 0xf5, 0xe0, 0x96, 0xeb, 0x81, 0x04, 0xa4, 0xc4, 0xf4, 0x20, 0x4e, 0xd0,
	0x83, 0x39, 0x41, 0xcf, 0x15, 0xe4, 0x04, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x88, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0x40, 0xcc, 0xc5, 0xd2, 0xc2, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TempsClient is the client API for Temps service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TempsClient interface {
	CreateTemp(ctx context.Context, in *Temp, opts ...grpc.CallOption) (*empty.Empty, error)
}

type tempsClient struct {
	cc *grpc.ClientConn
}

func NewTempsClient(cc *grpc.ClientConn) TempsClient {
	return &tempsClient{cc}
}

func (c *tempsClient) CreateTemp(ctx context.Context, in *Temp, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/telemetry.Temps/CreateTemp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TempsServer is the server API for Temps service.
type TempsServer interface {
	CreateTemp(context.Context, *Temp) (*empty.Empty, error)
}

func RegisterTempsServer(s *grpc.Server, srv TempsServer) {
	s.RegisterService(&_Temps_serviceDesc, srv)
}

func _Temps_CreateTemp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Temp)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TempsServer).CreateTemp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/telemetry.Temps/CreateTemp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TempsServer).CreateTemp(ctx, req.(*Temp))
	}
	return interceptor(ctx, in, info, handler)
}

var _Temps_serviceDesc = grpc.ServiceDesc{
	ServiceName: "telemetry.Temps",
	HandlerType: (*TempsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTemp",
			Handler:    _Temps_CreateTemp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "temperature.proto",
}
