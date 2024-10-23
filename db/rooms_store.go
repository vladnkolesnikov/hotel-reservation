package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel-reservation/types"
)

type RoomsStore interface {
	InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error)
}

type MongoRoomsStore struct {
	dbName     string
	client     *mongo.Client
	collection *mongo.Collection

	HotelStore
}

func (store *MongoRoomsStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	resp, err := store.collection.InsertOne(ctx, room)

	if err != nil {
		return nil, err
	}

	room.ID = resp.InsertedID.(primitive.ObjectID)

	filter := bson.M{"_id": room.HotelID}
	update := bson.M{"$push": bson.M{"rooms": room}}

	if err = store.HotelStore.Update(ctx, &filter, &update); err != nil {
		return nil, err
	}

	return room, nil
}

const roomsCollectionName = "rooms"

func NewMongoRoomsStore(client *mongo.Client, dbName string, hotelStore HotelStore) *MongoRoomsStore {
	return &MongoRoomsStore{
		dbName:     dbName,
		client:     client,
		collection: client.Database(dbName).Collection(roomsCollectionName),
		HotelStore: hotelStore,
	}
}
