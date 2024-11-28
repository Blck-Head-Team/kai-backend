package schedules

import (
	"context"
	"kai-backend/repository/mysql/trains"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	TrainId   uint
	Train     trains.Train
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Create(ctx context.Context, schedules Domain) (Domain, error)
	Get(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, schedules Domain) (Domain, error)
	Delete(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Create(ctx context.Context, schedules Domain) error
	Get(ctx context.Context) ([]Domain, error)
	GetById(ctx context.Context, id int) (Domain, error)
	Update(ctx context.Context, schedules Domain) (Domain, error)
	Delete(ctx context.Context, id int) (Domain, error)
}
