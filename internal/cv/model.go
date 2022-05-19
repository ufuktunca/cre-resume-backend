package cv

import (
	"context"
	"cre-resume-backend/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CVModel struct {
	MongoClient *mongo.Client
}

type CVModelInterface interface {
	CreateCV(cvData models.CV) error
	GetCVs(userID string) (*[]models.CV, error)
	GetCV(cvId string) (string, error)
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

func (cvr *CVModel) CreateCV(cvData models.CV) error {
	collection := cvr.MongoClient.Database("cre-resume").Collection("CV")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, cvData)

	if err != nil {
		return err
	}

	return nil
}

func (cvr *CVModel) GetCVs(userID string) (*[]models.CV, error) {
	collection := cvr.MongoClient.Database("cre-resume").Collection("CV")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{"ownerId": userID})

	CVs := []models.CV{}
	for cur.Next(ctx) {
		var CV models.CV
		err := cur.Decode(&CV)
		if err != nil {
			log.Fatal(err)
		}
		CVs = append(CVs, CV)
	}

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return &CVs, nil
}

func (cvr *CVModel) GetCV(cvId string) (string, error) {
	collection := cvr.MongoClient.Database("cre-resume").Collection("CV")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbResult := collection.FindOne(ctx, bson.M{"id": cvId})

	CVs := models.CV{}
	err := dbResult.Decode(&CVs)

	if err != nil {
		return "", err
	}

	return CVs.PDFCV, nil
}
