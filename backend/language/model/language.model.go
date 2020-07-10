package model

import (
	"time"

	ptypes "github.com/golang/protobuf/ptypes"
	pb "github.com/seidu626/audiobook/backend/language/proto/entities"
)

// Language model
type Language struct {
	ID           string    `json:id`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
	FlagSrc      string    `json:"flagSrc"`
	Words        []*Word   `json:"words"`
}

// Languages language collection
type Languages []*Language

// MarshalLanguageCollection convert proto collection to model collection
func MarshalLanguageCollection(languages []*pb.Language) []*Language {
	collection := make([]*Language, 0)
	for _, language := range languages {
		collection = append(collection, MarshalLanguage(language))
	}
	return collection
}

// UnmarshalLanguageCollection convert model collection to proto collection
func UnmarshalLanguageCollection(languages []*Language) []*pb.Language {
	collection := make([]*pb.Language, 0)
	for _, language := range languages {
		collection = append(collection, UnmarshalLanguage(language))
	}
	return collection
}

// UnmarshalLanguage convert model to proto
func UnmarshalLanguage(language *Language) *pb.Language {
	createdAt, _ := ptypes.TimestampProto(language.CreatedAt)
	updatedAt, _ := ptypes.TimestampProto(language.UpdatedAt)
	deletedAt, _ := ptypes.TimestampProto(language.DeletedAt)
	return &pb.Language{
		Id:           language.ID,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeletedAt:    deletedAt,
		Name:         language.Name,
		Abbreviation: language.Abbreviation,
		FlagSrc:      language.FlagSrc,
		Words:        UnmarshalWordCollection(language.Words),
	}
}

// MarshalLanguage convert proto to model
func MarshalLanguage(language *pb.Language) *Language {
	createdAt, _ := ptypes.Timestamp(language.CreatedAt)
	updatedAt, _ := ptypes.Timestamp(language.UpdatedAt)
	deletedAt, _ := ptypes.Timestamp(language.DeletedAt)
	return &Language{
		ID:           language.Id,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
		DeletedAt:    deletedAt,
		Name:         language.Name,
		Abbreviation: language.Abbreviation,
		FlagSrc:      language.FlagSrc,
		Words:        MarshalWordCollection(language.Words),
	}
}
