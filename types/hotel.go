package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name     string               `json:"name" bson:"name"`
	Location string               `json:"location" bson:"location"`
	Rooms    []primitive.ObjectID `json:"rooms" bson:"rooms"`
	Rating   int8                 `json:"rating" bson:"rating"`
}

type RoomType int8

const (
	_ RoomType = iota
	SingleRoomType
	DoubleRoomType
	SeaSideRoomType
	LuxuryRoomType
)

type Room struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Type      RoomType           `json:"type" bson:"type"`
	BasePrice float64            `json:"basePrice" bson:"basePrice"`
	Price     float64            `json:"price" bson:"price"`
	HotelID   primitive.ObjectID `json:"hotelId" bson:"hotelId"`
}

func CreateHotelFromParams(params Hotel) *Hotel {
	return &Hotel{
		ID:       primitive.NewObjectID(),
		Name:     params.Name,
		Location: params.Location,
		Rooms:    params.Rooms,
		Rating:   params.Rating,
	}
}

func CreateRoomFromParams(params Room) *Room {
	return &Room{
		ID:        primitive.NewObjectID(),
		Type:      params.Type,
		BasePrice: params.BasePrice,
		Price:     params.Price,
		HotelID:   primitive.ObjectID{},
	}
}
