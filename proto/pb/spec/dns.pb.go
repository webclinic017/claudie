// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: spec/dns.proto

package spec

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// DNS holds general information about the DNS records.
type DNS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DNS zone for the DNS records.
	DnsZone string `protobuf:"bytes,1,opt,name=dnsZone,proto3" json:"dnsZone,omitempty"`
	// User specified hostname. (optional)
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Provider for the DNS records.
	Provider *Provider `protobuf:"bytes,3,opt,name=provider,proto3" json:"provider,omitempty"`
	// The whole hostname of the DNS record.
	Endpoint string `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
}

func (x *DNS) Reset() {
	*x = DNS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spec_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNS) ProtoMessage() {}

func (x *DNS) ProtoReflect() protoreflect.Message {
	mi := &file_spec_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNS.ProtoReflect.Descriptor instead.
func (*DNS) Descriptor() ([]byte, []int) {
	return file_spec_dns_proto_rawDescGZIP(), []int{0}
}

func (x *DNS) GetDnsZone() string {
	if x != nil {
		return x.DnsZone
	}
	return ""
}

func (x *DNS) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *DNS) GetProvider() *Provider {
	if x != nil {
		return x.Provider
	}
	return nil
}

func (x *DNS) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

var File_spec_dns_proto protoreflect.FileDescriptor

var file_spec_dns_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x73, 0x70, 0x65, 0x63, 0x1a, 0x13, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x03,
	0x44, 0x4e, 0x53, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x70,
	0x65, 0x63, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x75, 0x64, 0x69, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spec_dns_proto_rawDescOnce sync.Once
	file_spec_dns_proto_rawDescData = file_spec_dns_proto_rawDesc
)

func file_spec_dns_proto_rawDescGZIP() []byte {
	file_spec_dns_proto_rawDescOnce.Do(func() {
		file_spec_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_spec_dns_proto_rawDescData)
	})
	return file_spec_dns_proto_rawDescData
}

var file_spec_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_spec_dns_proto_goTypes = []interface{}{
	(*DNS)(nil),      // 0: spec.DNS
	(*Provider)(nil), // 1: spec.Provider
}
var file_spec_dns_proto_depIdxs = []int32{
	1, // 0: spec.DNS.provider:type_name -> spec.Provider
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spec_dns_proto_init() }
func file_spec_dns_proto_init() {
	if File_spec_dns_proto != nil {
		return
	}
	file_spec_provider_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_spec_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNS); i {
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
			RawDescriptor: file_spec_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_spec_dns_proto_goTypes,
		DependencyIndexes: file_spec_dns_proto_depIdxs,
		MessageInfos:      file_spec_dns_proto_msgTypes,
	}.Build()
	File_spec_dns_proto = out.File
	file_spec_dns_proto_rawDesc = nil
	file_spec_dns_proto_goTypes = nil
	file_spec_dns_proto_depIdxs = nil
}