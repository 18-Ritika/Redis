package students

import (
	"Training/Redis/Redis/internal/models"
	"Training/Redis/Redis/internal/stores"
	"context"
	"strconv"
)

type Service struct {
	store stores.StudentsStores
}

func New(store stores.StudentsStores) *Service {
	return &Service{store: store}
}

func (s *Service) Get(ctx context.Context, ID string) (string, error) {
	res, err := s.store.Get(ctx, ID)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *Service) Post(ctx context.Context, student models.Student) (string, error) {
	id, err := s.store.Post(ctx, student)
	if err != nil {
		return "", err
	}

	res, err := s.store.Get(ctx, strconv.Itoa(id))
	if err != nil {
		return "", err
	}
	return res, nil
}

func (s *Service) Delete(ctx context.Context, ID string) error {
	_, err := s.store.Get(ctx, ID)
	if err != nil {
		return err
	}

	err = s.store.Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}
