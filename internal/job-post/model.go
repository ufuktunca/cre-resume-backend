package jobPost

import (
	"context"
	"cre-resume-backend/internal/models"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type JobPostModel struct {
	MongoClient *mongo.Client
}

type JobPostModelInterface interface {
	CreateJobPost(jobPost *models.JobPost) error
	GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error)
	GetJobPostByID(id string) (*models.JobPost, error)
	CreateApplyJobPost(applyJobPost *models.ApplyJobPost) error
	GetJobPostsWithUserID(id string, postType string) (*[]models.JobPost, error)
	GetJobApplyWithUserIDAndJobID(userId string, jobID string) (*models.ApplyJobPost, error)
	GetUserApplies(userId string) ([]models.ApplyJobPost, error)
	GetUserJobPosts(userId string) ([]models.JobPost, error)
	GetJobApplies(jobId string) (*[]models.ApplyJobPost, error)
}

func NewJobModel(uri string) *JobPostModel {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &JobPostModel{MongoClient: client}
}

func (r *JobPostModel) CreateJobPost(jobPost *models.JobPost) error {
	collection := r.MongoClient.Database("cre-resume").Collection("jobPosts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, jobPost)

	if err != nil {
		return err
	}

	return nil
}

func (r *JobPostModel) CreateApplyJobPost(applyJobPost *models.ApplyJobPost) error {
	collection := r.MongoClient.Database("cre-resume").Collection("apply-jobPost")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, applyJobPost)

	if err != nil {
		return err
	}

	return nil
}

func (r *JobPostModel) GetJobPostByID(id string) (*models.JobPost, error) {
	collection := r.MongoClient.Database("cre-resume").Collection("jobPosts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbResult := collection.FindOne(ctx, bson.M{"id": id})

	if dbResult.Err() != nil {
		return nil, dbResult.Err()
	}

	jobPost := &models.JobPost{}
	err := dbResult.Decode(&jobPost)

	if err != nil {
		return nil, err
	}
	return jobPost, nil
}

func (r *JobPostModel) GetJobPostsWithUserID(id string, postType string) (*[]models.JobPost, error) {
	collection := r.MongoClient.Database("cre-resume").Collection("jobPosts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{"ownerId": id, "type": postType})

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

	if err != nil {
		return nil, err
	}
	return &jobPosts, nil
}

func (r *JobPostModel) GetJobApplyWithUserIDAndJobID(userId string, jobID string) (*models.ApplyJobPost, error) {
	collection := r.MongoClient.Database("cre-resume").Collection("apply-jobPost")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbResult := collection.FindOne(ctx, bson.M{"applierId": userId, "jobPostId": jobID})

	if dbResult.Err() != nil {
		return nil, dbResult.Err()
	}

	applyJobPost := &models.ApplyJobPost{}
	err := dbResult.Decode(&applyJobPost)

	if err != nil {
		return nil, err
	}
	return applyJobPost, nil
}

func (r *JobPostModel) GetUserApplies(userId string) ([]models.ApplyJobPost, error) {
	collection := r.MongoClient.Database("cre-resume").Collection("apply-jobPost")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{"applierId": userId})

	applies := []models.ApplyJobPost{}
	for cur.Next(ctx) {
		var apply models.ApplyJobPost
		err := cur.Decode(&apply)
		if err != nil {
			log.Fatal(err)
		}
		applies = append(applies, apply)
	}

	if err != nil {
		return nil, err
	}
	return applies, nil
}

func (r *JobPostModel) GetUserJobPosts(userId string) ([]models.JobPost, error) {
	collection := r.MongoClient.Database("cre-resume").Collection("jobPosts")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{"ownerId": userId})

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
	return jobPosts, nil
}

func (r *JobPostModel) GetJobApplies(jobId string) (*[]models.ApplyJobPost, error) {
	collection := r.MongoClient.Database("cre-resume").Collection("apply-jobPost")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{"jobPostId": jobId})

	applies := []models.ApplyJobPost{}
	for cur.Next(ctx) {
		var apply models.ApplyJobPost
		err := cur.Decode(&apply)
		if err != nil {
			log.Fatal(err)
		}
		applies = append(applies, apply)
	}

	if err != nil {
		return nil, err
	}
	return &applies, nil
}

func (r *JobPostModel) GetJobPosts(jobPostType, category, from, to, sort string) (*[]models.JobPost, error) {
	collection := r.MongoClient.Database("cre-resume").Collection("jobPosts")
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
