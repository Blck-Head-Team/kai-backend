package newsletters

import (
	context "context"
	"kai-backend/business/paginations"
	"time"
)

type Domain struct {
	Id          int
	Title       string `validate:"required"`
	Description string `validate:"required"`
	Content     string `validate:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Usecase interface {
	GetAll(ctx context.Context, pagination paginations.Domain) ([]Domain, error)
	CountAll(ctx context.Context) (int, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
}

type Repository interface {
	GetAll(ctx context.Context, pagination paginations.Domain) ([]Domain, error)
	CountAll(ctx context.Context) (int, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, domain Domain) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
}
