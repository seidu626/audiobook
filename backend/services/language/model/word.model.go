package model

import (
	"time"

	ptypes "github.com/golang/protobuf/ptypes"
	timestamppb "github.com/golang/protobuf/ptypes/timestamp"
	go_uuid1 "github.com/satori/go.uuid"
	pb "github.com/seidu626/audiobook/backend/services/language/proto/entities"
)

// Word model
type Word struct {
	ID        string     `gorm:"primary_key;column:Id;type:uuid;unique;not null;" json:"id"`
	CreatedAt *time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"not null" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index:idx_words_deleted_at" json:"deleted_at"`
	Content   string     `gorm:"not null" json:"content"`
	AudioSrc  string     `json:"audio_src"`
	SkillID   string     `gorm:"not null" json:"skillId"`
}

// Words word collection
type Words []*Word

// MarshalWordCollection convert proto collection to model collection
func MarshalWordCollection(words []*pb.Word) ([]*Word, error) {
	collection := make([]*Word, 0)
	var err error
	for _, word := range words {
		var model *Word
		if model, err = MarshalWord(word); err != nil {
			return collection, err
		}
		collection = append(collection, model)
	}
	return collection, err
}

// UnmarshalWordCollection convert model collection to proto collection
func UnmarshalWordCollection(words []*Word) ([]*pb.Word, error) {
	collection := make([]*pb.Word, 0)
	var err error
	for _, word := range words {
		var model *pb.Word
		if model, err = UnmarshalWord(word); err != nil {
			return collection, err
		}
		collection = append(collection, model)
	}
	return collection, err
}

// UnmarshalWord convert model to proto
func UnmarshalWord(word *Word) (*pb.Word, error) {
	var err error

	to := &pb.Word{
		Content:  word.Content,
		AudioSrc: word.AudioSrc,
		Skill_Id: word.SkillID,
	}
	if word.ID != "" {
		id, err := go_uuid1.FromString(word.ID)
		if err != nil {
			return to, err
		}
		to.Id = id.String()
	} else {
		to.Id = ""
	}
	if word.CreatedAt != nil {
		var t *timestamppb.Timestamp
		if t, err = ptypes.TimestampProto(*word.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = t
	}
	if word.UpdatedAt != nil {
		var t *timestamppb.Timestamp
		if t, err = ptypes.TimestampProto(*word.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = t
	}
	if word.DeletedAt != nil {
		var t *timestamppb.Timestamp
		if t, err = ptypes.TimestampProto(*word.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = t
	}

	return to, err
}

// MarshalWord convert proto to model
func MarshalWord(word *pb.Word) (*Word, error) {
	var err error
	to := &Word{
		ID:       word.Id,
		Content:  word.Content,
		AudioSrc: word.AudioSrc,
		SkillID:  word.Skill_Id,
	}

	if word.CreatedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(word.CreatedAt); err != nil {
			return to, err
		}
		to.CreatedAt = &t
	}
	if word.UpdatedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(word.UpdatedAt); err != nil {
			return to, err
		}
		to.UpdatedAt = &t
	}
	if word.DeletedAt != nil {
		var t time.Time
		if t, err = ptypes.Timestamp(word.DeletedAt); err != nil {
			return to, err
		}
		to.DeletedAt = &t
	}
	return to, err
}
