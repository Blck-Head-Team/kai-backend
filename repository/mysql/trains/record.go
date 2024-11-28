package trains

import (
	"kai-backend/repository/mysql/stations"
	"time"

	"gorm.io/gorm"
)

type Train struct {
	Id                   int `gorm:"primaryKey"`
	Name                 string
	DestinationStationId uint             `gorm:"not null"`
	OriginStationId      uint             `gorm:"not null"`
	DestinationStation   stations.Station `gorm:"foreignkey:DestinationStationId"`
	OriginStation        stations.Station `gorm:"foreignkey:OriginStationId"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

func (t *Train) BeforeCreate(tx *gorm.DB) error {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	return nil
}
