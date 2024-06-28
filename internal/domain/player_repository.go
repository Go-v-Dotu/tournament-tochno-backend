package domain

import "context"

type PlayerRepository interface {
	GetByUserID(ctx context.Context, id string) (*Player, error)
	Save(ctx context.Context, player *Player) error
	Delete(ctx context.Context, player *Player) error
	NextID(ctx context.Context) string
}
