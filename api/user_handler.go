package api

import (
	"context"
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fiber.NewError(fiber.StatusNotFound, "User not found")
		}

		return err
	}

	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.userStore.GetUsers(c.Context())

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fiber.NewError(fiber.StatusNotFound, "Users not found")
		}

		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, err := types.CreateUserFromParams(params)

	if err != nil {
		return err
	}

	createdUser, err := h.userStore.InsertUser(c.Context(), user)

	if err != nil {
		return err
	}

	return c.JSON(createdUser)
}

func (h *UserHandler) HandlePutUser(ctx *fiber.Ctx) error {
	var (
		params types.UpdateUserParams
		userID = ctx.Params("id")
	)

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	if err := ctx.BodyParser(&params); err != nil {
		return err
	}

	filter := bson.M{"_id": id}

	if err := h.userStore.UpdateUser(ctx.Context(), filter, params); err != nil {
		return err
	}

	return ctx.JSON(fiber.Map{
		"status": "success",
		"userID": id,
	})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")

	if err := h.userStore.DeleteUser(c.Context(), userId); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"userId": userId,
	})
}
