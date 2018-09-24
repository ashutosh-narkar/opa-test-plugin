// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/appengine/v1/instance.proto

package appengine

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Availability of the instance.
type Instance_Availability int32

const (
	Instance_UNSPECIFIED Instance_Availability = 0
	Instance_RESIDENT    Instance_Availability = 1
	Instance_DYNAMIC     Instance_Availability = 2
)

var Instance_Availability_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "RESIDENT",
	2: "DYNAMIC",
}

var Instance_Availability_value = map[string]int32{
	"UNSPECIFIED": 0,
	"RESIDENT":    1,
	"DYNAMIC":     2,
}

func (x Instance_Availability) String() string {
	return proto.EnumName(Instance_Availability_name, int32(x))
}

func (Instance_Availability) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3b3f5aa565fc77c9, []int{0, 0}
}

// An Instance resource is the computing unit that App Engine uses to
// automatically scale an application.
type Instance struct {
	// Full path to the Instance resource in the API.
	// Example: `apps/myapp/services/default/versions/v1/instances/instance-1`.
	//
	// @OutputOnly
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Relative name of the instance within the version.
	// Example: `instance-1`.
	//
	// @OutputOnly
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// App Engine release this instance is running on.
	//
	// @OutputOnly
	AppEngineRelease string `protobuf:"bytes,3,opt,name=app_engine_release,json=appEngineRelease,proto3" json:"app_engine_release,omitempty"`
	// Availability of the instance.
	//
	// @OutputOnly
	Availability Instance_Availability `protobuf:"varint,4,opt,name=availability,proto3,enum=google.appengine.v1.Instance_Availability" json:"availability,omitempty"`
	// Name of the virtual machine where this instance lives. Only applicable
	// for instances in App Engine flexible environment.
	//
	// @OutputOnly
	VmName string `protobuf:"bytes,5,opt,name=vm_name,json=vmName,proto3" json:"vm_name,omitempty"`
	// Zone where the virtual machine is located. Only applicable for instances
	// in App Engine flexible environment.
	//
	// @OutputOnly
	VmZoneName string `protobuf:"bytes,6,opt,name=vm_zone_name,json=vmZoneName,proto3" json:"vm_zone_name,omitempty"`
	// Virtual machine ID of this instance. Only applicable for instances in
	// App Engine flexible environment.
	//
	// @OutputOnly
	VmId string `protobuf:"bytes,7,opt,name=vm_id,json=vmId,proto3" json:"vm_id,omitempty"`
	// Time that this instance was started.
	//
	// @OutputOnly
	StartTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Number of requests since this instance was started.
	//
	// @OutputOnly
	Requests int32 `protobuf:"varint,9,opt,name=requests,proto3" json:"requests,omitempty"`
	// Number of errors since this instance was started.
	//
	// @OutputOnly
	Errors int32 `protobuf:"varint,10,opt,name=errors,proto3" json:"errors,omitempty"`
	// Average queries per second (QPS) over the last minute.
	//
	// @OutputOnly
	Qps float32 `protobuf:"fixed32,11,opt,name=qps,proto3" json:"qps,omitempty"`
	// Average latency (ms) over the last minute.
	//
	// @OutputOnly
	AverageLatency int32 `protobuf:"varint,12,opt,name=average_latency,json=averageLatency,proto3" json:"average_latency,omitempty"`
	// Total memory in use (bytes).
	//
	// @OutputOnly
	MemoryUsage int64 `protobuf:"varint,13,opt,name=memory_usage,json=memoryUsage,proto3" json:"memory_usage,omitempty"`
	// Status of the virtual machine where this instance lives. Only applicable
	// for instances in App Engine flexible environment.
	//
	// @OutputOnly
	VmStatus string `protobuf:"bytes,14,opt,name=vm_status,json=vmStatus,proto3" json:"vm_status,omitempty"`
	// Whether this instance is in debug mode. Only applicable for instances in
	// App Engine flexible environment.
	//
	// @OutputOnly
	VmDebugEnabled       bool     `protobuf:"varint,15,opt,name=vm_debug_enabled,json=vmDebugEnabled,proto3" json:"vm_debug_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instance) Reset()         { *m = Instance{} }
func (m *Instance) String() string { return proto.CompactTextString(m) }
func (*Instance) ProtoMessage()    {}
func (*Instance) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b3f5aa565fc77c9, []int{0}
}

func (m *Instance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instance.Unmarshal(m, b)
}
func (m *Instance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instance.Marshal(b, m, deterministic)
}
func (m *Instance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instance.Merge(m, src)
}
func (m *Instance) XXX_Size() int {
	return xxx_messageInfo_Instance.Size(m)
}
func (m *Instance) XXX_DiscardUnknown() {
	xxx_messageInfo_Instance.DiscardUnknown(m)
}

var xxx_messageInfo_Instance proto.InternalMessageInfo

func (m *Instance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instance) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instance) GetAppEngineRelease() string {
	if m != nil {
		return m.AppEngineRelease
	}
	return ""
}

func (m *Instance) GetAvailability() Instance_Availability {
	if m != nil {
		return m.Availability
	}
	return Instance_UNSPECIFIED
}

func (m *Instance) GetVmName() string {
	if m != nil {
		return m.VmName
	}
	return ""
}

func (m *Instance) GetVmZoneName() string {
	if m != nil {
		return m.VmZoneName
	}
	return ""
}

func (m *Instance) GetVmId() string {
	if m != nil {
		return m.VmId
	}
	return ""
}

func (m *Instance) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Instance) GetRequests() int32 {
	if m != nil {
		return m.Requests
	}
	return 0
}

func (m *Instance) GetErrors() int32 {
	if m != nil {
		return m.Errors
	}
	return 0
}

func (m *Instance) GetQps() float32 {
	if m != nil {
		return m.Qps
	}
	return 0
}

func (m *Instance) GetAverageLatency() int32 {
	if m != nil {
		return m.AverageLatency
	}
	return 0
}

func (m *Instance) GetMemoryUsage() int64 {
	if m != nil {
		return m.MemoryUsage
	}
	return 0
}

func (m *Instance) GetVmStatus() string {
	if m != nil {
		return m.VmStatus
	}
	return ""
}

func (m *Instance) GetVmDebugEnabled() bool {
	if m != nil {
		return m.VmDebugEnabled
	}
	return false
}

func init() {
	proto.RegisterType((*Instance)(nil), "google.appengine.v1.Instance")
	proto.RegisterEnum("google.appengine.v1.Instance_Availability", Instance_Availability_name, Instance_Availability_value)
}

func init() { proto.RegisterFile("google/appengine/v1/instance.proto", fileDescriptor_3b3f5aa565fc77c9) }

var fileDescriptor_3b3f5aa565fc77c9 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x5d, 0x6b, 0xdb, 0x3c,
	0x14, 0x7e, 0x9d, 0xb6, 0xa9, 0x73, 0xe2, 0x26, 0x46, 0x85, 0xb7, 0x22, 0x1b, 0xcc, 0xcb, 0xcd,
	0xcc, 0x18, 0x36, 0xed, 0xae, 0xf6, 0x71, 0xd3, 0x36, 0x1e, 0x18, 0xb6, 0x10, 0x9c, 0xf6, 0x62,
	0xbd, 0x31, 0x4a, 0xac, 0x79, 0x02, 0x4b, 0x72, 0x2d, 0x45, 0x90, 0xfd, 0xc6, 0xfd, 0xa8, 0x61,
	0x39, 0x09, 0x2d, 0xf4, 0xce, 0xcf, 0xc7, 0x41, 0xcf, 0x79, 0x0e, 0x86, 0x69, 0x29, 0x65, 0x59,
	0xd1, 0x98, 0xd4, 0x35, 0x15, 0x25, 0x13, 0x34, 0x36, 0x97, 0x31, 0x13, 0x4a, 0x13, 0xb1, 0xa6,
	0x51, 0xdd, 0x48, 0x2d, 0xd1, 0x79, 0xe7, 0x89, 0x0e, 0x9e, 0xc8, 0x5c, 0x4e, 0x5e, 0x1f, 0x06,
	0x59, 0x4c, 0x84, 0x90, 0x9a, 0x68, 0x26, 0x85, 0xea, 0x46, 0x26, 0x6f, 0x76, 0xaa, 0x45, 0xab,
	0xcd, 0xaf, 0x58, 0x33, 0x4e, 0x95, 0x26, 0xbc, 0xee, 0x0c, 0xd3, 0xbf, 0xc7, 0xe0, 0xa6, 0xbb,
	0x67, 0x10, 0x82, 0x63, 0x41, 0x38, 0xc5, 0x4e, 0xe0, 0x84, 0x83, 0xcc, 0x7e, 0xa3, 0x11, 0xf4,
	0x58, 0x81, 0x7b, 0x96, 0xe9, 0xb1, 0x02, 0x7d, 0x00, 0x44, 0xea, 0x3a, 0xef, 0x02, 0xe4, 0x0d,
	0xad, 0x28, 0x51, 0x14, 0x1f, 0x59, 0xdd, 0x27, 0x75, 0x9d, 0x58, 0x21, 0xeb, 0x78, 0x34, 0x07,
	0x8f, 0x18, 0xc2, 0x2a, 0xb2, 0x62, 0x15, 0xd3, 0x5b, 0x7c, 0x1c, 0x38, 0xe1, 0xe8, 0xea, 0x7d,
	0xf4, 0xc2, 0x26, 0xd1, 0x3e, 0x46, 0x74, 0xfd, 0x64, 0x22, 0x7b, 0x36, 0x8f, 0x2e, 0xe0, 0xd4,
	0xf0, 0xdc, 0x86, 0x3c, 0xb1, 0x4f, 0xf6, 0x0d, 0x9f, 0xb7, 0x31, 0x03, 0xf0, 0x0c, 0xcf, 0xff,
	0x48, 0x41, 0x3b, 0xb5, 0x6f, 0x55, 0x30, 0xfc, 0x41, 0x0a, 0x6a, 0x1d, 0xe7, 0x70, 0x62, 0x78,
	0xce, 0x0a, 0x7c, 0xda, 0x6d, 0x67, 0x78, 0x5a, 0xa0, 0x4f, 0x00, 0x4a, 0x93, 0x46, 0xe7, 0x6d,
	0x2f, 0xd8, 0x0d, 0x9c, 0x70, 0x78, 0x35, 0xd9, 0xa7, 0xdb, 0x97, 0x16, 0xdd, 0xed, 0x4b, 0xcb,
	0x06, 0xd6, 0xdd, 0x62, 0x34, 0x01, 0xb7, 0xa1, 0x8f, 0x1b, 0xaa, 0xb4, 0xc2, 0x83, 0xc0, 0x09,
	0x4f, 0xb2, 0x03, 0x46, 0xff, 0x43, 0x9f, 0x36, 0x8d, 0x6c, 0x14, 0x06, 0xab, 0xec, 0x10, 0xf2,
	0xe1, 0xe8, 0xb1, 0x56, 0x78, 0x18, 0x38, 0x61, 0x2f, 0x6b, 0x3f, 0xd1, 0x3b, 0x18, 0x13, 0x43,
	0x1b, 0x52, 0xd2, 0xbc, 0x22, 0x9a, 0x8a, 0xf5, 0x16, 0x7b, 0x76, 0x64, 0xb4, 0xa3, 0xbf, 0x77,
	0x2c, 0x7a, 0x0b, 0x1e, 0xa7, 0x5c, 0x36, 0xdb, 0x7c, 0xa3, 0x48, 0x49, 0xf1, 0x59, 0xe0, 0x84,
	0x47, 0xd9, 0xb0, 0xe3, 0xee, 0x5b, 0x0a, 0xbd, 0x82, 0x81, 0xe1, 0xb9, 0xd2, 0x44, 0x6f, 0x14,
	0x1e, 0xd9, 0x2d, 0x5d, 0xc3, 0x97, 0x16, 0xa3, 0x10, 0x7c, 0xc3, 0xf3, 0x82, 0xae, 0x36, 0x65,
	0x4e, 0x05, 0x59, 0x55, 0xb4, 0xc0, 0xe3, 0xc0, 0x09, 0xdd, 0x6c, 0x64, 0xf8, 0xac, 0xa5, 0x93,
	0x8e, 0x9d, 0x7e, 0x06, 0xef, 0xe9, 0x05, 0xd0, 0x18, 0x86, 0xf7, 0xf3, 0xe5, 0x22, 0xb9, 0x4d,
	0xbf, 0xa5, 0xc9, 0xcc, 0xff, 0x0f, 0x79, 0xe0, 0x66, 0xc9, 0x32, 0x9d, 0x25, 0xf3, 0x3b, 0xdf,
	0x41, 0x43, 0x38, 0x9d, 0xfd, 0x9c, 0x5f, 0xff, 0x48, 0x6f, 0xfd, 0xde, 0xcd, 0x6f, 0xb8, 0x58,
	0x4b, 0xfe, 0xd2, 0x79, 0x6f, 0xce, 0xf6, 0xf7, 0x5d, 0xb4, 0xb5, 0x2e, 0x9c, 0x87, 0xaf, 0x3b,
	0x57, 0x29, 0x2b, 0x22, 0xca, 0x48, 0x36, 0x65, 0x5c, 0x52, 0x61, 0x4b, 0x8f, 0x3b, 0x89, 0xd4,
	0x4c, 0x3d, 0xfb, 0x23, 0xbe, 0x1c, 0xc0, 0xaa, 0x6f, 0x8d, 0x1f, 0xff, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x97, 0xe7, 0x7d, 0x88, 0x39, 0x03, 0x00, 0x00,
}
