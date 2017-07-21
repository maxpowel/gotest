// Code generated by protoc-gen-go. DO NOT EDIT.
// source: machinery.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TaskState int32

const (
	TaskState_UNKWNOWN TaskState = 0
	TaskState_PENDING  TaskState = 1
	TaskState_RECEIVED TaskState = 2
	TaskState_STARTED  TaskState = 3
	TaskState_RETRY    TaskState = 4
	TaskState_SUCCESS  TaskState = 5
	TaskState_FAILURE  TaskState = 6
)

var TaskState_name = map[int32]string{
	0: "UNKWNOWN",
	1: "PENDING",
	2: "RECEIVED",
	3: "STARTED",
	4: "RETRY",
	5: "SUCCESS",
	6: "FAILURE",
}
var TaskState_value = map[string]int32{
	"UNKWNOWN": 0,
	"PENDING":  1,
	"RECEIVED": 2,
	"STARTED":  3,
	"RETRY":    4,
	"SUCCESS":  5,
	"FAILURE":  6,
}

func (x TaskState) String() string {
	return proto1.EnumName(TaskState_name, int32(x))
}
func (TaskState) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type TaskStateResponse struct {
	State TaskState `protobuf:"varint,1,opt,name=State,enum=proto.TaskState" json:"State,omitempty"`
	ETA   int32     `protobuf:"varint,2,opt,name=ETA" json:"ETA,omitempty"`
	Uid   string    `protobuf:"bytes,3,opt,name=Uid" json:"Uid,omitempty"`
}

func (m *TaskStateResponse) Reset()                    { *m = TaskStateResponse{} }
func (m *TaskStateResponse) String() string            { return proto1.CompactTextString(m) }
func (*TaskStateResponse) ProtoMessage()               {}
func (*TaskStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TaskStateResponse) GetState() TaskState {
	if m != nil {
		return m.State
	}
	return TaskState_UNKWNOWN
}

func (m *TaskStateResponse) GetETA() int32 {
	if m != nil {
		return m.ETA
	}
	return 0
}

func (m *TaskStateResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func init() {
	proto1.RegisterType((*TaskStateResponse)(nil), "proto.TaskStateResponse")
	proto1.RegisterEnum("proto.TaskState", TaskState_name, TaskState_value)
}

func init() { proto1.RegisterFile("machinery.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4d, 0x4c, 0xce,
	0xc8, 0xcc, 0x4b, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a,
	0xf1, 0x5c, 0x82, 0x21, 0x89, 0xc5, 0xd9, 0xc1, 0x25, 0x89, 0x25, 0xa9, 0x41, 0xa9, 0xc5, 0x05,
	0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x6a, 0x5c, 0xac, 0x60, 0x01, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x3e,
	0x23, 0x01, 0x88, 0x16, 0x3d, 0x84, 0x42, 0x88, 0xb4, 0x90, 0x00, 0x17, 0xb3, 0x6b, 0x88, 0xa3,
	0x04, 0x93, 0x02, 0xa3, 0x06, 0x6b, 0x10, 0x88, 0x09, 0x12, 0x09, 0xcd, 0x4c, 0x91, 0x60, 0x56,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0xb5, 0xd2, 0xb8, 0x38, 0xe1, 0xfa, 0x84, 0x78, 0xb8, 0x38,
	0x42, 0xfd, 0xbc, 0xc3, 0xfd, 0xfc, 0xc3, 0xfd, 0x04, 0x18, 0x84, 0xb8, 0xb9, 0xd8, 0x03, 0x5c,
	0xfd, 0x5c, 0x3c, 0xfd, 0xdc, 0x05, 0x18, 0x41, 0x52, 0x41, 0xae, 0xce, 0xae, 0x9e, 0x61, 0xae,
	0x2e, 0x02, 0x4c, 0x20, 0xa9, 0xe0, 0x10, 0xc7, 0xa0, 0x10, 0x57, 0x17, 0x01, 0x66, 0x21, 0x4e,
	0x2e, 0xd6, 0x20, 0xd7, 0x90, 0xa0, 0x48, 0x01, 0x16, 0xb0, 0x78, 0xa8, 0xb3, 0xb3, 0x6b, 0x70,
	0xb0, 0x00, 0x2b, 0x88, 0xe3, 0xe6, 0xe8, 0xe9, 0x13, 0x1a, 0xe4, 0x2a, 0xc0, 0x96, 0xc4, 0x06,
	0x76, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xb0, 0x72, 0xea, 0xe9, 0x00, 0x00, 0x00,
}
