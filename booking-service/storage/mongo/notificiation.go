package mongo

import (
	bp "booking/genproto/booking"
	"context"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NotificationRepo struct {
	db *mongo.Database
}

func NewNotificationRepo(db *mongo.Database) *NotificationRepo {
	return &NotificationRepo{db: db}
}

func (r *NotificationRepo) Create(req *bp.NotificationRes) (*bp.Void, error) {
	collection := r.db.Collection("notifications")
	id := uuid.New().String()

	req.Id = id
	req.IsRead = "false"
	req.CreatedAt = time.Now().String()

	_, err := collection.InsertOne(context.TODO(), req)
	if err != nil {
		log.Println("Error while creating notification: ", err)
		return nil, err
	}

	log.Println("Successfully created notification")
	return nil, nil
}

func (r *NotificationRepo) GetById(req *bp.ById) (*bp.NotificationGetByIdRes, error) {
	collection := r.db.Collection("notifications")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var notification bp.NotificationGetByIdRes
	err := collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&notification.Notification)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("notification not found")
	} else if err != nil {
		return nil, err
	}

	return &notification, nil
}

func (r *NotificationRepo) GetAll(req *bp.NotificationGetAllReq) (*bp.NotificationGetAllRes, error) {
	collection := r.db.Collection("notifications")
	mongoFilter := bson.M{}

	if req.UserId != "" && req.UserId != "string" {
		mongoFilter["userid"] = req.UserId
	}
	if req.IsRead != "" && req.IsRead != "string" {
		mongoFilter["isread"] = req.IsRead
	}

	var defaultLimit int64 = 10
	var offset int64 = int64(req.Filter.Page - 1) * defaultLimit

	findOptions := options.Find()
	findOptions.SetLimit(defaultLimit)
	findOptions.SetSkip(offset)

	cursor, err := collection.Find(context.TODO(), mongoFilter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var notifications bp.NotificationGetAllRes
	for cursor.Next(context.TODO()) {
		var reading bp.NotificationRes
		if err := cursor.Decode(&reading); err != nil {
			return nil, err
		}
		notifications.Notification = append(notifications.Notification, &reading)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &notifications, nil
}

func (r *NotificationRepo) Update(req *bp.NotificationUpdateReq) (*bp.Void, error) {
	collection := r.db.Collection("notifications")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	updateDoc := bson.M{}
	if req.IsRead != "" && req.IsRead != "string" {
		updateDoc["isread"] = req.IsRead
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
