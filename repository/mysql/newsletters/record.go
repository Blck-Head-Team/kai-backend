package newsletters

import (
	"kai-backend/business/newsletters"
	"time"

	"gorm.io/gorm"
)

type Newsletter struct {
	Id          int `gorm:"primaryKey"`
	Title       string
	Description string
	Content     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (n *Newsletter) BeforeCreate(tx *gorm.DB) error {
	n.CreatedAt = time.Now()
	n.UpdatedAt = time.Now()
	return nil
}

func (n *Newsletter) ToDomain() newsletters.Domain {
	return newsletters.Domain{
		Id:          n.Id,
		Title:       n.Title,
		Description: n.Description,
		Content:     n.Content,
		CreatedAt:   n.CreatedAt,
		UpdatedAt:   n.UpdatedAt,
	}
}

func FromDomain(domain newsletters.Domain) Newsletter {
	return Newsletter{
		Id:          domain.Id,
		Title:       domain.Title,
		Description: domain.Description,
		Content:     domain.Content,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}

func ToListDomain(data []Newsletter) []newsletters.Domain {
	var listDomain []newsletters.Domain
	for _, item := range data {
		listDomain = append(listDomain, item.ToDomain())
	}
	return listDomain
}
