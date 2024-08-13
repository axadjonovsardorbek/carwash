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

type ServiceRepo struct {
	db *sql.DB
}

func NewServiceRepo(db *sql.DB) *ServiceRepo {
	return &ServiceRepo{db: db}
}

func (r *ServiceRepo) Create(req *bp.ServiceRes) (*bp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO services (
		id,
		name,
		description,
		price,
		duration
	) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(query, id, req.Name, req.Description, req.Price, req.Duration)

	if err != nil {
		log.Println("Error while creating comment: ", err)
		return nil, err
	}

	log.Println("Successfully created service")

	return nil, nil
}
func (r *ServiceRepo) GetById(req *bp.ById) (*bp.ServiceGetByIdRes, error) {
	service := bp.ServiceGetByIdRes{
		Service: &bp.ServiceRes{},
	}

	query := `
	SELECT 
		id,
		name,
		description,
		price,
		duration
	FROM 
		services
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&service.Service.Id,
		&service.Service.Name,
		&service.Service.Description,
		&service.Service.Price,
		&service.Service.Duration,
	)

	if err != nil {
		log.Println("Error while getting service by id: ", err)
		return nil, err
	}

	log.Println("Successfully got service")

	return &service, nil
}
func (r *ServiceRepo) GetAll(req *bp.ServiceGetAllReq) (*bp.ServiceGetAllRes, error) {
	services := bp.ServiceGetAllRes{}

	query := `
	SELECT
		id,
		name,
		description,
		price,
		duration
	FROM 
		services
	WHERE 
		deleted_at = 0	
	`

	var args []interface{}
	var conditions []string

	if req.Price >= 0 {
		conditions = append(conditions, " price <= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Price)
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
		log.Println("Services not found")
		return nil, errors.New("services list is empty")
	}

	if err != nil {
		log.Println("Error while retriving services: ", err)
		return nil, err
	}

	for rows.Next() {
		service := bp.ServiceRes{}

		err := rows.Scan(
			&service.Id,
			&service.Name,
			&service.Description,
			&service.Price,
			&service.Duration,
		)

		if err != nil {
			log.Println("Error while scanning all services: ", err)
			return nil, err
		}

		services.Services = append(services.Services, &service)
	}

	log.Println("Successfully fetched all services")

	return &services, nil
}
func (r *ServiceRepo) Update(req *bp.ServiceUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		services
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Service.Name != "" && req.Service.Name != "string" {
		conditions = append(conditions, " name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Service.Name)
	}
	if req.Service.Description != "" && req.Service.Description != "string" {
		conditions = append(conditions, " description = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Service.Description)
	}
	if req.Service.Duration > 0 {
		conditions = append(conditions, " duration = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Service.Duration)
	}
	if req.Service.Price >= 0 {
		conditions = append(conditions, " price = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Service.Price)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0"

	args = append(args, req.Id)

	_, err := r.db.Exec(query, args...)

	if err != nil {
		log.Println("Error while updating service: ", err)
		return nil, err
	}

	log.Println("Successfully updated service")

	return nil, nil
}
func (r *ServiceRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		services
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting service: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("service with id %s not found", req.Id)
	}

	log.Println("Successfully deleted service")

	return nil, nil
}
