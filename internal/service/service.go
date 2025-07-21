package service

import (
	"github.com/BOBAvov/sub_track"
	"github.com/BOBAvov/sub_track/internal/repository"
	"log/slog"
)

type Subscription interface {
	Create(input sub_track.Subscription) (int, error)
	GetAll() ([]sub_track.Subscription, error)
	GetById(id int) (sub_track.Subscription, error)
	Update(id int, input sub_track.SubscriptionUpdate) error
	Delete(id int) error
}

type Service struct {
	log *slog.Logger
	Subscription
}

func NewService(repos *repository.Repository, logger *slog.Logger) *Service {
	return &Service{
		Subscription: NewSubscriptionService(repos.Subscription),
		log:          logger,
	}
}
