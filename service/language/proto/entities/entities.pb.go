// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.4
// source: service/language/proto/entities/entities.proto

package entities

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
	types "github.com/infobloxopen/protoc-gen-gorm/types"
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

//
// A skill is a set of lessons focused on a particular topic, such as present
// tense verbs or sports vocabulary. https://duolingo.fandom.com/wiki/Skill
type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *types.UUID           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // primary key
	CreatedAt    *timestamp.Timestamp  `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamp.Timestamp  `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt    *timestamp.Timestamp  `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Dependencies *wrappers.StringValue `protobuf:"bytes,5,opt,name=dependencies,proto3" json:"dependencies,omitempty"`
	Disabled     bool                  `protobuf:"varint,6,opt,name=disabled,proto3" json:"disabled,omitempty"`
	LessonNumber int32                 `protobuf:"varint,7,opt,name=lesson_number,json=lessonNumber,proto3" json:"lesson_number,omitempty"`
	Description  *wrappers.StringValue `protobuf:"bytes,8,opt,name=description,proto3" json:"description,omitempty"`
	Locked       bool                  `protobuf:"varint,9,opt,name=locked,proto3" json:"locked,omitempty"`
	Type         *wrappers.StringValue `protobuf:"bytes,10,opt,name=type,proto3" json:"type,omitempty"`
	Title        string                `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	UrlTitle     string                `protobuf:"bytes,12,opt,name=url_title,json=urlTitle,proto3" json:"url_title,omitempty"`
	Category     *wrappers.StringValue `protobuf:"bytes,13,opt,name=category,proto3" json:"category,omitempty"`
	Index        int32                 `protobuf:"varint,14,opt,name=index,proto3" json:"index,omitempty"`
	Words        []*Word               `protobuf:"bytes,15,rep,name=words,proto3" json:"words,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_language_proto_entities_entities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_service_language_proto_entities_entities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_service_language_proto_entities_entities_proto_rawDescGZIP(), []int{0}
}

func (x *Skill) GetId() *types.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Skill) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Skill) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Skill) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Skill) GetDependencies() *wrappers.StringValue {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

func (x *Skill) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *Skill) GetLessonNumber() int32 {
	if x != nil {
		return x.LessonNumber
	}
	return 0
}

func (x *Skill) GetDescription() *wrappers.StringValue {
	if x != nil {
		return x.Description
	}
	return nil
}

func (x *Skill) GetLocked() bool {
	if x != nil {
		return x.Locked
	}
	return false
}

func (x *Skill) GetType() *wrappers.StringValue {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Skill) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Skill) GetUrlTitle() string {
	if x != nil {
		return x.UrlTitle
	}
	return ""
}

func (x *Skill) GetCategory() *wrappers.StringValue {
	if x != nil {
		return x.Category
	}
	return nil
}

func (x *Skill) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Skill) GetWords() []*Word {
	if x != nil {
		return x.Words
	}
	return nil
}

type Word struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        *types.UUID          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // primary key
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Content   string               `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	AudioSrc  string               `protobuf:"bytes,6,opt,name=audio_src,json=audioSrc,proto3" json:"audio_src,omitempty"`
	Skill     *Skill               `protobuf:"bytes,7,opt,name=Skill,proto3" json:"Skill,omitempty"`
}

func (x *Word) Reset() {
	*x = Word{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_language_proto_entities_entities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Word) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Word) ProtoMessage() {}

func (x *Word) ProtoReflect() protoreflect.Message {
	mi := &file_service_language_proto_entities_entities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Word.ProtoReflect.Descriptor instead.
func (*Word) Descriptor() ([]byte, []int) {
	return file_service_language_proto_entities_entities_proto_rawDescGZIP(), []int{1}
}

func (x *Word) GetId() *types.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Word) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Word) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Word) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Word) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Word) GetAudioSrc() string {
	if x != nil {
		return x.AudioSrc
	}
	return ""
}

func (x *Word) GetSkill() *Skill {
	if x != nil {
		return x.Skill
	}
	return nil
}

// Language Entity
type Language struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *types.UUID          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // primary key
	CreatedAt    *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    *timestamp.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt    *timestamp.Timestamp `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Name         string               `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Abbreviation string               `protobuf:"bytes,6,opt,name=abbreviation,proto3" json:"abbreviation,omitempty"`
	FlagSrc      string               `protobuf:"bytes,7,opt,name=flag_src,json=flagSrc,proto3" json:"flag_src,omitempty"`
	Skills       []*Skill             `protobuf:"bytes,9,rep,name=skills,proto3" json:"skills,omitempty"`
}

