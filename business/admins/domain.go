package admins

import (
	"context"
	"kai-backend/business/newsletters"
	"kai-backend/business/paginations"
	"kai-backend/business/superadmins"
	"time"
)

type Domain struct {
	Id              int `gorm:"primaryKey"`
	Username        string
	Password        string
	ChangedPassword string
	Token           string
	SuperadminID    int
	Newsletters     []newsletters.Domain
	// Video_contents []video_contents.Video_content
	Superadmin superadmins.Domain `gorm:"foreignKey:SuperadminID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Usecase interface {
	Register(ctx context.Context, admins Domain) (Domain, error)
	Login(ctx context.Context, admins Domain) (Domain, error)
	UpdatePassword(ctx context.Context, admins Domain) (Domain, error)
	GetAll(ctx context.Context, paginationDomain paginations.Domain) ([]Domain, error)
	CountAll(ctx context.Context) (int, error)
}

type Repository interface {
	Login(ctx context.Context, admins Domain) (Domain, error)
	Register(ctx context.Context, admins Domain) (Domain, error)
	GetByUsername(ctx context.Context, username string) (Domain, error)
	UpdatePassword(ctx context.Context, admins Domain) (Domain, error)
	GetAll(ctx context.Context, paginationDomain paginations.Domain) ([]Domain, error)
	CountAll(ctx context.Context) (int, error)
}
