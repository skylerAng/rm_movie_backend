package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/labstack/echo/v4"
	"github.com/parnurzeal/gorequest"
	
	"movie_api/database"
	"movie_api/models"
	"movie_api/utils"
)

var movie_collection *mongo.Collection = database.GetCollection(database.DB, "movie")

func SearchMovie(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	
	movie_name := c.QueryParam("movie_name")
	movie := models.Movie{}
	db_err := movie_collection.FindOne(ctx, bson.M{"title": movie_name}).Decode(&movie)
	if db_err != nil {
		if db_err != mongo.ErrNoDocuments {
			return models.HandleResponseSimple(c, http.StatusInternalServerError, db_err.Error())
		} else {
			return c.String(http.StatusOK, "No Users Detected")
		}
	}
	img_b64, convert_err := utils.ImagetoBase64(movie.ImagePath)
	if convert_err != nil {
		return models.HandleResponseSimple(c, http.StatusInternalServerError, convert_err.Error())
	}

	movie.ImageB64 = img_b64
	return c.JSON(http.StatusOK, movie)
}

func ListMovies(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	
	movies := []models.Movie{}
	opts := options.Find()
	opts.SetSort(bson.D{{"created_date", -1}})
	filterCursor, err := movie_collection.Find(ctx, bson.M{}, opts)

	if err != nil {
		return models.HandleResponseSimple(c, http.StatusInternalServerError, err.Error())
	}

	//reading from the db in an optimal way
	defer filterCursor.Close(ctx)
	if err = filterCursor.All(ctx, &movies); err != nil {
		return models.HandleResponseSimple(c, http.StatusInternalServerError, err.Error())
	}

	final_movies := []models.Movie{}
	for _, movie := range movies {
		img_b64, convert_err := utils.ImagetoBase64(movie.ImagePath)
		if convert_err != nil {
			return models.HandleResponseSimple(c, http.StatusInternalServerError, convert_err.Error())
		}
		movie.ImageB64 = img_b64
		final_movies = append(final_movies, movie)
	}

	return c.JSON(http.StatusOK, final_movies)
}

func CrawlMovies(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	request := gorequest.New()
	url := "https://api.themoviedb.org/3/movie/now_playing?language=en-US&page=1"

	resp, body, errs := request.Get(url).
		Set("Content-Type", "application/json").
		Set("Authorization","Bearer " + os.Getenv("API_KEY")).
		End()

	if errs != nil {
		log.Println("API Error")
		return models.HandleResponseSimple(c, http.StatusInternalServerError, errs[0].Error())
	}

	// Get DB Results
	movie_results := models.MovieDB{}
	if resp.StatusCode == 200 {
		json.Unmarshal([]byte(body), &movie_results)
	}

	// Get image
	utils.CreateFolder("image/")
	insert_doc := []interface{}{}
	insert_counter := 0
	for _, result := range movie_results.Results {
		// Check if file exists only save
		img_path := filepath.Join("image", result.OriginalTitle + ".jpg")
		file_exists, _ := utils.FileExits(img_path)
		original_link := ""

		if !file_exists {
			// Download Image first
			img, url_img, img_errs := utils.GetImage(result.PosterPath)
			original_link = url_img
			if img_errs != nil {
				return models.HandleResponseSimple(c, http.StatusInternalServerError, img_errs.Error())
			}

			utils.SaveImage(img, img_path)
		}

		// If record exists do not append to db
		movie := models.Movie{}
		db_err := movie_collection.FindOne(ctx, bson.M{"title": result.OriginalTitle}).Decode(&movie)
		if db_err == mongo.ErrNoDocuments {
			insert_doc = append(insert_doc, &models.Movie{
				MovieID: result.ID,
				Title: result.OriginalTitle,
				MovieLang: result.OrginalLang,
				Description: result.Overview,
				OriginalLink: original_link,
				ImagePath: img_path,
				CreatedDate: time.Now(),
			})
			insert_counter += 1
		}
	}

	if insert_counter > 0 {
		_, db_err := movie_collection.InsertMany(ctx, insert_doc)
		if db_err != nil {
			return models.HandleResponseSimple(c, http.StatusInternalServerError, db_err.Error())
		}
	}

	return c.JSON(http.StatusOK, insert_doc)
}