// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/language/proto/word/word_service.proto

package word

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	entities "github.com/seidu626/audiobook/backend/services/language/proto/entities"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// FIXME: https://github.com/envoyproxy/protoc-gen-validate/issues/223
// with Workaround in .override.go
type ExistRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              *wrappers.StringValue `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExistRequest) Reset()         { *m = ExistRequest{} }
func (m *ExistRequest) String() string { return proto.CompactTextString(m) }
func (*ExistRequest) ProtoMessage()    {}
func (*ExistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{0}
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

func (m *ExistRequest) GetContent() *wrappers.StringValue {
	if m != nil {
		return m.Content
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
	return fileDescriptor_9091488fa1aacfc8, []int{1}
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
	Content              *wrappers.StringValue `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	AudioSrc             *wrappers.StringValue `protobuf:"bytes,5,opt,name=audio_src,json=audioSrc,proto3" json:"audio_src,omitempty"`
	Language             *entities.Language    `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{2}
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

func (m *ListRequest) GetContent() *wrappers.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *ListRequest) GetAudioSrc() *wrappers.StringValue {
	if m != nil {
		return m.AudioSrc
	}
	return nil
}

func (m *ListRequest) GetLanguage() *entities.Language {
	if m != nil {
		return m.Language
	}
	return nil
}

type ListResponse struct {
	Results              []*entities.Word `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total                uint32           `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{3}
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

func (m *ListResponse) GetResults() []*entities.Word {
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
	Content              *wrappers.StringValue `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	AudioSrc             *wrappers.StringValue `protobuf:"bytes,3,opt,name=audio_src,json=audioSrc,proto3" json:"audio_src,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{4}
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

func (m *GetRequest) GetContent() *wrappers.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *GetRequest) GetAudioSrc() *wrappers.StringValue {
	if m != nil {
		return m.AudioSrc
	}
	return nil
}

type GetResponse struct {
	Result               *entities.Word `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{5}
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

func (m *GetResponse) GetResult() *entities.Word {
	if m != nil {
		return m.Result
	}
	return nil
}

type CreateRequest struct {
	Content              *wrappers.StringValue `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	AudioSrc             *wrappers.StringValue `protobuf:"bytes,2,opt,name=audio_src,json=audioSrc,proto3" json:"audio_src,omitempty"`
	SkillId              string                `protobuf:"bytes,3,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{6}
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

func (m *CreateRequest) GetContent() *wrappers.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *CreateRequest) GetAudioSrc() *wrappers.StringValue {
	if m != nil {
		return m.AudioSrc
	}
	return nil
}

func (m *CreateRequest) GetSkillId() string {
	if m != nil {
		return m.SkillId
	}
	return ""
}

