package models

import (
	"tournament_participation_service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Host struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserID   string             `bson:"user_id"`
	Username string             `bson:"username"`
}

func NewHost(host *domain.Host) *Host {
	id, err := primitive.ObjectIDFromHex(host.ID)
	if err != nil {
		panic(err)
	}
	return &Host{
		ID:       id,
		UserID:   host.UserID,
		Username: host.Username,
	}
}

func (m *Host) ToEntity() *domain.Host {
	return domain.NewHost(m.ID.Hex(), m.UserID, m.Username)
}
