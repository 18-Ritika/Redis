package services

import (
	"Training/Redis/Redis/internal/models"
	"context"
)

type StudentServices interface {
	Get(ctx context.Context, id string) (string, error)
	Post(ctx context.Context, student models.Student) (string, error)
	Delete(ctx context.Context, id string) error
}
