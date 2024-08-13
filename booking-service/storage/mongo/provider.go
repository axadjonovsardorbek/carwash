package mongo

import (
	bp "booking/genproto/booking"
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProviderRepo struct {
	db *mongo.Database
}

func NewProviderRepo(db *mongo.Database) *ProviderRepo {
	return &ProviderRepo{db: db}
}

func (r *ProviderRepo) Create(req *bp.ProviderRes) (*bp.Void, error) {
	collection := r.db.Collection("provider")
	id := uuid.New().String()

	req.Id = id

	res, err := collection.InsertOne(context.TODO(), req)

	fmt.Println(res)

	if err != nil {
		log.Println("Error while creating provider: ", err)
		return nil, err
	}

	log.Println("Successfully created provider")

	return nil, nil
}
func (r *ProviderRepo) GetById(req *bp.ById) (*bp.ProviderGetByIdRes, error) {
	collection := r.db.Collection("provider")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var provider bp.ProviderGetByIdRes
	err := collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&provider.Provider)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("provider not found")
	} else if err != nil {
		return nil, err
	}

	return &provider, nil
}
func (r *ProviderRepo) GetAll(req *bp.ProviderGetAllReq) (*bp.ProviderGetAllRes, error) {
	collection := r.db.Collection("provider")
	mongoFilter := bson.M{}
	if req.AverageRating > 0 {
		mongoFilter["average_rating"] = req.AverageRating
	}

	var defaultLimit int64
	var offset int64

	defaultLimit = 10
	offset = int64(req.Filter.Page) * defaultLimit

	findOptions := options.Find()
	findOptions.SetLimit(defaultLimit)
	findOptions.SetSkip(offset)

	cursor, err := collection.Find(context.TODO(), mongoFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var providers bp.ProviderGetAllRes

	for cursor.Next(context.TODO()) {
		var reading bp.ProviderRes
		err := cursor.Decode(&reading)
		if err != nil {
			return nil, err
		}

		provider := &bp.ProviderRes{
			Id:            reading.Id,
			UserId:        reading.UserId,
			CompanyName:   reading.CompanyName,
			Description:   reading.Description,
			Availability:  reading.Availability,
			AverageRating: reading.AverageRating,
			Location:      reading.Location,
		}

		providers.Providers = append(providers.Providers, provider)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &providers, nil
}
func (r *ProviderRepo) Update(req *bp.ProviderUpdateReq) (*bp.Void, error) {
	collection := r.db.Collection("provider")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	updateDoc := bson.M{}

	if req.Provider.UserId != "" && req.Provider.UserId != "string"{
		updateDoc["user_id"] = req.Provider.UserId
	}
	if req.Provider.CompanyName != "" && req.Provider.CompanyName != "string"{
		updateDoc["company_name"] = req.Provider.CompanyName
	}
	if req.Provider.Description != "" && req.Provider.Description != "string"{
		updateDoc["description"] = req.Provider.Description
	}
	if req.Provider.Availability != "" && req.Provider.Availability != "string"{
		updateDoc["availability"] = req.Provider.Availability
	}
	if req.Provider.Location != "" && req.Provider.Location != "string"{
		updateDoc["location"] = req.Provider.Location
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
func (r *ProviderRepo) Delete(req *bp.ById) (*bp.Void, error) {
	collection := r.db.Collection("provider")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": req.Id})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
