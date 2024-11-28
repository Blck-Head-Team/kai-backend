// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	gyms "kai-backend/business/gyms"
	context "context"

	mock "github.com/stretchr/testify/mock"

	paginations "kai-backend/business/paginations"
)

// DomainRepository is an autogenerated mock type for the DomainRepository type
type DomainRepository struct {
	mock.Mock
}

// CountAll provides a mock function with given fields: ctx
func (_m *DomainRepository) CountAll(ctx context.Context) (int, error) {
	ret := _m.Called(ctx)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context) int); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, domain
func (_m *DomainRepository) Create(ctx context.Context, domain gyms.Domain) (gyms.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 gyms.Domain
	if rf, ok := ret.Get(0).(func(context.Context, gyms.Domain) gyms.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(gyms.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, gyms.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *DomainRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx, paginationDomain
func (_m *DomainRepository) GetAll(ctx context.Context, paginationDomain paginations.Domain) ([]gyms.Domain, error) {
	ret := _m.Called(ctx, paginationDomain)

	var r0 []gyms.Domain
	if rf, ok := ret.Get(0).(func(context.Context, paginations.Domain) []gyms.Domain); ok {
		r0 = rf(ctx, paginationDomain)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]gyms.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, paginations.Domain) error); ok {
		r1 = rf(ctx, paginationDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: ctx, id
func (_m *DomainRepository) GetById(ctx context.Context, id string) (gyms.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 gyms.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) gyms.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(gyms.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, domain
func (_m *DomainRepository) Update(ctx context.Context, id string, domain gyms.Domain) (gyms.Domain, error) {
	ret := _m.Called(ctx, id, domain)

	var r0 gyms.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string, gyms.Domain) gyms.Domain); ok {
		r0 = rf(ctx, id, domain)
	} else {
		r0 = ret.Get(0).(gyms.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, gyms.Domain) error); ok {
		r1 = rf(ctx, id, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
