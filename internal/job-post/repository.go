package jobPost

import (
	"context"
	"cre-resume-backend/internal/models"
	"fmt"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobPostRepository struct {
	MongoClient *mongo.Client
}

type JobPostRepositoryInterface interface {
	CreateJobPost(jobPost *models.JobPost) error
	GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error)
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

func (r *JobPostRepository) GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error) {
	collection := r.MongoClient.Database("cre-resume").Collection("jobPosts")
	fmt.Println(jobPostType, category, from, to)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	typeFilter := bson.M{"type": jobPostType}
	categoryFilter := bson.M{}
	saleFilter := bson.M{}

	if category != "" {
		categoryFilter = bson.M{"category": category}
	}

	if from != "" && to != "" {
		fromInt, _ := strconv.Atoi(from)
		toInt, _ := strconv.Atoi(to)
		saleFilter = bson.M{"salary": bson.M{"$gte": fromInt, "$lte": toInt}}
	}

	options := options.Find()

	switch sort {
	case "location":
		options.SetSort(bson.M{"location": 1})
	case "salary":
		options.SetSort(bson.M{"salary": -1})
	case "company":
		options.SetSort(bson.M{"company": 1})
	default:
		options.SetSort(bson.M{"createdAt": -1})
	}

	cur, err := collection.Find(ctx, bson.M{
		"$and": []bson.M{
			typeFilter,
			categoryFilter,
			saleFilter,
		},
	}, options)

	jobPosts := []models.JobPost{}
	for cur.Next(ctx) {
		var jobPost models.JobPost
		err := cur.Decode(&jobPost)
		if err != nil {
			log.Fatal(err)
		}
		jobPosts = append(jobPosts, jobPost)
	}

	if err != nil {
		return nil, err
	}

	return &jobPosts, nil
}
