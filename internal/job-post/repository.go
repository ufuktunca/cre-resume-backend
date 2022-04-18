package jobPost

import (
	"context"
	"cre-resume-backend/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobPostRepository struct {
	MongoClient *mongo.Client
}

type JobPostRepositoryInterface interface {
	CreateJobPost(jobPost *models.JobPost) error
}

func CreateJobPostRepository(uri string) *JobPostRepository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &JobPostRepository{MongoClient: client}
}

func (r *JobPostRepository) CreateJobPost(jobPost *models.JobPost) error {
	collection := r.MongoClient.Database("cre-resume").Collection("jobPosts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, jobPost)

	if err != nil {
		return err
	}

	return nil
}
