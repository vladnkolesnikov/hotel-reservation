package hotels

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel-reservation/db"
	"hotel-reservation/types"
)

type HotelHandler struct {
	hotelStore db.HotelStore
}

func NewHotelHandler(store db.HotelStore) *HotelHandler {
	return &HotelHandler{
		hotelStore: store,
	}
}

func (handler *HotelHandler) HandlePostHotel(ctx *fiber.Ctx) error {
	var params types.Hotel

	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	hotel := types.CreateHotelFromParams(params)

	createdHotel, err := handler.hotelStore.InsertHotel(ctx.Context(), hotel)

	if err != nil {
		return err
	}

	return ctx.JSON(createdHotel)
}

type GetHotelsQuery struct {
	Rooms  bool
	Rating int8
}

func (handler *HotelHandler) HandleGetHotels(ctx *fiber.Ctx) error {
	var query GetHotelsQuery

	if err := ctx.QueryParser(&query); err != nil {
		return err
	}

	hotels, err := handler.hotelStore.GetAll(ctx.Context(), &bson.M{
		"rating": bson.M{
			"$gte": query.Rating,
		},
	})

	if err != nil {
		return err
	}

	return ctx.JSON(hotels)
}

func (handler *HotelHandler) HandleGetHotel(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	hotel, err := handler.hotelStore.GetHotelByID(ctx.Context(), id)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fiber.NewError(fiber.StatusNotFound, "Hotel not found")
		}

		return err
	}

	return ctx.JSON(hotel)
}
