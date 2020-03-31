// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog/blog.proto

package blogpb

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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{0}
}

func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (m *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(m, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

//create
type CreateBlogReq struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogReq) Reset()         { *m = CreateBlogReq{} }
func (m *CreateBlogReq) String() string { return proto.CompactTextString(m) }
func (*CreateBlogReq) ProtoMessage()    {}
func (*CreateBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{1}
}

func (m *CreateBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogReq.Unmarshal(m, b)
}
func (m *CreateBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogReq.Marshal(b, m, deterministic)
}
func (m *CreateBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogReq.Merge(m, src)
}
func (m *CreateBlogReq) XXX_Size() int {
	return xxx_messageInfo_CreateBlogReq.Size(m)
}
func (m *CreateBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogReq proto.InternalMessageInfo

func (m *CreateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRes) Reset()         { *m = CreateBlogRes{} }
func (m *CreateBlogRes) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRes) ProtoMessage()    {}
func (*CreateBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{2}
}

func (m *CreateBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRes.Unmarshal(m, b)
}
func (m *CreateBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRes.Marshal(b, m, deterministic)
}
func (m *CreateBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRes.Merge(m, src)
}
func (m *CreateBlogRes) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRes.Size(m)
}
func (m *CreateBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRes proto.InternalMessageInfo

func (m *CreateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

//read
type ReadBlogReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogReq) Reset()         { *m = ReadBlogReq{} }
func (m *ReadBlogReq) String() string { return proto.CompactTextString(m) }
func (*ReadBlogReq) ProtoMessage()    {}
func (*ReadBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{3}
}

func (m *ReadBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogReq.Unmarshal(m, b)
}
func (m *ReadBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogReq.Marshal(b, m, deterministic)
}
func (m *ReadBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogReq.Merge(m, src)
}
func (m *ReadBlogReq) XXX_Size() int {
	return xxx_messageInfo_ReadBlogReq.Size(m)
}
func (m *ReadBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogReq proto.InternalMessageInfo

func (m *ReadBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogRes) Reset()         { *m = ReadBlogRes{} }
func (m *ReadBlogRes) String() string { return proto.CompactTextString(m) }
func (*ReadBlogRes) ProtoMessage()    {}
func (*ReadBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{4}
}

func (m *ReadBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogRes.Unmarshal(m, b)
}
func (m *ReadBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogRes.Marshal(b, m, deterministic)
}
func (m *ReadBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogRes.Merge(m, src)
}
func (m *ReadBlogRes) XXX_Size() int {
	return xxx_messageInfo_ReadBlogRes.Size(m)
}
func (m *ReadBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogRes proto.InternalMessageInfo

func (m *ReadBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

//update
type UpdateBlogReq struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogReq) Reset()         { *m = UpdateBlogReq{} }
func (m *UpdateBlogReq) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogReq) ProtoMessage()    {}
func (*UpdateBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{5}
}

func (m *UpdateBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogReq.Unmarshal(m, b)
}
func (m *UpdateBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogReq.Marshal(b, m, deterministic)
}
func (m *UpdateBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogReq.Merge(m, src)
}
func (m *UpdateBlogReq) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogReq.Size(m)
}
func (m *UpdateBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogReq proto.InternalMessageInfo

func (m *UpdateBlogReq) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogRes) Reset()         { *m = UpdateBlogRes{} }
func (m *UpdateBlogRes) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogRes) ProtoMessage()    {}
func (*UpdateBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{6}
}

func (m *UpdateBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogRes.Unmarshal(m, b)
}
func (m *UpdateBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogRes.Marshal(b, m, deterministic)
}
func (m *UpdateBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogRes.Merge(m, src)
}
func (m *UpdateBlogRes) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogRes.Size(m)
}
func (m *UpdateBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogRes proto.InternalMessageInfo

func (m *UpdateBlogRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

//delete
type DeleteBlogReq struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogReq) Reset()         { *m = DeleteBlogReq{} }
func (m *DeleteBlogReq) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogReq) ProtoMessage()    {}
func (*DeleteBlogReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{7}
}

func (m *DeleteBlogReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogReq.Unmarshal(m, b)
}
func (m *DeleteBlogReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogReq.Marshal(b, m, deterministic)
}
func (m *DeleteBlogReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogReq.Merge(m, src)
}
func (m *DeleteBlogReq) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogReq.Size(m)
}
func (m *DeleteBlogReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogReq proto.InternalMessageInfo

func (m *DeleteBlogReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteBlogRes struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogRes) Reset()         { *m = DeleteBlogRes{} }
func (m *DeleteBlogRes) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogRes) ProtoMessage()    {}
func (*DeleteBlogRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{8}
}

func (m *DeleteBlogRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogRes.Unmarshal(m, b)
}
func (m *DeleteBlogRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogRes.Marshal(b, m, deterministic)
}
func (m *DeleteBlogRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogRes.Merge(m, src)
}
func (m *DeleteBlogRes) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogRes.Size(m)
}
func (m *DeleteBlogRes) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogRes.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogRes proto.InternalMessageInfo

