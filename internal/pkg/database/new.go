package database

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"

	"github.com/edjubert/tiny-url/internal/pkg/config"
)

type Database struct {
	Db *sqlx.DB
}

func New(config *config.Configuration) (*Database, error) {
	var db *sqlx.DB
	var err error

	database := config.Database.Dbname
	username := config.Database.Username
	password := config.Database.Password
	host := config.Database.Host
	port := config.Database.Port

	switch config.Database.Driver {
	case "postgres", "pgx":
		db, err = sqlx.Connect(
			"pgx", fmt.Sprintf(
				"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username,
				database, password,
			),
		)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", config.Database.Driver)
	}

	if db == nil {
		return nil, err
	}

	db.SetMaxIdleConns(config.Database.MaxIdleConns)
	db.SetMaxOpenConns(config.Database.MaxOpenConns)

	return &Database{Db: db}, err
}
