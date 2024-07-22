package models

type Skill struct {
	Name  string `bson:"name" json:"name"`
	Level string `bson:"levele" json:"level"`
}
