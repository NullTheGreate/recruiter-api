package repository

import (
	"context"
	"recruiter/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository interface {
	GetUsers() ([]models.Users, error)
	GetUser(id primitive.ObjectID) (models.Users, error)
	FindUserByUsername(string ) (models.Users, error)
	CreateUser(user models.Users) (models.Users, error)
	UpdateUser(id primitive.ObjectID, user models.Users) (models.Users, error)
	DeleteUser(id primitive.ObjectID) error
}

func (r *MongoRepository) GetUsers() ([]models.Users, error) {
	var users []models.Users
	cursor, err := r.db.Collection("users").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *MongoRepository) GetUser(id primitive.ObjectID) (models.Users, error) {
	var user models.Users
	err := r.db.Collection("users").FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (r *MongoRepository) CreateUser(user models.Users) (models.Users, error) {
	result, err := r.db.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return models.Users{}, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r *MongoRepository) UpdateUser(id primitive.ObjectID, user models.Users) (models.Users, error) {
	_, err := r.db.Collection("users").UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": user},
	)
	if err != nil {
		return models.Users{}, err
	}
	return r.GetUser(id)
}

func (r *MongoRepository) DeleteUser(id primitive.ObjectID) error {
	_, err := r.db.Collection("users").DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

func (r *MongoRepository) FindUserByUsername(username string) (models.Users, error) {
	var user models.Users
	err := r.db.Collection("users").FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return user, err
}
