package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	model "github.com/seidu626/audiobook/backend/services/langauge/model"
	log "github.com/sirupsen/logrus"
)

// WordRepository interface
type WordRepository interface {
	Exist(model *model.Word) bool
	List(limit, page uint32, sort string) (total uint32, words []*model.Word, err error)
	Get(id string) (*model.Word, error)
	Create(model *model.Word) error
	Update(id string, model *model.Word) error
	Delete(model *model.Word) error
}

// wordRepository struct
type wordRepository struct {
	db *gorm.DB
}

// NewWordRepository returns an instance of `WordRepository`.
func NewWordRepository(db *gorm.DB) WordRepository {
	return &wordRepository{
		db: db,
	}
}

// Exist
func (repo *wordRepository) Exist(model *model.Word) bool {
	log.Infof("Received wordRepository.Exist request %v", *model)
	var count int
	if model.Content != "" && len(model.Content) > 0 {
		repo.db.Model(&model.Word{}).Where("name = ?", model.Content).Count(&count)
		if count > 0 {
			return true
		}
	}
	if len(model.Id) > 0 {
		repo.db.Model(&model.Word{}).Where("id = ?", model.Id).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *wordRepository) List(limit, page uint32, sort string) (total uint32, words []*model.Word, err error) {
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
	if err = db.Set("gorm:auto_preload", true).Order(sort).Limit(limit).Offset(offset).Find(&words).Count(&total).Error; err != nil {
		log.WithError(err).Error("Error in WordRepository.List")
		return
	}
	return
}

// Find by Id
func (repo *wordRepository) Get(id string) (word *model.Word, err error) {
	u2, err := uuid.FromString(id)
	if err != nil {
		return
	}
	word = &model.Word{Id: u2}
	// enable auto preloading for `Profile`
	if err = repo.db.Set("gorm:auto_preload", true).First(word).Error; err != nil && err != gorm.ErrRecordNotFound {
		log.WithError(err).Error("Error in WordRepository.Get")
	}
	return
}

// Create
func (repo *wordRepository) Create(model *model.Word) error {
	if exist := repo.Exist(model); exist {
		return errors.New("word already exist")
	}
	// if err := repo.db.Set("gorm:association_autoupdate", false).Create(model).Error; err != nil {
	if err := repo.db.Create(model).Error; err != nil {
		log.WithError(err).Error("Error in WordRepository.Create")
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *wordRepository) Update(id string, model *model.Word) error {
	u2, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	word := &model.Word{
		Id: u2,
	}
	// result := repo.db.Set("gorm:association_autoupdate", false).Save(model)
	result := repo.db.Model(word).Updates(model)
	if err := result.Error; err != nil {
		log.WithError(err).Error("Error in WordRepository.Update")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Errorf("Error in WordRepository.Update, rowsAffected: %v", rowsAffected)
		return errors.New("no records updated, No match was found")
	}
	return nil
}

// Delete
func (repo *wordRepository) Delete(model *model.Word) error {
	result := repo.db.Delete(model)
	if err := result.Error; err != nil {
		log.WithError(err).Error("Error in WordRepository.Delete")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		log.Errorf("Error in WordRepository.Delete, rowsAffected: %v", rowsAffected)
		return errors.New("no records deleted, No match was found")
	}
	return nil
}
