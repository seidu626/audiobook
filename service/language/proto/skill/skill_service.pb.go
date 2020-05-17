// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/language/proto/skill/skill_service.proto

package skill

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
	Title                *wrappers.StringValue `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ExistRequest) Reset()         { *m = ExistRequest{} }
func (m *ExistRequest) String() string { return proto.CompactTextString(m) }
func (*ExistRequest) ProtoMessage()    {}
func (*ExistRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{0}
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

func (m *ExistRequest) GetTitle() *wrappers.StringValue {
	if m != nil {
		return m.Title
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
	return fileDescriptor_e027081fdca39ddd, []int{1}
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
	Title                *wrappers.StringValue `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	UrlTitle             *wrappers.StringValue `protobuf:"bytes,5,opt,name=url_title,json=urlTitle,proto3" json:"url_title,omitempty"`
	LessonNumber         *wrappers.StringValue `protobuf:"bytes,6,opt,name=lesson_number,json=lessonNumber,proto3" json:"lesson_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{2}
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

func (m *ListRequest) GetTitle() *wrappers.StringValue {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *ListRequest) GetUrlTitle() *wrappers.StringValue {
	if m != nil {
		return m.UrlTitle
	}
	return nil
}

func (m *ListRequest) GetLessonNumber() *wrappers.StringValue {
	if m != nil {
		return m.LessonNumber
	}
	return nil
}

type ListResponse struct {
	Results              []*entities.Skill `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Total                uint32            `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{3}
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

func (m *ListResponse) GetResults() []*entities.Skill {
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
	Title                *wrappers.StringValue `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	UrlTitle             *wrappers.StringValue `protobuf:"bytes,3,opt,name=url_title,json=urlTitle,proto3" json:"url_title,omitempty"`
	LessonNumber         *wrappers.StringValue `protobuf:"bytes,4,opt,name=lesson_number,json=lessonNumber,proto3" json:"lesson_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{4}
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

func (m *GetRequest) GetTitle() *wrappers.StringValue {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *GetRequest) GetUrlTitle() *wrappers.StringValue {
	if m != nil {
		return m.UrlTitle
	}
	return nil
}

func (m *GetRequest) GetLessonNumber() *wrappers.StringValue {
	if m != nil {
		return m.LessonNumber
	}
	return nil
}

type GetResponse struct {
	Result               *entities.Skill `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{5}
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

func (m *GetResponse) GetResult() *entities.Skill {
	if m != nil {
		return m.Result
	}
	return nil
}

