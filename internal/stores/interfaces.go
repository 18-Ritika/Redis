package stores

import (
	"Training/Redis/Redis/internal/models"
	"context"
)

type StudentsStores interface {
	Get(ctx context.Context, id string) (string, error)
	Post(ctx context.Context, student models.Student) (int, error)
	Delete(ctx context.Context, id string) error
}
