// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/language/proto/language/language_service.proto

package language

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	entities "github.com/seidu626/audiobook/service/language/proto/entities"
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

// FIXME: https://github.com/envoyproxy/protoc-gen-validate/issues/223
// with Workaround in .override.go
type ExistRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Abbreviation         *wrappers.StringValue `protobuf:"bytes,3,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExistRequest) Reset()         { *m = ExistRequest{} }
func (m *ExistRequest) String() string { return proto.CompactTextString(m) }
func (*ExistRequest) ProtoMessage()    {}
func (*ExistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{0}
}

func (m *ExistRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistRequest.Unmarshal(m, b)
}
func (m *ExistRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistRequest.Marshal(b, m, deterministic)
}
func (m *ExistRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistRequest.Merge(m, src)
}
func (m *ExistRequest) XXX_Size() int {
	return xxx_messageInfo_ExistRequest.Size(m)
}
func (m *ExistRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExistRequest proto.InternalMessageInfo

func (m *ExistRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ExistRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ExistRequest) GetAbbreviation() *wrappers.StringValue {
	if m != nil {
		return m.Abbreviation
	}
	return nil
}

type ExistResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExistResponse) Reset()         { *m = ExistResponse{} }
func (m *ExistResponse) String() string { return proto.CompactTextString(m) }
func (*ExistResponse) ProtoMessage()    {}
func (*ExistResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{1}
}

func (m *ExistResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExistResponse.Unmarshal(m, b)
}
func (m *ExistResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExistResponse.Marshal(b, m, deterministic)
}
func (m *ExistResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExistResponse.Merge(m, src)
}
func (m *ExistResponse) XXX_Size() int {
	return xxx_messageInfo_ExistResponse.Size(m)
}
func (m *ExistResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExistResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExistResponse proto.InternalMessageInfo

func (m *ExistResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type ListRequest struct {
	Limit                *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Page                 *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Sort                 *wrappers.StringValue `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Abbreviation         *wrappers.StringValue `protobuf:"bytes,5,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	FlagSrc              *wrappers.StringValue `protobuf:"bytes,6,opt,name=flag_src,json=flagSrc,proto3" json:"flag_src,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{2}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetLimit() *wrappers.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *ListRequest) GetPage() *wrappers.UInt32Value {
	if m != nil {
		return m.Page
	}
	return nil
}

func (m *ListRequest) GetSort() *wrappers.StringValue {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *ListRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *ListRequest) GetAbbreviation() *wrappers.StringValue {
	if m != nil {
		return m.Abbreviation
	}
	return nil
}

func (m *ListRequest) GetFlagSrc() *wrappers.StringValue {
	if m != nil {
		return m.FlagSrc
	}
	return nil
}

type ListResponse struct {
	Results              []*entities.Language `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total                uint32               `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{3}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetResults() []*entities.Language {
	if m != nil {
		return m.Results
	}
	return nil
}

func (m *ListResponse) GetTotal() uint32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type GetRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Abbreviation         *wrappers.StringValue `protobuf:"bytes,3,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	FlagSrc              *wrappers.StringValue `protobuf:"bytes,4,opt,name=flag_src,json=flagSrc,proto3" json:"flag_src,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{4}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GetRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *GetRequest) GetAbbreviation() *wrappers.StringValue {
	if m != nil {
		return m.Abbreviation
	}
	return nil
}

func (m *GetRequest) GetFlagSrc() *wrappers.StringValue {
	if m != nil {
		return m.FlagSrc
	}
	return nil
}

type GetResponse struct {
	Result               *entities.Language `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{5}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetResult() *entities.Language {
	if m != nil {
		return m.Result
	}
	return nil
}

type CreateRequest struct {
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Abbreviation         *wrappers.StringValue `protobuf:"bytes,3,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	FlagSrc              *wrappers.StringValue `protobuf:"bytes,4,opt,name=flag_src,json=flagSrc,proto3" json:"flag_src,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{6}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *CreateRequest) GetAbbreviation() *wrappers.StringValue {
	if m != nil {
		return m.Abbreviation
	}
	return nil
}

