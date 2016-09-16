// Code generated by protoc-gen-go.
// source: agent.proto
// DO NOT EDIT!

/*
Package agent is a generated protocol buffer package.

It is generated from these files:
	agent.proto

It has these top-level messages:
	AppInfo
	Process
	ProcRequest
	ProcList
	RegReply
	DeregRequest
	DeregReply
*/
package agent

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

type AppInfo struct {
	// Application's id
	Id string `protobuf:"bytes,1,opt,name=Id,json=id" json:"Id,omitempty"`
	// Framework used
	Frameworks []string `protobuf:"bytes,2,rep,name=Frameworks,json=frameworks" json:"Frameworks,omitempty"`
	// The JobIds given by corresponding frameworks
	JobIds []string `protobuf:"bytes,3,rep,name=JobIds,json=jobIds" json:"JobIds,omitempty"`
	// Process pattern used to match processes
	Pattern string `protobuf:"bytes,4,opt,name=Pattern,json=pattern" json:"Pattern,omitempty"`
	// Pid of the application's main process.
	// If it's set, then the process on the same node as
	// the target agent, otherwise not.
	Pid int32 `protobuf:"varint,5,opt,name=Pid,json=pid" json:"Pid,omitempty"`
}

func (m *AppInfo) Reset()                    { *m = AppInfo{} }
func (m *AppInfo) String() string            { return proto.CompactTextString(m) }
func (*AppInfo) ProtoMessage()               {}
func (*AppInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Process struct {
	// Process's operating system pid
	Pid int32 `protobuf:"varint,1,opt,name=Pid,json=pid" json:"Pid,omitempty"`
	// Process name
	Name string `protobuf:"bytes,2,opt,name=Name,json=name" json:"Name,omitempty"`
	// Command line to run the process
	Cmd string `protobuf:"bytes,3,opt,name=Cmd,json=cmd" json:"Cmd,omitempty"`
}

func (m *Process) Reset()                    { *m = Process{} }
func (m *Process) String() string            { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()               {}
func (*Process) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ProcRequest struct {
	AppId string `protobuf:"bytes,1,opt,name=AppId,json=appId" json:"AppId,omitempty"`
}

func (m *ProcRequest) Reset()                    { *m = ProcRequest{} }
func (m *ProcRequest) String() string            { return proto.CompactTextString(m) }
func (*ProcRequest) ProtoMessage()               {}
func (*ProcRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ProcList struct {
	Procs []*Process `protobuf:"bytes,1,rep,name=Procs,json=procs" json:"Procs,omitempty"`
}

func (m *ProcList) Reset()                    { *m = ProcList{} }
func (m *ProcList) String() string            { return proto.CompactTextString(m) }
func (*ProcList) ProtoMessage()               {}
func (*ProcList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ProcList) GetProcs() []*Process {
	if m != nil {
		return m.Procs
	}
	return nil
}

// Reply message of Register()
//
type RegReply struct {
}

func (m *RegReply) Reset()                    { *m = RegReply{} }
func (m *RegReply) String() string            { return proto.CompactTextString(m) }
func (*RegReply) ProtoMessage()               {}
func (*RegReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Request message of Deregister()
type DeregRequest struct {
	// Application's Id
	AppId string `protobuf:"bytes,1,opt,name=AppId,json=appId" json:"AppId,omitempty"`
}

func (m *DeregRequest) Reset()                    { *m = DeregRequest{} }
func (m *DeregRequest) String() string            { return proto.CompactTextString(m) }
func (*DeregRequest) ProtoMessage()               {}
func (*DeregRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// Reply message of Deregister()
type DeregReply struct {
}

func (m *DeregReply) Reset()                    { *m = DeregReply{} }
func (m *DeregReply) String() string            { return proto.CompactTextString(m) }
func (*DeregReply) ProtoMessage()               {}
func (*DeregReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func init() {
	proto.RegisterType((*AppInfo)(nil), "agent.AppInfo")
	proto.RegisterType((*Process)(nil), "agent.Process")
	proto.RegisterType((*ProcRequest)(nil), "agent.ProcRequest")
	proto.RegisterType((*ProcList)(nil), "agent.ProcList")
	proto.RegisterType((*RegReply)(nil), "agent.RegReply")
	proto.RegisterType((*DeregRequest)(nil), "agent.DeregRequest")
	proto.RegisterType((*DeregReply)(nil), "agent.DeregReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for MonitorProcs service

type MonitorProcsClient interface {
	// Manager registers an application, and agent responds
	// with the app's processes' info on the node.
	Register(ctx context.Context, in *AppInfo, opts ...grpc.CallOption) (*RegReply, error)
	// Manager deregisters an application, and agent responds
	// with a success indicator.
	Deregister(ctx context.Context, in *DeregRequest, opts ...grpc.CallOption) (*DeregReply, error)
	// Manager send application info, and agent responds
	// with the app's processes' info on the node.
	GetProcesses(ctx context.Context, in *ProcRequest, opts ...grpc.CallOption) (*ProcList, error)
}

type monitorProcsClient struct {
	cc *grpc.ClientConn
}

func NewMonitorProcsClient(cc *grpc.ClientConn) MonitorProcsClient {
	return &monitorProcsClient{cc}
}

func (c *monitorProcsClient) Register(ctx context.Context, in *AppInfo, opts ...grpc.CallOption) (*RegReply, error) {
	out := new(RegReply)
	err := grpc.Invoke(ctx, "/agent.MonitorProcs/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorProcsClient) Deregister(ctx context.Context, in *DeregRequest, opts ...grpc.CallOption) (*DeregReply, error) {
	out := new(DeregReply)
	err := grpc.Invoke(ctx, "/agent.MonitorProcs/Deregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *monitorProcsClient) GetProcesses(ctx context.Context, in *ProcRequest, opts ...grpc.CallOption) (*ProcList, error) {
	out := new(ProcList)
	err := grpc.Invoke(ctx, "/agent.MonitorProcs/GetProcesses", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MonitorProcs service

type MonitorProcsServer interface {
	// Manager registers an application, and agent responds
	// with the app's processes' info on the node.
	Register(context.Context, *AppInfo) (*RegReply, error)
	// Manager deregisters an application, and agent responds
	// with a success indicator.
	Deregister(context.Context, *DeregRequest) (*DeregReply, error)
	// Manager send application info, and agent responds
	// with the app's processes' info on the node.
	GetProcesses(context.Context, *ProcRequest) (*ProcList, error)
}

func RegisterMonitorProcsServer(s *grpc.Server, srv MonitorProcsServer) {
	s.RegisterService(&_MonitorProcs_serviceDesc, srv)
}

func _MonitorProcs_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorProcsServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.MonitorProcs/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorProcsServer).Register(ctx, req.(*AppInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorProcs_Deregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeregRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorProcsServer).Deregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.MonitorProcs/Deregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorProcsServer).Deregister(ctx, req.(*DeregRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MonitorProcs_GetProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MonitorProcsServer).GetProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/agent.MonitorProcs/GetProcesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MonitorProcsServer).GetProcesses(ctx, req.(*ProcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MonitorProcs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "agent.MonitorProcs",
	HandlerType: (*MonitorProcsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _MonitorProcs_Register_Handler,
		},
		{
			MethodName: "Deregister",
			Handler:    _MonitorProcs_Deregister_Handler,
		},
		{
			MethodName: "GetProcesses",
			Handler:    _MonitorProcs_GetProcesses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x14, 0xa4, 0x2d, 0xa5, 0xf0, 0x68, 0x40, 0x9f, 0xc6, 0x6c, 0x38, 0x98, 0xa6, 0x72, 0xe8, 0x45,
	0x62, 0x30, 0x7a, 0x27, 0x1a, 0x0d, 0x46, 0x0d, 0xe9, 0x1f, 0x14, 0xba, 0x90, 0x2a, 0xed, 0xae,
	0xbb, 0x6b, 0x0c, 0x07, 0xff, 0xc8, 0x8f, 0x34, 0xbb, 0xdd, 0x12, 0xbc, 0x78, 0xeb, 0xbc, 0x99,
	0x37, 0xb3, 0xf3, 0x52, 0xe8, 0x67, 0x1b, 0x5a, 0xa9, 0x09, 0x17, 0x4c, 0x31, 0xf4, 0x0d, 0x88,
	0xbf, 0x21, 0x98, 0x71, 0x3e, 0xaf, 0xd6, 0x0c, 0x07, 0xe0, 0xce, 0x73, 0xe2, 0x44, 0x4e, 0xd2,
	0x4b, 0xdd, 0x22, 0xc7, 0x73, 0x80, 0x07, 0x91, 0x95, 0xf4, 0x8b, 0x89, 0x77, 0x49, 0xdc, 0xc8,
	0x4b, 0x7a, 0x29, 0xac, 0xf7, 0x13, 0x3c, 0x83, 0xce, 0x13, 0x5b, 0xce, 0x73, 0x49, 0x3c, 0xc3,
	0x75, 0xde, 0x0c, 0x42, 0x02, 0xc1, 0x22, 0x53, 0x8a, 0x8a, 0x8a, 0xb4, 0x8d, 0x59, 0xc0, 0x6b,
	0x88, 0x47, 0xe0, 0x2d, 0x8a, 0x9c, 0xf8, 0x91, 0x93, 0xf8, 0xa9, 0xc7, 0x8b, 0x3c, 0x9e, 0x41,
	0xb0, 0x10, 0x6c, 0x45, 0xa5, 0x6c, 0x48, 0x67, 0x4f, 0x22, 0x42, 0xfb, 0x35, 0x2b, 0x29, 0x71,
	0x8d, 0x4b, 0xbb, 0xca, 0x4a, 0xaa, 0x55, 0x77, 0x65, 0x4e, 0x3c, 0x33, 0xf2, 0x56, 0x65, 0x1e,
	0x5f, 0x40, 0x5f, 0x5b, 0xa4, 0xf4, 0xe3, 0x93, 0x4a, 0x85, 0xa7, 0xe0, 0xeb, 0x42, 0x4d, 0x11,
	0x3f, 0xd3, 0x20, 0xbe, 0x82, 0xae, 0x16, 0x3d, 0x17, 0x52, 0xe1, 0x18, 0x7c, 0xfd, 0x2d, 0x89,
	0x13, 0x79, 0x49, 0x7f, 0x3a, 0x98, 0xd4, 0x67, 0xb1, 0xef, 0x48, 0x7d, 0xae, 0xc9, 0x18, 0xa0,
	0x9b, 0xd2, 0x4d, 0x4a, 0xf9, 0x76, 0x17, 0x8f, 0x21, 0xbc, 0xa7, 0x42, 0xa3, 0xff, 0x32, 0x42,
	0x00, 0xab, 0xe2, 0xdb, 0xdd, 0xf4, 0xc7, 0x81, 0xf0, 0x85, 0x55, 0x85, 0x62, 0xc2, 0xa4, 0xe1,
	0xa5, 0x31, 0x2c, 0xa4, 0xa2, 0x02, 0x9b, 0x4c, 0x7b, 0xfa, 0xd1, 0xd0, 0xe2, 0x7d, 0x62, 0x0b,
	0x6f, 0xad, 0x5b, 0xbd, 0x70, 0x62, 0x05, 0x87, 0xcf, 0x18, 0x1d, 0xff, 0x1d, 0xd6, 0x7b, 0x37,
	0x10, 0x3e, 0x52, 0x65, 0xcb, 0x50, 0x89, 0x78, 0x50, 0xaf, 0x59, 0x1c, 0x1e, 0xcc, 0xf4, 0x49,
	0xe2, 0xd6, 0xb2, 0x63, 0xfe, 0x8a, 0xeb, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb6, 0xeb, 0xeb,
	0xd9, 0x24, 0x02, 0x00, 0x00,
}
