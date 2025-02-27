package classes

import (
	context "context"
	"kai-backend/business/gyms"
	"kai-backend/business/paginations"
	"kai-backend/exceptions"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type Usecase struct {
	Repo           DomainRepository
	GymRepo        gyms.DomainRepository
	ContextTimeout time.Duration
}

func NewUsecase(repo DomainRepository, gymRepo gyms.DomainRepository, timeout time.Duration) *Usecase {
	return &Usecase{
		Repo:           repo,
		GymRepo:        gymRepo,
		ContextTimeout: timeout,
	}
}

func (u *Usecase) GetAll(ctx context.Context, pagination paginations.Domain, domain Domain) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.GetAll(ctx, pagination, domain)
}

func (u *Usecase) CountAll(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	return u.Repo.CountAll(ctx)
}

func (u *Usecase) GetById(ctx context.Context, id string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}
	domain, _ := u.Repo.GetById(ctx, id)
	return domain, nil
}

func (u *Usecase) Create(ctx context.Context, domain Domain, gymId string) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	// check for gymId
	gym, gymErr := u.GymRepo.GetById(ctx, gymId)
	if (gymErr != nil) || (gym.Id == 0) {
		return Domain{}, exceptions.ErrGymNotFound
	}

	log.Printf("%+v", domain)

	validate := validator.New()
	err := validate.Struct(domain)
	if err != nil {
		return Domain{}, exceptions.ErrValidationFailed
	}

	return u.Repo.Create(ctx, domain, gymId)
}

func (u *Usecase) Update(ctx context.Context, id string, domain Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if id == "" {
		return Domain{}, exceptions.ErrEmptyInput
	}

	validate := validator.New()
	err := validate.Struct(domain)
	if err != nil {
		return Domain{}, exceptions.ErrValidationFailed
	}

	return u.Repo.Update(ctx, id, domain)
}

func (u *Usecase) Delete(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	if id == "" {
		return exceptions.ErrEmptyInput
	}

	return u.Repo.Delete(ctx, id)
}
