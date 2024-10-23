package main

import (
	"fmt"
	"hotel-reservation/types"
)

func main() {
	hotel := types.Hotel{Name: "Lloret de Mar Resort", Location: "Barcelona, Spain"}

	room := types.Room{
		Type:      types.SingleRoomType,
		BasePrice: 99.9,
	}

	fmt.Println("seeding the database", room, hotel)
}
