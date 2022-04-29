package user

import (
	"context"
	"cre-resume-backend/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Respository struct {
	MongoClient *mongo.Client
}

type UserModelInterface interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	Activation(userID string) error
}

func NewUserModel(uri string) *Respository {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Respository{MongoClient: client}
}

func (r *Respository) GetUserByEmail(email string) (*models.User, error) {
	collection := r.MongoClient.Database("cre-resume").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbResult := collection.FindOne(ctx, bson.M{"email": email})

	if dbResult.Err() != nil {
		return nil, dbResult.Err()
	}

	user := &models.User{}
	err := dbResult.Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Respository) CreateUser(user *models.User) error {
	collection := r.MongoClient.Database("cre-resume").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (r *Respository) Activation(userID string) error {
	collection := r.MongoClient.Database("MedicalCaseDB").Collection("Users")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"userID": userID, "userActivate": false}

	_, err := collection.UpdateOne(ctx,
		filter,
		bson.M{
			"$set": bson.M{
				"userActivate": true,
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}
