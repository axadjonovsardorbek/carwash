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

type BookingRepo struct {
	db *mongo.Database
}

func NewBookingRepo(db *mongo.Database) *BookingRepo {
	return &BookingRepo{db: db}
}

func (r *BookingRepo) Create(req *bp.BookingRes) (*bp.Void, error) {
	collection := r.db.Collection("booking")
	id := uuid.New().String()

	req.Id = id

	res, err := collection.InsertOne(context.TODO(), req)

	fmt.Println(res)

	if err != nil {
		log.Println("Error while creating booking: ", err)
		return nil, err
	}

	log.Println("Successfully created booking")

	return nil, nil
}
func (r *BookingRepo) GetById(req *bp.ById) (*bp.BookingGetByIdRes, error) {
	collection := r.db.Collection("booking")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var booking bp.BookingGetByIdRes
	err := collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&booking.Booking)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("booking not found")
	} else if err != nil {
		return nil, err
	}

	return &booking, nil
}
func (r *BookingRepo) GetAll(req *bp.BookingGetAllReq) (*bp.BookingGetAllRes, error) {
	collection := r.db.Collection("booking")
	mongoFilter := bson.M{}
	if req.UserId != "" && req.UserId != "string" {
		mongoFilter["user_id"] = req.UserId
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

	var bookings bp.BookingGetAllRes

	for cursor.Next(context.TODO()) {
		var reading bp.BookingRes
		err := cursor.Decode(&reading)
		if err != nil {
			return nil, err
		}

		booking := &bp.BookingRes{
			Id:            reading.Id,
			UserId:        reading.UserId,
			ProviderId:    reading.ProviderId,
			ServiceId:     reading.ServiceId,
			Status:        reading.Status,
			ScheduledTime: reading.ScheduledTime,
			Location:      reading.Location,
			TotalPrice:    reading.TotalPrice,
		}

		bookings.Bookings = append(bookings.Bookings, booking)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &bookings, nil
}
func (r *BookingRepo) Update(req *bp.BookingUpdateReq) (*bp.Void, error) {
	collection := r.db.Collection("booking")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	updateDoc := bson.M{}

	if req.Status != "" && req.Status != "string" {
		updateDoc["status"] = req.Status
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
func (r *BookingRepo) Delete(req *bp.ById) (*bp.Void, error) {
	collection := r.db.Collection("booking")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": req.Id})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
