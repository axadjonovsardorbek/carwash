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

type PaymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (r *PaymentRepo) Create(req *bp.PaymentRes) (*bp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO payments (
		id,
		user_id,
		booking_id,
		amount,
		status,
		payment_method
	) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(query, id, req.UserId, req.BookingId, req.Amount, req.Status, req.PaymentMethod)
	if err != nil {
		log.Println("Error while creating payment: ", err)
		return nil, err
	}

	log.Println("Successfully created payment")
	return nil, nil
}

func (r *PaymentRepo) GetById(req *bp.ById) (*bp.PaymentGetByIdRes, error) {
	payment := bp.PaymentGetByIdRes{
		Payment: &bp.PaymentRes{},
	}

	query := `
	SELECT 
		id,
		user_id,
		booking_id,
		amount,
		status,
		payment_method
	FROM 
		payments
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&payment.Payment.Id,
		&payment.Payment.UserId,
		&payment.Payment.BookingId,
		&payment.Payment.Amount,
		&payment.Payment.Status,
		&payment.Payment.PaymentMethod,
	)

	if err != nil {
		log.Println("Error while getting payment by id: ", err)
		return nil, err
	}

	log.Println("Successfully got payment")
	return &payment, nil
}

func (r *PaymentRepo) GetAll(req *bp.PaymentGetAllReq) (*bp.PaymentGetAllRes, error) {
	payments := bp.PaymentGetAllRes{}

	query := `
	SELECT 
		id,
		user_id,
		booking_id,
		amount,
		status,
		payment_method
	FROM 
		payments
	WHERE 
		deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.Status != "" && req.Status != "string" {
		conditions = append(conditions, " status = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Status)
	}
	if req.UserId != "" && req.UserId != "string" {
		conditions = append(conditions, " user_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.UserId)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var limit int32 = 10
	var offset int32 = (req.Filter.Page - 1) * limit

	args = append(args, limit, offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)
	if err == sql.ErrNoRows {
		log.Println("Payments not found")
		return nil, errors.New("payments list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving payments: ", err)
		return nil, err
	}

	for rows.Next() {
		payment := bp.PaymentRes{}

		err := rows.Scan(
			&payment.Id,
			&payment.UserId,
			&payment.BookingId,
			&payment.Amount,
			&payment.Status,
			&payment.PaymentMethod,
		)

		if err != nil {
			log.Println("Error while scanning all payments: ", err)
			return nil, err
		}

		payments.Payments = append(payments.Payments, &payment)
	}

	log.Println("Successfully fetched all payments")
	return &payments, nil
}

func (r *PaymentRepo) Update(req *bp.PaymentUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		payments
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Status != "" && req.Status != "string" {
		conditions = append(conditions, " status = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Status)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	args = append(args, req.Id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating payment: ", err)
		return nil, err
	}

	log.Println("Successfully updated payment")
	return nil, nil
}

func (r *PaymentRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		bookings
	SET 
		status = 'cancelled'
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	AND 
		status <> 'completed'
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting payment: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("payment with id %s not found", req.Id)
	}

	log.Println("Successfully deleted payment")
	return nil, nil
}
