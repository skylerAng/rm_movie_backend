package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MovieDB struct {
	Results []MovieDBResults `json:"results,omitempty""`
}

type MovieDBResults struct {
	ID int64 `json:"id,omitempty"`
	OrginalLang string `json:"original_language,omitempty"`
	OriginalTitle string `json:"original_title,omitempty"`
	Overview string `json:"overview,omitempty"`
	PosterPath string `json:"poster_path,omitempty"`
	Date string `json:"release_date,omitempty"`
}

type Movie struct {
	// Map native go structure
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	MovieID	     int64				`json:"movie_id,omitempty" bson:"movie_id,omitempty"`
	MovieLang	 string				`json:"movie_lang,omitempty" bson:"movie_lang,omitempty"`
	Title        string             `json:"title,omitempty" bson:"title,omitempty"`
	Description  string             `json:"description,omitempty" bson:"description,omitempty"`
	Filename     string             `json:"filename,omitempty" bson:"filename,omitempty"`
	OriginalLink string             `json:"originallink,omitempty" bson:"originallink,omitempty"`
	ImagePath 	 string             `json:"image_path,omitempty" bson:"image_path,omitempty"`
	ImageB64	 string             `json:"base64,omitempty"`
	CreatedDate time.Time `bson:"created_date,omitempty" json:"created_date,omitempty"`
	UpdatedDate time.Time `bson:"updated_date,omitempty" json:"updated_date,omitempty"`
}
