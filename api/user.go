package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	users := [3]types.User{}

	for i := 0; i < 3; i++ {
		users[i] = types.User{
			ID:        uuid.New(),
			FirstName: "Ivan",
			LastName:  "Ivanov",
		}
	}

	return c.JSON(users)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(types.User{
		ID:        uuid.New(),
		FirstName: "Joe",
		LastName:  "Don",
	})
}
