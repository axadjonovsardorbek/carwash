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

type PaymentRepo struct {
	db *mongo.Database
}

func NewPaymentRepo(db *mongo.Database) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) Create(req *bp.PaymentRes) (*bp.Void, error) {
	collection := r.db.Collection("payment")
	id := uuid.New().String()

	req.Id = id

	_, err := collection.InsertOne(context.TODO(), req)
	if err != nil {
		log.Println("Error while creating payment: ", err)
		return nil, err
	}

	log.Println("Successfully created payment")
	return nil, nil
}

func (r *PaymentRepo) GetById(req *bp.ById) (*bp.PaymentGetByIdRes, error) {
	collection := r.db.Collection("payment")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var payment bp.PaymentGetByIdRes
	err := collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&payment.Payment)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("payment not found")
	} else if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *PaymentRepo) GetAll(req *bp.PaymentGetAllReq) (*bp.PaymentGetAllRes, error) {
	collection := r.db.Collection("payment")
	mongoFilter := bson.M{}

	if req.Status != "" && req.Status != "string" {
		mongoFilter["status"] = req.Status
	}
	if req.UserId != "" && req.UserId != "string" {
		mongoFilter["user_id"] = req.UserId
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

	var payments bp.PaymentGetAllRes
	for cursor.Next(context.TODO()) {
		var reading bp.PaymentRes
		if err := cursor.Decode(&reading); err != nil {
			return nil, err
		}
		payments.Payments = append(payments.Payments, &reading)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &payments, nil
}

func (r *PaymentRepo) Update(req *bp.PaymentUpdateReq) (*bp.Void, error) {
	collection := r.db.Collection("payment")
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

func (r *PaymentRepo) Delete(req *bp.ById) (*bp.Void, error) {
	collection := r.db.Collection("payment")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": req.Id})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
