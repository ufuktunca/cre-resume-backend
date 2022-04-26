package cv

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CVRepository struct {
	MongoClient *mongo.Client
}

type CVRepositoryInterface interface {
}

func CreateCVRepository(uri string) *CVRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &CVRepository{MongoClient: client}
}

func (cvr *CVRepository) CreateCV(test string) error {
	return nil
}