func (m *DeleteBlogRes) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

//listing
type ListBlogsReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogsReq) Reset()         { *m = ListBlogsReq{} }
func (m *ListBlogsReq) String() string { return proto.CompactTextString(m) }
func (*ListBlogsReq) ProtoMessage()    {}
func (*ListBlogsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{9}
}

func (m *ListBlogsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogsReq.Unmarshal(m, b)
}
func (m *ListBlogsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogsReq.Marshal(b, m, deterministic)
}
func (m *ListBlogsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogsReq.Merge(m, src)
}
func (m *ListBlogsReq) XXX_Size() int {
	return xxx_messageInfo_ListBlogsReq.Size(m)
}
func (m *ListBlogsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogsReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogsReq proto.InternalMessageInfo

type ListBlogsRes struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogsRes) Reset()         { *m = ListBlogsRes{} }
func (m *ListBlogsRes) String() string { return proto.CompactTextString(m) }
func (*ListBlogsRes) ProtoMessage()    {}
func (*ListBlogsRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e737bdbda5671783, []int{10}
}

func (m *ListBlogsRes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogsRes.Unmarshal(m, b)
}
func (m *ListBlogsRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogsRes.Marshal(b, m, deterministic)
}
func (m *ListBlogsRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogsRes.Merge(m, src)
}
func (m *ListBlogsRes) XXX_Size() int {
	return xxx_messageInfo_ListBlogsRes.Size(m)
}
func (m *ListBlogsRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogsRes.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogsRes proto.InternalMessageInfo

func (m *ListBlogsRes) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogReq)(nil), "blog.CreateBlogReq")
	proto.RegisterType((*CreateBlogRes)(nil), "blog.CreateBlogRes")
	proto.RegisterType((*ReadBlogReq)(nil), "blog.ReadBlogReq")
	proto.RegisterType((*ReadBlogRes)(nil), "blog.ReadBlogRes")
	proto.RegisterType((*UpdateBlogReq)(nil), "blog.UpdateBlogReq")
	proto.RegisterType((*UpdateBlogRes)(nil), "blog.UpdateBlogRes")
	proto.RegisterType((*DeleteBlogReq)(nil), "blog.DeleteBlogReq")
	proto.RegisterType((*DeleteBlogRes)(nil), "blog.DeleteBlogRes")
	proto.RegisterType((*ListBlogsReq)(nil), "blog.ListBlogsReq")
	proto.RegisterType((*ListBlogsRes)(nil), "blog.ListBlogsRes")
}

func init() {
	proto.RegisterFile("blog/blog.proto", fileDescriptor_e737bdbda5671783)
}

