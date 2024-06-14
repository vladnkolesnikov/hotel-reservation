package main

import (
	"flag"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"hotel-reservation/api"
)

const defaultPort = 3000

func main() {
	port := flag.String("port", fmt.Sprint(defaultPort), "Port of the server")
	flag.Parse()

	app := fiber.New()
	app.Use(helmet.New())

	apiV1 := app.Group("/api/v1")
	apiV1.Get("/users", api.HandleGetUsers)
	apiV1.Get("/users/:id", api.HandleGetUser)

	err := app.Listen(fmt.Sprintf(":%s", *port))

	if err != nil {
		panic(err)
	}
}
