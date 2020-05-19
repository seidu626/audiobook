package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	entities "github.com/seidu626/audiobook/service/language/proto/entities"
	log "github.com/sirupsen/logrus"
)

// SkillRepository interface
type SkillRepository interface {
	Exist(model *entities.SkillORM) bool
	List(limit, page uint32, sort string) (total uint32, skills []*entities.SkillORM, err error)
	Get(id string) (*entities.SkillORM, error)
	Create(model *entities.SkillORM) error
	Update(id string, model *entities.SkillORM) error
	Delete(model *entities.SkillORM) error
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
func (repo *skillRepository) Exist(model *entities.SkillORM) bool {
	log.Infof("Received skillRepository.Exist request %v", *model)
	var count int
	if model.Title != "" && len(model.Title) > 0 {
		repo.db.Model(&entities.SkillORM{}).Where("name = ?", model.Title).Count(&count)
		if count > 0 {
			return true
		}
	}
	if len(model.Id) > 0 {
		repo.db.Model(&entities.SkillORM{}).Where("id = ?", model.Id).Count(&count)
		if count > 0 {
			return true
		}
	}
	if model.UrlTitle != "" {
		repo.db.Model(&entities.SkillORM{}).Where("code = ?", model.UrlTitle).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *skillRepository) List(limit, page uint32, sort string) (total uint32, skills []*entities.SkillORM, err error) {
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
func (repo *skillRepository) Get(id string) (skill *entities.SkillORM, err error) {
	u2, err := uuid.FromString(id)
	if err != nil {
		return
	}
	skill = &entities.SkillORM{Id: u2.String()}
	// enable auto preloading for `Profile`
	if err = repo.db.Set("gorm:auto_preload", true).First(skill).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("Error in SkillRepository.Get")
	}
	return
}

// Create
func (repo *skillRepository) Create(model *entities.SkillORM) error {
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
func (repo *skillRepository) Update(id string, model *entities.SkillORM) error {
	u2, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	skill := &entities.SkillORM{
		Id: u2.String(),
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
func (repo *skillRepository) Delete(model *entities.SkillORM) error {
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
