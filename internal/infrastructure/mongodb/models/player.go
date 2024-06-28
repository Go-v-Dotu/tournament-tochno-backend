package models

import (
	"tournament_participation_service/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Player struct {
	ID       primitive.ObjectID `bson:"_id"`
	UserID   string             `bson:"user_id"`
	Username string             `bson:"username"`
}

func NewPlayer(player *domain.Player) *Player {
	id, err := primitive.ObjectIDFromHex(player.ID)
	if err != nil {
		panic(err)
	}
	return &Player{
		ID:       id,
		UserID:   player.UserID,
		Username: player.Username,
	}
}

func (m *Player) ToEntity() *domain.Player {
	return domain.NewPlayer(m.ID.Hex(), m.UserID, m.Username)
}
