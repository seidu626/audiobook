package handler

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/errors"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"

	entities "github.com/seidu626/audiobook/service/language/proto/entities"
	languagePB "github.com/seidu626/audiobook/service/language/proto/language"
	"github.com/seidu626/audiobook/service/language/repository"
	myErrors "github.com/seidu626/audiobook/shared/errors"
)

// LanguageHandler struct
type languageHandler struct {
	languageRepository repository.LanguageRepository
	Event              micro.Event
}

// NewLanguageHandler returns an instance of `LanguageServiceHandler`.
func NewLanguageHandler(repo repository.LanguageRepository, eve micro.Event) languagePB.LanguageServiceHandler {
	return &languageHandler{
		languageRepository: repo,
		Event:              eve,
	}
}

func (h *languageHandler) Exist(ctx context.Context, req *languagePB.ExistRequest, rsp *languagePB.ExistResponse) error {
	log.Info("Received LanguageHandler.Exist request")
	model := entities.LanguageORM{}
	model.Id = uuid.FromStringOrNil(req.Id.GetValue()).String()
	model.Name = req.Name.GetValue()
	model.Abbreviation = req.Abbreviation.GetValue()

	exists := h.languageRepository.Exist(&model)
	log.Infof("language exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *languageHandler) List(ctx context.Context, req *languagePB.ListRequest, rsp *languagePB.ListResponse) error {
	log.Info("Received LanguageHandler.List request")

	total, languages, err := h.languageRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue())
	if err != nil {
		return errors.NotFound("micro.service.language.language.list", "Error %v", err.Error())
	}
	rsp.Total = total
	results := funk.Map(languages, func(language *entities.LanguageORM) *entities.Language {
		tmpModel, _ := language.ToPB(ctx)
		return &tmpModel
	}).([]*entities.Language)

	rsp.Results = results
	return nil
}

func (h *languageHandler) Get(ctx context.Context, req *languagePB.GetRequest, rsp *languagePB.GetResponse) error {
	log.Info("Received LanguageHandler.Get request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("micro.service.language.language.get", "validation error: Missing Id")
	}
	language, err := h.languageRepository.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return myErrors.AppError(myErrors.DBE, err)
	}
	result, _ := language.ToPB(ctx)

	rsp.Result = &result

	return nil
}

func (h *languageHandler) Create(ctx context.Context, req *languagePB.CreateRequest, rsp *languagePB.CreateResponse) error {
	log.Info("Received LanguageHandler.Create request")

	model := entities.LanguageORM{}
	model.Name = req.Name.GetValue()
	model.Abbreviation = req.Abbreviation.GetValue()
	model.FlagSrc = req.FlagSrc.GetValue()

	if err := h.languageRepository.Create(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *languageHandler) Update(ctx context.Context, req *languagePB.UpdateRequest, rsp *languagePB.UpdateResponse) error {
	log.Info("Received LanguageHandler.Update request")
	// Identify the language
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("micro.service.language.language.update", "A valid auth token is required")
	}
	log.Infof("Caller Account: %v", acc)

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("micro.service.language.language.update", "validation error: Missing Id")
	}

	model := entities.LanguageORM{}
	model.Id = id
	model.Name = req.Name.GetValue()
	model.Abbreviation = req.Abbreviation.GetValue()
	model.FlagSrc = req.FlagSrc.GetValue()

	if err := h.languageRepository.Update(id, &model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *languageHandler) Delete(ctx context.Context, req *languagePB.DeleteRequest, rsp *languagePB.DeleteResponse) error {
	log.Info("Received LanguageHandler.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("micro.service.language.language.update", "validation error: Missing Id")
	}

	model := entities.LanguageORM{}
	model.Id = uuid.FromStringOrNil(id).String()

	if err := h.languageRepository.Delete(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}
