package queries

import (
	"context"
	"fmt"
)

type AllTournamentHandler struct {
	tournamentQueryService TournamentQueryService
}

func NewAllTournamentHandler(tournamentQueryService TournamentQueryService) *AllTournamentHandler {
	return &AllTournamentHandler{
		tournamentQueryService: tournamentQueryService,
	}
}

func (h *AllTournamentHandler) Execute(ctx context.Context) ([]*Tournament, error) {
	tournamentResp, err := h.tournamentQueryService.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting all tournaments: %w", err)
	}

	return tournamentResp, nil
}
