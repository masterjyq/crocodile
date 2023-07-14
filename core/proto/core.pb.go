// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/proto/core.proto

package crocodile_task

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

// task req
type TaskReq struct {
	TaskId               string   `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TaskType             int32    `protobuf:"varint,2,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	TaskData             []byte   `protobuf:"bytes,3,opt,name=task_data,json=taskData,proto3" json:"task_data,omitempty"`
	Params               []string `protobuf:"bytes,4,rep,name=params,json=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskReq) Reset()         { *m = TaskReq{} }
func (m *TaskReq) String() string { return proto.CompactTextString(m) }
func (*TaskReq) ProtoMessage()    {}
func (*TaskReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{0}
}

func (m *TaskReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskReq.Unmarshal(m, b)
}
func (m *TaskReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskReq.Marshal(b, m, deterministic)
}
func (m *TaskReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskReq.Merge(m, src)
}
func (m *TaskReq) XXX_Size() int {
	return xxx_messageInfo_TaskReq.Size(m)
}
func (m *TaskReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskReq.DiscardUnknown(m)
}

var xxx_messageInfo_TaskReq proto.InternalMessageInfo

func (m *TaskReq) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskReq) GetTaskType() int32 {
	if m != nil {
		return m.TaskType
	}
	return 0
}

func (m *TaskReq) GetTaskData() []byte {
	if m != nil {
		return m.TaskData
	}
	return nil
}

func (m *TaskReq) GetParams() []string {
	if m != nil {
		return m.Params
	}
	return nil
}

// task reso stream
type TaskResp struct {
	Resp                 []byte   `protobuf:"bytes,3,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskResp) Reset()         { *m = TaskResp{} }
func (m *TaskResp) String() string { return proto.CompactTextString(m) }
func (*TaskResp) ProtoMessage()    {}
func (*TaskResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{1}
}

func (m *TaskResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResp.Unmarshal(m, b)
}
func (m *TaskResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResp.Marshal(b, m, deterministic)
}
func (m *TaskResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResp.Merge(m, src)
}
func (m *TaskResp) XXX_Size() int {
	return xxx_messageInfo_TaskResp.Size(m)
}
func (m *TaskResp) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResp.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResp proto.InternalMessageInfo

func (m *TaskResp) GetResp() []byte {
	if m != nil {
		return m.Resp
	}
	return nil
}

type TaskRespOld struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	ErrMsg               []byte   `protobuf:"bytes,2,opt,name=err_msg,json=errMsg,proto3" json:"err_msg,omitempty"`
	RespData             []byte   `protobuf:"bytes,3,opt,name=resp_data,json=respData,proto3" json:"resp_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskRespOld) Reset()         { *m = TaskRespOld{} }
func (m *TaskRespOld) String() string { return proto.CompactTextString(m) }
func (*TaskRespOld) ProtoMessage()    {}
func (*TaskRespOld) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{2}
}

func (m *TaskRespOld) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskRespOld.Unmarshal(m, b)
}
func (m *TaskRespOld) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskRespOld.Marshal(b, m, deterministic)
}
func (m *TaskRespOld) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskRespOld.Merge(m, src)
}
func (m *TaskRespOld) XXX_Size() int {
	return xxx_messageInfo_TaskRespOld.Size(m)
}
func (m *TaskRespOld) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskRespOld.DiscardUnknown(m)
}

var xxx_messageInfo_TaskRespOld proto.InternalMessageInfo

func (m *TaskRespOld) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TaskRespOld) GetErrMsg() []byte {
	if m != nil {
		return m.ErrMsg
	}
	return nil
}

func (m *TaskRespOld) GetRespData() []byte {
	if m != nil {
		return m.RespData
	}
	return nil
}

type RegistryReq struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Weight               int32    `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Hostname             string   `protobuf:"bytes,4,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Version              string   `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Hostgroup            string   `protobuf:"bytes,6,opt,name=hostgroup,proto3" json:"hostgroup,omitempty"`
	Remark               string   `protobuf:"bytes,7,opt,name=remark,proto3" json:"remark,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegistryReq) Reset()         { *m = RegistryReq{} }
func (m *RegistryReq) String() string { return proto.CompactTextString(m) }
func (*RegistryReq) ProtoMessage()    {}
func (*RegistryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{3}
}

func (m *RegistryReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegistryReq.Unmarshal(m, b)
}
func (m *RegistryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegistryReq.Marshal(b, m, deterministic)
}
func (m *RegistryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegistryReq.Merge(m, src)
}
func (m *RegistryReq) XXX_Size() int {
	return xxx_messageInfo_RegistryReq.Size(m)
}
func (m *RegistryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegistryReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegistryReq proto.InternalMessageInfo

func (m *RegistryReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *RegistryReq) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RegistryReq) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *RegistryReq) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *RegistryReq) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *RegistryReq) GetHostgroup() string {
	if m != nil {
		return m.Hostgroup
	}
	return ""
}

func (m *RegistryReq) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

type HeartbeatReq struct {
	// string ip = 1;
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	RunningTask          []string `protobuf:"bytes,3,rep,name=running_task,json=runningTask,proto3" json:"running_task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatReq) Reset()         { *m = HeartbeatReq{} }
func (m *HeartbeatReq) String() string { return proto.CompactTextString(m) }
func (*HeartbeatReq) ProtoMessage()    {}
func (*HeartbeatReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{4}
}

func (m *HeartbeatReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatReq.Unmarshal(m, b)
}
func (m *HeartbeatReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatReq.Marshal(b, m, deterministic)
}
func (m *HeartbeatReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatReq.Merge(m, src)
}
func (m *HeartbeatReq) XXX_Size() int {
	return xxx_messageInfo_HeartbeatReq.Size(m)
}
func (m *HeartbeatReq) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatReq.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatReq proto.InternalMessageInfo

func (m *HeartbeatReq) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *HeartbeatReq) GetRunningTask() []string {
	if m != nil {
		return m.RunningTask
	}
	return nil
}

