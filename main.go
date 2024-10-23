package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel-reservation/api/hotels"
	"hotel-reservation/api/users"
	"hotel-reservation/db"
	"log"
	"os"
)

const defaultPort = 3000

func main() {
	port := flag.String("port", fmt.Sprint(defaultPort), "Port of the server")
	flag.Parse()

	log.SetFlags(log.LstdFlags)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(db.EnvDbURI)))
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return ctx.Status(code).JSON(struct {
				Message string `json:"message"`
				Status  int    `json:"status"`
			}{
				Message: err.Error(),
				Status:  code,
			})
		},
	})

	app.Use(helmet.New())

	var (
		apiV1        = app.Group("/api/v1")
		dbName       = os.Getenv(db.EnvDbName)
		userHandler  = users.NewUserHandler(db.NewMongoDBUserStore(client, dbName))
		hotelHandler = hotels.NewHotelHandler(db.NewMongoDBHotelStore(client, dbName))
	)

	usersAPI := apiV1.Group("/users")

	usersAPI.Get("/", userHandler.HandleGetUsers)
	usersAPI.Post("/", userHandler.HandlePostUser)
	usersAPI.Get("/:id", userHandler.HandleGetUser)
	usersAPI.Delete("/:id", userHandler.HandleDeleteUser)
	usersAPI.Put("/:id", userHandler.HandlePutUser)

	hotelsAPI := apiV1.Group("/hotels")
	hotelsAPI.Post("/", hotelHandler.HandlePostHotel)
	hotelsAPI.Get("/", hotelHandler.HandleGetHotels)
	hotelsAPI.Get("/:id", hotelHandler.HandleGetHotel)

	if appError := app.Listen(fmt.Sprintf(":%s", *port)); appError != nil {
		log.Fatal(appError)
	}
}
