package commands

import (
	"context"
	"fmt"
	"tournament_participation_service/internal/domain"
)

type CreateTournamentHandler struct {
	tournamentRepo domain.TournamentRepository
}

func NewCreateTournamentHandler(tournamentRepo domain.TournamentRepository) *CreateTournamentHandler {
	return &CreateTournamentHandler{tournamentRepo: tournamentRepo}
}

func (h *CreateTournamentHandler) Execute(ctx context.Context, tournament *domain.Tournament) error {
	if err := h.tournamentRepo.Save(ctx, tournament); err != nil {
		return fmt.Errorf("error creating tournament: %w", err)
	}

	return nil
}
