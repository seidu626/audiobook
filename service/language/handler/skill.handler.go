package handler

import (
	"context"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/auth"
	"github.com/micro/go-micro/v2/errors"
	uuid "github.com/satori/go.uuid"
	language_models "github.com/seidu626/audiobook/service/language/model"
	skillPB "github.com/seidu626/audiobook/service/language/proto/skill"
	"github.com/seidu626/audiobook/service/language/repository"
	myErrors "github.com/seidu626/audiobook/shared/errors"
	log "github.com/sirupsen/logrus"
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
	model := language_models.Skill{}
	model.ID = uuid.FromStringOrNil(req.Id.GetValue()).String()
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

	rsp.Results = language_models.UnmarshalSkillCollection(skills)
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

	rsp.Result = language_models.UnmarshalSkill(skill)

	return nil
}

func (h *skillHandler) Create(ctx context.Context, req *skillPB.CreateRequest, rsp *skillPB.CreateResponse) error {
	log.Info("Received SkillHandler.Create request")

	model := language_models.Skill{}
	model.Title = req.Title.GetValue()
	model.URLTitle = req.UrlTitle.GetValue()
	model.LessonNumber = req.LessonNumber
	model.Dependencies = strings.Join(req.Dependencies, ",")
	model.Disabled = req.Disabled
	model.Locked = req.Locked
	model.Type = req.Type
	model.Category = req.Category
	model.Index = int32(req.Index)
	model.Description = req.Description
	model.LanguageID = req.LanguageId.GetValue()

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

	model := language_models.Skill{}
	model.ID = id
	model.Title = req.Title.GetValue()
	model.URLTitle = req.UrlTitle.GetValue()
	model.LessonNumber = req.LessonNumber
	model.Dependencies = strings.Join(req.Dependencies, ",")
	model.Disabled = req.Disabled
	model.Locked = req.Locked
	model.Type = req.Type
	model.Category = req.Category
	model.Index = int32(req.Index)
	model.Description = req.Description
	model.LanguageID = req.LanguageId.GetValue()

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

	model := language_models.Skill{}
	model.ID = uuid.FromStringOrNil(id).String()

	if err := h.skillRepository.Delete(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}
