package handler

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/errors"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	language_models "github.com/seidu626/audiobook/service/language/model"
	wordPB "github.com/seidu626/audiobook/service/language/proto/word"
	"github.com/seidu626/audiobook/service/language/repository"
	myErrors "github.com/seidu626/audiobook/shared/errors"
)

// WordHandler struct
type wordHandler struct {
	wordRepository repository.WordRepository
	Event          micro.Event
}

// NewWordHandler returns an instance of `WordServiceHandler`.
func NewWordHandler(repo repository.WordRepository, eve micro.Event) wordPB.WordServiceHandler {
	return &wordHandler{
		wordRepository: repo,
		Event:          eve,
	}
}

func (h *wordHandler) Exist(ctx context.Context, req *wordPB.ExistRequest, rsp *wordPB.ExistResponse) error {
	log.Info("Received WordHandler.Exist request")
	model := language_models.Word{}
	model.ID = uuid.FromStringOrNil(req.Id.GetValue()).String()
	model.Content = req.Content.GetValue()

	exists := h.wordRepository.Exist(&model)
	log.Infof("word exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *wordHandler) List(ctx context.Context, req *wordPB.ListRequest, rsp *wordPB.ListResponse) error {
	log.Info("Received WordHandler.List request")
	total, words, err := h.wordRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue())
	if err != nil {
		return errors.NotFound("micro.service.word.word.list", "Error %v", err.Error())
	}
	rsp.Total = total

	rsp.Results = language_models.UnmarshalWordCollection(words)
	return nil
}

func (h *wordHandler) Get(ctx context.Context, req *wordPB.GetRequest, rsp *wordPB.GetResponse) error {
	log.Info("Received WordHandler.Get request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("micro.service.word.word.get", "validation error: Missing Id")
	}
	word, err := h.wordRepository.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return myErrors.AppError(myErrors.DBE, err)
	}

	rsp.Result = language_models.UnmarshalWord(word)

	return nil
}

func (h *wordHandler) Create(ctx context.Context, req *wordPB.CreateRequest, rsp *wordPB.CreateResponse) error {
	log.Info("Received WordHandler.Create request")

	model := language_models.Word{}
	model.Content = req.Content.GetValue()
	model.AudioSrc = req.AudioSrc.GetValue()
	model.LanguageID = req.LanguageId.GetValue()

	if err := h.wordRepository.Create(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *wordHandler) Update(ctx context.Context, req *wordPB.UpdateRequest, rsp *wordPB.UpdateResponse) error {
	log.Info("Received WordHandler.Update request")
	// Identify the word
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("micro.service.word.word.update", "A valid auth token is required")
	}
	log.Infof("Caller Account: %v", acc)

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("micro.service.word.word.update", "validation error: Missing Id")
	}

	model := language_models.Word{}
	model.ID = id
	model.Content = req.Content.GetValue()
	model.AudioSrc = req.AudioSrc.GetValue()
	model.LanguageID = req.LanguageId.GetValue()

	if err := h.wordRepository.Update(id, &model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *wordHandler) Delete(ctx context.Context, req *wordPB.DeleteRequest, rsp *wordPB.DeleteResponse) error {
	log.Info("Received WordHandler.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("micro.service.word.word.update", "validation error: Missing Id")
	}

	model := language_models.Word{}
	model.ID = uuid.FromStringOrNil(id).String()

	if err := h.wordRepository.Delete(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}