func (x *Language) Reset() {
	*x = Language{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_language_proto_entities_entities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_service_language_proto_entities_entities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_service_language_proto_entities_entities_proto_rawDescGZIP(), []int{2}
}

func (x *Language) GetId() *types.UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Language) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Language) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Language) GetDeletedAt() *timestamp.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

func (x *Language) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Language) GetAbbreviation() string {
	if x != nil {
		return x.Abbreviation
	}
	return ""
}

func (x *Language) GetFlagSrc() string {
	if x != nil {
		return x.FlagSrc
	}
	return ""
}

func (x *Language) GetSkills() []*Skill {
	if x != nil {
		return x.Skills
	}
	return nil
}

var File_service_language_proto_entities_entities_proto protoreflect.FileDescriptor

var file_service_language_proto_entities_entities_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x22, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x34, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x74, 0x68, 0x69,
	0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x72, 0x6d, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x29, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x06, 0x0a, 0x05,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55,
	0x55, 0x49, 0x44, 0x42, 0x12, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x12, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x28, 0x01, 0x30, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19,
	0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x43, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x58, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x1d, 0xba, 0xb9, 0x19, 0x19, 0x0a, 0x17, 0x52, 0x15, 0x69,
	0x64, 0x78, 0x5f, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x48, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x0a, 0x00, 0x52, 0x0c, 0x64, 0x65, 0x70,
	0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x08, 0x64, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x42, 0x06, 0xba, 0xb9, 0x19,
	0x02, 0x0a, 0x00, 0x52, 0x08, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2d, 0x0a,
	0x0d, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x0c,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x00, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x00,
	0x52, 0x06, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x00, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x18, 0x32, 0x40, 0x01, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a,
	0x04, 0x18, 0x64, 0x40, 0x01, 0x52, 0x08, 0x75, 0x72, 0x6c, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x44, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x18, 0x64, 0x40, 0x00, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1e, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x48, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x0f,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x42, 0x08,
	0xba, 0xb9, 0x19, 0x04, 0x2a, 0x02, 0x48, 0x01, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x3a,
	0x09, 0xf8, 0x42, 0x01, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0xc3, 0x03, 0x0a, 0x04, 0x57,
	0x6f, 0x72, 0x64, 0x12, 0x34, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x55, 0x49,
	0x44, 0x42, 0x12, 0xba, 0xb9, 0x19, 0x0e, 0x0a, 0x0c, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28,
	0x01, 0x30, 0x01, 0x40, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a,
	0x02, 0x40, 0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x43,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08,
	0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x57, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x1c, 0xba, 0xb9, 0x19, 0x18, 0x0a, 0x16, 0x52, 0x14, 0x69, 0x64, 0x78,
	0x5f, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba,
	0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x28, 0x0a, 0x09, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x73, 0x72, 0x63, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x18, 0xff, 0x01, 0x40, 0x01,
	0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x72, 0x63, 0x12, 0x49, 0x0a, 0x05, 0x53, 0x6b,
	0x69, 0x6c, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x1a, 0x02, 0x38, 0x00, 0x52, 0x05,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x3a, 0x09, 0xf8, 0x42, 0x01, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01,
	0x22, 0xf7, 0x03, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x72, 0x6d,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x42, 0x12, 0xba, 0xb9, 0x19,
	0x0e, 0x0a, 0x0c, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x30, 0x01, 0x40, 0x01, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02, 0x40, 0x01, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x43, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xba, 0xb9, 0x19, 0x04, 0x0a, 0x02,
	0x40, 0x01, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x5b, 0x0a,
	0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x20, 0xba,
	0xb9, 0x19, 0x1c, 0x0a, 0x1a, 0x52, 0x18, 0x69, 0x64, 0x78, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x52,
	0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04,
	0x18, 0x64, 0x40, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x0c, 0x61, 0x62,
	0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0a, 0xba, 0xb9, 0x19, 0x06, 0x0a, 0x04, 0x18, 0x0a, 0x40, 0x01, 0x52, 0x0c, 0x61, 0x62,
	0x62, 0x72, 0x65, 0x76, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x08, 0x66, 0x6c,
	0x61, 0x67, 0x5f, 0x73, 0x72, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xba, 0xb9,
	0x19, 0x07, 0x0a, 0x05, 0x18, 0xff, 0x01, 0x40, 0x01, 0x52, 0x07, 0x66, 0x6c, 0x61, 0x67, 0x53,
	0x72, 0x63, 0x12, 0x4b, 0x0a, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x08, 0xba,
	0xb9, 0x19, 0x04, 0x2a, 0x02, 0x48, 0x01, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x3a,
	0x09, 0xf8, 0x42, 0x01, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x65, 0x69, 0x64, 0x75, 0x36, 0x32,
	0x36, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x62, 0x6f, 0x6f, 0x6b, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_language_proto_entities_entities_proto_rawDescOnce sync.Once
	file_service_language_proto_entities_entities_proto_rawDescData = file_service_language_proto_entities_entities_proto_rawDesc
)