func (m *CreateRequest) GetFlagSrc() *wrappers.StringValue {
	if m != nil {
		return m.FlagSrc
	}
	return nil
}

type CreateResponse struct {
	Result               *entities.Language `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{7}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetResult() *entities.Language {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Abbreviation         *wrappers.StringValue `protobuf:"bytes,3,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	FlagSrc              *wrappers.StringValue `protobuf:"bytes,4,opt,name=flag_src,json=flagSrc,proto3" json:"flag_src,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{8}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *UpdateRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *UpdateRequest) GetAbbreviation() *wrappers.StringValue {
	if m != nil {
		return m.Abbreviation
	}
	return nil
}

func (m *UpdateRequest) GetFlagSrc() *wrappers.StringValue {
	if m != nil {
		return m.FlagSrc
	}
	return nil
}

type UpdateResponse struct {
	Result               *entities.Language `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{9}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetResult() *entities.Language {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 *wrappers.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{10}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DeleteRequest) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

type DeleteResponse struct {
	Result               *entities.Language `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9745f0f97b1ecfc5, []int{11}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetResult() *entities.Language {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ExistRequest)(nil), "micro.service.language.language.v1.ExistRequest")
	proto.RegisterType((*ExistResponse)(nil), "micro.service.language.language.v1.ExistResponse")
	proto.RegisterType((*ListRequest)(nil), "micro.service.language.language.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "micro.service.language.language.v1.ListResponse")
	proto.RegisterType((*GetRequest)(nil), "micro.service.language.language.v1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "micro.service.language.language.v1.GetResponse")
	proto.RegisterType((*CreateRequest)(nil), "micro.service.language.language.v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "micro.service.language.language.v1.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "micro.service.language.language.v1.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "micro.service.language.language.v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "micro.service.language.language.v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "micro.service.language.language.v1.DeleteResponse")
}

func init() {
	proto.RegisterFile("service/language/proto/language/language_service.proto", fileDescriptor_9745f0f97b1ecfc5)
}

var fileDescriptor_9745f0f97b1ecfc5 = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x4d, 0x6b, 0x13, 0x41,
	0x18, 0x66, 0x37, 0x9b, 0xb4, 0xbe, 0x69, 0xa4, 0x0c, 0x22, 0x21, 0x7e, 0x50, 0x82, 0x87, 0x58,
	0xda, 0xdd, 0x64, 0x8b, 0x81, 0x0a, 0x85, 0xb2, 0x56, 0x4b, 0xa1, 0x07, 0xd9, 0x52, 0x11, 0x8b,
	0x86, 0x49, 0x76, 0xba, 0x1d, 0xdc, 0xec, 0x6c, 0x67, 0x67, 0xe3, 0x17, 0x82, 0xbf, 0x43, 0x4f,
	0x1e, 0x04, 0x8f, 0xfe, 0x20, 0x7f, 0x84, 0x47, 0xc9, 0x49, 0xb2, 0x1f, 0x49, 0xb6, 0x10, 0x3a,
	0x91, 0x1c, 0x8a, 0x78, 0x9b, 0xd9, 0xbc, 0xcf, 0xfb, 0xce, 0xfb, 0x3c, 0xf3, 0xcc, 0x4c, 0xa0,
	0x1d, 0x12, 0x3e, 0xa0, 0x3d, 0x62, 0x78, 0xd8, 0x77, 0x23, 0xec, 0x12, 0x23, 0xe0, 0x4c, 0xb0,
	0xc9, 0x34, 0x1b, 0x74, 0xd2, 0x40, 0x3d, 0xfe, 0x1d, 0xd5, 0xfb, 0xb4, 0xc7, 0x99, 0x9e, 0x7d,
	0xcc, 0xa2, 0x26, 0x83, 0x41, 0xab, 0x76, 0xd7, 0x65, 0xcc, 0xf5, 0xd2, 0x8c, 0xdd, 0xe8, 0xd4,
	0x78, 0xc3, 0x71, 0x10, 0x10, 0x1e, 0x26, 0x39, 0x6a, 0xfa, 0x8c, 0xda, 0xc4, 0x17, 0x54, 0x50,
	0x12, 0x8e, 0x07, 0x69, 0xfc, 0x7d, 0x71, 0x46, 0xb9, 0xd3, 0x09, 0x30, 0x17, 0xef, 0xd2, 0xd0,
	0x01, 0xf6, 0xa8, 0x83, 0x05, 0x19, 0x0f, 0x92, 0xd0, 0xfa, 0x4f, 0x05, 0x56, 0x1e, 0xbf, 0xa5,
	0xa1, 0xb0, 0xc9, 0x79, 0x44, 0x42, 0x81, 0xda, 0xa0, 0x52, 0xa7, 0xaa, 0xac, 0x29, 0x8d, 0xb2,
	0x79, 0x5b, 0x4f, 0x16, 0xa6, 0x67, 0x0b, 0xd3, 0x8f, 0x04, 0xa7, 0xbe, 0xfb, 0x0c, 0x7b, 0x11,
	0xb1, 0x96, 0x87, 0x56, 0x91, 0x17, 0x7e, 0x28, 0x8a, 0xad, 0x52, 0x07, 0xd9, 0xa0, 0xf9, 0xb8,
	0x4f, 0xaa, 0xaa, 0x04, 0x72, 0x6d, 0x68, 0xdd, 0xe1, 0xb7, 0x56, 0xb5, 0xea, 0x6a, 0xe3, 0x93,
	0x6a, 0xa2, 0x57, 0x27, 0x78, 0xf3, 0x7d, 0x73, 0x73, 0xbb, 0xb3, 0xf9, 0xf2, 0xc3, 0xd6, 0x46,
	0xeb, 0xc1, 0xc7, 0x7b, 0x76, 0x9c, 0x0b, 0xed, 0xc2, 0x0a, 0xee, 0x76, 0x39, 0x19, 0x50, 0x2c,
	0x28, 0xf3, 0xab, 0x85, 0xcb, 0x73, 0xdb, 0x39, 0x44, 0x7d, 0x03, 0x2a, 0x69, 0x77, 0x61, 0xc0,
	0xfc, 0x90, 0xa0, 0x9b, 0x50, 0xe2, 0x24, 0x8c, 0x3c, 0x11, 0xb7, 0xb8, 0x6c, 0xa7, 0xb3, 0x87,
	0x85, 0xdf, 0x96, 0x52, 0xff, 0x5c, 0x80, 0xf2, 0xe1, 0x14, 0x17, 0x3b, 0x50, 0xf4, 0x68, 0x9f,
	0x8a, 0x99, 0x74, 0x1c, 0x1f, 0xf8, 0x62, 0xcb, 0x4c, 0x9a, 0xba, 0x36, 0xb4, 0x4a, 0xeb, 0x5a,
	0xd5, 0x69, 0x28, 0x76, 0x82, 0x42, 0xdb, 0xa0, 0x05, 0xd8, 0x9d, 0x4d, 0xc9, 0x34, 0x7a, 0x69,
	0x68, 0x69, 0xeb, 0x6a, 0x43, 0xb1, 0x63, 0x08, 0x6a, 0x82, 0x16, 0x32, 0x2e, 0xa4, 0x3a, 0x8e,
	0x23, 0xc7, 0xfc, 0x6b, 0x0b, 0xe4, 0xff, 0xe0, 0x02, 0xff, 0x45, 0x89, 0xdc, 0xa3, 0x46, 0xb8,
	0xba, 0x5a, 0xc8, 0x0b, 0x81, 0x76, 0x61, 0xf9, 0xd4, 0xc3, 0x6e, 0x27, 0xe4, 0xbd, 0x6a, 0x69,
	0x9e, 0x34, 0x4b, 0x23, 0xd8, 0x11, 0xef, 0xd5, 0xcf, 0x61, 0xe5, 0x70, 0x5a, 0xc9, 0x27, 0xb0,
	0x94, 0x68, 0x17, 0x56, 0x95, 0xb5, 0x42, 0xa3, 0x6c, 0x6e, 0xe8, 0x33, 0xac, 0x36, 0x76, 0xc7,
	0xa0, 0xa5, 0x1f, 0xa6, 0x1f, 0xed, 0x0c, 0x8c, 0x6e, 0x40, 0x51, 0x30, 0x81, 0xbd, 0x58, 0xa6,
	0x8a, 0x9d, 0x4c, 0x92, 0xfd, 0xf0, 0x4d, 0x05, 0xd8, 0x27, 0x57, 0xd2, 0x1a, 0x07, 0xf3, 0x5b,
	0x43, 0x46, 0x1a, 0xed, 0xaf, 0xa4, 0x79, 0x0e, 0xe5, 0x98, 0xa6, 0x54, 0x99, 0xbd, 0x9c, 0xc7,
	0xe6, 0x15, 0x26, 0xe7, 0xc8, 0x5f, 0x0a, 0x54, 0x1e, 0x71, 0x82, 0x05, 0xc9, 0x44, 0xf8, 0xe7,
	0xc9, 0x3c, 0x81, 0xeb, 0x59, 0xc7, 0x8b, 0xe7, 0xf3, 0xbb, 0x0a, 0x95, 0xe3, 0xc0, 0x99, 0xe2,
	0xf3, 0xff, 0xa6, 0x9e, 0xad, 0x43, 0xc6, 0xd4, 0xe2, 0x75, 0xf8, 0xa2, 0x40, 0x65, 0x8f, 0x78,
	0xe4, 0x4a, 0xea, 0x30, 0x6a, 0x3d, 0x5b, 0xdc, 0xc2, 0x5b, 0x37, 0xbf, 0x16, 0xa1, 0x7c, 0x1c,
	0x12, 0x7e, 0x94, 0x40, 0x91, 0x07, 0xc5, 0xf8, 0x8a, 0x46, 0x4d, 0xfd, 0xf2, 0xa7, 0x92, 0x3e,
	0xfd, 0x56, 0xa9, 0xb5, 0xe6, 0x40, 0xa4, 0x8d, 0x50, 0xd0, 0x46, 0xb7, 0x08, 0x32, 0x64, 0xa0,
	0x53, 0x6f, 0x81, 0x5a, 0x53, 0x1e, 0x90, 0x96, 0x3a, 0x85, 0xc2, 0x3e, 0x11, 0x48, 0x97, 0x01,
	0x4e, 0x6e, 0x99, 0x9a, 0x21, 0x1d, 0x9f, 0xd6, 0x61, 0x50, 0x4a, 0x0e, 0x0c, 0x24, 0xc5, 0x47,
	0xee, 0x38, 0xad, 0x99, 0xf3, 0x40, 0x26, 0x05, 0x13, 0x67, 0xc8, 0x15, 0xcc, 0x9d, 0x37, 0x72,
	0x05, 0x2f, 0x18, 0x8f, 0x41, 0x29, 0xd9, 0x8f, 0x72, 0x05, 0x73, 0xc6, 0x92, 0x2b, 0x98, 0xdf,
	0xee, 0x16, 0x01, 0x89, 0x67, 0xfb, 0x53, 0xe5, 0xc5, 0x8e, 0x4b, 0xc5, 0x59, 0xd4, 0xd5, 0x7b,
	0xac, 0x6f, 0x84, 0x84, 0x3a, 0x51, 0xdb, 0x6c, 0x1b, 0x38, 0x72, 0x28, 0xeb, 0x32, 0xf6, 0xda,
	0xb8, 0xe4, 0x2f, 0x43, 0xb7, 0x14, 0xcf, 0xb7, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x5d,
	0x2b, 0x6b, 0x5c, 0x0c, 0x00, 0x00,
}
