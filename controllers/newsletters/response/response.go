package request

import (
	"kai-backend/business/newsletters"
	"time"
)

type NewsResponse struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Content     string     `json:"content"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

func FromDomain(domain newsletters.Domain) NewsResponse {
	return NewsResponse{
		Id:          domain.Id,
		Title:       domain.Title,
		Description: domain.Description,
		Content:     domain.Content,
		CreatedAt:   &domain.CreatedAt,
		UpdatedAt:   &domain.UpdatedAt,
	}
}
