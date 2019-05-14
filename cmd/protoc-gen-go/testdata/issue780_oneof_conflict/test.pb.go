// Code generated by protoc-gen-go. DO NOT EDIT.
// source: issue780_oneof_conflict/test.proto

package oneoftest

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoregistry "google.golang.org/protobuf/reflect/protoregistry"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

type Foo struct {
	// Types that are valid to be assigned to Bar:
	//	*Foo_GetBar
	Bar                  isFoo_Bar               `protobuf_oneof:"bar"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     protoimpl.UnknownFields `json:"-"`
	XXX_sizecache        protoimpl.SizeCache     `json:"-"`
}

func (x *Foo) Reset() {
	*x = Foo{}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	return file_issue780_oneof_conflict_test_proto_msgTypes[0].MessageOf(x)
}

func (m *Foo) XXX_Methods() *protoiface.Methods {
	return file_issue780_oneof_conflict_test_proto_msgTypes[0].Methods()
}

// Deprecated: Use Foo.ProtoReflect.Type instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_issue780_oneof_conflict_test_proto_rawDescGZIP(), []int{0}
}

func (m *Foo) GetBar() isFoo_Bar {
	if m != nil {
		return m.Bar
	}
	return nil
}

func (x *Foo) GetGetBar() string {
	if x, ok := x.GetBar().(*Foo_GetBar); ok {
		return x.GetBar
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Foo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Foo_GetBar)(nil),
	}
}

type isFoo_Bar interface {
	isFoo_Bar()
}

type Foo_GetBar struct {
	GetBar string `protobuf:"bytes,1,opt,name=get_bar,json=getBar,oneof"`
}

func (*Foo_GetBar) isFoo_Bar() {}

var File_issue780_oneof_conflict_test_proto protoreflect.FileDescriptor

var file_issue780_oneof_conflict_test_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x73, 0x73, 0x75, 0x65, 0x37, 0x38, 0x30, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6c, 0x69, 0x63, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x74, 0x65, 0x73, 0x74, 0x22,
	0x27, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x19, 0x0a, 0x07, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x61,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x67, 0x65, 0x74, 0x42, 0x61,
	0x72, 0x42, 0x05, 0x0a, 0x03, 0x62, 0x61, 0x72,
}

var (
	file_issue780_oneof_conflict_test_proto_rawDescOnce sync.Once
	file_issue780_oneof_conflict_test_proto_rawDescData = file_issue780_oneof_conflict_test_proto_rawDesc
)

func file_issue780_oneof_conflict_test_proto_rawDescGZIP() []byte {
	file_issue780_oneof_conflict_test_proto_rawDescOnce.Do(func() {
		file_issue780_oneof_conflict_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_issue780_oneof_conflict_test_proto_rawDescData)
	})
	return file_issue780_oneof_conflict_test_proto_rawDescData
}

var file_issue780_oneof_conflict_test_proto_msgTypes = make([]protoimpl.MessageType, 1)
var file_issue780_oneof_conflict_test_proto_goTypes = []interface{}{
	(*Foo)(nil), // 0: oneoftest.Foo
}
var file_issue780_oneof_conflict_test_proto_depIdxs = []int32{}

func init() { file_issue780_oneof_conflict_test_proto_init() }
func file_issue780_oneof_conflict_test_proto_init() {
	if File_issue780_oneof_conflict_test_proto != nil {
		return
	}
	File_issue780_oneof_conflict_test_proto = protoimpl.FileBuilder{
		RawDescriptor:      file_issue780_oneof_conflict_test_proto_rawDesc,
		GoTypes:            file_issue780_oneof_conflict_test_proto_goTypes,
		DependencyIndexes:  file_issue780_oneof_conflict_test_proto_depIdxs,
		MessageOutputTypes: file_issue780_oneof_conflict_test_proto_msgTypes,
		FilesRegistry:      protoregistry.GlobalFiles,
		TypesRegistry:      protoregistry.GlobalTypes,
	}.Init()
	file_issue780_oneof_conflict_test_proto_rawDesc = nil
	file_issue780_oneof_conflict_test_proto_goTypes = nil
	file_issue780_oneof_conflict_test_proto_depIdxs = nil
}
