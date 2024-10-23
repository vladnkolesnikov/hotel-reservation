package hotels

import (
	"github.com/gofiber/fiber/v2"
	"hotel-reservation/db"
	"hotel-reservation/types"
)

type RoomsHandler struct {
	roomsStore db.RoomsStore
}

func NewRoomsHandler(store db.RoomsStore) *RoomsHandler {
	return &RoomsHandler{
		roomsStore: store,
	}
}

func (handler *RoomsHandler) HandlePostRoom(c *fiber.Ctx) error {
	var params types.Room

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	room := types.CreateRoomFromParams(params)

	createdHotel, err := handler.roomsStore.InsertRoom(c.Context(), room)

	if err != nil {
		return err
	}

	return c.JSON(createdHotel)
}
