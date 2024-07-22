package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Applicants struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Age        int                `bson:"age" json:"age"`
	Skills     []Skill            `bson:"skills" json:"skills"`
	Experience []Experience       `bson:"experience" json:"experience"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}

type Experience struct {
	Company   string    `bson:"company" json:"company"`
	Position  string    `bson:"position" json:"position"`
	StartDate time.Time `bson:"start_date" json:"start_date"`
	EndDate   time.Time `bson:"end_date" json:"end_date"`
}
