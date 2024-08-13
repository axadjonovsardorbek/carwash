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

type ProviderServiceRepo struct {
	db *mongo.Database
}

func NewProviderServiceRepo(db *mongo.Database) *ProviderServiceRepo {
	return &ProviderServiceRepo{db: db}
}

func (r *ProviderServiceRepo) Create(req *bp.ProviderServiceRes) (*bp.Void, error) {
	collection := r.db.Collection("provider_services")
	id := uuid.New().String()

	req.Id = id

	_, err := collection.InsertOne(context.TODO(), req)
	if err != nil {
		log.Println("Error while creating provider service: ", err)
		return nil, err
	}

	log.Println("Successfully created provider service")
	return nil, nil
}

func (r *ProviderServiceRepo) GetById(req *bp.ById) (*bp.ProviderServiceGetByIdRes, error) {
	collection := r.db.Collection("provider_services")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var providerService bp.ProviderServiceGetByIdRes
	err := collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&providerService.ProviderService)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("provider service not found")
	} else if err != nil {
		return nil, err
	}

	log.Println("Successfully retrieved provider service")
	return &providerService, nil
}

func (r *ProviderServiceRepo) GetAll(req *bp.ProviderServiceGetAllReq) (*bp.ProviderServiceGetAllRes, error) {
	collection := r.db.Collection("provider_services")
	mongoFilter := bson.M{}
	if req.ProviderId != "" && req.ProviderId != "string" {
		mongoFilter["provider_id"] = req.ProviderId
	}
	if req.UserId != "" && req.UserId != "string" {
		mongoFilter["user_id"] = req.UserId
	}

	var defaultLimit int64 = 10
	var offset int64 = int64(req.Filter.Page-1) * defaultLimit

	findOptions := options.Find()
	findOptions.SetLimit(defaultLimit)
	findOptions.SetSkip(offset)

	cursor, err := collection.Find(context.TODO(), mongoFilter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var providerServices bp.ProviderServiceGetAllRes

	for cursor.Next(context.TODO()) {
		var providerService bp.ProviderServiceRes
		if err := cursor.Decode(&providerService); err != nil {
			return nil, err
		}
		providerServices.ProviderServices = append(providerServices.ProviderServices, &providerService)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	log.Println("Successfully retrieved all provider services")
	return &providerServices, nil
}

func (r *ProviderServiceRepo) Delete(req *bp.ById) (*bp.Void, error) {
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
