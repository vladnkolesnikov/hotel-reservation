package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel-reservation/api"
	"hotel-reservation/db"
	"log"
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

	app := fiber.New()
	app.Use(helmet.New())
	apiV1 := app.Group("/api/v1")

	userHandler := api.NewUserHandler(db.NewMongoDBUserStore(client))

	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Get("/users/:id", userHandler.HandleGetUser)

	if appError := app.Listen(fmt.Sprintf(":%s", *port)); appError != nil {
		log.Fatal(appError)
	}
}