var fileDescriptor_e737bdbda5671783 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x0d, 0x15, 0xb1, 0x0c, 0x82, 0x71, 0xf4, 0xd0, 0xd4, 0xf8, 0x91, 0x9e, 0xf4, 0x20, 0x10,
	0x8c, 0xfe, 0x00, 0xf4, 0x62, 0xe2, 0xa9, 0xc6, 0x8b, 0x17, 0x03, 0xdd, 0x09, 0x6e, 0xd2, 0xb0,
	0xd0, 0x59, 0xfc, 0x39, 0xfe, 0x56, 0xb3, 0xbb, 0x59, 0xdb, 0x02, 0x09, 0x7a, 0x69, 0x3a, 0x6f,
	0xde, 0xeb, 0x7b, 0x7d, 0x9b, 0x85, 0xa3, 0x69, 0xae, 0x66, 0x03, 0xf3, 0xe8, 0x2f, 0x0a, 0xa5,
	0x15, 0x36, 0xcd, 0x7b, 0x92, 0x41, 0x73, 0x9c, 0xab, 0x19, 0xf6, 0x20, 0x90, 0x22, 0x6a, 0x5c,
	0x35, 0xae, 0xdb, 0x69, 0x20, 0x05, 0x9e, 0x41, 0x7b, 0xb2, 0xd2, 0x9f, 0xaa, 0xf8, 0x90, 0x22,
	0x0a, 0x2c, 0x1c, 0x3a, 0xe0, 0x59, 0xe0, 0x29, 0xec, 0x6b, 0xa9, 0x73, 0x8a, 0xf6, 0xec, 0xc2,
	0x0d, 0x18, 0xc1, 0x41, 0xa6, 0xe6, 0x9a, 0xe6, 0x3a, 0x6a, 0x5a, 0xdc, 0x8f, 0xc9, 0x00, 0xba,
	0x8f, 0x05, 0x4d, 0x34, 0x19, 0xab, 0x94, 0x96, 0x78, 0x01, 0xd6, 0xdd, 0xfa, 0x75, 0x46, 0xd0,
	0xb7, 0xb1, 0xec, 0xd2, 0xa5, 0x5a, 0x13, 0xf0, 0x4e, 0xc1, 0x39, 0x74, 0x52, 0x9a, 0x08, 0xff,
	0xfd, 0xb5, 0xbf, 0x49, 0x6e, 0xab, 0x6b, 0xfe, 0x8b, 0xfd, 0xdb, 0x42, 0xfc, 0x2f, 0x6f, 0x55,
	0xb0, 0xdb, 0xe1, 0x12, 0xba, 0x4f, 0x94, 0x53, 0xe9, 0xb0, 0x9e, 0xf8, 0xa6, 0x4e, 0x60, 0xd3,
	0x2e, 0xaf, 0xb2, 0x8c, 0x98, 0x2d, 0x2b, 0x4c, 0xfd, 0x98, 0xf4, 0xe0, 0xf0, 0x45, 0xb2, 0x36,
	0x44, 0x4e, 0x69, 0x99, 0xf4, 0x6b, 0xf3, 0xce, 0x2c, 0xa3, 0xef, 0x00, 0x3a, 0x66, 0x7c, 0xa5,
	0xe2, 0x4b, 0x66, 0x84, 0x0f, 0x00, 0x65, 0xf9, 0x78, 0xe2, 0xf8, 0xb5, 0xf3, 0x8b, 0xb7, 0x80,
	0x8c, 0x43, 0x08, 0x7d, 0xc9, 0x78, 0xec, 0x08, 0x95, 0x33, 0x89, 0x37, 0x20, 0x36, 0x4e, 0x65,
	0x6d, 0xde, 0xa9, 0xd6, 0x7c, 0xbc, 0x05, 0xb4, 0xba, 0xb2, 0x1c, 0xaf, 0xab, 0xf5, 0x19, 0x6f,
	0x01, 0x19, 0xef, 0xa1, 0xfd, 0xdb, 0x0c, 0xa2, 0x63, 0x54, 0xab, 0x8b, 0x37, 0x31, 0x1e, 0x36,
	0xc6, 0xe1, 0x7b, 0xcb, 0xc0, 0x8b, 0xe9, 0xb4, 0x65, 0xaf, 0xce, 0xdd, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x8a, 0xed, 0x20, 0x4e, 0x4d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error)
	ReadBlog(ctx context.Context, in *ReadBlogReq, opts ...grpc.CallOption) (*ReadBlogRes, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error)
	DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogRes, error)
	ListBlogs(ctx context.Context, in *ListBlogsReq, opts ...grpc.CallOption) (BlogService_ListBlogsClient, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogReq, opts ...grpc.CallOption) (*CreateBlogRes, error) {
	out := new(CreateBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *ReadBlogReq, opts ...grpc.CallOption) (*ReadBlogRes, error) {
	out := new(ReadBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *UpdateBlogReq, opts ...grpc.CallOption) (*UpdateBlogRes, error) {
	out := new(UpdateBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlog(ctx context.Context, in *DeleteBlogReq, opts ...grpc.CallOption) (*DeleteBlogRes, error) {
	out := new(DeleteBlogRes)
	err := c.cc.Invoke(ctx, "/blog.BlogService/DeleteBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListBlogs(ctx context.Context, in *ListBlogsReq, opts ...grpc.CallOption) (BlogService_ListBlogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlogService_serviceDesc.Streams[0], "/blog.BlogService/ListBlogs", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServiceListBlogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogService_ListBlogsClient interface {
	Recv() (*ListBlogsRes, error)
	grpc.ClientStream
}

type blogServiceListBlogsClient struct {
	grpc.ClientStream
}

func (x *blogServiceListBlogsClient) Recv() (*ListBlogsRes, error) {
	m := new(ListBlogsRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogReq) (*CreateBlogRes, error)
	ReadBlog(context.Context, *ReadBlogReq) (*ReadBlogRes, error)
	UpdateBlog(context.Context, *UpdateBlogReq) (*UpdateBlogRes, error)
	DeleteBlog(context.Context, *DeleteBlogReq) (*DeleteBlogRes, error)
	ListBlogs(*ListBlogsReq, BlogService_ListBlogsServer) error
}

// UnimplementedBlogServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (*UnimplementedBlogServiceServer) CreateBlog(ctx context.Context, req *CreateBlogReq) (*CreateBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ReadBlog(ctx context.Context, req *ReadBlogReq) (*ReadBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlog not implemented")
}
func (*UnimplementedBlogServiceServer) UpdateBlog(ctx context.Context, req *UpdateBlogReq) (*UpdateBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlog not implemented")
}
func (*UnimplementedBlogServiceServer) DeleteBlog(ctx context.Context, req *DeleteBlogReq) (*DeleteBlogRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlog not implemented")
}
func (*UnimplementedBlogServiceServer) ListBlogs(req *ListBlogsReq, srv BlogService_ListBlogsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListBlogs not implemented")
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*ReadBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*UpdateBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlog(ctx, req.(*DeleteBlogReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListBlogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlogsReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServiceServer).ListBlogs(m, &blogServiceListBlogsServer{stream})
}

type BlogService_ListBlogsServer interface {
	Send(*ListBlogsRes) error
	grpc.ServerStream
}

type blogServiceListBlogsServer struct {
	grpc.ServerStream
}

func (x *blogServiceListBlogsServer) Send(m *ListBlogsRes) error {
	return x.ServerStream.SendMsg(m)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _BlogService_DeleteBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBlogs",
			Handler:       _BlogService_ListBlogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blog/blog.proto",
}
