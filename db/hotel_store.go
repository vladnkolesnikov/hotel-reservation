package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel-reservation/types"
)

type HotelStore interface {
	InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error)
	Update(ctx context.Context, filter *bson.M, update *bson.M) error
	GetAll(ctx context.Context, filter *bson.M) ([]*types.Hotel, error)
	GetHotelByID(ctx context.Context, id string) (*types.Hotel, error)
}

type MongoHotelStore struct {
	dbName     string
	client     *mongo.Client
	collection *mongo.Collection
}

func (store *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	resp, err := store.collection.InsertOne(ctx, hotel)

	if err != nil {
		return nil, err
	}

	hotel.ID = resp.InsertedID.(primitive.ObjectID)

	return hotel, nil
}

func (store *MongoHotelStore) Update(ctx context.Context, filter *bson.M, update *bson.M) error {
	_, err := store.collection.UpdateOne(ctx, filter, update)

	return err
}

func (store *MongoHotelStore) GetAll(ctx context.Context, filter *bson.M) ([]*types.Hotel, error) {
	res, err := store.collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	var hotels []*types.Hotel

	if err := res.All(ctx, &hotels); err != nil {
		return nil, err
	}

	return hotels, nil
}

func (store *MongoHotelStore) GetHotelByID(ctx context.Context, id string) (*types.Hotel, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var hotel types.Hotel

	if err := store.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&hotel); err != nil {
		return nil, err
	}

	return &hotel, nil
}

const hotelCollectionName = "hotels"

func NewMongoDBHotelStore(client *mongo.Client, dbName string) *MongoHotelStore {
	return &MongoHotelStore{
		dbName:     dbName,
		client:     client,
		collection: client.Database(dbName).Collection(hotelCollectionName),
	}
}