func file_service_language_proto_entities_entities_proto_rawDescGZIP() []byte {
	file_service_language_proto_entities_entities_proto_rawDescOnce.Do(func() {
		file_service_language_proto_entities_entities_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_language_proto_entities_entities_proto_rawDescData)
	})
	return file_service_language_proto_entities_entities_proto_rawDescData
}

var file_service_language_proto_entities_entities_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_language_proto_entities_entities_proto_goTypes = []interface{}{
	(*Skill)(nil),                // 0: micro.service.language.entities.v1.Skill
	(*Word)(nil),                 // 1: micro.service.language.entities.v1.Word
	(*Language)(nil),             // 2: micro.service.language.entities.v1.Language
	(*types.UUID)(nil),           // 3: gorm.types.UUID
	(*timestamp.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*wrappers.StringValue)(nil), // 5: google.protobuf.StringValue
}
var file_service_language_proto_entities_entities_proto_depIdxs = []int32{
	3,  // 0: micro.service.language.entities.v1.Skill.id:type_name -> gorm.types.UUID
	4,  // 1: micro.service.language.entities.v1.Skill.created_at:type_name -> google.protobuf.Timestamp
	4,  // 2: micro.service.language.entities.v1.Skill.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 3: micro.service.language.entities.v1.Skill.deleted_at:type_name -> google.protobuf.Timestamp
	5,  // 4: micro.service.language.entities.v1.Skill.dependencies:type_name -> google.protobuf.StringValue
	5,  // 5: micro.service.language.entities.v1.Skill.description:type_name -> google.protobuf.StringValue
	5,  // 6: micro.service.language.entities.v1.Skill.type:type_name -> google.protobuf.StringValue
	5,  // 7: micro.service.language.entities.v1.Skill.category:type_name -> google.protobuf.StringValue
	1,  // 8: micro.service.language.entities.v1.Skill.words:type_name -> micro.service.language.entities.v1.Word
	3,  // 9: micro.service.language.entities.v1.Word.id:type_name -> gorm.types.UUID
	4,  // 10: micro.service.language.entities.v1.Word.created_at:type_name -> google.protobuf.Timestamp
	4,  // 11: micro.service.language.entities.v1.Word.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 12: micro.service.language.entities.v1.Word.deleted_at:type_name -> google.protobuf.Timestamp
	0,  // 13: micro.service.language.entities.v1.Word.Skill:type_name -> micro.service.language.entities.v1.Skill
	3,  // 14: micro.service.language.entities.v1.Language.id:type_name -> gorm.types.UUID
	4,  // 15: micro.service.language.entities.v1.Language.created_at:type_name -> google.protobuf.Timestamp
	4,  // 16: micro.service.language.entities.v1.Language.updated_at:type_name -> google.protobuf.Timestamp
	4,  // 17: micro.service.language.entities.v1.Language.deleted_at:type_name -> google.protobuf.Timestamp
	0,  // 18: micro.service.language.entities.v1.Language.skills:type_name -> micro.service.language.entities.v1.Skill
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_service_language_proto_entities_entities_proto_init() }
func file_service_language_proto_entities_entities_proto_init() {
	if File_service_language_proto_entities_entities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_language_proto_entities_entities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
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
		file_service_language_proto_entities_entities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Word); i {
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
		file_service_language_proto_entities_entities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Language); i {
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
			RawDescriptor: file_service_language_proto_entities_entities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_service_language_proto_entities_entities_proto_goTypes,
		DependencyIndexes: file_service_language_proto_entities_entities_proto_depIdxs,
		MessageInfos:      file_service_language_proto_entities_entities_proto_msgTypes,
	}.Build()
	File_service_language_proto_entities_entities_proto = out.File
	file_service_language_proto_entities_entities_proto_rawDesc = nil
	file_service_language_proto_entities_entities_proto_goTypes = nil
	file_service_language_proto_entities_entities_proto_depIdxs = nil
}
