package integration_events

import (
	"context"

	"tournament_participation_service/internal/usecases"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type HostCreatedEvent struct {
	ID       string
	UserID   string
	Username string
}

type HostCreatedHandler struct {
	uc *usecases.UseCases
}

var (
	_ cqrs.GroupEventHandler = (*HostCreatedHandler)(nil)
)

func NewHostCreatedHandler(uc *usecases.UseCases) *HostCreatedHandler {
	return &HostCreatedHandler{uc: uc}
}

func (h *HostCreatedHandler) NewEvent() any {
	return &HostCreatedEvent{}
}

func (h *HostCreatedHandler) Handle(ctx context.Context, e any) error {
	event := e.(*HostCreatedEvent)

	h.uc.Commands.CreateHostHandler.Execute(ctx, event.ID, event.UserID, event.Username)
	return nil
}
