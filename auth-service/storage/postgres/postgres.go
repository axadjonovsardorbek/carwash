package postgres

import (
	"database/sql"
	"fmt"

	"auth/config"
	"auth/storage"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

type Storage struct {
	Db    *sql.DB
	UserS storage.UserI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	rAddr := fmt.Sprintf("%s%s", config.REDIS_HOST, config.REDIS_PORT)

	rdb := redis.NewClient(&redis.Options{
		Addr: rAddr,
	})

	user := NewUsersRepo(db, rdb)

	return &Storage{
		Db:    db,
		UserS: user,
	}, nil
}
