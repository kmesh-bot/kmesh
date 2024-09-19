// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: api/endpoint/endpoint.proto

package endpoint

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	core "kmesh.net/kmesh/api/v2/core"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address *core.SocketAddress `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_endpoint_endpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_api_endpoint_endpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_api_endpoint_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *Endpoint) GetAddress() *core.SocketAddress {
	if x != nil {
		return x.Address
	}
	return nil
}

type LocalityLbEndpoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LbEndpoints         []*Endpoint `protobuf:"bytes,1,rep,name=lb_endpoints,json=lbEndpoints,proto3" json:"lb_endpoints,omitempty"`
	LoadBalancingWeight uint32      `protobuf:"varint,3,opt,name=load_balancing_weight,json=loadBalancingWeight,proto3" json:"load_balancing_weight,omitempty"`
	Priority            uint32      `protobuf:"varint,5,opt,name=priority,proto3" json:"priority,omitempty"`
	ConnectNum          uint32      `protobuf:"varint,11,opt,name=connect_num,json=connectNum,proto3" json:"connect_num,omitempty"`
}

func (x *LocalityLbEndpoints) Reset() {
	*x = LocalityLbEndpoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_endpoint_endpoint_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocalityLbEndpoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocalityLbEndpoints) ProtoMessage() {}

func (x *LocalityLbEndpoints) ProtoReflect() protoreflect.Message {
	mi := &file_api_endpoint_endpoint_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocalityLbEndpoints.ProtoReflect.Descriptor instead.
func (*LocalityLbEndpoints) Descriptor() ([]byte, []int) {
	return file_api_endpoint_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *LocalityLbEndpoints) GetLbEndpoints() []*Endpoint {
	if x != nil {
		return x.LbEndpoints
	}
	return nil
}

func (x *LocalityLbEndpoints) GetLoadBalancingWeight() uint32 {
	if x != nil {
		return x.LoadBalancingWeight
	}
	return 0
}

func (x *LocalityLbEndpoints) GetPriority() uint32 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *LocalityLbEndpoints) GetConnectNum() uint32 {
	if x != nil {
		return x.ConnectNum
	}
	return 0
}

type ClusterLoadAssignment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClusterName string                 `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	Endpoints   []*LocalityLbEndpoints `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
}

func (x *ClusterLoadAssignment) Reset() {
	*x = ClusterLoadAssignment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_endpoint_endpoint_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterLoadAssignment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterLoadAssignment) ProtoMessage() {}

func (x *ClusterLoadAssignment) ProtoReflect() protoreflect.Message {
	mi := &file_api_endpoint_endpoint_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterLoadAssignment.ProtoReflect.Descriptor instead.
func (*ClusterLoadAssignment) Descriptor() ([]byte, []int) {
	return file_api_endpoint_endpoint_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterLoadAssignment) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

func (x *ClusterLoadAssignment) GetEndpoints() []*LocalityLbEndpoints {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

var File_api_endpoint_endpoint_proto protoreflect.FileDescriptor

var file_api_endpoint_endpoint_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x39, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x13, 0x4c,
	0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x35, 0x0a, 0x0c, 0x6c, 0x62, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0b, 0x6c, 0x62,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x15, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x77, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4e, 0x75, 0x6d, 0x22, 0x77, 0x0a, 0x15, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x61, 0x64, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x65, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x4c, 0x62, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x6b, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x6e, 0x65, 0x74,
	0x2f, 0x6b, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x3b, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_endpoint_endpoint_proto_rawDescOnce sync.Once
	file_api_endpoint_endpoint_proto_rawDescData = file_api_endpoint_endpoint_proto_rawDesc
)

func file_api_endpoint_endpoint_proto_rawDescGZIP() []byte {
	file_api_endpoint_endpoint_proto_rawDescOnce.Do(func() {
		file_api_endpoint_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_endpoint_endpoint_proto_rawDescData)
	})
	return file_api_endpoint_endpoint_proto_rawDescData
}

var file_api_endpoint_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_endpoint_endpoint_proto_goTypes = []any{
	(*Endpoint)(nil),              // 0: endpoint.Endpoint
	(*LocalityLbEndpoints)(nil),   // 1: endpoint.LocalityLbEndpoints
	(*ClusterLoadAssignment)(nil), // 2: endpoint.ClusterLoadAssignment
	(*core.SocketAddress)(nil),    // 3: core.SocketAddress
}
var file_api_endpoint_endpoint_proto_depIdxs = []int32{
	3, // 0: endpoint.Endpoint.address:type_name -> core.SocketAddress
	0, // 1: endpoint.LocalityLbEndpoints.lb_endpoints:type_name -> endpoint.Endpoint
	1, // 2: endpoint.ClusterLoadAssignment.endpoints:type_name -> endpoint.LocalityLbEndpoints
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_endpoint_endpoint_proto_init() }
func file_api_endpoint_endpoint_proto_init() {
	if File_api_endpoint_endpoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_endpoint_endpoint_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Endpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_endpoint_endpoint_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LocalityLbEndpoints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_endpoint_endpoint_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ClusterLoadAssignment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_endpoint_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_endpoint_endpoint_proto_goTypes,
		DependencyIndexes: file_api_endpoint_endpoint_proto_depIdxs,
		MessageInfos:      file_api_endpoint_endpoint_proto_msgTypes,
	}.Build()
	File_api_endpoint_endpoint_proto = out.File
	file_api_endpoint_endpoint_proto_rawDesc = nil
	file_api_endpoint_endpoint_proto_goTypes = nil
	file_api_endpoint_endpoint_proto_depIdxs = nil
}
