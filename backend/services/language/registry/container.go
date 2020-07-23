package registry

import (
	"github.com/jinzhu/gorm"
	"github.com/sarulabs/di/v2"

	"github.com/seidu626/audiobook/backend/services/language/handler"
	entities "github.com/seidu626/audiobook/backend/services/language/proto/entities"
	"github.com/seidu626/audiobook/backend/services/language/repository"
	"github.com/seidu626/audiobook/backend/shared/database"
	configPB "github.com/seidu626/audiobook/backend/shared/proto/config"
	log "github.com/sirupsen/logrus"
)

// Container - provide di Container
type Container struct {
	ctn di.Container
}

// NewContainer - create new Container
func NewContainer(cfg configPB.Configuration) (*Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal().Err(err).Msg("")
		return nil, err
	}

	if err := builder.Add([]di.Def{
		{
			Name:  "config",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return cfg, nil
			},
		},
		{
			Name:  "language-repository",
			Scope: di.App,
			Build: buildLanguageRepository,
		},
		{
			Name:  "skill-repository",
			Scope: di.App,
			Build: buildSkillRepository,
		},
		{
			Name:  "word-repository",
			Scope: di.App,
			Build: buildWordRepository,
		},
		{
			Name:  "language-handler",
			Scope: di.App,
			Build: buildLanguageHandler,
		},
		{
			Name:  "skill-handler",
			Scope: di.App,
			Build: buildSkillHandler,
		},
		{
			Name:  "word-handler",
			Scope: di.App,
			Build: buildWordHandler,
		},
		{
			Name:  "database",
			Scope: di.App,
			Build: func(ctn di.Container) (interface{}, error) {
				return database.GetDatabaseConnection(cfg.Database, cfg.Log)
			},
			Close: func(obj interface{}) error {
				return obj.(*gorm.DB).Close()
			},
		},
	}...); err != nil {
		return nil, err
	}

	return &Container{
		ctn: builder.Build(),
	}, nil
}

// Resolve object
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

// Clean Container
func (c *Container) Clean() error {
	return c.ctn.Clean()
}

// Delete Container
func (c *Container) Delete() error {
	return c.ctn.Delete()
}

func buildLanguageRepository(ctn di.Container) (interface{}, error) {
	db := ctn.Get("database").(*gorm.DB)
	db.AutoMigrate(&entities.LanguageORM{})
	return repository.NewLanguageRepository(db), nil
}

func buildSkillRepository(ctn di.Container) (interface{}, error) {
	db := ctn.Get("database").(*gorm.DB)
	db.AutoMigrate(&entities.SkillORM{})
	return repository.NewSkillRepository(db), nil
}

func buildWordRepository(ctn di.Container) (interface{}, error) {
	db := ctn.Get("database").(*gorm.DB)
	db.AutoMigrate(&entities.WordORM{})
	return repository.NewWordRepository(db), nil
}

func buildLanguageHandler(ctn di.Container) (interface{}, error) {
	repo := ctn.Get("language-repository").(repository.LanguageRepository)
	return handler.NewLanguageHandler(repo, nil), nil // FIXME inject Publisher, and greeter service
}

func buildSkillHandler(ctn di.Container) (interface{}, error) {
	repo := ctn.Get("skill-repository").(repository.SkillRepository)
	return handler.NewSkillHandler(repo, nil), nil // FIXME inject Publisher, and greeter service
}

func buildWordHandler(ctn di.Container) (interface{}, error) {
	repo := ctn.Get("word-repository").(repository.WordRepository)
	return handler.NewWordHandler(repo, nil), nil // FIXME inject Publisher, and greeter service
}
