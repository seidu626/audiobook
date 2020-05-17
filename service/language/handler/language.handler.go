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

	language_models "github.com/seidu626/audiobook/service/language/models"
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
	model := language_models.Language{}
	model.ID = uuid.FromStringOrNil(req.Id.GetValue())
	name := req.Name.GetValue()
	model.Name = &name
	model.Abbreviation = req.Abbreviation.GetValue()

	exists := h.languageRepository.Exist(&model)
	log.Infof("language exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *languageHandler) List(ctx context.Context, req *languagePB.ListRequest, rsp *languagePB.ListResponse) error {
	log.Info("Received LanguageHandler.List request")
	model := language_models.Language{}
	name := req.Name.GetValue()
	model.Name = &name
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	total, languages, err := h.languageRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue(), &model)
	if err != nil {
		return errors.NotFound("mkit.service.account.language.list", "Error %v", err.Error())
	}
	rsp.Total = total

	// newLanguages := make([]*accountPB.Language, len(languages))
	// for index, language := range languages {
	// 	tmpLanguage, _ := language.ToPB(ctx)
	// 	newLanguages[index] = &tmpLanguage
	// 	// *newLanguages[index], _ = language.ToPB(ctx) ???
	// }
	newLanguages := funk.Map(languages, func(language *language_models.Language) *language_models.Language {
		tmpLanguage, _ := language.ToPB(ctx)
		return &tmpLanguage
	}).([]*language_models.Language)

	rsp.Results = newLanguages
	return nil
}

func (h *languageHandler) Get(ctx context.Context, req *languagePB.GetRequest, rsp *languagePB.GetResponse) error {
	log.Info("Received LanguageHandler.Get request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.account.language.get", "validation error: Missing Id")
	}
	language, err := h.languageRepository.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return myErrors.AppError(myErrors.DBE, err)
	}

	tempLanguage, _ := language.ToPB(ctx)
	rsp.Result = &tempLanguage

	return nil
}

func (h *languageHandler) Create(ctx context.Context, req *languagePB.CreateRequest, rsp *languagePB.CreateResponse) error {
	log.Info("Received LanguageHandler.Create request")

	model := language_models.Language{}
	name := req.Name.GetValue()
	model.Name = &name
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	if err := h.languageRepository.Create(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	// send email (TODO: async `go h.Event.Publish(...)`)
	if err := h.Event.Publish(ctx, &emailerPB.Message{To: req.Email.GetValue()}); err != nil {
		log.WithError(err).Error("Received Event.Publish request error")
		return myErrors.AppError(myErrors.PSE, err)
	}

	// call greeter
	// if res, err := h.greeterSrvClient.Hello(ctx, &greeterPB.Request{Name: req.GetFirstName().GetValue()}); err != nil {
	if res, err := h.greeterSrvClient.Hello(ctx, &greeterPB.HelloRequest{Name: req.GetFirstName().GetValue()}); err != nil {
		log.WithError(err).Error("Received greeterService.Hello request error")
		return myErrors.AppError(myErrors.PSE, err)
	} else {
		log.Infof("Got greeterService responce %s", res.Msg)
	}

	return nil
}

func (h *languageHandler) Update(ctx context.Context, req *languagePB.UpdateRequest, rsp *languagePB.UpdateResponse) error {
	log.Info("Received LanguageHandler.Update request")
	// Identify the language
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("mkit.service.account.language.update", "A valid auth token is required")
	}
	log.Infof("Caller Account: %v", acc)

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.account.language.update", "validation error: Missing Id")
	}

	model := language_models.Language{}
	name := req.Name.GetValue()
	model.Name = &name
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	if err := h.languageRepository.Update(id, &model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *languageHandler) Delete(ctx context.Context, req *languagePB.DeleteRequest, rsp *languagePB.DeleteResponse) error {
	log.Info("Received LanguageHandler.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.account.language.update", "validation error: Missing Id")
	}

	model := language_models.Language{}
	model.Id = uuid.FromStringOrNil(id)

	if err := h.languageRepository.Delete(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}
