package repositories

import "context"

type Repository[T any] interface {
	Create(ctx context.Context, entity *T) (*T, error)
	FindAll(ctx context.Context) ([]*T, error)
	Update(ctx context.Context, id string, entity *T) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (*T, error)
}
