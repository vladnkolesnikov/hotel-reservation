package types

import "github.com/google/uuid"

type User struct {
	// To omit field from the JSON we can use _ or omitempty in the value, e.g. json:"id,omitempty"
	ID        uuid.UUID `bson:"_id" json:"id"`
	FirstName string    `bson:"firstName" json:"firstName"`
	LastName  string    `bson:"lastName" json:"lastName"`
	Email     string    `bson:"email" json:"email,omitempty"`
}
