package model

import (
	"time"

	ptypes "github.com/golang/protobuf/ptypes"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/golang/protobuf/ptypes/wrappers"
	go_uuid1 "github.com/satori/go.uuid"
	pb "github.com/seidu626/audiobook/backend/services/language/proto/entities"
)

// Skill model
type Skill struct {
	ID           string     `gorm:"primary_key;column:Id;type:uuid;unique;not null;" json:"id"`
	CreatedAt    *time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"not null" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index:idx_skills_deleted_at" json:"deleted_at"`
	Title        string     `gorm:"not null" json:"title"`
	URLTitle     string     `json:"url_title"`
	LessonNumber int32      `gorm:"not null" json:"lesson_number"`
	Dependencies string     `json:"dependencies"`
	Disabled     bool       `json:"disabled"`
	Locked       bool       `json:"locked"`
	Type         string     `json:"type"`
	Category     string     `json:"category"`
	Index        int32      `json:"index"`
	LanguageID   string     `gorm:"not null" json:"language_id"`
	Description  string     `json:"description"`
	Words        []*Word    `gorm:"foreignkey:SkillID;association_foreignkey:ID;preload:true" json:"words"`
}

// Skills model
type Skills []*Skill

// MarshalSkillCollection convert proto collection to model collection
func MarshalSkillCollection(skills []*pb.Skill) ([]*Skill, error) {
	collection := make([]*Skill, 0)
	var err error
	for _, skill := range skills {
		var model *Skill
		if model, err = MarshalSkill(skill); err != nil {
			return collection, err
		}
		collection = append(collection, model)
	}
	return collection, err
}

// UnmarshalSkillCollection convert model collection to proto collection
func UnmarshalSkillCollection(skills []*Skill) ([]*pb.Skill, error) {
	collection := make([]*pb.Skill, 0)
	var err error
	for _, skill := range skills {
		var model *pb.Skill
		if model, err = UnmarshalSkill(skill); err != nil {
			return collection, err
		}
		collection = append(collection, model)
	}
	return collection, err
}

// UnmarshalSkill convert model to proto
func UnmarshalSkill(skill *Skill) (*pb.Skill, error) {
	var err error
	to := &pb.Skill{
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

	if skill.ID != "" {
		id, err := go_uuid1.FromString(skill.ID)
		if err != nil {
			return to, err
		}
		to.Id = id.String()
	} else {
		to.Id = ""
	}
	if skill.CreatedAt != nil {
		var t *timestamppb.Timestamp
		if t, err = ptypes.TimestampProto(*skill.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = t
	}
	if skill.UpdatedAt != nil {
		var t *timestamppb.Timestamp
		if t, err = ptypes.TimestampProto(*skill.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = t
	}
	if skill.DeletedAt != nil {
		var t *timestamppb.Timestamp
		if t, err = ptypes.TimestampProto(*skill.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = t
	}
	if skill.Words != nil {
		words, err := UnmarshalWordCollection(skill.Words)
		if err != nil {
			return to, err
		}
		to.Words = words
	}

	return to, err
}

// MarshalSkill convert proto to model
func MarshalSkill(skill *pb.Skill) (*Skill, error) {
	var err error
	to := &Skill{
		ID:           skill.Id,
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

	if skill.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(skill.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if skill.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(skill.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	if skill.DeletedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(skill.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = &t
	}
	if skill.Words != nil {
		words, err := MarshalWordCollection(skill.Words)
		if err != nil {
			return to, err
		}
		to.Words = words
	}

	return to, err
}
