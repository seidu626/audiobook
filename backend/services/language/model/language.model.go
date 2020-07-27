package model

import (
	time "time"

	ptypes "github.com/golang/protobuf/ptypes"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	go_uuid1 "github.com/satori/go.uuid"
	pb "github.com/seidu626/audiobook/backend/services/language/proto/entities"
	pbReq "github.com/seidu626/audiobook/backend/services/language/proto/language"
)

// Language model
type Language struct {
	ID           string     `gorm:"primary_key;column:Id;type:uuid;unique;not null;" json:"id"`
	CreatedAt    *time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"not null" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index:idx_languages_deleted_at" json:"deleted_at"`
	Name         string     `gorm:"size:100;not null" json:"name"`
	Abbreviation string     `gorm:"size:100;" json:"abbreviation"`
	FlagSrc      string     `json:"flagSrc"`
	Skills       []*Skill   `gorm:"foreignkey:LanguageID;association_foreignkey:ID;preload:true" json:"skills"`
}

// Languages language collection
type Languages []*Language

// MarshalLanguageCollection convert proto collection to model collection
func MarshalLanguageCollection(languages []*pb.Language) ([]*Language, error) {
	collection := make([]*Language, 0)
	var err error
	for _, language := range languages {
		var model *Language
		if model, err = MarshalLanguage(language); err != nil {
			return collection, err
		}
		collection = append(collection, model)
	}
	return collection, err
}

// UnmarshalLanguageCollection convert model collection to proto collection
func UnmarshalLanguageCollection(languages []*Language) ([]*pb.Language, error) {
	collection := make([]*pb.Language, 0)
	var err error
	for _, language := range languages {
		var model *pb.Language
		if model, err = UnmarshalLanguage(language); err != nil {
			return collection, err
		}
		collection = append(collection, model)
	}
	return collection, err
}

// UnmarshalLanguage convert model to proto
func UnmarshalLanguage(language *Language) (*pb.Language, error) {
	to := &pb.Language{
		Name:         language.Name,
		Abbreviation: language.Abbreviation,
		FlagSrc:      language.FlagSrc,
	}
	var err error

	if language.ID != "" {
		id, err := go_uuid1.FromString(language.ID)
		if err != nil {
			return to, err
		}
		to.Id = id.String()
	} else {
		to.Id = ""
	}
	if language.CreatedAt != nil {
		var t *timestamppb.Timestamp
		if t, err = ptypes.TimestampProto(*language.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = t
	}
	if language.UpdatedAt != nil {
		var t *timestamppb.Timestamp
		if t, err = ptypes.TimestampProto(*language.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = t
	}
	if language.DeletedAt != nil {
		var t *timestamppb.Timestamp
		if t, err = ptypes.TimestampProto(*language.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = t
	}

	if language.Skills != nil {
		skills, err := UnmarshalSkillCollection(language.Skills)
		if err != nil {
			return to, err
		}
		to.Skills = skills
	}

	return to, err
}

func (req *pbReq.CreateRequest) MarshalLanguage() (*Language, error) {
	to := &Language{
		Name:         req.Name.Value,
		Abbreviation: req.Abbreviation.Value,
		FlagSrc:      req.FlagSrc.Value,
	}
	var err error

	if req.Skills != nil {
		skills, err := MarshalSkillCollection(req.Skills)
		if err != nil {
			return to, err
		}
		to.Skills = skills
	}

	return to, err
}

// MarshalLanguage convert proto to model
func (language *pb.Language) MarshalLanguage() (*Language, error) {
	to := &Language{
		ID:           language.Id,
		Name:         language.Name,
		Abbreviation: language.Abbreviation,
		FlagSrc:      language.FlagSrc,
	}
	var err error

	if language.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(language.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if language.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(language.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	if language.DeletedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(language.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = &t
	}

	if language.Skills != nil {
		skills, err := MarshalSkillCollection(language.Skills)
		if err != nil {
			return to, err
		}
		to.Skills = skills
	}

	return to, err
}
