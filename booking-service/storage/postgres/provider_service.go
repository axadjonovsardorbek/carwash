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

type ProviderServiceRepo struct {
	db *sql.DB
}

func NewProviderServiceRepo(db *sql.DB) *ProviderServiceRepo {
	return &ProviderServiceRepo{db: db}
}

func (r *ProviderServiceRepo) Create(req *bp.ProviderServiceRes) (*bp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO provider_services (
		id,
		user_id,
		service_id,
		provider_id
	) VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(query, id, req.UserId, req.ServiceId, req.ProviderId)

	if err != nil {
		log.Println("Error while creating provider service: ", err)
		return nil, err
	}

	log.Println("Successfully created provider service")

	return nil, nil
}

func (r *ProviderServiceRepo) GetById(req *bp.ById) (*bp.ProviderServiceGetByIdRes, error) {
	providerService := bp.ProviderServiceGetByIdRes{
		ProviderService: &bp.ProviderServiceRes{},
	}

	query := `
	SELECT 
		id,
		user_id,
		service_id,
		provider_id
	FROM 
		provider_services
	WHERE 
		id = $1
	AND 
		deleted_at = 0	
	`

	row := r.db.QueryRow(query, req.Id)

	err := row.Scan(
		&providerService.ProviderService.Id,
		&providerService.ProviderService.UserId,
		&providerService.ProviderService.ServiceId,
		&providerService.ProviderService.ProviderId,
	)

	if err != nil {
		log.Println("Error while getting provider service by id: ", err)
		return nil, err
	}

	log.Println("Successfully got provider service")

	return &providerService, nil
}

func (r *ProviderServiceRepo) GetAll(req *bp.ProviderServiceGetAllReq) (*bp.ProviderServiceGetAllRes, error) {
	providerServices := bp.ProviderServiceGetAllRes{}

	query := `
	SELECT 
		id,
		user_id,
		service_id,
		provider_id
	FROM 
		provider_services
	WHERE 
		deleted_at = 0	
	`

	var args []interface{}
	var conditions []string

	if req.ProviderId != "" && req.ProviderId != "string" {
		conditions = append(conditions, " provider_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.ProviderId)
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
		log.Println("Provider services not found")
		return nil, errors.New("provider services list is empty")
	}

	if err != nil {
		log.Println("Error while retrieving provider services: ", err)
		return nil, err
	}

	for rows.Next() {
		providerService := bp.ProviderServiceRes{}

		err := rows.Scan(
			&providerService.Id,
			&providerService.UserId,
			&providerService.ServiceId,
			&providerService.ProviderId,
		)

		if err != nil {
			log.Println("Error while scanning provider services: ", err)
			return nil, err
		}

		providerServices.ProviderServices = append(providerServices.ProviderServices, &providerService)
	}

	log.Println("Successfully fetched all provider services")

	return &providerServices, nil
}

func (r *ProviderServiceRepo) Delete(req *bp.ById) (*bp.Void, error) {
	query := `
	UPDATE 
		provider_services
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := r.db.Exec(query, req.Id)

	if err != nil {
		log.Println("Error while deleting provider service: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("provider service with id %s not found", req.Id)
	}

	log.Println("Successfully deleted provider service")

	return nil, nil
}
