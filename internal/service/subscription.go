package service

import (
	"fmt"
	"github.com/BOBAvov/sub_track"
	"github.com/BOBAvov/sub_track/internal/repository"
)

type SubscriptionService struct {
	repo repository.Subscription
}

func NewSubscriptionService(repo repository.Subscription) *SubscriptionService {
	return &SubscriptionService{repo: repo}
}

func (s *SubscriptionService) Create(input sub_track.Subscription) (id int, err error) {
	return s.repo.Create(input)
}

func (s *SubscriptionService) GetAll() ([]sub_track.Subscription, error) {
	return s.repo.GetAll()
}

func (s *SubscriptionService) GetById(id int) (sub_track.Subscription, error) {
	return s.repo.GetById(id)
}

func (s *SubscriptionService) Update(id int, input sub_track.SubscriptionUpdate) error {
	if input.Price == 0 && input.EndDate == "" {
		return fmt.Errorf("invalid input")
	}

	return s.repo.Update(id, input)
}
func (s *SubscriptionService) Delete(id int) error {
	return s.repo.Delete(id)
}
