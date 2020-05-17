package models

import (
	"time"

	ptypes "github.com/golang/protobuf/ptypes"
	pb "github.com/seidu626/audiobook/service/language/proto/entities"
)

// Word model
type Word struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	Content    string    `json:"content"`
	AudioSrc   string    `json:"audio_src"`
	LanguageID string    `json:"language_id"`
}

// Words word collection
type Words []*Word

// MarshalWordCollection convert proto collection to model collection
func MarshalWordCollection(words []*pb.Word) []*Word {
	collection := make([]*Word, 0)
	for _, word := range words {
		collection = append(collection, MarshalWord(word))
	}
	return collection
}

// UnmarshalWordCollection convert model collection to proto collection
func UnmarshalWordCollection(words []*Word) []*pb.Word {
	collection := make([]*pb.Word, 0)
	for _, word := range words {
		collection = append(collection, UnmarshalWord(word))
	}
	return collection
}

// UnmarshalWord convert model to proto
func UnmarshalWord(word *Word) *pb.Word {
	createdAt, _ := ptypes.TimestampProto(word.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(word.UpdatedAt)
	deletedAt, _ := ptypes.TimestampProto(word.DeletedAt)
	return &pb.Word{
		Id:         word.ID,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeletedAt:  deletedAt,
		Content:    word.Content,
		AudioSrc:   word.AudioSrc,
		LanguageID: word.LanguageID,
	}
}

// MarshalWord convert proto to model
func MarshalWord(word *pb.Word) *Word {
	createdAt, _ := ptypes.Timestamp(word.CreatedAt)
	updatedAt, _ := ptypes.Timestamp(word.UpdatedAt)
	deletedAt, _ := ptypes.Timestamp(word.DeletedAt)
	return &Word{
		ID:         word.Id,
		CreatedAt:  createdAt,
		UpdatedAt:  updatedAt,
		DeletedAt:  deletedAt,
		Content:    word.Content,
		AudioSrc:   word.AudioSrc,
		LanguageID: word.LanguageID,
	}
}