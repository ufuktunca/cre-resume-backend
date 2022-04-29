package cv

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CVModel struct {
	MongoClient *mongo.Client
}

type CVModelInterface interface {
}

func CreateCVModel(uri string) *CVModel {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &CVModel{MongoClient: client}
}

func (cvr *CVModel) CreateCV(test string) error {
	return nil
}
