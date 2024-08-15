package postgres

import (
	bp "booking/genproto/booking"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type ReviewRepo struct {
	db *sql.DB
}

func NewReviewRepo(db *sql.DB) *ReviewRepo {
	return &ReviewRepo{db: db}
}

func (r *ReviewRepo) Create(req *bp.ReviewRes) (*bp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO reviews (
		id,
		booking_id,
		user_id,
		provider_id,
		rating,
		comment
	) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(query, id, req.BookingId, req.UserId, req.ProviderId, req.Rating, req.Comment)

	if err != nil {
		log.Println("Error while creating review: ", err)
		return nil, err
	}

	log.Println("Successfully created review")

	return nil, nil
}
func (r *ReviewRepo) GetById(req *bp.ById) (*bp.ReviewGetByIdRes, error) {
	review := bp.ReviewGetByIdRes{
		Review: &bp.ReviewRes{},
	}

	query := `
	SELECT 
		id,
		booking_id,
		user_id,
		provider_id,
		rating,
		comment
	FROM 
		reviews
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&review.Review.Id,
		&review.Review.BookingId,
		&review.Review.UserId,
		&review.Review.ProviderId,
		&review.Review.Rating,
		&review.Review.Comment,
	)

	if err != nil {
		log.Println("Error while getting review by id: ", err)
		return nil, err
	}

	log.Println("Successfully got review")

	return &review, nil
}
func (r *ReviewRepo) GetAll(req *bp.ReviewGetAllReq) (*bp.ReviewGetAllRes, error) {
	reviews := bp.ReviewGetAllRes{}

	query := `
	SELECT 
		id,
		booking_id,
		user_id,
		provider_id,
		rating,
		comment
	FROM 
		reviews
	WHERE 
		deleted_at = 0	
	`

	var args []interface{}
	var conditions []string

	if req.UserId != "" && req.UserId != "string" {
		conditions = append(conditions, " user_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.UserId)
	}
	if req.ProviderId != "" && req.ProviderId != "string" {
		conditions = append(conditions, " provider_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.ProviderId)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var limit int32
	var offset int32

	limit = 10
	offset = (req.Filter.Page - 1) * limit

	args = append(args, limit, offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)

	if err == sql.ErrNoRows {
		log.Println("Reviews not found")
		return nil, errors.New("reviews list is empty")
	}

	if err != nil {
		log.Println("Error while retriving reviews: ", err)
		return nil, err
	}

	for rows.Next() {
		review := bp.ReviewRes{}

		err := rows.Scan(
			&review.Id,
			&review.BookingId,
			&review.UserId,
			&review.ProviderId,
			&review.Rating,
			&review.Comment,
		)

		if err != nil {
			log.Println("Error while scanning all reviews: ", err)
			return nil, err
		}

		reviews.Reviews = append(reviews.Reviews, &review)
	}

	log.Println("Successfully fetched all reviews")

	return &reviews, nil
}
func (r *ReviewRepo) Update(req *bp.ReviewUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		reviews
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Comment != "" && req.Comment != "string" {
		conditions = append(conditions, " comment = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Comment)
	}
	if req.Rating > 0 {
		conditions = append(conditions, " rating = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Rating)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 AND user_id = $" + strconv.Itoa(len(args)+2)

	args = append(args, req.Id, req.UserId)

	_, err := r.db.Exec(query, args...)

	if err != nil {
		log.Println("Error while updating review: ", err)
		return nil, err
	}

	log.Println("Successfully updated review")

	return nil, nil
}
func (r *ReviewRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		reviews
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting review: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("review with id %s not found", req.Id)
	}

	log.Println("Successfully deleted review")

	return nil, nil
}

func (r *ReviewRepo) GetProviderRating(req *bp.ById) (*bp.GetProviderRatingRes, error){
	query := `
	SELECT
		SUM(rating),
		COUNT(rating)
	FROM
		reviews
	WHERE
		provider_id = $1
	AND 
		deleted_at = 0
	`

	var rating int64
	var count int64

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&rating,
		&count,
	)

	if err != nil {
		log.Println("Error while getting rating: ", err)
		return nil, err
	}

	log.Println("Successfully got rating")

	return &bp.GetProviderRatingRes{Rating: rating, Count: count}, nil
}