package schedules

import (
	"context"
	"kai-backend/business/schedules"

	"gorm.io/gorm"
)

type ScheduleRepository struct {
	Conn *gorm.DB
}

func NewScheduleRepository(conn *gorm.DB) schedules.Repository {
	return &ScheduleRepository{Conn: conn}
}

// Delete implements schedules.Repository.
func (s *ScheduleRepository) Delete(ctx context.Context, id int) (schedules.Domain, error) {
	panic("unimplemented")
}

// Get implements schedules.Repository.
func (s *ScheduleRepository) Get(ctx context.Context) ([]schedules.Domain, error) {
	panic("unimplemented")
}

// GetById implements schedules.Repository.
func (s *ScheduleRepository) GetById(ctx context.Context, id int) (schedules.Domain, error) {
	panic("unimplemented")
}

// Create implements schedules.Repository.
func (s *ScheduleRepository) Create(ctx context.Context, schedule schedules.Domain) error {
	newSchedule := Schedules{
		Name:      schedule.Name,
		TrainId:   schedule.TrainId,
		StartDate: schedule.StartDate,
		EndDate:   schedule.EndDate,
	}

	createError := s.Conn.Create(&newSchedule).Error
	if createError != nil {
		return createError
	}

	return nil
}

// Update implements schedules.Repository.
func (s *ScheduleRepository) Update(ctx context.Context, schedules schedules.Domain) (schedules.Domain, error) {
	panic("unimplemented")
}
