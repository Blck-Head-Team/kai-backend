package classes

import (
	context "context"
	"kai-backend/business/paginations"
	"kai-backend/business/schedules"
	"time"
)

type Domain struct {
	Id                 uint
	Name               string `validate:"required"`
	Description        string `validate:"required"`
	Banner_picture_url string `validate:"required"`
	Card_picture_url   string `validate:"required"`
	Online             bool
	Link               string
	Category           string `validate:"required"`
	Status             string `validate:"required"`
	GymID              uint   `validate:"required"`
	GymName            string
	Price              int
	Membership_typeID  uint `validate:"required"`
	// Booking_details    []booking_details.Domain
	Schedules []schedules.Domain
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DomainRepository interface {
	GetAll(ctx context.Context, pagination paginations.Domain, domain Domain) ([]Domain, error)
	CountAll(ctx context.Context) (int, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, domain Domain, gymId string) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
}

type DomainService interface {
	GetAll(ctx context.Context, pagination paginations.Domain, domain Domain) ([]Domain, error)
	CountAll(ctx context.Context) (int, error)
	GetById(ctx context.Context, id string) (Domain, error)
	Create(ctx context.Context, domain Domain, gymId string) (Domain, error)
	Update(ctx context.Context, id string, domain Domain) (Domain, error)
	Delete(ctx context.Context, id string) error
}
