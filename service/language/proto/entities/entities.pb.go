// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/language/proto/entities/entities.proto

package entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

//
// A skill is a set of lessons focused on a particular topic, such as present
// tense verbs or sports vocabulary. https://duolingo.fandom.com/wiki/Skill
type Skill struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Dependencies         []string             `protobuf:"bytes,5,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	Disabled             bool                 `protobuf:"varint,6,opt,name=disabled,proto3" json:"disabled,omitempty"`
	LessonNumber         int32                `protobuf:"varint,7,opt,name=lesson_number,json=lessonNumber,proto3" json:"lesson_number,omitempty"`
	Description          string               `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Locked               bool                 `protobuf:"varint,9,opt,name=locked,proto3" json:"locked,omitempty"`
	Type                 string               `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Title                string               `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	UrlTitle             string               `protobuf:"bytes,12,opt,name=url_title,json=urlTitle,proto3" json:"url_title,omitempty"`
	Category             string               `protobuf:"bytes,13,opt,name=category,proto3" json:"category,omitempty"`
	Index                int32                `protobuf:"varint,14,opt,name=index,proto3" json:"index,omitempty"`
	Language             *Language            `protobuf:"bytes,15,opt,name=language,proto3" json:"language,omitempty"`
	Words                []*Word              `protobuf:"bytes,16,rep,name=words,proto3" json:"words,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Skill) Reset()         { *m = Skill{} }
func (m *Skill) String() string { return proto.CompactTextString(m) }
func (*Skill) ProtoMessage()    {}
func (*Skill) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2734188b7083b7d, []int{0}
}

func (m *Skill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Skill.Unmarshal(m, b)
}
func (m *Skill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Skill.Marshal(b, m, deterministic)
}
func (m *Skill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Skill.Merge(m, src)
}
func (m *Skill) XXX_Size() int {
	return xxx_messageInfo_Skill.Size(m)
}
func (m *Skill) XXX_DiscardUnknown() {
	xxx_messageInfo_Skill.DiscardUnknown(m)
}

var xxx_messageInfo_Skill proto.InternalMessageInfo

func (m *Skill) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Skill) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Skill) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Skill) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Skill) GetDependencies() []string {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

func (m *Skill) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Skill) GetLessonNumber() int32 {
	if m != nil {
		return m.LessonNumber
	}
	return 0
}

func (m *Skill) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Skill) GetLocked() bool {
	if m != nil {
		return m.Locked
	}
	return false
}

func (m *Skill) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Skill) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Skill) GetUrlTitle() string {
	if m != nil {
		return m.UrlTitle
	}
	return ""
}

func (m *Skill) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Skill) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Skill) GetLanguage() *Language {
	if m != nil {
		return m.Language
	}
	return nil
}

func (m *Skill) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

type Word struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Content              string               `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	AudioSrc             string               `protobuf:"bytes,6,opt,name=audio_src,json=audioSrc,proto3" json:"audio_src,omitempty"`
	Language             *Language            `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Word) Reset()         { *m = Word{} }
func (m *Word) String() string { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()    {}
func (*Word) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2734188b7083b7d, []int{1}
}

func (m *Word) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Word.Unmarshal(m, b)
}
func (m *Word) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Word.Marshal(b, m, deterministic)
}
func (m *Word) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Word.Merge(m, src)
}
func (m *Word) XXX_Size() int {
	return xxx_messageInfo_Word.Size(m)
}
func (m *Word) XXX_DiscardUnknown() {
	xxx_messageInfo_Word.DiscardUnknown(m)
}

var xxx_messageInfo_Word proto.InternalMessageInfo

func (m *Word) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Word) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Word) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Word) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Word) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Word) GetAudioSrc() string {
	if m != nil {
		return m.AudioSrc
	}
	return ""
}

func (m *Word) GetLanguage() *Language {
	if m != nil {
		return m.Language
	}
	return nil
}

// Language Entity
type Language struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Name                 string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Abbreviation         string               `protobuf:"bytes,6,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	FlagSrc              string               `protobuf:"bytes,7,opt,name=flag_src,json=flagSrc,proto3" json:"flag_src,omitempty"`
	Words                []*Word              `protobuf:"bytes,8,rep,name=words,proto3" json:"words,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Language) Reset()         { *m = Language{} }
func (m *Language) String() string { return proto.CompactTextString(m) }
func (*Language) ProtoMessage()    {}
func (*Language) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2734188b7083b7d, []int{2}
}

func (m *Language) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Language.Unmarshal(m, b)
}
func (m *Language) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Language.Marshal(b, m, deterministic)
}
func (m *Language) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Language.Merge(m, src)
}
func (m *Language) XXX_Size() int {
	return xxx_messageInfo_Language.Size(m)
}
func (m *Language) XXX_DiscardUnknown() {
	xxx_messageInfo_Language.DiscardUnknown(m)
}

var xxx_messageInfo_Language proto.InternalMessageInfo

func (m *Language) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Language) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Language) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Language) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Language) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Language) GetAbbreviation() string {
	if m != nil {
		return m.Abbreviation
	}
	return ""
}

func (m *Language) GetFlagSrc() string {
	if m != nil {
		return m.FlagSrc
	}
	return ""
}

func (m *Language) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

func init() {
	proto.RegisterType((*Skill)(nil), "micro.service.language.entities.v1.Skill")
	proto.RegisterType((*Word)(nil), "micro.service.language.entities.v1.Word")
	proto.RegisterType((*Language)(nil), "micro.service.language.entities.v1.Language")
}

func init() {
	proto.RegisterFile("service/language/proto/entities/entities.proto", fileDescriptor_f2734188b7083b7d)
}

var fileDescriptor_f2734188b7083b7d = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x54, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x55, 0xbb, 0xcd, 0x36, 0x99, 0xee, 0x2e, 0xc8, 0x42, 0xc8, 0x94, 0x03, 0x51, 0xb9, 0xe4,
	0x80, 0x12, 0x51, 0xa4, 0x95, 0x38, 0x00, 0x5a, 0x4e, 0x1c, 0x10, 0x87, 0xec, 0x4a, 0x48, 0x5c,
	0x2a, 0x27, 0x9e, 0x0d, 0x56, 0x9d, 0x38, 0xb2, 0x9d, 0xc2, 0xfe, 0x21, 0x5f, 0xc1, 0x07, 0xf0,
	0x15, 0x28, 0x76, 0x52, 0xd1, 0x53, 0x11, 0x9c, 0xf6, 0xe6, 0xf7, 0x66, 0xde, 0x4b, 0xf2, 0x66,
	0x1c, 0x48, 0x0d, 0xea, 0x9d, 0x28, 0x31, 0x93, 0xac, 0xa9, 0x3a, 0x56, 0x61, 0xd6, 0x6a, 0x65,
	0x55, 0x86, 0x8d, 0x15, 0x56, 0xa0, 0xd9, 0x1f, 0x52, 0xc7, 0x93, 0x55, 0x2d, 0x4a, 0xad, 0x46,
	0x55, 0x3a, 0xaa, 0xd2, 0x7d, 0xdb, 0xee, 0xe5, 0xf2, 0x59, 0xa5, 0x54, 0x25, 0x07, 0xa7, 0xa2,
	0xbb, 0xcd, 0xac, 0xa8, 0xd1, 0x58, 0x56, 0xb7, 0xde, 0x64, 0xf5, 0x6b, 0x06, 0xc1, 0xf5, 0x56,
	0x48, 0x49, 0x2e, 0x60, 0x2a, 0x38, 0x9d, 0xc4, 0x93, 0x24, 0xca, 0xa7, 0x82, 0x93, 0xd7, 0x00,
	0xa5, 0x46, 0x66, 0x91, 0x6f, 0x98, 0xa5, 0xd3, 0x78, 0x92, 0x2c, 0xd6, 0xcb, 0xd4, 0xfb, 0xa5,
	0xa3, 0x5f, 0x7a, 0x33, 0xfa, 0xe5, 0xd1, 0xd0, 0x7d, 0x65, 0x7b, 0x69, 0xd7, 0xf2, 0x51, 0x7a,
	0x72, 0x5c, 0x3a, 0x74, 0x7b, 0x29, 0x47, 0x89, 0x83, 0x74, 0x76, 0x5c, 0x3a, 0x74, 0x5f, 0x59,
	0xb2, 0x82, 0x33, 0x8e, 0x2d, 0x36, 0x1c, 0x9b, 0x52, 0xa0, 0xa1, 0x41, 0x7c, 0x92, 0x44, 0xf9,
	0x01, 0x47, 0x96, 0x10, 0x72, 0x61, 0x58, 0x21, 0x91, 0xd3, 0xd3, 0x78, 0x92, 0x84, 0xf9, 0x1e,
	0x93, 0xe7, 0x70, 0x2e, 0xd1, 0x18, 0xd5, 0x6c, 0x9a, 0xae, 0x2e, 0x50, 0xd3, 0x79, 0x3c, 0x49,
	0x82, 0xfc, 0xcc, 0x93, 0x9f, 0x1c, 0x47, 0x62, 0x58, 0x70, 0x34, 0xa5, 0x16, 0xad, 0x15, 0xaa,
	0xa1, 0xa1, 0x8b, 0xeb, 0x4f, 0x8a, 0x3c, 0x86, 0x53, 0xa9, 0xca, 0x2d, 0x72, 0x1a, 0xb9, 0x07,
	0x0c, 0x88, 0x10, 0x98, 0xd9, 0xbb, 0x16, 0x29, 0x38, 0x89, 0x3b, 0x93, 0x47, 0x10, 0x58, 0x61,
	0x25, 0xd2, 0x85, 0x23, 0x3d, 0x20, 0x4f, 0x21, 0xea, 0xb4, 0xdc, 0xf8, 0xca, 0x99, 0xab, 0x84,
	0x9d, 0x96, 0x37, 0xae, 0xb8, 0x84, 0xb0, 0x64, 0x16, 0x2b, 0xa5, 0xef, 0xe8, 0xb9, 0xaf, 0x8d,
	0xb8, 0xb7, 0x13, 0x0d, 0xc7, 0xef, 0xf4, 0xc2, 0xbd, 0xb9, 0x07, 0xe4, 0x03, 0x84, 0xe3, 0x6e,
	0xd0, 0x07, 0x2e, 0xd0, 0x17, 0xe9, 0xf1, 0xd5, 0x49, 0x3f, 0x0e, 0x64, 0xbe, 0x57, 0x93, 0xb7,
	0x10, 0x7c, 0x53, 0x9a, 0x1b, 0xfa, 0x30, 0x3e, 0x49, 0x16, 0xeb, 0xe4, 0x6f, 0x6c, 0x3e, 0x2b,
	0xcd, 0x73, 0x2f, 0x5b, 0xfd, 0x98, 0xc2, 0xac, 0xc7, 0xf7, 0x7f, 0xd7, 0x28, 0xcc, 0x4b, 0xd5,
	0x58, 0x6c, 0x2c, 0x0d, 0xdc, 0x57, 0x8c, 0xb0, 0x1f, 0x1e, 0xeb, 0xb8, 0x50, 0x1b, 0xa3, 0x4b,
	0xb7, 0x62, 0x51, 0x1e, 0x3a, 0xe2, 0x5a, 0x97, 0x07, 0xa3, 0x98, 0xff, 0xcf, 0x28, 0x56, 0x3f,
	0xa7, 0x10, 0x8e, 0xf4, 0xfd, 0x8f, 0x93, 0xc0, 0xac, 0x61, 0x35, 0x0e, 0x59, 0xba, 0x73, 0x7f,
	0x9d, 0x59, 0x51, 0x68, 0xdc, 0x09, 0xe6, 0xae, 0x9a, 0xcf, 0xf2, 0x80, 0x23, 0x4f, 0x20, 0xbc,
	0x95, 0xac, 0x72, 0x59, 0xcf, 0xfd, 0x1c, 0x7a, 0xdc, 0x47, 0xbd, 0xdf, 0xd5, 0xf0, 0x9f, 0x76,
	0xf5, 0xfd, 0xbb, 0x2f, 0x6f, 0x2a, 0x61, 0xbf, 0x76, 0x45, 0x5a, 0xaa, 0x3a, 0x33, 0x28, 0x78,
	0x77, 0xb9, 0xbe, 0xcc, 0xdc, 0x28, 0x0b, 0xa5, 0xb6, 0xd9, 0x91, 0xbf, 0x75, 0x71, 0xea, 0xf0,
	0xab, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1f, 0xb3, 0x85, 0xc0, 0xd7, 0x05, 0x00, 0x00,
}