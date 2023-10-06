package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"name" json:"name"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password"`
	CreatedDate time.Time          `bson:"created_date,omitempty" json:"created_date,omitempty"`
	UpdatedDate time.Time          `bson:"updated_date,omitempty" json:"updated_date,omitempty"`
}
