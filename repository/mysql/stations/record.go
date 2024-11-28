package stations

import (
	"time"

	"gorm.io/gorm"
)

type Station struct {
	Id        int    `gorm:"primaryKey"`
	Name      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Station) BeforeCreate(tx *gorm.DB) error {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	return nil
}
