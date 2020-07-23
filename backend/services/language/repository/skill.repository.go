package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	model "github.com/seidu626/audiobook/backend/services/langauge/model"
	log "github.com/sirupsen/logrus"
)

// SkillRepository interface
type SkillRepository interface {
	Exist(model *model.Skill) bool
	List(limit, page uint32, sort string) (total uint32, skills []*model.Skill, err error)
	Get(id string) (*model.Skill, error)
	Create(model *model.Skill) error
	Update(id string, model *model.Skill) error
	Delete(model *model.Skill) error
}

// skillRepository struct
type skillRepository struct {
	db *gorm.DB
}

// NewSkillRepository returns an instance of `SkillRepository`.
func NewSkillRepository(db *gorm.DB) SkillRepository {
	return &skillRepository{
		db: db,
	}
}

// Exist
func (repo *skillRepository) Exist(model *model.Skill) bool {
	log.Infof("Received skillRepository.Exist request %v", *model)
	var count int
	if model.Title != "" && len(model.Title) > 0 {
		repo.db.Model(&model.Skill{}).Where("name = ?", model.Title).Count(&count)
		if count > 0 {
			return true
		}
	}
	if len(model.Id) > 0 {
		repo.db.Model(&model.Skill{}).Where("id = ?", model.Id).Count(&count)
		if count > 0 {
			return true
		}
	}
	if model.UrlTitle != "" {
		repo.db.Model(&model.Skill{}).Where("code = ?", model.UrlTitle).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *skillRepository) List(limit, page uint32, sort string) (total uint32, skills []*model.Skill, err error) {
	db := repo.db

	if limit == 0 {
		limit = 10
	}
	var offset uint32
	if page > 1 {
		offset = (page - 1) * limit
	} else {
		offset = 0
	}
	if sort == "" {
		sort = "created_at desc"
	}

	// enable auto preloading for `Profile`
	if err = db.Set("gorm:auto_preload", true).Order(sort).Limit(limit).Offset(offset).Find(&skills).Count(&total).Error; err != nil {
		log.WithError(err).Error("Error in SkillRepository.List")
		return
	}
	return
}

// Find by Id
func (repo *skillRepository) Get(id string) (skill *model.Skill, err error) {
	u2, err := uuid.FromString(id)
	if err != nil {
		return
	}
	skill = &model.Skill{Id: u2}
	// enable auto preloading for `Profile`
	if err = repo.db.Set("gorm:auto_preload", true).First(skill).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("Error in SkillRepository.Get")
	}
	return
}

// Create
func (repo *skillRepository) Create(model *model.Skill) error {
	if exist := repo.Exist(model); exist {
		return errors.New("skill already exist")
	}
	// if err := repo.db.Set("gorm:association_autoupdate", false).Create(model).Error; err != nil {
	if err := repo.db.Create(model).Error; err != nil {
		log.WithError(err).Error("Error in SkillRepository.Create")
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *skillRepository) Update(id string, model *model.Skill) error {
	u2, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	skill := &model.Skill{
		Id: u2,
	}
	// result := repo.db.Set("gorm:association_autoupdate", false).Save(model)
	result := repo.db.Model(skill).Updates(model)
	if err := result.Error; err != nil {
		log.WithError(err).Error("Error in SkillRepository.Update")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Errorf("Error in SkillRepository.Update, rowsAffected: %v", rowsAffected)
		return errors.New("no records updated, No match was found")
	}
	return nil
}

// Delete
func (repo *skillRepository) Delete(model *model.Skill) error {
	result := repo.db.Delete(model)
	if err := result.Error; err != nil {
		log.WithError(err).Error("Error in SkillRepository.Delete")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Errorf("Error in SkillRepository.Delete, rowsAffected: %v", rowsAffected)
		return errors.New("no records deleted, No match was found")
	}
	return nil
}
