package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel-reservation/types"
	"os"
)

const usersCollectionName = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct {
	dbName     string
	client     *mongo.Client
	collection *mongo.Collection
}

func (store *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var user types.User

	if err := store.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func NewMongoDBUserStore(client *mongo.Client) *MongoUserStore {
	dbName := os.Getenv(envDBName)

	return &MongoUserStore{
		dbName:     dbName,
		client:     client,
		collection: client.Database(dbName).Collection(usersCollectionName),
	}
}
