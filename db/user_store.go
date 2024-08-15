package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel-reservation/types"
)

const usersCollectionName = "users"

type UserStore interface {
	Dropper

	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error
}

type MongoUserStore struct {
	dbName     string
	client     *mongo.Client
	collection *mongo.Collection
}

func (store *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("dropping User collection")

	return store.collection.Drop(ctx)
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

func (store *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := store.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	var users []*types.User

	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (store *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	result, err := store.collection.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (store *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	if _, err := store.collection.DeleteOne(ctx, bson.M{"_id": oid}); err != nil {
		return err

	}

	return nil
}

func (store *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, params types.UpdateUserParams) error {
	values := params.ToBson()

	update := bson.D{{
		"$set", values,
	}}

	_, err := store.collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}

func NewMongoDBUserStore(client *mongo.Client, dbName string) *MongoUserStore {
	return &MongoUserStore{
		dbName:     dbName,
		client:     client,
		collection: client.Database(dbName).Collection(usersCollectionName),
	}
}
