package handler

import (
	"context"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/errors"
	uuid "github.com/satori/go.uuid"
	entities "github.com/seidu626/audiobook/service/language/proto/entities"
	skillPB "github.com/seidu626/audiobook/service/language/proto/skill"
	"github.com/seidu626/audiobook/service/language/repository"
	myErrors "github.com/seidu626/audiobook/shared/errors"
	log "github.com/sirupsen/logrus"
	"github.com/thoas/go-funk"
)

// SkillHandler struct
type skillHandler struct {
	skillRepository repository.SkillRepository
	Event           micro.Event
}

// NewSkillHandler returns an instance of `SkillServiceHandler`.
func NewSkillHandler(repo repository.SkillRepository, eve micro.Event) skillPB.SkillServiceHandler {
	return &skillHandler{
		skillRepository: repo,
		Event:           eve,
	}
}

func (h *skillHandler) Exist(ctx context.Context, req *skillPB.ExistRequest, rsp *skillPB.ExistResponse) error {
	log.Info("Received SkillHandler.Exist request")
	model := entities.SkillORM{}
	model.Id = uuid.FromStringOrNil(req.Id.GetValue())
	model.Title = req.Title.GetValue()

	exists := h.skillRepository.Exist(&model)
	log.Infof("skill exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *skillHandler) List(ctx context.Context, req *skillPB.ListRequest, rsp *skillPB.ListResponse) error {
	log.Info("Received SkillHandler.List request")
	total, skills, err := h.skillRepository.List(req.Limit.GetValue(), req.Page.GetValue(), req.Sort.GetValue())
	if err != nil {
		return errors.NotFound("micro.service.skill.skill.list", "Error %v", err.Error())
	}
	rsp.Total = total
	results := funk.Map(skills, func(skill *entities.SkillORM) *entities.Skill {
		tmpModel, _ := skill.ToPB(ctx)
		return &tmpModel
	}).([]*entities.Skill)

	rsp.Results = results
	return nil
}

func (h *skillHandler) Get(ctx context.Context, req *skillPB.GetRequest, rsp *skillPB.GetResponse) error {
	log.Info("Received SkillHandler.Get request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("micro.service.skill.skill.get", "validation error: Missing Id")
	}
	skill, err := h.skillRepository.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return myErrors.AppError(myErrors.DBE, err)
	}

	result, _ := skill.ToPB(ctx)
	rsp.Result = &result

	return nil
}

func (h *skillHandler) Create(ctx context.Context, req *skillPB.CreateRequest, rsp *skillPB.CreateResponse) error {
	log.Info("Received SkillHandler.Create request")
	dependencySlice := []string{}
	for _, dep := range req.Dependencies {
		rec := dep.GetValue()
		dependencySlice = append(dependencySlice, rec)
	}
	dependencies := strings.Join(dependencySlice, ",")
	skillType := req.Type.GetValue()
	category := req.Category.GetValue()
	description := req.Description.GetValue()

	model := entities.SkillORM{}
	model.Title = req.Title.GetValue()
	model.UrlTitle = req.UrlTitle.GetValue()
	model.LessonNumber = req.LessonNumber
	model.Dependencies = &dependencies
	model.Disabled = req.Disabled
	model.Locked = req.Locked
	model.Type = &skillType
	model.Category = &category
	model.Index = int32(req.Index)
	model.Description = &description
	languageId := uuid.FromStringOrNil(req.LanguageId.GetValue())
	model.LanguageId = &languageId

	if err := h.skillRepository.Create(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *skillHandler) Update(ctx context.Context, req *skillPB.UpdateRequest, rsp *skillPB.UpdateResponse) error {
	log.Info("Received SkillHandler.Update request")
	// Identify the skill
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("micro.service.skill.skill.update", "A valid auth token is required")
	}
	log.Infof("Caller Account: %v", acc)

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("micro.service.skill.skill.update", "validation error: Missing Id")
	}

	dependencySlice := []string{}
	for _, dep := range req.Dependencies {
		rec := dep.GetValue()
		dependencySlice = append(dependencySlice, rec)
	}
	dependencies := strings.Join(dependencySlice, ",")
	skillType := req.Type.GetValue()
	category := req.Category.GetValue()
	description := req.Description.GetValue()

	model := entities.SkillORM{}
	model.Id = uuid.FromStringOrNil(id)
	model.Title = req.Title.GetValue()
	model.UrlTitle = req.UrlTitle.GetValue()
	model.LessonNumber = req.LessonNumber
	model.Dependencies = &dependencies
	model.Disabled = req.Disabled
	model.Locked = req.Locked
	model.Type = &skillType
	model.Category = &category
	model.Index = int32(req.Index)
	model.Description = &description
	languageId := uuid.FromStringOrNil(req.LanguageId.GetValue())
	model.LanguageId = &languageId

	if err := h.skillRepository.Update(id, &model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *skillHandler) Delete(ctx context.Context, req *skillPB.DeleteRequest, rsp *skillPB.DeleteResponse) error {
	log.Info("Received SkillHandler.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("micro.service.skill.skill.update", "validation error: Missing Id")
	}

	model := entities.SkillORM{}
	model.Id = uuid.FromStringOrNil(id)

	if err := h.skillRepository.Delete(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}
