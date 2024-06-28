package events_handler

import (
	"context"
	"time"

	"tournament_participation_service/internal/application"
	"tournament_participation_service/internal/infrastructure/integration_events"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
	wla "github.com/ma-hartma/watermill-logrus-adapter"
	"github.com/sirupsen/logrus"
)

type EventsHandler struct {
	router    *message.Router
	processor *cqrs.EventGroupProcessor
}

func NewEventsHandler(app *application.App, rabbitAddr string, log *logrus.Logger) (*EventsHandler, error) {
	logger := wla.NewLogrusLogger(log)
	cqrsMarshaler := cqrs.JSONMarshaler{GenerateName: cqrs.StructName}

	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		return nil, err
	}

	processor, err := cqrs.NewEventGroupProcessorWithConfig(
		router,
		cqrs.EventGroupProcessorConfig{
			GenerateSubscribeTopic: func(params cqrs.EventGroupProcessorGenerateSubscribeTopicParams) (string, error) {
				return "events", nil
			},
			SubscriberConstructor: func(params cqrs.EventGroupProcessorSubscriberConstructorParams) (message.Subscriber, error) {
				config := amqp.NewDurablePubSubConfig(
					rabbitAddr,
					amqp.GenerateQueueNameTopicNameWithSuffix(params.EventGroupName),
				)

				return amqp.NewSubscriber(config, logger)
			},

			AckOnUnknownEvent: true,
			OnHandle: func(params cqrs.EventGroupProcessorOnHandleParams) error {
				start := time.Now()

				err := params.Handler.Handle(params.Message.Context(), params.Event)

				logger.Info("Event handled", watermill.LogFields{
					"event_name": params.EventName,
					"duration":   time.Since(start),
					"err":        err,
				})

				return err
			},

			Marshaler: cqrsMarshaler,
			Logger:    logger,
		},
	)
	if err != nil {
		return nil, err
	}

	if err := processor.AddHandlersGroup(
		"tournament_participation",
		integration_events.NewHostCreatedHandler(app.UseCases),
		integration_events.NewPlayerCreatedHandler(app.UseCases),
		integration_events.NewTournamentHostedHandler(app.UseCases),
		integration_events.NewTournamentUpdatedHandler(app.UseCases),
	); err != nil {
		return nil, err
	}

	return &EventsHandler{
		router:    router,
		processor: processor,
	}, nil
}

func (h *EventsHandler) Start(ctx context.Context) {
	if err := h.router.Run(ctx); err != nil {
		panic(err)
	}
}
