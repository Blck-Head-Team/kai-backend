package admins

import (
	"context"
	"kai-backend/app/middlewares"
	"kai-backend/business/paginations"
	"kai-backend/exceptions"
	"kai-backend/helpers"
	"time"
)

type OperationalAdminsUsecase struct {
	Repo           Repository
	ContextTimeout time.Duration
	JWTAuth        *middlewares.ConfigJWT
}

func NewOperationaldminsUsecase(repo Repository, timeout time.Duration, jwt *middlewares.ConfigJWT) Usecase {
	return &OperationalAdminsUsecase{
		Repo:           repo,
		ContextTimeout: timeout,
		JWTAuth:        jwt,
	}
}

func (oa *OperationalAdminsUsecase) Register(ctx context.Context, opadmin Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, oa.ContextTimeout)
	defer cancel()
	if opadmin.Username == "" || opadmin.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}

	var err error
	opadmin.Password, err = helpers.Hash(opadmin.Password)
	if err != nil {
		return Domain{}, err
	}
	res, err := oa.Repo.Register(ctx, opadmin)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (oa *OperationalAdminsUsecase) Login(ctx context.Context, opadmin Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, oa.ContextTimeout)
	defer cancel()
	if opadmin.Username == "" || opadmin.Password == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	res, err := oa.Repo.GetByUsername(ctx, opadmin.Username)
	if err != nil {
		return Domain{}, err
	}
	if !helpers.ValidateHash(opadmin.Password, res.Password) {
		return Domain{}, exceptions.ErrValidationFailed
	}
	res.Token, _ = oa.JWTAuth.GenerateToken(res.Id, res.Username, false, false, true)
	return res, nil
}

func (oa *OperationalAdminsUsecase) UpdatePassword(ctx context.Context, opadmin Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, oa.ContextTimeout)
	defer cancel()
	if opadmin.Username == "" || opadmin.Password == "" || opadmin.ChangedPassword == "" {
		return Domain{}, exceptions.ErrInvalidCredentials
	}
	res, err := oa.Repo.Login(ctx, opadmin)
	if err != nil {
		return Domain{}, err
	}
	if !helpers.ValidateHash(opadmin.Password, res.Password) {
		return Domain{}, exceptions.ErrValidationFailed
	}
	opadmin.Password, err = helpers.Hash(opadmin.ChangedPassword)
	if err != nil {
		return Domain{}, err
	}
	res, err = oa.Repo.UpdatePassword(ctx, opadmin)
	if err != nil {
		return Domain{}, err
	}
	return res, nil
}

func (oa *OperationalAdminsUsecase) CountAll(ctx context.Context) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, oa.ContextTimeout)
	defer cancel()

	return oa.Repo.CountAll(ctx)
}
func (oa *OperationalAdminsUsecase) GetAll(ctx context.Context, pagination paginations.Domain) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, oa.ContextTimeout)
	defer cancel()

	return oa.Repo.GetAll(ctx, pagination)
}
