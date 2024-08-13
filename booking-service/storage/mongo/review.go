package mongo

import (
	bp "booking/genproto/booking"
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ReviewRepo struct {
	db *mongo.Database
}

func NewReviewRepo(db *mongo.Database) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (r *ReviewRepo) Create(req *bp.ReviewRes) (*bp.Void, error) {
	collection := r.db.Collection("review")
	id := uuid.New().String()

	req.Id = id

	_, err := collection.InsertOne(context.TODO(), req)
	if err != nil {
		log.Println("Error while creating review: ", err)
		return nil, err
	}

	log.Println("Successfully created review")
	return nil, nil
}

func (r *ReviewRepo) GetById(req *bp.ById) (*bp.ReviewGetByIdRes, error) {
	collection := r.db.Collection("review")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var review bp.ReviewGetByIdRes
	err := collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&review.Review)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("review not found")
	} else if err != nil {
		return nil, err
	}

	return &review, nil
}

func (r *ReviewRepo) GetAll(req *bp.ReviewGetAllReq) (*bp.ReviewGetAllRes, error) {
	collection := r.db.Collection("review")
	mongoFilter := bson.M{}

	if req.UserId != "" && req.UserId != "string" {
		mongoFilter["user_id"] = req.UserId
	}

	if req.ProviderId != "" && req.ProviderId != "string" {
		mongoFilter["provider_id"] = req.ProviderId
	}

	var defaultLimit int64 = 10
	var offset int64 = int64(req.Filter.Page) * defaultLimit

	findOptions := options.Find()
	findOptions.SetLimit(defaultLimit)
	findOptions.SetSkip(offset)

	cursor, err := collection.Find(context.TODO(), mongoFilter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var reviews bp.ReviewGetAllRes
	for cursor.Next(context.TODO()) {
		var reading bp.ReviewRes
		if err := cursor.Decode(&reading); err != nil {
			return nil, err
		}
		reviews.Reviews = append(reviews.Reviews, &reading)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &reviews, nil
}

func (r *ReviewRepo) Update(req *bp.ReviewUpdateReq) (*bp.Void, error) {
	collection := r.db.Collection("review")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	updateDoc := bson.M{}
	if req.Rating != 0 {
		updateDoc["rating"] = req.Rating
	}
	if req.Comment != "" && req.Comment != "string" {
		updateDoc["comment"] = req.Comment
	}

	if len(updateDoc) == 0 {
		return nil, errors.New("no fields to update")
	}

	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"id": req.Id},
		bson.M{"$set": updateDoc},
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *ReviewRepo) Delete(req *bp.ById) (*bp.Void, error) {
	collection := r.db.Collection("review")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": req.Id})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
