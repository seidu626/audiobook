package main

import (
	"path/filepath"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config"
	"github.com/seidu626/audiobook/logger/log"

	// "github.com/micro/go-micro/v2/service/grpc"
	"github.com/seidu626/audiobook/service/language/handler"
	languagePB "github.com/seidu626/audiobook/service/language/proto/language"
	skillPB "github.com/seidu626/audiobook/service/language/proto/skill"
	"github.com/seidu626/audiobook/service/language/registry"
	"github.com/seidu626/audiobook/service/language/repository"
	myConfig "github.com/seidu626/audiobook/shared/config"
	"github.com/seidu626/audiobook/shared/constants"
	"github.com/seidu626/audiobook/shared/logger"
	"github.com/seidu626/audiobook/shared/util"
	logWrapper "github.com/seidu626/audiobook/shared/wrapper/log"
	transWrapper "github.com/seidu626/audiobook/shared/wrapper/transaction"
	validatorWrapper "github.com/seidu626/audiobook/shared/wrapper/validator"
)

const (
	serviceName = constants.ACCOUNT_SERVICE
	configDir   = "/config"
	configFile  = "config.yaml"
)

var (
	cfg myConfig.ServiceConfiguration
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version(myConfig.Version),
	)

	// Initialize service
	service.Init(
		micro.Action(func(c *cli.Context) (err error) {
			// load config
			myConfig.InitConfig(configDir, configFile)
			err = config.Scan(&cfg)
			logger.InitLogger(cfg.Log)
			return
		}),
	)

	// Initialize Features
	var options []micro.Option
	if cfg.Features["mtls"].Enabled {
		// if tlsConf, err := util.GetSelfSignedTLSConfig("localhost"); err != nil {
		if tlsConf, err := util.GetTLSConfig(
			filepath.Join(configDir, config.Get("features", "mtls", "certfile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "keyfile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "cafile").String("")),
			filepath.Join(configDir, config.Get("features", "mtls", "servername").String("")),
		); err != nil {
			log.WithError(err).Error("unable to load certs")
		} else {
			println(tlsConf)
			options = append(options,
				util.WithTLS(tlsConf),
			)
		}
	}
	// Wrappers are invoked in the order as they added
	if cfg.Features["reqlogs"].Enabled {
		options = append(options,
			micro.WrapHandler(logWrapper.NewHandlerWrapper()),
			micro.WrapClient(logWrapper.NewClientWrapper()),
		)
	}
	if cfg.Features["validator"].Enabled {
		options = append(options,
			micro.WrapHandler(validatorWrapper.NewHandlerWrapper()),
		)
	}
	if cfg.Features["translogs"].Enabled {
		topic := config.Get("features", "translogs", "topic").String("mkit.service.recorder")
		publisher := micro.NewEvent(topic, service.Client())
		options = append(options,
			micro.WrapHandler(transWrapper.NewHandlerWrapper(publisher)),
		)
	}

	// Initialize Features
	service.Init(
		options...,
	)

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatalf("failed to build container: %v", err)
	}

	log.Debugf("Client type: grpc or regular? %T\n", service.Client()) // FIXME: expected *grpc.grpcClient but got *micro.clientWrapper

	// Publisher publish to "mkit.service.emailer"
	publisher := micro.NewEvent(constants.EMAILER_SERVICE, service.Client())

	// // Handlers
	languageHandler := handler.NewLanguageHandler(ctn.Resolve("language-repository").(repository.LanguageRepository), publisher)
	skillHandler := ctn.Resolve("skill-handler").(skillPB.SkillServiceHandler)

	// Register Handlers
	languagePB.RegisterLanguageServiceHandler(service.Server(), languageHandler)
	skillPB.RegisterSkillServiceHandler(service.Server(), skillHandler)

	myConfig.PrintBuildInfo()
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
