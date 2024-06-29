package mongodb

import (
	"context"
	"fmt"

	"tournament_participation_service/internal/infrastructure/mongodb/models"
	"tournament_participation_service/internal/usecases/queries"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type playerQueryService struct {
	client         Client
	tournamentColl *mongo.Collection
	playerColl     *mongo.Collection
}

var (
	_ queries.PlayerQueryService = (*playerQueryService)(nil)
)

func NewPlayerQueryService(client Client) queries.PlayerQueryService {
	tournamentColl := client.Database("tournament_participation").Collection("tournaments")
	playerColl := client.Database("tournament_participation").Collection("players")
	return &playerQueryService{
		client:         client,
		tournamentColl: tournamentColl,
		playerColl:     playerColl,
	}
}

func (r *playerQueryService) GetByTournamentID(ctx context.Context, tournamentID string) ([]*queries.Player, error) {
	oID, err := primitive.ObjectIDFromHex(tournamentID)
	if err != nil {
		return nil, fmt.Errorf("error getting players by tournament id: %w", err)
	}

	f := bson.D{{"_id", oID}}
	res := r.tournamentColl.FindOne(ctx, f)

	var tournamentModel models.Tournament
	if err := res.Decode(&tournamentModel); err != nil {
		return nil, fmt.Errorf("tournament not found: %w", err)
	}

	players := make([]*queries.Player, 0, len(tournamentModel.Players))
	for _, enrolledPlayer := range tournamentModel.Players {
		f := bson.D{{"_id", enrolledPlayer.PlayerID}}
		res := r.playerColl.FindOne(ctx, f)

		var playerModel models.Player
		if err := res.Decode(&playerModel); err != nil {
			return nil, fmt.Errorf("player not found: %w", err)
		}

		players = append(players, &queries.Player{
			ID:       enrolledPlayer.PlayerID.Hex(),
			UserID:   playerModel.UserID,
			Username: playerModel.Username,
			Dropped:  enrolledPlayer.Dropped,
		})
	}

	return players, nil
}
