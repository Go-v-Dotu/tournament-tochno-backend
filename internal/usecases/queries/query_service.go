package queries

import (
	"context"
)

type TournamentQueryService interface {
	GetByID(ctx context.Context, id string) (*Tournament, error)
	GetAll(ctx context.Context) ([]*Tournament, error)
	GetByPlayerID(ctx context.Context, playerID string) ([]*Tournament, error)
}

type PlayerQueryService interface {
	GetByTournamentID(ctx context.Context, tournamentID string) ([]*Player, error)
}
