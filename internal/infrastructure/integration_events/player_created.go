package integration_events

import (
	"context"

	"tournament_participation_service/internal/usecases"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type PlayerCreatedEvent struct {
	ID       string
	UserID   string
	Username string
}

type PlayerCreatedHandler struct {
	uc *usecases.UseCases
}

var (
	_ cqrs.GroupEventHandler = (*PlayerCreatedHandler)(nil)
)

func NewPlayerCreatedHandler(uc *usecases.UseCases) *PlayerCreatedHandler {
	return &PlayerCreatedHandler{uc: uc}
}

func (h *PlayerCreatedHandler) NewEvent() any {
	return &PlayerCreatedEvent{}
}

func (h *PlayerCreatedHandler) Handle(ctx context.Context, e any) error {
	event := e.(*PlayerCreatedEvent)

	h.uc.Commands.CreatePlayerHandler.Execute(ctx, event.ID, event.UserID, event.Username)
	return nil
}
