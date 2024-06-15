package api

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"hotel-reservation/db"
	"hotel-reservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)

	user, err := h.userStore.GetUserByID(ctx, id)

	if err != nil {
		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users := [3]types.User{}

	for i := 0; i < 3; i++ {
		users[i] = types.User{
			FirstName: "Ivan",
			LastName:  "Ivanov",
		}
	}

	return c.JSON(users)
}
