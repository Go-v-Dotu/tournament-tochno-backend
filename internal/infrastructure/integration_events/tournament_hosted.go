package integration_events

import (
	"context"

	"tournament_participation_service/internal/domain"
	"tournament_participation_service/internal/usecases"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type TournamentHostedEvent struct {
	*domain.Tournament
}

type TournamentHostedHandler struct {
	uc *usecases.UseCases
}

var (
	_ cqrs.GroupEventHandler = (*TournamentHostedHandler)(nil)
)

func NewTournamentHostedHandler(uc *usecases.UseCases) *TournamentHostedHandler {
	return &TournamentHostedHandler{uc: uc}
}

func (h *TournamentHostedHandler) NewEvent() any {
	return &TournamentHostedEvent{}
}

func (h *TournamentHostedHandler) Handle(ctx context.Context, e any) error {
	event := e.(*TournamentHostedEvent)

	h.uc.Commands.CreateTournamentHandler.Execute(ctx, event.Tournament)
	return nil
}
