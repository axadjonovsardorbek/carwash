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

type ProviderRepo struct {
	db *sql.DB
}

func NewProviderRepo(db *sql.DB) *ProviderRepo {
	return &ProviderRepo{db: db}
}

func (r *ProviderRepo) Create(req *bp.ProviderRes) (*bp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO providers (
		id,
		user_id,
		company_name,
		description,
		availability,
		average_rating,
		location
	) VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := r.db.Exec(query, id, req.UserId, req.CompanyName, req.Description, req.Availability, req.AverageRating, req.Location)

	if err != nil {
		log.Println("Error while creating provider: ", err)
		return nil, err
	}

	log.Println("Successfully created provider")

	return nil, nil
}

func (r *ProviderRepo) GetById(req *bp.ById) (*bp.ProviderGetByIdRes, error) {
	provider := bp.ProviderGetByIdRes{
		Provider: &bp.ProviderRes{},
	}

	query := `
	SELECT 
		id,
		user_id,
		company_name,
		description,
		availability,
		average_rating,
		location
	FROM 
		providers
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&provider.Provider.Id,
		&provider.Provider.UserId,
		&provider.Provider.CompanyName,
		&provider.Provider.Description,
		&provider.Provider.Availability,
		&provider.Provider.AverageRating,
		&provider.Provider.Location,
	)

	if err != nil {
		log.Println("Error while getting provider by id: ", err)
		return nil, err
	}

	log.Println("Successfully got provider")

	return &provider, nil
}

func (r *ProviderRepo) GetAll(req *bp.ProviderGetAllReq) (*bp.ProviderGetAllRes, error) {
	providers := bp.ProviderGetAllRes{}

	query := `
	SELECT 
		id,
		user_id,
		company_name,
		description,
		availability,
		average_rating,
		location
	FROM 
		providers
	WHERE 
		deleted_at = 0	
	`

	var args []interface{}
	var conditions []string

	if req.AverageRating > 0 {
		conditions = append(conditions, " average_rating >= $"+strconv.Itoa(len(args)+1))
		args = append(args, req.AverageRating)
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
		log.Println("Providers not found")
		return nil, errors.New("providers list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving providers: ", err)
		return nil, err
	}

	for rows.Next() {
		provider := bp.ProviderRes{}

		err := rows.Scan(
			&provider.Id,
			&provider.UserId,
			&provider.CompanyName,
			&provider.Description,
			&provider.Availability,
			&provider.AverageRating,
			&provider.Location,
		)

		if err != nil {
			log.Println("Error while scanning all providers: ", err)
			return nil, err
		}

		providers.Providers = append(providers.Providers, &provider)
	}

	log.Println("Successfully fetched all providers")

	return &providers, nil
}

func (r *ProviderRepo) Update(req *bp.ProviderUpdateReq) (*bp.Void, error) {
	query := `
	UPDATE
		providers
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Provider.CompanyName != "" && req.Provider.CompanyName != "string" {
		conditions = append(conditions, " company_name = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Provider.CompanyName)
	}
	if req.Provider.Description != "" && req.Provider.Description != "string" {
		conditions = append(conditions, " description = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Provider.Description)
	}
	if req.Provider.Availability != "" && req.Provider.Availability != "string" {
		conditions = append(conditions, " availability = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Provider.Availability)
	}
	if req.Provider.AverageRating > 0 {
		conditions = append(conditions, " average_rating = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Provider.AverageRating)
	}
	if req.Provider.Location != "" && req.Provider.Location != "string" {
		conditions = append(conditions, " location = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Provider.Location)
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
		log.Println("Error while updating provider: ", err)
		return nil, err
	}

	log.Println("Successfully updated provider")

	return nil, nil
}

func (r *ProviderRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		providers
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting provider: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("provider with id %s not found", req.Id)
	}

	log.Println("Successfully deleted provider")

	return nil, nil
}

func (r *ProviderRepo) GetProviderId(req *bp.ById) (*bp.ById, error){
	query := `
	SELECT
		id
	FROM
		providers
	WHERE
		user_id = $1
	AND 
		deleted_at = 0
	`

	var id string

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&id,
	)

	if err != nil {
		log.Println("Error while getting provider id: ", err)
		return nil, err
	}

	log.Println("Successfully got provider id")

	return &bp.ById{Id: id}, nil
}