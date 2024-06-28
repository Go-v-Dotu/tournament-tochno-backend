package commands

import (
	"context"
	"fmt"

	"tournament_participation_service/internal/domain"
)

type CreatePlayerHandler struct {
	playerRepo domain.PlayerRepository
}

func NewCreatePlayerHandler(playerRepo domain.PlayerRepository) *CreatePlayerHandler {
	return &CreatePlayerHandler{playerRepo: playerRepo}
}

func (h *CreatePlayerHandler) Execute(ctx context.Context, id string, userID string, username string) (string, error) {
	player := domain.NewPlayer(id, userID, username)

	if err := h.playerRepo.Save(ctx, player); err != nil {
		return "", fmt.Errorf("error creating player: %w", err)
	}

	return id, nil
}
