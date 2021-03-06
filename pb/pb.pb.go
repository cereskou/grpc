// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: pb.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

//Empty -
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{0}
}

//BrowseInfo -
type BrowseInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Initialdir string `protobuf:"bytes,2,opt,name=initialdir,proto3" json:"initialdir,omitempty"`
}

func (x *BrowseInfo) Reset() {
	*x = BrowseInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BrowseInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrowseInfo) ProtoMessage() {}

func (x *BrowseInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrowseInfo.ProtoReflect.Descriptor instead.
func (*BrowseInfo) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{1}
}

func (x *BrowseInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *BrowseInfo) GetInitialdir() string {
	if x != nil {
		return x.Initialdir
	}
	return ""
}

//Folder - BrowseForFolder response
type Folder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Folder) Reset() {
	*x = Folder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Folder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Folder) ProtoMessage() {}

func (x *Folder) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Folder.ProtoReflect.Descriptor instead.
func (*Folder) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{2}
}

func (x *Folder) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

//AliveResponse - alive response
type AliveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processid int32  `protobuf:"varint,1,opt,name=processid,proto3" json:"processid,omitempty"`
	Homedir   string `protobuf:"bytes,2,opt,name=homedir,proto3" json:"homedir,omitempty"`
	Accesskey string `protobuf:"bytes,3,opt,name=accesskey,proto3" json:"accesskey,omitempty"`
	Secretkey string `protobuf:"bytes,4,opt,name=secretkey,proto3" json:"secretkey,omitempty"`
	Region    string `protobuf:"bytes,5,opt,name=region,proto3" json:"region,omitempty"`
}

func (x *AliveResponse) Reset() {
	*x = AliveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AliveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AliveResponse) ProtoMessage() {}

func (x *AliveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AliveResponse.ProtoReflect.Descriptor instead.
func (*AliveResponse) Descriptor() ([]byte, []int) {
	return file_pb_proto_rawDescGZIP(), []int{3}
}

func (x *AliveResponse) GetProcessid() int32 {
	if x != nil {
		return x.Processid
	}
	return 0
}

func (x *AliveResponse) GetHomedir() string {
	if x != nil {
		return x.Homedir
	}
	return ""
}

func (x *AliveResponse) GetAccesskey() string {
	if x != nil {
		return x.Accesskey
	}
	return ""
}

func (x *AliveResponse) GetSecretkey() string {
	if x != nil {
		return x.Secretkey
	}
	return ""
}

func (x *AliveResponse) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

var File_pb_proto protoreflect.FileDescriptor

var file_pb_proto_rawDesc = []byte{
	0x0a, 0x08, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x42, 0x0a, 0x0a, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x61, 0x6c, 0x64, 0x69, 0x72, 0x22, 0x1c, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x9b, 0x01, 0x0a, 0x0d, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x6f, 0x6d, 0x65, 0x64, 0x69, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x68, 0x6f, 0x6d, 0x65, 0x64, 0x69, 0x72, 0x12,
	0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6b, 0x65, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x32, 0x5c, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x0f, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x46, 0x6f, 0x72,
	0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x42, 0x72, 0x6f, 0x77, 0x73, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x1a, 0x07, 0x2e, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x22, 0x00, 0x12, 0x21,
	0x0a, 0x05, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0e, 0x2e, 0x41, 0x6c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_proto_rawDescOnce sync.Once
	file_pb_proto_rawDescData = file_pb_proto_rawDesc
)

func file_pb_proto_rawDescGZIP() []byte {
	file_pb_proto_rawDescOnce.Do(func() {
		file_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_proto_rawDescData)
	})
	return file_pb_proto_rawDescData
}

var file_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_proto_goTypes = []interface{}{
	(*Empty)(nil),         // 0: Empty
	(*BrowseInfo)(nil),    // 1: BrowseInfo
	(*Folder)(nil),        // 2: Folder
	(*AliveResponse)(nil), // 3: AliveResponse
}
var file_pb_proto_depIdxs = []int32{
	1, // 0: LocalService.BrowseForFolder:input_type -> BrowseInfo
	0, // 1: LocalService.Alive:input_type -> Empty
	2, // 2: LocalService.BrowseForFolder:output_type -> Folder
	3, // 3: LocalService.Alive:output_type -> AliveResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_proto_init() }
func file_pb_proto_init() {
	if File_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BrowseInfo); i {
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
		file_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Folder); i {
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
		file_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AliveResponse); i {
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
			RawDescriptor: file_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_proto_goTypes,
		DependencyIndexes: file_pb_proto_depIdxs,
		MessageInfos:      file_pb_proto_msgTypes,
	}.Build()
	File_pb_proto = out.File
	file_pb_proto_rawDesc = nil
	file_pb_proto_goTypes = nil
	file_pb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LocalServiceClient is the client API for LocalService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocalServiceClient interface {
	BrowseForFolder(ctx context.Context, in *BrowseInfo, opts ...grpc.CallOption) (*Folder, error)
	Alive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AliveResponse, error)
}

type localServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocalServiceClient(cc grpc.ClientConnInterface) LocalServiceClient {
	return &localServiceClient{cc}
}

func (c *localServiceClient) BrowseForFolder(ctx context.Context, in *BrowseInfo, opts ...grpc.CallOption) (*Folder, error) {
	out := new(Folder)
	err := c.cc.Invoke(ctx, "/LocalService/BrowseForFolder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localServiceClient) Alive(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AliveResponse, error) {
	out := new(AliveResponse)
	err := c.cc.Invoke(ctx, "/LocalService/Alive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalServiceServer is the server API for LocalService service.
type LocalServiceServer interface {
	BrowseForFolder(context.Context, *BrowseInfo) (*Folder, error)
	Alive(context.Context, *Empty) (*AliveResponse, error)
}

// UnimplementedLocalServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLocalServiceServer struct {
}

func (*UnimplementedLocalServiceServer) BrowseForFolder(context.Context, *BrowseInfo) (*Folder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BrowseForFolder not implemented")
}
func (*UnimplementedLocalServiceServer) Alive(context.Context, *Empty) (*AliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Alive not implemented")
}

func RegisterLocalServiceServer(s *grpc.Server, srv LocalServiceServer) {
	s.RegisterService(&_LocalService_serviceDesc, srv)
}

func _LocalService_BrowseForFolder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BrowseInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServiceServer).BrowseForFolder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LocalService/BrowseForFolder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServiceServer).BrowseForFolder(ctx, req.(*BrowseInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalService_Alive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalServiceServer).Alive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LocalService/Alive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalServiceServer).Alive(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _LocalService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LocalService",
	HandlerType: (*LocalServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BrowseForFolder",
			Handler:    _LocalService_BrowseForFolder_Handler,
		},
		{
			MethodName: "Alive",
			Handler:    _LocalService_Alive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb.proto",
}
