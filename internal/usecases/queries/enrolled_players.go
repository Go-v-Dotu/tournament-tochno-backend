package queries

import (
	"context"
	"fmt"

	"tournament_participation_service/internal/domain"
)

type EnrolledPlayersHandler struct {
	playerRepo         domain.PlayerRepository
	tournamentRepo     domain.TournamentRepository
	playerQueryService PlayerQueryService
}

func NewEnrolledPlayersHandler(
	playerRepo domain.PlayerRepository,
	tournamentRepo domain.TournamentRepository,
	playerQueryService PlayerQueryService,
) *EnrolledPlayersHandler {
	return &EnrolledPlayersHandler{
		playerRepo:         playerRepo,
		tournamentRepo:     tournamentRepo,
		playerQueryService: playerQueryService,
	}
}

func (h *EnrolledPlayersHandler) Execute(ctx context.Context, playerUserID string, tournamentID string) ([]*Player, error) {
	player, err := h.playerRepo.GetByUserID(ctx, playerUserID)
	if err != nil {
		return nil, fmt.Errorf("error getting enrolled players: %w", err)
	}

	tournament, err := h.tournamentRepo.Get(ctx, tournamentID)
	if err != nil {
		return nil, fmt.Errorf("error getting enrolled players: %w", err)
	}

	if !tournament.IsParticipant(player) {
		return nil, fmt.Errorf("error getting enrolled players: you aren't participant of the tournament")
	}

	enrolledPlayers, err := h.playerQueryService.GetByTournamentID(ctx, tournament.ID)
	if err != nil {
		return nil, fmt.Errorf("error getting enrolled players: %w", err)
	}

	return enrolledPlayers, nil
}
