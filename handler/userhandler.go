package handler

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"movie_api/database"
	"movie_api/models"
)

var user_collection *mongo.Collection = database.GetCollection(database.DB, "user")

func RegisterUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Read form fields
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	search_param := bson.M{"email": email}
	sel_user := models.User{}
	db_err := user_collection.FindOne(ctx, search_param).Decode(&sel_user)
	if db_err != nil {
		if db_err != mongo.ErrNoDocuments {
			return models.HandleResponseSimple(c, http.StatusInternalServerError, db_err.Error())
		}
	}

	// Check if sel_user is not empty
	if sel_user != (models.User{}) {
		return models.HandleResponseSimple(
			c, http.StatusBadRequest,
			"Duplicate User Detected Please Try Again",
		)
	}

	result, err := user_collection.InsertOne(ctx, &models.User{
		Name: name,
		Email: email,
		Password: password,
	})
	if err != nil {
		return models.HandleResponseSimple(c, http.StatusInternalServerError, err.Error())
	}

	log.Println(result)
	return c.String(http.StatusOK, "User Created")
}

func LoginUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	email := c.FormValue("email")
	password := c.FormValue("password")

	// Search for 1 user
	search_params := bson.M{"email": email, "password": password}
	user := models.User{}
	db_err := user_collection.FindOne(ctx, search_params).Decode(&user)
	// Check for other db errors
	if db_err != nil {
		if db_err != mongo.ErrNoDocuments {
			return models.HandleResponseSimple(c, http.StatusBadRequest, db_err.Error())
		} else {
			return models.HandleResponseSimple(c, http.StatusBadRequest, "No User Detected")
		}
	}

	return c.JSON(http.StatusOK, user)
}