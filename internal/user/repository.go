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

type UserRepositoryInterface interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
}

func NewUserRepository(uri string) *Respository {
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
