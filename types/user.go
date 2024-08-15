package types

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (params UpdateUserParams) ToBson() bson.M {
	values := bson.M{}

	if len(params.FirstName) > 0 {
		values["firstName"] = params.FirstName
	}

	if len(params.LastName) > 0 {
		values["lastName"] = params.LastName
	}

	return values
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	// To omit field from the JSON we can use _ or omitempty in the value, e.g. json:"id,omitempty"
	ID                primitive.ObjectID `bson:"_id" json:"id"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lastName" json:"lastName"`
	Email             string             `bson:"email" json:"email,omitempty"`
	EncryptedPassword string             `bson:"EncryptedPassword" json:"-"`
}

func CreateUserFromParams(params CreateUserParams) (*User, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:                primitive.NewObjectID(),
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encryptedPassword),
	}, nil
}
