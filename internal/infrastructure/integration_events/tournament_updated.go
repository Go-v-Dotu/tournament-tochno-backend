package integration_events

import (
	"context"

	"tournament_participation_service/internal/domain"
	"tournament_participation_service/internal/usecases"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type TournamentUpdatedEvent struct {
	*domain.Tournament
}

type TournamentUpdatedHandler struct {
	uc *usecases.UseCases
}

var (
	_ cqrs.GroupEventHandler = (*TournamentUpdatedHandler)(nil)
)

func NewTournamentUpdatedHandler(uc *usecases.UseCases) *TournamentUpdatedHandler {
	return &TournamentUpdatedHandler{uc: uc}
}

func (h *TournamentUpdatedHandler) NewEvent() any {
	return &TournamentUpdatedEvent{}
}

func (h *TournamentUpdatedHandler) Handle(ctx context.Context, e any) error {
	event := e.(*TournamentUpdatedEvent)

	h.uc.Commands.CreateTournamentHandler.Execute(ctx, event.Tournament)
	return nil
}
