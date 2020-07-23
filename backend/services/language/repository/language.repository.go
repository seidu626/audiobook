package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	entities "github.com/seidu626/audiobook/backend/services/language/proto/entities"
	log "github.com/sirupsen/logrus"
)

// LanguageRepository interface
type LanguageRepository interface {
	Exist(model *entities.LanguageORM) bool
	List(limit, page uint32, sort string) (total uint32, languages []*entities.LanguageORM, err error)
	Get(id string) (*entities.LanguageORM, error)
	Create(model *entities.LanguageORM) error
	Update(id string, model *entities.LanguageORM) error
	Delete(model *entities.LanguageORM) error
}

// languageRepository struct
type languageRepository struct {
	db *gorm.DB
}

// NewLanguageRepository returns an instance of `LanguageRepository`.
func NewLanguageRepository(db *gorm.DB) LanguageRepository {
	return &languageRepository{
		db: db,
	}
}

// Exist
func (repo *languageRepository) Exist(model *entities.LanguageORM) bool {
	log.Infof("Received languageRepository.Exist request %v", *model)
	var count int
	if model.Name != "" && len(model.Name) > 0 {
		repo.db.Model(&entities.LanguageORM{}).Where("name = ?", model.Name).Count(&count)
		if count > 0 {
			return true
		}
	}
	if len(model.Id) > 0 {
		repo.db.Model(&entities.LanguageORM{}).Where("id = ?", model.Id).Count(&count)
		if count > 0 {
			return true
		}
	}
	if model.Abbreviation != "" {
		repo.db.Model(&entities.LanguageORM{}).Where("abbreviation = ?", model.Abbreviation).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *languageRepository) List(limit, page uint32, sort string) (total uint32, languages []*entities.LanguageORM, err error) {
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
	if err = db.Set("gorm:auto_preload", true).Order(sort).Limit(limit).Offset(offset).Find(&languages).Count(&total).Error; err != nil {
		log.WithError(err).Error("Error in LanguageRepository.List")
		return
	}
	log.Error(languages)
	return
}

// Find by Id
func (repo *languageRepository) Get(id string) (language *entities.LanguageORM, err error) {
	u2, err := uuid.FromString(id)
	if err != nil {
		return
	}
	language = &entities.LanguageORM{Id: u2}
	// enable auto preloading for `Profile`
	if err = repo.db.Set("gorm:auto_preload", true).First(language).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("Error in LanguageRepository.Get")
	}
	return
}

// Create
func (repo *languageRepository) Create(model *entities.LanguageORM) error {
	if exist := repo.Exist(model); exist {
		return errors.New("language already exist")
	}
	// if err := repo.db.Set("gorm:association_autoupdate", false).Create(model).Error; err != nil {
	if err := repo.db.Create(model).Error; err != nil {
		log.WithError(err).Error("Error in LanguageRepository.Create")
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *languageRepository) Update(id string, model *entities.LanguageORM) error {
	u2, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	language := &entities.LanguageORM{
		Id: u2,
	}
	// result := repo.db.Set("gorm:association_autoupdate", false).Save(model)
	result := repo.db.Model(language).Updates(model)
	if err := result.Error; err != nil {
		log.WithError(err).Error("Error in LanguageRepository.Update")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Errorf("Error in LanguageRepository.Update, rowsAffected: %v", rowsAffected)
		return errors.New("no records updated, No match was found")
	}
	return nil
}

// Delete
func (repo *languageRepository) Delete(model *entities.LanguageORM) error {
	result := repo.db.Delete(model)
	if err := result.Error; err != nil {
		log.WithError(err).Error("Error in LanguageRepository.Delete")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Errorf("Error in LanguageRepository.Delete, rowsAffected: %v", rowsAffected)
		return errors.New("no records deleted, No match was found")
	}
	return nil
}