func (m *HeartbeatReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_80ea9561f1d738ba, []int{5}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TaskReq)(nil), "crocodile.task.TaskReq")
	proto.RegisterType((*TaskResp)(nil), "crocodile.task.TaskResp")
	proto.RegisterType((*TaskRespOld)(nil), "crocodile.task.TaskRespOld")
	proto.RegisterType((*RegistryReq)(nil), "crocodile.task.RegistryReq")
	proto.RegisterType((*HeartbeatReq)(nil), "crocodile.task.HeartbeatReq")
	proto.RegisterType((*Empty)(nil), "crocodile.task.Empty")
}

func init() { proto.RegisterFile("core/proto/core.proto", fileDescriptor_80ea9561f1d738ba) }

var fileDescriptor_80ea9561f1d738ba = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0x6c, 0xc7, 0x13, 0xab, 0x87, 0x95, 0x4a, 0x57, 0x69, 0x85, 0x82, 0x4f, 0x39,
	0xa5, 0xa8, 0xdc, 0xe1, 0xd2, 0xa2, 0x70, 0x40, 0x48, 0x4b, 0x25, 0x2e, 0x48, 0xd1, 0x26, 0x1e,
	0xb9, 0x56, 0x1a, 0xef, 0xb2, 0xbb, 0x01, 0xe5, 0x37, 0xf8, 0x17, 0xfe, 0x0f, 0xcd, 0xd8, 0x09,
	0x6d, 0x94, 0xde, 0x66, 0xde, 0x1b, 0xcf, 0x7b, 0xf3, 0xbc, 0x70, 0xbe, 0x32, 0x0e, 0xaf, 0xad,
	0x33, 0xc1, 0x5c, 0x53, 0x39, 0xe3, 0x52, 0x9c, 0xad, 0x9c, 0x59, 0x99, 0xb2, 0x7e, 0xc4, 0x59,
	0xd0, 0x7e, 0x5d, 0xfc, 0x80, 0xf4, 0x5e, 0xfb, 0xb5, 0xc2, 0x9f, 0xe2, 0x02, 0x52, 0x82, 0x16,
	0x75, 0x29, 0xa3, 0x49, 0x34, 0xcd, 0x54, 0x42, 0xed, 0xe7, 0x52, 0x5c, 0x42, 0xc6, 0x44, 0xd8,
	0x59, 0x94, 0xbd, 0x49, 0x34, 0x8d, 0xd5, 0x90, 0x80, 0xfb, 0x9d, 0xc5, 0x03, 0x59, 0xea, 0xa0,
	0x65, 0x7f, 0x12, 0x4d, 0xf3, 0x96, 0xbc, 0xd5, 0x41, 0x17, 0x6f, 0x60, 0xd8, 0x6e, 0xf7, 0x56,
	0x08, 0x18, 0x38, 0xf4, 0xb6, 0x9b, 0xe1, 0xba, 0xf8, 0x0e, 0xa3, 0x3d, 0xff, 0xf5, 0xb1, 0xa4,
	0x91, 0x95, 0x29, 0x91, 0xe5, 0x63, 0xc5, 0x35, 0xb9, 0x42, 0xe7, 0x16, 0x1b, 0x5f, 0xb1, 0x74,
	0xae, 0x12, 0x74, 0xee, 0x8b, 0xaf, 0x48, 0x98, 0x76, 0x3c, 0x13, 0x26, 0x80, 0x85, 0xff, 0x46,
	0x30, 0x52, 0x58, 0xd5, 0x3e, 0xb8, 0x1d, 0xdd, 0x76, 0x06, 0xbd, 0xda, 0x76, 0x67, 0xf5, 0x6a,
	0x36, 0x63, 0x8d, 0x0b, 0xdd, 0x35, 0x5c, 0x8b, 0xd7, 0x90, 0xfc, 0xc6, 0xba, 0x7a, 0x08, 0xbc,
	0x2d, 0x56, 0x5d, 0x27, 0xc6, 0x30, 0x7c, 0x30, 0x3e, 0x34, 0x7a, 0x83, 0x72, 0xc0, 0x1b, 0x0e,
	0xbd, 0x90, 0x90, 0xfe, 0x42, 0xe7, 0x6b, 0xd3, 0xc8, 0x98, 0xa9, 0x7d, 0x2b, 0xae, 0x20, 0xa3,
	0xa9, 0xca, 0x99, 0xad, 0x95, 0x09, 0x73, 0xff, 0x01, 0xd2, 0x72, 0xb8, 0xd1, 0x6e, 0x2d, 0xd3,
	0x36, 0xea, 0xb6, 0x2b, 0xee, 0x20, 0x9f, 0xa3, 0x76, 0x61, 0x89, 0x3a, 0x90, 0xef, 0x53, 0x3e,
	0xdf, 0x42, 0xee, 0xb6, 0x4d, 0x53, 0x37, 0xd5, 0x82, 0x82, 0x96, 0xfd, 0x49, 0x7f, 0x9a, 0xa9,
	0x51, 0x87, 0x51, 0x9e, 0x45, 0x0a, 0xf1, 0xdd, 0xc6, 0x86, 0xdd, 0xcd, 0x27, 0x18, 0x10, 0x20,
	0x3e, 0x40, 0xaa, 0xb6, 0x0d, 0x97, 0x17, 0xb3, 0xe7, 0x4f, 0x60, 0xd6, 0xfd, 0xff, 0xb1, 0x3c,
	0x4d, 0x78, 0xfb, 0x2e, 0xba, 0xf9, 0x13, 0x41, 0x76, 0x30, 0x26, 0x6e, 0x21, 0xdf, 0x87, 0x3b,
	0x37, 0x3e, 0x88, 0xcb, 0xe3, 0x2f, 0x9f, 0x44, 0x3f, 0x3e, 0x3f, 0x26, 0xd9, 0x59, 0xf1, 0x4a,
	0x7c, 0x84, 0xe4, 0x1b, 0x36, 0xe5, 0x7c, 0x29, 0xae, 0x8e, 0x47, 0x9e, 0x66, 0xf0, 0xe2, 0x82,
	0x65, 0xc2, 0x4f, 0xfa, 0xfd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xc4, 0xc1, 0x9d, 0xeb,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TaskClient is the client API for Task service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TaskClient interface {
	// run task return stream bytes
	RunTask(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (Task_RunTaskClient, error)
}

type taskClient struct {
	cc *grpc.ClientConn
}

func NewTaskClient(cc *grpc.ClientConn) TaskClient {
	return &taskClient{cc}
}

func (c *taskClient) RunTask(ctx context.Context, in *TaskReq, opts ...grpc.CallOption) (Task_RunTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Task_serviceDesc.Streams[0], "/crocodile.task.Task/RunTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskRunTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Task_RunTaskClient interface {
	Recv() (*TaskResp, error)
	grpc.ClientStream
}

type taskRunTaskClient struct {
	grpc.ClientStream
}

func (x *taskRunTaskClient) Recv() (*TaskResp, error) {
	m := new(TaskResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TaskServer is the server API for Task service.
type TaskServer interface {
	// run task return stream bytes
	RunTask(*TaskReq, Task_RunTaskServer) error
}

// UnimplementedTaskServer can be embedded to have forward compatible implementations.
type UnimplementedTaskServer struct {
}

func (*UnimplementedTaskServer) RunTask(req *TaskReq, srv Task_RunTaskServer) error {
	return status.Errorf(codes.Unimplemented, "method RunTask not implemented")
}

func RegisterTaskServer(s *grpc.Server, srv TaskServer) {
	s.RegisterService(&_Task_serviceDesc, srv)
}

func _Task_RunTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TaskReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskServer).RunTask(m, &taskRunTaskServer{stream})
}

type Task_RunTaskServer interface {
	Send(*TaskResp) error
	grpc.ServerStream
}

type taskRunTaskServer struct {
	grpc.ServerStream
}

func (x *taskRunTaskServer) Send(m *TaskResp) error {
	return x.ServerStream.SendMsg(m)
}

var _Task_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crocodile.task.Task",
	HandlerType: (*TaskServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunTask",
			Handler:       _Task_RunTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "core/proto/core.proto",
}

// HeartbeatClient is the client API for Heartbeat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HeartbeatClient interface {
	// registry host
	RegistryHost(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*Empty, error)
	// SendHb send to server req to itself alive
	SendHb(ctx context.Context, in *HeartbeatReq, opts ...grpc.CallOption) (*Empty, error)
}