type CreateRequest struct {
	Title                *wrappers.StringValue `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	UrlTitle             *wrappers.StringValue `protobuf:"bytes,3,opt,name=url_title,json=urlTitle,proto3" json:"url_title,omitempty"`
	LessonNumber         *wrappers.StringValue `protobuf:"bytes,4,opt,name=lesson_number,json=lessonNumber,proto3" json:"lesson_number,omitempty"`
	Dependencies         []string              `protobuf:"bytes,5,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Disabled             bool                  `protobuf:"varint,6,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Description          string                `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Locked               bool                  `protobuf:"varint,8,opt,name=locked,proto3" json:"locked,omitempty"`
	Type                 string                `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	Category             string                `protobuf:"bytes,10,opt,name=category,proto3" json:"category,omitempty"`
	Index                int32                 `protobuf:"varint,11,opt,name=index,proto3" json:"index,omitempty"`
	LanguageId           *wrappers.StringValue `protobuf:"bytes,12,opt,name=language_id,json=languageId,proto3" json:"language_id,omitempty"`
	Language             *entities.Language    `protobuf:"bytes,13,opt,name=language,proto3" json:"language,omitempty"`
	Words                []*entities.Word      `protobuf:"bytes,14,rep,name=words,proto3" json:"words,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{6}
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

func (m *CreateRequest) GetTitle() *wrappers.StringValue {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *CreateRequest) GetUrlTitle() *wrappers.StringValue {
	if m != nil {
		return m.UrlTitle
	}
	return nil
}

func (m *CreateRequest) GetLessonNumber() *wrappers.StringValue {
	if m != nil {
		return m.LessonNumber
	}
	return nil
}

func (m *CreateRequest) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *CreateRequest) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *CreateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateRequest) GetLocked() bool {
	if m != nil {
		return m.Locked
	}
	return false
}

func (m *CreateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CreateRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *CreateRequest) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *CreateRequest) GetLanguageId() *wrappers.StringValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

func (m *CreateRequest) GetLanguage() *entities.Language {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *CreateRequest) GetWords() []*entities.Word {
	if m != nil {
		return m.Words
	}
	return nil
}

type CreateResponse struct {
	Result               *entities.Skill `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{7}
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

func (m *CreateResponse) GetResult() *entities.Skill {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                *wrappers.StringValue `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	UrlTitle             *wrappers.StringValue `protobuf:"bytes,3,opt,name=url_title,json=urlTitle,proto3" json:"url_title,omitempty"`
	LessonNumber         *wrappers.StringValue `protobuf:"bytes,4,opt,name=lesson_number,json=lessonNumber,proto3" json:"lesson_number,omitempty"`
	Dependencies         []string              `protobuf:"bytes,5,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Disabled             bool                  `protobuf:"varint,6,opt,name=disabled,proto3" json:"disabled,omitempty"`
	Description          string                `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	Locked               bool                  `protobuf:"varint,8,opt,name=locked,proto3" json:"locked,omitempty"`
	Type                 string                `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	Category             string                `protobuf:"bytes,10,opt,name=category,proto3" json:"category,omitempty"`
	Index                int32                 `protobuf:"varint,11,opt,name=index,proto3" json:"index,omitempty"`
	LanguageId           *wrappers.StringValue `protobuf:"bytes,12,opt,name=language_id,json=languageId,proto3" json:"language_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{8}
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

func (m *UpdateRequest) GetTitle() *wrappers.StringValue {
	if m != nil {
		return m.Title
	}
	return nil
}

func (m *UpdateRequest) GetUrlTitle() *wrappers.StringValue {
	if m != nil {
		return m.UrlTitle
	}
	return nil
}

func (m *UpdateRequest) GetLessonNumber() *wrappers.StringValue {
	if m != nil {
		return m.LessonNumber
	}
	return nil
}

func (m *UpdateRequest) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *UpdateRequest) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *UpdateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateRequest) GetLocked() bool {
	if m != nil {
		return m.Locked
	}
	return false
}

func (m *UpdateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdateRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *UpdateRequest) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *UpdateRequest) GetLanguageId() *wrappers.StringValue {
	if m != nil {
		return m.LanguageId
	}
	return nil
}

type UpdateResponse struct {
	Result               *entities.Skill `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{9}
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

func (m *UpdateResponse) GetResult() *entities.Skill {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteRequest struct {
	Id                   *wrappers.StringValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                *wrappers.StringValue `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{10}
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

func (m *DeleteRequest) GetTitle() *wrappers.StringValue {
	if m != nil {
		return m.Title
	}
	return nil
}

type DeleteResponse struct {
	Result               *entities.Skill `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e027081fdca39ddd, []int{11}
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

func (m *DeleteResponse) GetResult() *entities.Skill {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*ExistRequest)(nil), "micro.service.language.skill.v1.ExistRequest")
	proto.RegisterType((*ExistResponse)(nil), "micro.service.language.skill.v1.ExistResponse")
	proto.RegisterType((*ListRequest)(nil), "micro.service.language.skill.v1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "micro.service.language.skill.v1.ListResponse")
	proto.RegisterType((*GetRequest)(nil), "micro.service.language.skill.v1.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "micro.service.language.skill.v1.GetResponse")
	proto.RegisterType((*CreateRequest)(nil), "micro.service.language.skill.v1.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "micro.service.language.skill.v1.CreateResponse")
	proto.RegisterType((*UpdateRequest)(nil), "micro.service.language.skill.v1.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "micro.service.language.skill.v1.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "micro.service.language.skill.v1.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "micro.service.language.skill.v1.DeleteResponse")
}

func init() {
	proto.RegisterFile("service/language/proto/skill/skill_service.proto", fileDescriptor_e027081fdca39ddd)
}

var fileDescriptor_e027081fdca39ddd = []byte{
	// 875 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x57, 0x5f, 0x8b, 0x23, 0x45,
	0x10, 0x67, 0x32, 0x33, 0xd9, 0xa4, 0x92, 0x1c, 0x4b, 0x23, 0x32, 0xc4, 0x7f, 0x21, 0xf8, 0x30,
	0x77, 0x26, 0x93, 0xdd, 0x2c, 0x2e, 0xdc, 0xc1, 0x09, 0xe6, 0x94, 0xf3, 0x64, 0x11, 0x99, 0xbd,
	0x55, 0x51, 0x34, 0x4e, 0xd2, 0x75, 0xb9, 0x66, 0x67, 0xa7, 0xc7, 0xee, 0x9e, 0xbd, 0x5b, 0x45,
	0xf0, 0x83, 0x28, 0xbe, 0x8b, 0x0f, 0x7e, 0x3c, 0xc9, 0xc3, 0x21, 0xe9, 0x9e, 0xc9, 0x26, 0x42,
	0xcc, 0x9c, 0xe4, 0xe1, 0xe0, 0xee, 0x65, 0xe9, 0xea, 0xad, 0x5f, 0x55, 0x57, 0xd5, 0xaf, 0x2a,
	0x35, 0x70, 0x20, 0x51, 0x5c, 0xb2, 0x29, 0x0e, 0xe2, 0x28, 0x99, 0x65, 0xd1, 0x0c, 0x07, 0xa9,
	0xe0, 0x8a, 0x0f, 0xe4, 0x39, 0x8b, 0x63, 0xf3, 0x77, 0x9c, 0xab, 0x04, 0xfa, 0x3f, 0xe4, 0x9d,
	0x0b, 0x36, 0x15, 0x3c, 0x28, 0x2e, 0x0b, 0x5c, 0xa0, 0x75, 0x83, 0xcb, 0xc3, 0xf6, 0xdb, 0x33,
	0xce, 0x67, 0x71, 0x6e, 0x68, 0x92, 0x3d, 0x1a, 0x3c, 0x11, 0x51, 0x9a, 0xa2, 0x90, 0xc6, 0x40,
	0x3b, 0xd8, 0xe0, 0x12, 0x13, 0xc5, 0x14, 0x43, 0xb9, 0x3c, 0xe4, 0xfa, 0x37, 0xd5, 0x63, 0x26,
	0xe8, 0x38, 0x8d, 0x84, 0xba, 0xca, 0x55, 0x2f, 0xa3, 0x98, 0xd1, 0x48, 0xe1, 0xf2, 0x60, 0x54,
	0xbb, 0xbf, 0x5a, 0xd0, 0xfc, 0xf8, 0x29, 0x93, 0x2a, 0xc4, 0x1f, 0x32, 0x94, 0x8a, 0x1c, 0x43,
	0x85, 0x51, 0xcf, 0xea, 0x58, 0x7e, 0x63, 0xf8, 0x66, 0x60, 0x1e, 0x16, 0x14, 0x0f, 0x0b, 0x4e,
	0x95, 0x60, 0xc9, 0xec, 0x8b, 0x28, 0xce, 0x70, 0x54, 0x9b, 0x8f, 0x5c, 0x61, 0xff, 0x65, 0x59,
	0x61, 0x85, 0x51, 0xf2, 0x10, 0x5c, 0xc5, 0x54, 0x8c, 0x5e, 0xa5, 0x04, 0xb4, 0x33, 0x1f, 0xbd,
	0x25, 0xde, 0xd8, 0x77, 0xbc, 0x7d, 0xff, 0x97, 0xca, 0x90, 0x7c, 0xf7, 0x4d, 0xd4, 0xff, 0xf1,
	0xa0, 0x7f, 0x7b, 0xdc, 0xff, 0xf6, 0xa7, 0xa3, 0xde, 0xe1, 0xfb, 0x3f, 0xbf, 0x1b, 0x1a, 0x63,
	0xdd, 0x1e, 0xb4, 0xf2, 0xd7, 0xc9, 0x94, 0x27, 0x12, 0xc9, 0xeb, 0x50, 0x15, 0x28, 0xb3, 0x58,
	0xe9, 0x27, 0xd6, 0xc2, 0x5c, 0xba, 0x63, 0xff, 0x3d, 0xb2, 0xba, 0xbf, 0xdb, 0xd0, 0x38, 0x59,
	0x89, 0xe5, 0x2e, 0xb8, 0x31, 0xbb, 0x60, 0x6a, 0x63, 0x38, 0x67, 0x0f, 0x12, 0x75, 0x34, 0x34,
	0x6f, 0xaa, 0xcf, 0x47, 0xd5, 0x5b, 0x8e, 0x47, 0x7d, 0x2b, 0x34, 0x28, 0x72, 0x1b, 0x9c, 0x34,
	0x9a, 0x6d, 0x8e, 0x68, 0x15, 0xbd, 0x37, 0x1f, 0x39, 0xb7, 0x2a, 0xbe, 0x15, 0x6a, 0x08, 0x39,
	0x00, 0x47, 0x72, 0xa1, 0x3c, 0x7b, 0x7b, 0x32, 0x42, 0xad, 0x79, 0x9d, 0x3f, 0x67, 0x87, 0xf9,
	0x23, 0x23, 0xa8, 0x67, 0x22, 0x1e, 0x1b, 0xcb, 0x6e, 0x09, 0xcb, 0x8b, 0x38, 0x44, 0x65, 0xdf,
	0x0e, 0x6b, 0x99, 0x88, 0x1f, 0x6a, 0x1b, 0x9f, 0x42, 0x2b, 0x46, 0x29, 0x79, 0x32, 0x4e, 0xb2,
	0x8b, 0x09, 0x0a, 0xaf, 0xfa, 0x3c, 0x76, 0x9a, 0x06, 0xfb, 0x99, 0x86, 0x76, 0x13, 0x68, 0x9e,
	0xac, 0x96, 0xf3, 0x1e, 0xec, 0x99, 0x02, 0x4a, 0xcf, 0xea, 0xd8, 0x7e, 0x63, 0x78, 0x33, 0xd8,
	0xd0, 0x2c, 0x4b, 0x8a, 0x5f, 0x1e, 0x06, 0xa7, 0x8b, 0xc6, 0x09, 0x0b, 0x24, 0x79, 0x0d, 0x5c,
	0xc5, 0x55, 0x14, 0xeb, 0x42, 0xb5, 0x42, 0x23, 0x18, 0x46, 0xfc, 0x51, 0x01, 0xb8, 0x8f, 0x2f,
	0x26, 0xb9, 0xd7, 0x8b, 0x63, 0xef, 0xa8, 0x38, 0xce, 0xff, 0x2f, 0xce, 0x19, 0x34, 0x74, 0xae,
	0xf2, 0xda, 0x7c, 0xb8, 0xd6, 0x6a, 0xcf, 0x55, 0x9a, 0xb5, 0xae, 0x7c, 0xe6, 0x40, 0xeb, 0x9e,
	0xc0, 0x48, 0x61, 0x51, 0x86, 0x97, 0x22, 0x9d, 0xa4, 0x0b, 0x4d, 0x8a, 0x29, 0x26, 0x14, 0x93,
	0x29, 0x43, 0xe9, 0xb9, 0x1d, 0xdb, 0xaf, 0x87, 0x6b, 0x77, 0xa4, 0x0d, 0x35, 0xca, 0x64, 0x34,
	0x89, 0x91, 0xea, 0xb6, 0xaa, 0x85, 0x4b, 0x99, 0x74, 0xa0, 0x41, 0x51, 0x4e, 0x05, 0x4b, 0x15,
	0xe3, 0x89, 0xb7, 0xd7, 0xb1, 0xfc, 0x7a, 0xb8, 0x7a, 0xb5, 0x18, 0x86, 0x31, 0x9f, 0x9e, 0x23,
	0xf5, 0x6a, 0x66, 0x18, 0x1a, 0x89, 0x10, 0x70, 0xd4, 0x55, 0x8a, 0x5e, 0x5d, 0x43, 0xf4, 0x79,
	0xe1, 0x69, 0x1a, 0x29, 0x9c, 0x71, 0x71, 0xe5, 0x81, 0xbe, 0x5f, 0xca, 0x8b, 0x06, 0x62, 0x09,
	0xc5, 0xa7, 0x5e, 0xa3, 0x63, 0xf9, 0x6e, 0x68, 0x04, 0x72, 0x17, 0x1a, 0x45, 0x89, 0xc7, 0x8c,
	0x7a, 0xcd, 0x12, 0xa3, 0x0c, 0x0a, 0xc0, 0x03, 0x4a, 0x3e, 0x81, 0x5a, 0x21, 0x79, 0x2d, 0x8d,
	0xed, 0x95, 0x21, 0xd0, 0x49, 0x7e, 0x19, 0x2e, 0xd1, 0xe4, 0x03, 0x70, 0x9f, 0x70, 0x41, 0xa5,
	0x77, 0x43, 0x8f, 0x08, 0xbf, 0x8c, 0x99, 0x2f, 0xb9, 0xa0, 0xa1, 0x81, 0x75, 0xbf, 0x82, 0x1b,
	0x05, 0xff, 0x76, 0x4c, 0xed, 0x3f, 0x1d, 0x68, 0x9d, 0xa5, 0x74, 0x85, 0xda, 0xaf, 0x26, 0xcc,
	0xab, 0x96, 0xd8, 0xd4, 0x12, 0x0b, 0x22, 0x16, 0x6c, 0xd9, 0x31, 0x11, 0x7f, 0xb3, 0xa0, 0xf5,
	0x11, 0xc6, 0xf8, 0x82, 0x12, 0x71, 0x11, 0x79, 0xf1, 0xbc, 0xdd, 0x46, 0x3e, 0x7c, 0xe6, 0x40,
	0x53, 0x5f, 0x9f, 0x1a, 0x20, 0x79, 0x04, 0xae, 0x5e, 0x19, 0x49, 0x3f, 0xd8, 0xb2, 0x77, 0x07,
	0xab, 0x8b, 0x6f, 0x3b, 0x28, 0xab, 0x9e, 0x07, 0x30, 0x05, 0x67, 0xb1, 0xca, 0x90, 0xde, 0x56,
	0xdc, 0xca, 0x4a, 0xda, 0xee, 0x97, 0xd4, 0xce, 0x9d, 0x7c, 0x0f, 0xf6, 0x7d, 0x54, 0xe4, 0xbd,
	0xad, 0xa8, 0xeb, 0x25, 0xa7, 0xdd, 0x2b, 0xa7, 0x9c, 0x7b, 0x60, 0x50, 0x35, 0xc3, 0x91, 0x6c,
	0x4f, 0xc0, 0xda, 0xaf, 0x78, 0x7b, 0x50, 0x5a, 0xff, 0xda, 0x95, 0xa1, 0x7f, 0x09, 0x57, 0x6b,
	0x53, 0xb5, 0x84, 0xab, 0x7f, 0xf5, 0x15, 0x83, 0xaa, 0xe1, 0x5b, 0x09, 0x57, 0x6b, 0x7d, 0x53,
	0xc2, 0xd5, 0x3a, 0x91, 0x47, 0x63, 0xd8, 0xf6, 0x7d, 0xf7, 0xb9, 0xf5, 0xf5, 0x9d, 0x19, 0x53,
	0x8f, 0xb3, 0x49, 0x30, 0xe5, 0x17, 0x03, 0x89, 0x8c, 0x66, 0xc7, 0xc3, 0xe3, 0x41, 0x94, 0x51,
	0xc6, 0x27, 0x9c, 0x9f, 0x0f, 0xfe, 0xeb, 0x93, 0x72, 0x52, 0xd5, 0xc2, 0xd1, 0x3f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x5d, 0xbb, 0x39, 0x9d, 0x79, 0x0e, 0x00, 0x00,
}
