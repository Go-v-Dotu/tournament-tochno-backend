package domain

import "context"

type HostRepository interface {
	GetByUserID(ctx context.Context, userID string) (*Host, error)
	Save(ctx context.Context, host *Host) error
	Delete(ctx context.Context, host *Host) error
	NextID(ctx context.Context) string
}
