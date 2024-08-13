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

type ServiceRepo struct {
	db *mongo.Database
}

func NewServiceRepo(db *mongo.Database) *ServiceRepo {
	return &ServiceRepo{db: db}
}

func (r *ServiceRepo) Create(req *bp.ServiceRes) (*bp.Void, error) {
	collection := r.db.Collection("service")
	id := uuid.New().String()

	req.Id = id

	res, err := collection.InsertOne(context.TODO(), req)

	fmt.Println(res)

	if err != nil {
		log.Println("Error while creating service: ", err)
		return nil, err
	}

	log.Println("Successfully created service")

	return nil, nil
}
func (r *ServiceRepo) GetById(req *bp.ById) (*bp.ServiceGetByIdRes, error) {
	collection := r.db.Collection("service")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var service bp.ServiceGetByIdRes
	err := collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&service.Service)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("service not found")
	} else if err != nil {
		return nil, err
	}

	return &service, nil
}
func (r *ServiceRepo) GetAll(req *bp.ServiceGetAllReq) (*bp.ServiceGetAllRes, error) {
	collection := r.db.Collection("service")
	mongoFilter := bson.M{}
	if req.Price > 0 {
		mongoFilter["average_rating"] = req.Price
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

	var services bp.ServiceGetAllRes

	for cursor.Next(context.TODO()) {
		var reading bp.ServiceRes
		err := cursor.Decode(&reading)
		if err != nil {
			return nil, err
		}

		service := &bp.ServiceRes{
			Id:          reading.Id,
			Name:        reading.Name,
			Description: reading.Description,
			Price:       reading.Price,
			Duration:    reading.Duration,
		}

		services.Services = append(services.Services, service)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &services, nil
}
func (r *ServiceRepo) Update(req *bp.ServiceUpdateReq) (*bp.Void, error) {
	collection := r.db.Collection("service")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	updateDoc := bson.M{}

	if req.Service.Name != "" && req.Service.Name != "string" {
		updateDoc["name"] = req.Service.Name
	}
	if req.Service.Duration > 0 {
		updateDoc["duration"] = req.Service.Duration
	}
	if req.Service.Description != "" && req.Service.Description != "string" {
		updateDoc["description"] = req.Service.Description
	}
	if req.Service.Price > 0 {
		updateDoc["price"] = req.Service.Price
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
func (r *ServiceRepo) Delete(req *bp.ById) (*bp.Void, error) {
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