type heartbeatClient struct {
	cc *grpc.ClientConn
}

func NewHeartbeatClient(cc *grpc.ClientConn) HeartbeatClient {
	return &heartbeatClient{cc}
}

func (c *heartbeatClient) RegistryHost(ctx context.Context, in *RegistryReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/crocodile.task.Heartbeat/RegistryHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *heartbeatClient) SendHb(ctx context.Context, in *HeartbeatReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/crocodile.task.Heartbeat/SendHb", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServer is the server API for Heartbeat service.
type HeartbeatServer interface {
	// registry host
	RegistryHost(context.Context, *RegistryReq) (*Empty, error)
	// SendHb send to server req to itself alive
	SendHb(context.Context, *HeartbeatReq) (*Empty, error)
}

// UnimplementedHeartbeatServer can be embedded to have forward compatible implementations.
type UnimplementedHeartbeatServer struct {
}

func (*UnimplementedHeartbeatServer) RegistryHost(ctx context.Context, req *RegistryReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegistryHost not implemented")
}
func (*UnimplementedHeartbeatServer) SendHb(ctx context.Context, req *HeartbeatReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHb not implemented")
}

func RegisterHeartbeatServer(s *grpc.Server, srv HeartbeatServer) {
	s.RegisterService(&_Heartbeat_serviceDesc, srv)
}

func _Heartbeat_RegistryHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegistryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).RegistryHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crocodile.task.Heartbeat/RegistryHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).RegistryHost(ctx, req.(*RegistryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Heartbeat_SendHb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServer).SendHb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/crocodile.task.Heartbeat/SendHb",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServer).SendHb(ctx, req.(*HeartbeatReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Heartbeat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "crocodile.task.Heartbeat",
	HandlerType: (*HeartbeatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegistryHost",
			Handler:    _Heartbeat_RegistryHost_Handler,
		},
		{
			MethodName: "SendHb",
			Handler:    _Heartbeat_SendHb_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "core/proto/core.proto",
}
