package queries

import (
	"context"
	"fmt"
	"tournament_participation_service/internal/domain"
)

type TournamentByIDHandler struct {
	playerRepo             domain.PlayerRepository
	tournamentRepo         domain.TournamentRepository
	tournamentQueryService TournamentQueryService
}

func NewTournamentByIDHandler(
	playerRepo domain.PlayerRepository,
	tournamentRepo domain.TournamentRepository,
	tournamentQueryService TournamentQueryService,
) *TournamentByIDHandler {
	return &TournamentByIDHandler{
		playerRepo:             playerRepo,
		tournamentRepo:         tournamentRepo,
		tournamentQueryService: tournamentQueryService,
	}
}

func (h *TournamentByIDHandler) Execute(ctx context.Context, playerUserID string, id string) (*Tournament, error) {
	player, err := h.playerRepo.GetByUserID(ctx, playerUserID)
	if playerUserID != "" && err != nil {
		return nil, fmt.Errorf("error getting tournament by id: %w", err)
	}

	tournament, err := h.tournamentRepo.Get(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting tournament by id: %w", err)
	}

	tournamentResp, err := h.tournamentQueryService.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting tournament by id: %w", err)
	}

	tournamentResp.PlayerEnrolled = tournament.IsParticipant(player)

	return tournamentResp, nil
}
