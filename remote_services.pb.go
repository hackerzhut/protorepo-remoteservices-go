// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote_services.proto

/*
Package remoteservices is a generated protocol buffer package.

It is generated from these files:
	remote_services.proto

It has these top-level messages:
	Request
	Response
	Result
	Command
	Error
*/
package remoteservices

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

// Common Request for all remote services
type Request struct {
	Brand         string `protobuf:"bytes,1,opt,name=brand" json:"brand,omitempty"`
	Channel       string `protobuf:"bytes,2,opt,name=channel" json:"channel,omitempty"`
	Authorization string `protobuf:"bytes,3,opt,name=authorization" json:"authorization,omitempty"`
	Vin           string `protobuf:"bytes,4,opt,name=vin" json:"vin,omitempty"`
	RequestId     string `protobuf:"bytes,5,opt,name=requestId" json:"requestId,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Request) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Request) GetAuthorization() string {
	if m != nil {
		return m.Authorization
	}
	return ""
}

func (m *Request) GetVin() string {
	if m != nil {
		return m.Vin
	}
	return ""
}

func (m *Request) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

type Response struct {
	Result  *Result  `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Command *Command `protobuf:"bytes,2,opt,name=command" json:"command,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *Response) GetCommand() *Command {
	if m != nil {
		return m.Command
	}
	return nil
}

type Result struct {
	Code        int64  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Result) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Command struct {
	IssuedTime int64  `protobuf:"varint,1,opt,name=issuedTime" json:"issuedTime,omitempty"`
	StartTime  int64  `protobuf:"varint,2,opt,name=startTime" json:"startTime,omitempty"`
	Status     string `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Action     string `protobuf:"bytes,4,opt,name=action" json:"action,omitempty"`
	Progress   string `protobuf:"bytes,5,opt,name=progress" json:"progress,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Command) GetIssuedTime() int64 {
	if m != nil {
		return m.IssuedTime
	}
	return 0
}

func (m *Command) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Command) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Command) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Command) GetProgress() string {
	if m != nil {
		return m.Progress
	}
	return ""
}

type Error struct {
	Code    string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "remoteservices.Request")
	proto.RegisterType((*Response)(nil), "remoteservices.Response")
	proto.RegisterType((*Result)(nil), "remoteservices.Result")
	proto.RegisterType((*Command)(nil), "remoteservices.Command")
	proto.RegisterType((*Error)(nil), "remoteservices.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RemoteServices service

type RemoteServicesClient interface {
	GetStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	GetStatusStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (RemoteServices_GetStatusStreamClient, error)
}

type remoteServicesClient struct {
	cc *grpc.ClientConn
}

func NewRemoteServicesClient(cc *grpc.ClientConn) RemoteServicesClient {
	return &remoteServicesClient{cc}
}

func (c *remoteServicesClient) GetStatus(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/remoteservices.RemoteServices/GetStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServicesClient) GetStatusStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (RemoteServices_GetStatusStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RemoteServices_serviceDesc.Streams[0], c.cc, "/remoteservices.RemoteServices/GetStatusStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteServicesGetStatusStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RemoteServices_GetStatusStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type remoteServicesGetStatusStreamClient struct {
	grpc.ClientStream
}

func (x *remoteServicesGetStatusStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RemoteServices service

type RemoteServicesServer interface {
	GetStatus(context.Context, *Request) (*Response, error)
	GetStatusStream(*Request, RemoteServices_GetStatusStreamServer) error
}

func RegisterRemoteServicesServer(s *grpc.Server, srv RemoteServicesServer) {
	s.RegisterService(&_RemoteServices_serviceDesc, srv)
}

func _RemoteServices_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServicesServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remoteservices.RemoteServices/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServicesServer).GetStatus(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteServices_GetStatusStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RemoteServicesServer).GetStatusStream(m, &remoteServicesGetStatusStreamServer{stream})
}

type RemoteServices_GetStatusStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type remoteServicesGetStatusStreamServer struct {
	grpc.ServerStream
}

func (x *remoteServicesGetStatusStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

var _RemoteServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remoteservices.RemoteServices",
	HandlerType: (*RemoteServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _RemoteServices_GetStatus_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStatusStream",
			Handler:       _RemoteServices_GetStatusStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remote_services.proto",
}

func init() { proto.RegisterFile("remote_services.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x25, 0xed, 0x36, 0xbb, 0x99, 0x15, 0x0b, 0x1a, 0xc1, 0x6e, 0x54, 0xa1, 0xaa, 0x8a, 0x38,
	0x70, 0x8a, 0xa0, 0x88, 0x2b, 0x42, 0x20, 0x04, 0xdc, 0x90, 0xcb, 0x1d, 0xb9, 0xc9, 0xa8, 0x8d,
	0xd4, 0xc4, 0xc1, 0x76, 0x2a, 0xc1, 0x9f, 0xe0, 0xc0, 0x95, 0x1f, 0x8b, 0x33, 0x76, 0xfa, 0x05,
	0x27, 0x6e, 0x79, 0x6f, 0xc6, 0x33, 0xf3, 0xde, 0x0b, 0x3c, 0xd6, 0x54, 0x2b, 0x4b, 0x5f, 0x0d,
	0xe9, 0x5d, 0x55, 0x90, 0xc9, 0x5b, 0xad, 0xac, 0xc2, 0x1b, 0x4f, 0x0f, 0x6c, 0xf6, 0x33, 0x82,
	0x4b, 0x41, 0xdf, 0x3a, 0x32, 0x16, 0x1f, 0xc1, 0x64, 0xa5, 0x65, 0x53, 0xa6, 0xd1, 0x3c, 0x7a,
	0x96, 0x08, 0x0f, 0x30, 0x85, 0xcb, 0x62, 0x23, 0x9b, 0x86, 0xb6, 0xe9, 0x88, 0xf9, 0x01, 0xe2,
	0x53, 0xb8, 0x2f, 0x3b, 0xbb, 0x51, 0xba, 0xfa, 0x21, 0x6d, 0xa5, 0x9a, 0x74, 0xcc, 0xf5, 0x53,
	0x12, 0x1f, 0xc2, 0x78, 0x57, 0x35, 0xe9, 0x05, 0xd7, 0xfa, 0x4f, 0x7c, 0x02, 0x89, 0xf6, 0x2b,
	0x3f, 0x95, 0xe9, 0x84, 0xf9, 0x03, 0x91, 0xd5, 0x70, 0x25, 0xc8, 0xb4, 0xaa, 0x31, 0x84, 0x39,
	0xc4, 0x9a, 0x4c, 0xb7, 0xb5, 0x7c, 0xd2, 0xf5, 0xe2, 0x36, 0x3f, 0x3d, 0x3f, 0x17, 0x5c, 0x15,
	0xa1, 0x0b, 0x5f, 0xb8, 0x5b, 0x55, 0x5d, 0xf7, 0x1a, 0x46, 0xfc, 0xe0, 0xee, 0xfc, 0xc1, 0x3b,
	0x5f, 0x16, 0x43, 0x5f, 0xf6, 0x1a, 0x62, 0x3f, 0x04, 0x11, 0x2e, 0x0a, 0x55, 0x12, 0xaf, 0x1a,
	0x0b, 0xfe, 0xc6, 0x39, 0x5c, 0x97, 0x64, 0x0a, 0x5d, 0xb5, 0x2c, 0xd0, 0x1b, 0x70, 0x4c, 0x65,
	0xbf, 0x9c, 0x81, 0x61, 0x28, 0xce, 0x00, 0x2a, 0x63, 0x3a, 0x2a, 0xbf, 0x54, 0xf5, 0x30, 0xe7,
	0x88, 0xe9, 0x85, 0x1b, 0x2b, 0xb5, 0xe5, 0xf2, 0x88, 0xcb, 0x07, 0x02, 0x6f, 0x21, 0x76, 0xc0,
	0x76, 0x26, 0xf8, 0x18, 0x50, 0xcf, 0xcb, 0x82, 0xd7, 0x7b, 0x0f, 0x03, 0xc2, 0x29, 0x5c, 0xb9,
	0x4c, 0xd7, 0x4e, 0xba, 0x09, 0x2e, 0xee, 0x71, 0xf6, 0x0a, 0x26, 0xef, 0xb5, 0x56, 0xfa, 0x44,
	0x54, 0x12, 0x44, 0xb9, 0x44, 0x6b, 0xd7, 0x24, 0xd7, 0x34, 0x24, 0x1a, 0xe0, 0xe2, 0x77, 0x04,
	0x37, 0x82, 0x0d, 0x5b, 0x06, 0xc3, 0xf0, 0x0d, 0x24, 0x1f, 0xc8, 0x2e, 0xfd, 0x29, 0x77, 0x7f,
	0xfb, 0xcf, 0xb1, 0x4d, 0xd3, 0x7f, 0x04, 0xc3, 0x11, 0x66, 0xf7, 0xf0, 0x23, 0x3c, 0xd8, 0x4f,
	0x58, 0x5a, 0x4d, 0xb2, 0xfe, 0xaf, 0x39, 0xcf, 0xa3, 0xb7, 0x19, 0xcc, 0x5c, 0x6c, 0xb9, 0x55,
	0xdf, 0x95, 0x95, 0xe7, 0x9d, 0x6b, 0xdd, 0x16, 0x9f, 0xa3, 0x55, 0xcc, 0xff, 0xf9, 0xcb, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xdd, 0x2d, 0xc6, 0x00, 0x03, 0x00, 0x00,
}
