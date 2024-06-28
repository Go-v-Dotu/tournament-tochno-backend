package queries

import (
	"context"
	"fmt"

	"tournament_participation_service/internal/domain"
)

type ParticipatedTournamentsHandler struct {
	playerRepo             domain.PlayerRepository
	tournamentQueryService TournamentQueryService
}

func NewParticipatedTournamentsHandler(
	playerRepo domain.PlayerRepository,
	tournamentQueryService TournamentQueryService,
) *ParticipatedTournamentsHandler {
	return &ParticipatedTournamentsHandler{
		playerRepo:             playerRepo,
		tournamentQueryService: tournamentQueryService,
	}
}

func (h *ParticipatedTournamentsHandler) Execute(ctx context.Context, playerUserID string) ([]*Tournament, error) {
	player, err := h.playerRepo.GetByUserID(ctx, playerUserID)
	if err != nil {
		return nil, fmt.Errorf("error getting participated tournaments: %w", err)
	}

	tournaments, err := h.tournamentQueryService.GetByPlayerID(ctx, player.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting participated tournaments: %w", err)
	}

	return tournaments, nil
}
