package commands

import (
	"context"
	"fmt"

	"tournament_participation_service/internal/domain"
)

type CreateHostHandler struct {
	hostRepo domain.HostRepository
}

func NewCreateHostHandler(hostRepo domain.HostRepository) *CreateHostHandler {
	return &CreateHostHandler{hostRepo: hostRepo}
}

func (h *CreateHostHandler) Execute(ctx context.Context, id string, userID string, username string) (string, error) {
	host := domain.NewHost(id, userID, username)

	if err := h.hostRepo.Save(ctx, host); err != nil {
		return "", fmt.Errorf("error creating tournament: %w", err)
	}

	return id, nil
}