type CreateResponse struct {
	Result               *entities.Word `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{7}
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

func (m *CreateResponse) GetResult() *entities.Word {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              *wrappers.StringValue `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	AudioSrc             *wrappers.StringValue `protobuf:"bytes,3,opt,name=audio_src,json=audioSrc,proto3" json:"audio_src,omitempty"`
	SkillId              string                `protobuf:"bytes,4,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{8}
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

func (m *UpdateRequest) GetContent() *wrappers.StringValue {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *UpdateRequest) GetAudioSrc() *wrappers.StringValue {
	if m != nil {
		return m.AudioSrc
	}
	return nil
}

func (m *UpdateRequest) GetSkillId() string {
	if m != nil {
		return m.SkillId
	}
	return ""
}

type UpdateResponse struct {
	Result               *entities.Word `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{9}
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

func (m *UpdateResponse) GetResult() *entities.Word {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{10}
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

type DeleteResponse struct {
	Result               *entities.Word `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9091488fa1aacfc8, []int{11}
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

func (m *DeleteResponse) GetResult() *entities.Word {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ExistRequest)(nil), "audiobook.language.word.v1.ExistRequest")
	proto.RegisterType((*ExistResponse)(nil), "audiobook.language.word.v1.ExistResponse")
	proto.RegisterType((*ListRequest)(nil), "audiobook.language.word.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "audiobook.language.word.v1.ListResponse")
	proto.RegisterType((*GetRequest)(nil), "audiobook.language.word.v1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "audiobook.language.word.v1.GetResponse")
	proto.RegisterType((*CreateRequest)(nil), "audiobook.language.word.v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "audiobook.language.word.v1.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "audiobook.language.word.v1.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "audiobook.language.word.v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "audiobook.language.word.v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "audiobook.language.word.v1.DeleteResponse")
}

func init() {
	proto.RegisterFile("services/language/proto/word/word_service.proto", fileDescriptor_9091488fa1aacfc8)
}

var fileDescriptor_9091488fa1aacfc8 = []byte{
	// 754 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0x45, 0x96, 0xfc, 0x93, 0xeb, 0xf8, 0xfb, 0xc2, 0xf4, 0xcf, 0x88, 0x50, 0x82, 0x28, 0x8d,
	0x6b, 0x8a, 0xd4, 0x38, 0x10, 0x68, 0xe8, 0x0f, 0xa8, 0x29, 0x21, 0x90, 0x45, 0x91, 0x49, 0x0b,
	0xdd, 0x04, 0xd9, 0x9a, 0xba, 0x83, 0x15, 0x8d, 0x3b, 0x33, 0x72, 0x0a, 0xa5, 0x9b, 0x52, 0xe8,
	0x03, 0xf4, 0x45, 0xba, 0xec, 0xba, 0xbb, 0xae, 0xfb, 0x0a, 0x7d, 0x80, 0x2e, 0x4b, 0x56, 0xc5,
	0x23, 0x8d, 0x23, 0x41, 0x2c, 0x9b, 0x24, 0x50, 0xba, 0x31, 0xa3, 0x99, 0x7b, 0xef, 0x39, 0xf7,
	0xcc, 0x99, 0x19, 0x83, 0xc3, 0x31, 0x1b, 0x93, 0x3e, 0xe6, 0x4e, 0xe8, 0x47, 0x83, 0xd8, 0x1f,
	0x60, 0x67, 0xc4, 0xa8, 0xa0, 0xce, 0x31, 0x65, 0x81, 0xfc, 0x39, 0x4c, 0x23, 0x6c, 0x39, 0x8f,
	0x4c, 0x3f, 0x0e, 0x08, 0xed, 0x51, 0x3a, 0xb4, 0x55, 0x86, 0x3d, 0x09, 0xb3, 0xc7, 0x1b, 0xe6,
	0xcd, 0x01, 0xa5, 0x83, 0x30, 0xad, 0xd0, 0x8b, 0x5f, 0x39, 0xc7, 0xcc, 0x1f, 0x8d, 0x30, 0xe3,
	0x49, 0xae, 0xb9, 0x9a, 0xae, 0xfb, 0x23, 0xe2, 0xf8, 0x51, 0x44, 0x85, 0x2f, 0x08, 0x8d, 0xd4,
	0xea, 0x4c, 0x2a, 0x38, 0x12, 0x44, 0x10, 0xcc, 0xa7, 0x83, 0x34, 0xe1, 0xc6, 0xd8, 0x0f, 0x49,
	0xe0, 0x0b, 0xec, 0xa8, 0x41, 0xb2, 0x60, 0x7d, 0xd2, 0x60, 0xf9, 0xe9, 0x5b, 0xc2, 0x85, 0x87,
	0xdf, 0xc4, 0x98, 0x0b, 0xb4, 0x05, 0x25, 0x12, 0x34, 0xb5, 0x35, 0xad, 0x55, 0xef, 0xac, 0xda,
	0x09, 0x0b, 0x5b, 0xb1, 0xb4, 0xbb, 0x82, 0x91, 0x68, 0xf0, 0xdc, 0x0f, 0x63, 0xec, 0xd6, 0x4e,
	0xdc, 0x32, 0xd3, 0xbf, 0x68, 0x9a, 0x57, 0x22, 0x01, 0x7a, 0x0c, 0xd5, 0x3e, 0x8d, 0x04, 0x8e,
	0x44, 0xb3, 0xb4, 0x40, 0x72, 0xf5, 0xc4, 0x35, 0x58, 0x69, 0xc5, 0xf0, 0x54, 0x96, 0x75, 0x17,
	0x1a, 0x29, 0x11, 0x3e, 0xa2, 0x11, 0xc7, 0xe8, 0x3a, 0x54, 0x18, 0xe6, 0x71, 0x28, 0x24, 0x9b,
	0x9a, 0x97, 0x7e, 0x6d, 0xeb, 0xbf, 0x5d, 0xcd, 0xfa, 0xa8, 0x43, 0x7d, 0x3f, 0x43, 0xfb, 0x21,
	0x94, 0x43, 0x72, 0x44, 0xc4, 0x4c, 0xe6, 0x07, 0x7b, 0x91, 0xd8, 0xec, 0x24, 0xe0, 0x4b, 0x27,
	0x6e, 0xa5, 0x6d, 0x34, 0x83, 0x96, 0xe6, 0x25, 0x59, 0xe8, 0x3e, 0x18, 0x23, 0x7f, 0x80, 0x67,
	0x52, 0xcf, 0x66, 0x4f, 0xa8, 0xb7, 0x4b, 0x2d, 0xcd, 0x93, 0x29, 0xe8, 0x1e, 0x18, 0x9c, 0x32,
	0xd1, 0xd4, 0xe7, 0x77, 0xed, 0xc9, 0xc8, 0xac, 0x54, 0xc6, 0x79, 0xa4, 0x42, 0x2e, 0x2c, 0x49,
	0x6b, 0x1d, 0x72, 0xd6, 0x6f, 0x96, 0x17, 0x2f, 0xa1, 0x7b, 0x35, 0x99, 0xd7, 0x65, 0x7d, 0xb4,
	0x03, 0x35, 0xe5, 0x9d, 0x66, 0x45, 0x96, 0x68, 0xd9, 0x67, 0xf8, 0x75, 0xea, 0xa3, 0xf1, 0x86,
	0xbd, 0x9f, 0x4e, 0x7a, 0xd3, 0x4c, 0x8b, 0xc0, 0xf2, 0x7e, 0x76, 0xcf, 0x1e, 0x41, 0x35, 0xd9,
	0x25, 0xde, 0xd4, 0xd6, 0xf4, 0x56, 0xbd, 0x73, 0x6b, 0x5e, 0xd1, 0x17, 0x94, 0x05, 0x9e, 0x4a,
	0x42, 0x57, 0xa1, 0x2c, 0xa8, 0xf0, 0x43, 0xb9, 0x11, 0x0d, 0x2f, 0xf9, 0x48, 0x76, 0xfc, 0xbb,
	0x06, 0xb0, 0x8b, 0xff, 0xba, 0x4f, 0xf3, 0xe2, 0xeb, 0xe7, 0x12, 0xdf, 0x7a, 0x06, 0x75, 0xd9,
	0x4a, 0xaa, 0xda, 0x83, 0x9c, 0xd3, 0x17, 0x15, 0x2d, 0x77, 0x1e, 0xbe, 0x6a, 0xd0, 0x78, 0xc2,
	0xb0, 0x2f, 0xb0, 0x12, 0x28, 0xd3, 0xa8, 0x76, 0xf1, 0x46, 0x4b, 0xe7, 0x73, 0x99, 0x05, 0x35,
	0x3e, 0x24, 0x61, 0x78, 0x48, 0x02, 0xa9, 0xd5, 0xd2, 0x69, 0x50, 0x55, 0x2e, 0xec, 0x05, 0x56,
	0x17, 0xfe, 0x53, 0xcc, 0x2f, 0x4f, 0x8f, 0x5f, 0x1a, 0x34, 0x0e, 0x46, 0x41, 0x46, 0x8f, 0x7f,
	0xd9, 0x30, 0x39, 0x1d, 0x8d, 0xd9, 0x3a, 0xaa, 0x8e, 0x2f, 0x4f, 0xc7, 0x5d, 0x68, 0xec, 0xe0,
	0x10, 0x5f, 0x58, 0xc6, 0x09, 0x3b, 0x55, 0xe8, 0xd2, 0xd8, 0x75, 0xbe, 0x95, 0xa1, 0x3e, 0x99,
	0xed, 0x26, 0xcf, 0x21, 0xe2, 0x50, 0x96, 0x6f, 0x08, 0x3a, 0xf3, 0x2e, 0x4b, 0xdf, 0x5e, 0x3b,
	0xfb, 0xde, 0x99, 0x77, 0x16, 0x88, 0x4c, 0x08, 0x5b, 0xd7, 0x3e, 0xfc, 0xf8, 0xf9, 0xb9, 0xf4,
	0x3f, 0x6a, 0x38, 0xe3, 0x0d, 0x67, 0x02, 0xeb, 0xbc, 0x23, 0xc1, 0x7b, 0x34, 0x04, 0x63, 0x72,
	0x07, 0xa2, 0xf5, 0xa2, 0x4a, 0x99, 0xb7, 0xca, 0x6c, 0xcd, 0x0f, 0x4c, 0x11, 0x57, 0x24, 0x22,
	0xa0, 0x9a, 0x42, 0x44, 0x47, 0xa0, 0xef, 0x62, 0x81, 0x6e, 0x17, 0x95, 0x38, 0xbd, 0x25, 0xcd,
	0xf5, 0xb9, 0x71, 0xc5, 0xbd, 0x09, 0xa8, 0x24, 0x67, 0x13, 0x15, 0xea, 0x94, 0xbb, 0x79, 0xcc,
	0xf6, 0x22, 0xa1, 0x29, 0xee, 0x15, 0x89, 0xdb, 0xb0, 0xa6, 0x1d, 0x6e, 0x6b, 0xed, 0x09, 0x6a,
	0xe2, 0xe4, 0x62, 0xd4, 0xdc, 0xf9, 0x2e, 0x46, 0xcd, 0x1f, 0x0c, 0x85, 0x6a, 0xe6, 0x50, 0xc7,
	0x50, 0x49, 0x1c, 0x5a, 0x8c, 0x9a, 0x3b, 0x0e, 0xc5, 0xa8, 0x79, 0xc3, 0x2b, 0x8d, 0xdb, 0x79,
	0x8d, 0xdd, 0x9d, 0x97, 0xee, 0x80, 0x88, 0xd7, 0x71, 0xcf, 0xee, 0xd3, 0x23, 0x87, 0x63, 0x12,
	0xc4, 0x5b, 0x9d, 0x2d, 0x67, 0x5a, 0xd7, 0xe9, 0xf9, 0xfd, 0x21, 0x8e, 0x82, 0xc2, 0xff, 0x9f,
	0xbd, 0x8a, 0x1c, 0x6f, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x93, 0x67, 0x1c, 0x4e, 0xa6, 0x0a,
	0x00, 0x00,
}
