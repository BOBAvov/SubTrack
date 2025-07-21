package repository

import (
	"github.com/BOBAvov/sub_track"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type Subscription interface {
	Create(input sub_track.Subscription) (id int, err error)
	GetAll() ([]sub_track.Subscription, error)
	GetById(id int) (sub_track.Subscription, error)
	Update(id int, input sub_track.SubscriptionUpdate) error
	Delete(id int) error
}

type Repository struct {
	log *slog.Logger
	Subscription
}

func NewRepository(db *sqlx.DB, logger *slog.Logger) *Repository {
	return &Repository{
		Subscription: NewSubPostgres(db),
		log:          logger,
	}
}
