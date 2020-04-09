// Code generated by protoc-gen-go. DO NOT EDIT.
// source: processor_nhan_vien.proto

package pb

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

type NhanVien struct {
	DmNhanVienId         uint32   `protobuf:"varint,1,opt,name=dm_nhan_vien_id,json=dmNhanVienId,proto3" json:"dm_nhan_vien_id,omitempty"`
	MaNhanVien           string   `protobuf:"bytes,2,opt,name=ma_nhan_vien,json=maNhanVien,proto3" json:"ma_nhan_vien,omitempty"`
	TenNhanVien          string   `protobuf:"bytes,3,opt,name=ten_nhan_vien,json=tenNhanVien,proto3" json:"ten_nhan_vien,omitempty"`
	FullName             string   `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	LastName             string   `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PhongId              uint32   `protobuf:"varint,6,opt,name=phong_id,json=phongId,proto3" json:"phong_id,omitempty"`
	DmChucDanhId         uint32   `protobuf:"varint,7,opt,name=dm_chuc_danh_id,json=dmChucDanhId,proto3" json:"dm_chuc_danh_id,omitempty"`
	DmChucVuId           uint32   `protobuf:"varint,8,opt,name=dm_chuc_vu_id,json=dmChucVuId,proto3" json:"dm_chuc_vu_id,omitempty"`
	SoDienThoai          string   `protobuf:"bytes,9,opt,name=so_dien_thoai,json=soDienThoai,proto3" json:"so_dien_thoai,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NhanVien) Reset()         { *m = NhanVien{} }
func (m *NhanVien) String() string { return proto.CompactTextString(m) }
func (*NhanVien) ProtoMessage()    {}
func (*NhanVien) Descriptor() ([]byte, []int) {
	return fileDescriptor_c946787ad119af67, []int{0}
}

func (m *NhanVien) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NhanVien.Unmarshal(m, b)
}
func (m *NhanVien) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NhanVien.Marshal(b, m, deterministic)
}
func (m *NhanVien) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NhanVien.Merge(m, src)
}
func (m *NhanVien) XXX_Size() int {
	return xxx_messageInfo_NhanVien.Size(m)
}
func (m *NhanVien) XXX_DiscardUnknown() {
	xxx_messageInfo_NhanVien.DiscardUnknown(m)
}

var xxx_messageInfo_NhanVien proto.InternalMessageInfo

func (m *NhanVien) GetDmNhanVienId() uint32 {
	if m != nil {
		return m.DmNhanVienId
	}
	return 0
}

func (m *NhanVien) GetMaNhanVien() string {
	if m != nil {
		return m.MaNhanVien
	}
	return ""
}

func (m *NhanVien) GetTenNhanVien() string {
	if m != nil {
		return m.TenNhanVien
	}
	return ""
}

func (m *NhanVien) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *NhanVien) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *NhanVien) GetPhongId() uint32 {
	if m != nil {
		return m.PhongId
	}
	return 0
}

func (m *NhanVien) GetDmChucDanhId() uint32 {
	if m != nil {
		return m.DmChucDanhId
	}
	return 0
}

func (m *NhanVien) GetDmChucVuId() uint32 {
	if m != nil {
		return m.DmChucVuId
	}
	return 0
}

func (m *NhanVien) GetSoDienThoai() string {
	if m != nil {
		return m.SoDienThoai
	}
	return ""
}

// Request data to read todo task
type ReadRequest struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Unique integer identifier of the todo task
	Id                   int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c946787ad119af67, []int{1}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Contains todo task data specified in by ID request
type ReadResponse struct {
	// API versioning: it is my best practice to specify version explicitly
	Api string `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	// Task entity read by ID
	NhanVien             []*NhanVien `protobuf:"bytes,2,rep,name=nhanVien,proto3" json:"nhanVien,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c946787ad119af67, []int{2}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetApi() string {
	if m != nil {
		return m.Api
	}
	return ""
}

func (m *ReadResponse) GetNhanVien() []*NhanVien {
	if m != nil {
		return m.NhanVien
	}
	return nil
}

func init() {
	proto.RegisterType((*NhanVien)(nil), "proto.NhanVien")
	proto.RegisterType((*ReadRequest)(nil), "proto.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "proto.ReadResponse")
}

func init() {
	proto.RegisterFile("processor_nhan_vien.proto", fileDescriptor_c946787ad119af67)
}

var fileDescriptor_c946787ad119af67 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x6a, 0xe3, 0x30,
	0x14, 0x85, 0x89, 0x9d, 0x1f, 0xfb, 0x3a, 0x99, 0x0c, 0x9a, 0x8d, 0x33, 0xb3, 0xf1, 0x18, 0x0a,
	0x81, 0x42, 0x0a, 0x29, 0xf4, 0x01, 0xda, 0x40, 0xf1, 0xa2, 0x59, 0xb8, 0x25, 0x8b, 0x6e, 0x8c,
	0x62, 0xdd, 0x56, 0x82, 0x58, 0x72, 0x2d, 0x3b, 0x0f, 0xd7, 0xa7, 0x2b, 0x92, 0x63, 0x27, 0x85,
	0xae, 0x6c, 0x9d, 0xf3, 0x5d, 0x71, 0xcf, 0x11, 0x2c, 0xca, 0x4a, 0xe5, 0xa8, 0xb5, 0xaa, 0x32,
	0xc9, 0xa9, 0xcc, 0x8e, 0x02, 0xe5, 0xaa, 0xac, 0x54, 0xad, 0xc8, 0xc8, 0x7e, 0xe2, 0x4f, 0x07,
	0xbc, 0x2d, 0xa7, 0x72, 0x27, 0x50, 0x92, 0x2b, 0x98, 0xb3, 0xe2, 0x4c, 0x66, 0x82, 0x85, 0x83,
	0x68, 0xb0, 0x9c, 0xa5, 0x53, 0x56, 0x74, 0x50, 0xc2, 0x48, 0x04, 0xd3, 0x82, 0x9e, 0xb1, 0xd0,
	0x89, 0x06, 0x4b, 0x3f, 0x85, 0x82, 0xf6, 0x17, 0xc5, 0x30, 0xab, 0x51, 0x5e, 0x20, 0xae, 0x45,
	0x82, 0x1a, 0x65, 0xcf, 0xfc, 0x03, 0xff, 0xad, 0x39, 0x1c, 0x32, 0x49, 0x0b, 0x0c, 0x87, 0xd6,
	0xf7, 0x8c, 0xb0, 0xa5, 0x05, 0x1a, 0xf3, 0x40, 0x75, 0xdd, 0x9a, 0xa3, 0xd6, 0x34, 0x82, 0x35,
	0x17, 0xe0, 0x95, 0x5c, 0xc9, 0x77, 0xb3, 0xdf, 0xd8, 0xee, 0x37, 0xb1, 0xe7, 0x84, 0x9d, 0x12,
	0xe4, 0xbc, 0xc9, 0x33, 0x46, 0x25, 0x37, 0xc4, 0xa4, 0x4b, 0xf0, 0xc0, 0x9b, 0x7c, 0x43, 0x25,
	0x4f, 0x18, 0xf9, 0x0f, 0xb3, 0x0e, 0x3b, 0x36, 0x06, 0xf2, 0x2c, 0x04, 0x2d, 0xb4, 0x6b, 0x12,
	0x66, 0x22, 0x68, 0x95, 0x31, 0x53, 0x43, 0xcd, 0x15, 0x15, 0xa1, 0xdf, 0x46, 0xd0, 0x6a, 0x23,
	0x50, 0xbe, 0x18, 0x29, 0xbe, 0x81, 0x20, 0x45, 0xca, 0x52, 0xfc, 0x68, 0x50, 0xd7, 0xe4, 0x37,
	0xb8, 0xb4, 0x14, 0xb6, 0x32, 0x3f, 0x35, 0xbf, 0xe4, 0x17, 0x38, 0x82, 0xd9, 0x7e, 0xdc, 0xd4,
	0x11, 0x2c, 0x7e, 0x82, 0x69, 0x3b, 0xa0, 0x4b, 0x25, 0x35, 0xfe, 0x30, 0x71, 0x0d, 0x9e, 0x3c,
	0x35, 0x14, 0x3a, 0x91, 0xbb, 0x0c, 0xd6, 0xf3, 0xf6, 0xc1, 0x56, 0x5d, 0x71, 0x69, 0x0f, 0xac,
	0x13, 0x98, 0x77, 0xea, 0x33, 0x56, 0x47, 0x91, 0x23, 0xb9, 0x83, 0xe0, 0x11, 0xeb, 0xbe, 0x64,
	0x72, 0x1a, 0xbe, 0x58, 0xf3, 0xef, 0x9f, 0x6f, 0x5a, 0xbb, 0xc9, 0xfd, 0xf0, 0xd5, 0x29, 0xf7,
	0xfb, 0xb1, 0x75, 0x6e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x33, 0x7e, 0x32, 0xe2, 0x38, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NhanVienServiceClient is the client API for NhanVienService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NhanVienServiceClient interface {
	GetNhanVien(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
}

type nhanVienServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNhanVienServiceClient(cc grpc.ClientConnInterface) NhanVienServiceClient {
	return &nhanVienServiceClient{cc}
}

func (c *nhanVienServiceClient) GetNhanVien(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/proto.NhanVienService/GetNhanVien", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NhanVienServiceServer is the server API for NhanVienService service.
type NhanVienServiceServer interface {
	GetNhanVien(context.Context, *ReadRequest) (*ReadResponse, error)
}

// UnimplementedNhanVienServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNhanVienServiceServer struct {
}

func (*UnimplementedNhanVienServiceServer) GetNhanVien(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNhanVien not implemented")
}

func RegisterNhanVienServiceServer(s *grpc.Server, srv NhanVienServiceServer) {
	s.RegisterService(&_NhanVienService_serviceDesc, srv)
}

func _NhanVienService_GetNhanVien_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NhanVienServiceServer).GetNhanVien(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.NhanVienService/GetNhanVien",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NhanVienServiceServer).GetNhanVien(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NhanVienService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NhanVienService",
	HandlerType: (*NhanVienServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNhanVien",
			Handler:    _NhanVienService_GetNhanVien_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "processor_nhan_vien.proto",
}
