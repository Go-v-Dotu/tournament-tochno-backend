package main

import (
	"context"
	"os"

	"tournament_participation_service/internal/application"
	"tournament_participation_service/internal/infrastructure/api"
	"tournament_participation_service/internal/infrastructure/events_handler"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.TODO()

	logger := logrus.New()
	rabbitAddr := os.Getenv("RABBIT_ADDR")
	if rabbitAddr == "" {
		logger.Fatal("Error during initialization: RABBIT_ADDR env not set")
	}

	mongoAddr := os.Getenv("MONGO_ADDR")
	if mongoAddr == "" {
		logger.Fatal("Error during initialization: MONGO_ADDR env not set")
	}

	app, err := application.NewApp(
		ctx,
		mongoAddr,
		rabbitAddr,
		logger,
	)
	if err != nil {
		logger.Fatalf("Error during initialization: %v", err)
		os.Exit(1)
	}

	eventsHandler, err := events_handler.NewEventsHandler(app, rabbitAddr, logger)
	if err != nil {
		logger.Fatalf("Error during initialization: %v", err)
		return
	}

	srv := api.NewServer(app, logger)

	go eventsHandler.Start(ctx)

	srv.Start()
}
