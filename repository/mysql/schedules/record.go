package schedules

import (
	"kai-backend/repository/mysql/trains"
	"time"

	"gorm.io/gorm"
)

type Schedules struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	TrainId   uint         `gorm:"not null"`
	Train     trains.Train `gorm:"foreignkey:TrainId"`
	StartDate time.Time
	EndDate   time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Schedules) BeforeCreate(tx *gorm.DB) error {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	return nil
}
