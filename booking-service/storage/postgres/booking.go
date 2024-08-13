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

type BookingRepo struct {
	db *sql.DB
}

func NewBookingRepo(db *sql.DB) *BookingRepo {
	return &BookingRepo{db: db}
}

func (r *BookingRepo) Create(req *bp.BookingRes) (*bp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO bookings (
		id,
		user_id,
		provider_id,
		service_id,
		status,
		scheduled_time,
		location,
		total_price
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.db.Exec(query, id, req.UserId, req.ProviderId, req.ServiceId, req.Status, req.ScheduledTime, req.Location, req.TotalPrice)
	if err != nil {
		log.Println("Error while creating booking: ", err)
		return nil, err
	}

	log.Println("Successfully created booking")
	return nil, nil
}
func (r *BookingRepo) GetById(req *bp.ById) (*bp.BookingGetByIdRes, error) {
	booking := bp.BookingGetByIdRes{
		Booking: &bp.BookingRes{},
	}

	query := `
	SELECT 
		id,
		user_id,
		provider_id,
		service_id,
		status,
		scheduled_time,
		location,
		total_price
	FROM 
		bookings
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&booking.Booking.Id,
		&booking.Booking.UserId,
		&booking.Booking.ProviderId,
		&booking.Booking.ServiceId,
		&booking.Booking.Status,
		&booking.Booking.ScheduledTime,
		&booking.Booking.Location,
		&booking.Booking.TotalPrice,
	)

	if err != nil {
		log.Println("Error while getting booking by id: ", err)
		return nil, err
	}

	log.Println("Successfully got booking")
	return &booking, nil
}
func (r *BookingRepo) GetAll(req *bp.BookingGetAllReq) (*bp.BookingGetAllRes, error) {
	bookings := bp.BookingGetAllRes{}

	query := `
	SELECT 
		id,
		user_id,
		provider_id,
		service_id,
		status,
		scheduled_time,
		location,
		total_price
	FROM 
		bookings
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

	var limit int32 = 10
	var offset int32 = (req.Filter.Page - 1) * limit

	args = append(args, limit, offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := r.db.Query(query, args...)
	if err == sql.ErrNoRows {
		log.Println("Bookings not found")
		return nil, errors.New("bookings list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving bookings: ", err)
		return nil, err
	}

	for rows.Next() {
		booking := bp.BookingRes{}

		err := rows.Scan(
			&booking.Id,
			&booking.UserId,
			&booking.ProviderId,
			&booking.ServiceId,
			&booking.Status,
			&booking.ScheduledTime,
			&booking.Location,
			&booking.TotalPrice,
		)

		if err != nil {
			log.Println("Error while scanning all bookings: ", err)
			return nil, err
		}

		bookings.Bookings = append(bookings.Bookings, &booking)
	}

	log.Println("Successfully fetched all bookings")
	return &bookings, nil
}
func (r *BookingRepo) Update(req *bp.BookingUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		bookings
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
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 AND status <> 'cancelled' "

	args = append(args, req.Id)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating booking: ", err)
		return nil, err
	}

	log.Println("Successfully updated booking")
	return nil, nil
}
func (r *BookingRepo) Delete(req *bp.ById) (*bp.Void, error) {
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
		log.Println("Error while deleting booking: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("book with id %s not found", req.Id)
	}

	log.Println("Successfully deleted book")

	return nil, nil
}
