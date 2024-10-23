package users

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel-reservation/db"
	"hotel-reservation/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testDB struct {
	db.UserStore
}

func (testDb *testDB) teardown(t *testing.T) {
	if err := testDb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal()
	}
}

const (
	testDBURI = "mongodb://localhost:27017"
	dbName    = "hotel-reservation-testdb"
)

func setup(t *testing.T) *testDB {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testDBURI))

	if err != nil {
		t.Fatal(err)
	}

	return &testDB{
		UserStore: db.NewMongoDBUserStore(client, dbName),
	}
}

func TestUserHandler(t *testing.T) {
	testDB := setup(t)
	defer testDB.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(testDB)
	app.Post("/", userHandler.HandlePostUser)
	app.Get("/:id", userHandler.HandleGetUser)
	app.Get("/", userHandler.HandleGetUsers)
	app.Put("/:id", userHandler.HandlePutUser)

	postUser := func() (types.User, types.CreateUserParams) {
		params := types.CreateUserParams{
			FirstName: "Test",
			LastName:  "Test",
			Email:     "test@live.com",
			Password:  "123123",
		}

		b, _ := json.Marshal(params)

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)

		var user types.User
		json.NewDecoder(resp.Body).Decode(&user)

		return user, params
	}

	t.Run("HandlePostUser", func(t *testing.T) {
		user, params := postUser()

		utils.AssertEqual(t, user.FirstName, params.FirstName)
		utils.AssertEqual(t, user.LastName, params.LastName)
		utils.AssertEqual(t, user.Email, params.Email)
	})

	t.Run("HandleGetUser", func(t *testing.T) {
		postUser, _ := postUser()

		req := httptest.NewRequest(http.MethodGet, fmt.Sprint("/", postUser.ID.Hex()), nil)
		resp, _ := app.Test(req)

		var user types.User
		json.NewDecoder(resp.Body).Decode(&user)

		utils.AssertEqual(t, user.FirstName, postUser.FirstName)
		utils.AssertEqual(t, user.LastName, postUser.LastName)
		utils.AssertEqual(t, user.Email, postUser.Email)
	})

	t.Run("HandleGetUsers", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		resp, _ := app.Test(req)

		var users []types.User
		json.NewDecoder(resp.Body).Decode(&users)

		utils.AssertEqual(t, len(users), 2)
	})

	t.Run("HandlePutUser", func(t *testing.T) {
		postUser, _ := postUser()

		b, _ := json.Marshal(types.UpdateUserParams{
			FirstName: "New Test",
		})

		req := httptest.NewRequest(http.MethodPut, fmt.Sprint("/", postUser.ID.Hex()), bytes.NewReader(b))
		req.Header.Set("Content-Type", "application/json")
		resp, _ := app.Test(req)

		var statusResponse map[string]string

		json.NewDecoder(resp.Body).Decode(&statusResponse)

		utils.AssertEqual(t, statusResponse["status"], "success")

	})
}
