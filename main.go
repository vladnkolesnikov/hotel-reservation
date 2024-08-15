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
	"hotel-reservation/api"
	"hotel-reservation/db"
	"log"
	"os"
)

const defaultPort = 3000
const dbURI = "mongodb://localhost:27017"

func main() {
	port := flag.String("port", fmt.Sprint(defaultPort), "Port of the server")
	flag.Parse()

	log.SetFlags(log.LstdFlags)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbURI))
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
	apiV1 := app.Group("/api/v1")
	dbName := os.Getenv(db.EnvDBName)

	userHandler := api.NewUserHandler(db.NewMongoDBUserStore(client, dbName))

	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Post("/users", userHandler.HandlePostUser)
	apiV1.Get("/users/:id", userHandler.HandleGetUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)
	apiV1.Put("/users/:id", userHandler.HandlePutUser)

	if appError := app.Listen(fmt.Sprintf(":%s", *port)); appError != nil {
		log.Fatal(appError)
	}
}
