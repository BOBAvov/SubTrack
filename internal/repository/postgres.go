package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func PostgresNormalDate(date string, isEndDate bool) (string, error) {
	if date == "" && isEndDate {
		return "9999-12-31", nil
	}

	if len(date) == 7 {
		if isEndDate {
			return date + "-28", nil
		}
		return date + "-01", nil
	}

	return "", fmt.Errorf("invalid date format")
}
