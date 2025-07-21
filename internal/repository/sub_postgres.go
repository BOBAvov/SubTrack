package repository

import (
	"fmt"
	"github.com/BOBAvov/sub_track"
	"github.com/jmoiron/sqlx"
)

type SubscriptionRepository struct {
	db *sqlx.DB
}

const (
	subsTable = "subs"
)

func NewSubPostgres(db *sqlx.DB) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (s *SubscriptionRepository) Create(sub sub_track.Subscription) (id int, err error) {
	const op = "repository.sub_postgres.create"

	sub.StartDate, err = PostgresNormalDate(sub.StartDate, false)
	if err != nil {
		return 0, err
	}

	sub.EndDate, err = PostgresNormalDate(sub.EndDate, true)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO %s (user_id,service_name,price,start_date,end_date) VALUES ($1,$2,$3,$4,$5) RETURNING id", subsTable)
	row := s.db.QueryRow(query, sub.Userid, sub.ServiceName, sub.Price, sub.StartDate, sub.EndDate)

	if err = row.Scan(&id); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *SubscriptionRepository) GetAll() ([]sub_track.Subscription, error) {
	var subs []sub_track.Subscription
	const op = "repository.sub_postgres.getAll"

	query := fmt.Sprintf("SELECT * FROM %s", subsTable)
	err := s.db.Select(&subs, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return subs, nil
}

func (s *SubscriptionRepository) GetById(id int) (sub sub_track.Subscription, err error) {
	const op = "repository.sub_postgres.getById"
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", subsTable)
	err = s.db.Get(&sub, query, id)

	if err != nil {
		return sub, fmt.Errorf("%s: %w", op, err)
	}

	return sub, nil
}

func (s *SubscriptionRepository) Update(id int, input sub_track.SubscriptionUpdate) (err error) {
	const op = "repository.sub_postgres.update"

	if input.EndDate != "" {
		input.EndDate, err = PostgresNormalDate(input.EndDate, true)
		if err != nil {
			return err
		}
		query := fmt.Sprintf("UPDATE %s SET end_date = $1 WHERE id = $2", subsTable)
		_, err := s.db.Exec(query, input.EndDate, id)

		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	if input.Price > 0 {
		query := fmt.Sprintf("UPDATE %s SET price = $1 WHERE id = $2", subsTable)
		_, err := s.db.Exec(query, input.Price, id)

		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	return nil

}
func (s *SubscriptionRepository) Delete(id int) error {
	const op = "repository.sub_postgres.delete"
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", subsTable)
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
