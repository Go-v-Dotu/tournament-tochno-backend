package mongodb

import (
	"context"
	"fmt"

	"tournament_participation_service/internal/domain"
	"tournament_participation_service/internal/infrastructure/mongodb/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type hostRepository struct {
	client Client
	coll   *mongo.Collection
}

var (
	_ domain.HostRepository = (*hostRepository)(nil)
)

func NewHostRepository(client Client) domain.HostRepository {
	coll := client.Database("tournament_participation").Collection("hosts")
	return &hostRepository{client: client, coll: coll}
}

func (r *hostRepository) GetByUserID(ctx context.Context, userID string) (*domain.Host, error) {
	f := bson.D{{Key: "user_id", Value: userID}}
	res := r.coll.FindOne(ctx, f)

	var hostModel models.Host
	if err := res.Decode(&hostModel); err != nil {
		return nil, fmt.Errorf("host not found: %w", err)
	}

	return hostModel.ToEntity(), nil
}

func (r *hostRepository) Save(ctx context.Context, host *domain.Host) error {
	hostModel := models.NewHost(host)

	if _, err := r.coll.InsertOne(ctx, hostModel); err != nil {
		return fmt.Errorf("error saving host: %w", err)
	}

	return nil
}

func (r *hostRepository) Delete(ctx context.Context, host *domain.Host) error {
	return nil
}

func (r *hostRepository) NextID(_ context.Context) string {
	return primitive.NewObjectID().Hex()
}
