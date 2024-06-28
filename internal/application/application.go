package application

import (
	"context"
	"fmt"
	"time"

	"tournament_participation_service/internal/infrastructure/mongodb"
	"tournament_participation_service/internal/usecases"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	wla "github.com/ma-hartma/watermill-logrus-adapter"
	"github.com/sirupsen/logrus"
)

type App struct {
	UseCases *usecases.UseCases
}

func NewApp(
	ctx context.Context,
	mongoAddr string,
	rabbitAddr string,
	logger *logrus.Logger,
) (*App, error) {
	dbClient, err := mongodb.NewClient(ctx, mongoAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to create Mongo client: %v", err)
	}

	hostRepo := mongodb.NewHostRepository(*dbClient)
	playerRepo := mongodb.NewPlayerRepository(*dbClient)
	tournamentRepo := mongodb.NewTournamentRepository(*dbClient)
	tournamentQueryService := mongodb.NewTournamentQueryService(*dbClient)
	playerQueryService := mongodb.NewPlayerQueryService(*dbClient)

	eventBus, err := newEventBus(rabbitAddr, logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create event bus: %v", err)
	}

	app := App{
		UseCases: usecases.NewUseCases(
			eventBus,
			hostRepo,
			playerRepo,
			tournamentRepo,
			tournamentQueryService,
			playerQueryService,
		),
	}

	return &app, nil
}

func newEventBus(rabbitAddr string, log *logrus.Logger) (*cqrs.EventBus, error) {
	logger := wla.NewLogrusLogger(log)
	cqrsMarshaler := cqrs.JSONMarshaler{GenerateName: cqrs.StructName}

	pub, err := amqp.NewPublisher(amqp.NewDurablePubSubConfig(rabbitAddr, nil), logger)
	if err != nil {
		return nil, err
	}

	return cqrs.NewEventBusWithConfig(
		pub,
		cqrs.EventBusConfig{
			GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
				return "events", nil
			},

			OnPublish: func(params cqrs.OnEventSendParams) error {
				logger.Info("Publishing event", watermill.LogFields{
					"event_name": params.EventName,
				})

				params.Message.Metadata.Set("published_at", time.Now().String())

				return nil
			},

			Marshaler: cqrsMarshaler,
			Logger:    logger,
		},
	)
}
