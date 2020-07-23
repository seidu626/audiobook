package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
	"github.com/rs/zerolog/log"

	"github.com/seidu626/audiobook/backend/services/language/handler"
	languagePB "github.com/seidu626/audiobook/backend/services/language/proto/language"
	skillPB "github.com/seidu626/audiobook/backend/services/language/proto/skill"
	wordPB "github.com/seidu626/audiobook/backend/services/language/proto/word"
	"github.com/seidu626/audiobook/backend/services/language/registry"
	"github.com/seidu626/audiobook/backend/services/language/repository"
	"github.com/seidu626/audiobook/backend/shared/config"
	"github.com/seidu626/audiobook/backend/shared/constants"
	myMicro "github.com/seidu626/audiobook/backend/shared/util/micro"
	logWrapper "github.com/seidu626/audiobook/backend/shared/wrapper/log"
	transWrapper "github.com/seidu626/audiobook/backend/shared/wrapper/transaction"
	validatorWrapper "github.com/seidu626/audiobook/backend/shared/wrapper/validator"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://1c23b49583a64ceebb5af93ab9974377@o336647.ingest.sentry.io/5303538",
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
	})

	if err != nil {
		log.Fatal().Msgf("sentry.Init: %s", err)
	}

	cfg := config.GetConfig()

	// Initialize Features
	var clientWrappers []client.Wrapper
	var handlerWrappers []server.HandlerWrapper
	var subscriberWrappers []server.SubscriberWrapper

	// Wrappers are invoked in the order as they added
	if cfg.Features.Reqlogs.Enabled {
		clientWrappers = append(clientWrappers, logWrapper.NewClientWrapper())
		handlerWrappers = append(handlerWrappers, logWrapper.NewHandlerWrapper())
		subscriberWrappers = append(subscriberWrappers, logWrapper.NewSubscriberWrapper())
	}
	//if cfg.Features.Translogs.Enabled {
	//    topic := cfg.Features.Translogs.Topic
	//    publisher := micro.NewEvent(topic, client.DefaultClient) // service.Client())
	//    handlerWrappers = append(handlerWrappers, transWrapper.NewHandlerWrapper(publisher))
	//    subscriberWrappers = append(subscriberWrappers, transWrapper.NewSubscriberWrapper(publisher))
	//}
	if cfg.Features.Validator.Enabled {
		handlerWrappers = append(handlerWrappers, validatorWrapper.NewHandlerWrapper())
		subscriberWrappers = append(subscriberWrappers, validatorWrapper.NewSubscriberWrapper())
	}

	service := micro.NewService(
		micro.Name(constants.ACCOUNT_SERVICE),
		micro.Version(config.Version),
		myMicro.WithTLS(),
		// Wrappers are applied in reverse order so the last is executed first.
		micro.WrapClient(clientWrappers...),
		// Adding some optional lifecycle actions
		micro.BeforeStart(func() (err error) {
			log.Debug().Msg("called BeforeStart")
			return
		}),
		micro.BeforeStop(func() (err error) {
			log.Debug().Msg("called BeforeStop")
			return
		}),
	)

	if cfg.Features.Translogs.Enabled {
		topic := cfg.Features.Translogs.Topic
		publisher := micro.NewEvent(topic, service.Client())
		handlerWrappers = append(handlerWrappers, transWrapper.NewHandlerWrapper(publisher))
		subscriberWrappers = append(subscriberWrappers, transWrapper.NewSubscriberWrapper(publisher))
	}

	service.Init(
		micro.WrapHandler(handlerWrappers...),
		micro.WrapSubscriber(subscriberWrappers...),
	)

	// Initialize DI Container
	ctn, err := registry.NewContainer(cfg)
	defer ctn.Clean()
	if err != nil {
		log.Fatal().Msgf("failed to build container: %v", err)
	}

	// Publisher publish to "mkit.service.emailer"
	publisher := micro.NewEvent(constants.EMAILER_SERVICE, service.Client())

	// // Handlers
	languageHandler := handler.NewLanguageHandler(ctn.Resolve("language-repository").(repository.LanguageRepository), publisher)
	skillHandler := ctn.Resolve("skill-handler").(skillPB.SkillServiceHandler)
	wordHandler := ctn.Resolve("word-handler").(wordPB.WordServiceHandler)

	// Register Handlers
	languagePB.RegisterLanguageServiceHandler(service.Server(), languageHandler)
	skillPB.RegisterSkillServiceHandler(service.Server(), skillHandler)
	wordPB.RegisterWordServiceHandler(service.Server(), wordHandler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal().Err(err).Send()
	}
}
