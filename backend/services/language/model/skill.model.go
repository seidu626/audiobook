package model

import (
	"time"

	ptypes "github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/seidu626/audiobook/backend/services/language/proto/entities"
)

// Skill model
type Skill struct {
	ID           string    `gorm:"primary_key;column:Id;type:STRING;" json:"id" db:"Id" protobuf:"string,0,opt,name=id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	Title        string    `json:"title"`
	URLTitle     string    `json:"url_title"`
	LessonNumber int32     `json:"lesson_number"`
	Dependencies string    `json:"dependencies"`
	Disabled     bool      `json:"disabled"`
	Locked       bool      `json:"locked"`
	Type         string    `json:"type"`
	Category     string    `json:"category"`
	Index        int32     `json:"index"`
	LanguageID   string    `json:"language_id"`
	Description  string    `json:"description"`
	Language     *Language `json:"language"`
}

// Skills model
type Skills []*Skill

// MarshalSkillCollection convert proto collection to model collection
func MarshalSkillCollection(skills []*pb.Skill) []*Skill {
	collection := make([]*Skill, 0)
	for _, skill := range skills {
		collection = append(collection, MarshalSkill(skill))
	}
	return collection
}

// UnmarshalSkillCollection convert model collection to proto collection
func UnmarshalSkillCollection(skills []*Skill) []*pb.Skill {
	collection := make([]*pb.Skill, 0)
	for _, skill := range skills {
		collection = append(collection, UnmarshalSkill(skill))
	}
	return collection
}

// UnmarshalSkill convert model to proto
func UnmarshalSkill(skill *Skill) *pb.Skill {
	createdAt, _ := ptypes.TimestampProto(skill.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(skill.UpdatedAt)
	deletedAt, _ := ptypes.TimestampProto(skill.DeletedAt)
	//languageID := &wrappers.StringValue{Value: skill.LanguageID} //https://iximiuz.com/en/posts/truly-optional-scalar-types-in-protobuf3/
	return &pb.Skill{
		Id:           skill.ID,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeletedAt:    deletedAt,
		Title:        skill.Title,
		UrlTitle:     skill.URLTitle,
		LessonNumber: skill.LessonNumber,
		Dependencies: &wrappers.StringValue{Value: skill.Dependencies},
		Locked:       skill.Locked,
		Type:         &wrappers.StringValue{Value: skill.Type},
		Category:     &wrappers.StringValue{Value: skill.Category},
		Index:        skill.Index,
		Description:  &wrappers.StringValue{Value: skill.Description},
	}
}

// MarshalSkill convert proto to model
func MarshalSkill(skill *pb.Skill) *Skill {
	createdAt, _ := ptypes.Timestamp(skill.CreatedAt)
	updatedAt, _ := ptypes.Timestamp(skill.UpdatedAt)
	deletedAt, _ := ptypes.Timestamp(skill.DeletedAt)
	// strings.Join(skill.Dependencies.GetValue(), ","),
	return &Skill{
		ID:           skill.Id,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeletedAt:    deletedAt,
		Title:        skill.Title,
		URLTitle:     skill.UrlTitle,
		LessonNumber: skill.LessonNumber,
		Dependencies: skill.Dependencies.Value,
		//		Disabled:     skill.Disabled,
		Locked:      skill.Locked,
		Type:        skill.Type.GetValue(),
		Category:    skill.Category.GetValue(),
		Index:       skill.Index,
		Description: skill.Description.GetValue(),
	}
}
