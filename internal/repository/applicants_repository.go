package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"recruiter/internal/models"
)

type ApplicantRepository interface {
	GetApplicants() ([]models.Applicants, error)
	GetApplicant(id primitive.ObjectID) (models.Applicants, error)
	CreateApplicant(applicants models.Applicants) (models.Applicants, error)
	UpdateApplicant(id primitive.ObjectID, applicant models.Applicants) (models.Applicants, error)
	DeleteApplicant(id primitive.ObjectID) error
}

func (r *MongoRepository) GetApplicants() ([]models.Applicants, error) {
	var applicants []models.Applicants
	cursor, err := r.db.Collection("applicants").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err = cursor.All(context.Background(), &applicants); err != nil {
		return nil, err
	}

	return applicants, nil
}

func (r *MongoRepository) GetApplicant(id primitive.ObjectID) (models.Applicants, error) {
	var applicant models.Applicants
	err := r.db.Collection("applicants").FindOne(context.Background(), bson.M{"_id": id}).Decode(&applicant)
	return applicant, err
}

func (r *MongoRepository) CreateApplicant(applicant models.Applicants) (models.Applicants, error) {
	result, err := r.db.Collection("applicants").InsertOne(context.Background(), applicant)
	if err != nil {
		return models.Applicants{}, err
	}
	applicant.ID = result.InsertedID.(primitive.ObjectID)
	return applicant, nil
}

func (r *MongoRepository) UpdateApplicant(id primitive.ObjectID, applicant models.Applicants) (models.Applicants, error) {
	_, err := r.db.Collection("applicants").UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": applicant},
	)
	if err != nil {
		return models.Applicants{}, err
	}
	return r.GetApplicant(id)
}

func (r *MongoRepository) DeleteApplicant(id primitive.ObjectID) error {
	_, err := r.db.Collection("applicants").DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}
