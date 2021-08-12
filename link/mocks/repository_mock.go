package mocks

import (
	"context"

	"github.com/goserg/links/domain"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (r *Repository) Create(ctx context.Context, link domain.Link) error {
	ret := r.Called(ctx, link)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Link) error); ok {
		r0 = rf(ctx, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (r *Repository) Get(ctx context.Context, h string) (*domain.Link, error) {
	ret := r.Called(ctx, h)

	var r0 *domain.Link
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.Link); ok {
		r0 = rf(ctx, h)
	} else {
		r0 = ret.Get(0).(*domain.Link)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, h)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (r *Repository) Delete(ctx context.Context, link domain.Link) error {
	ret := r.Called(ctx, link)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Link) error); ok {
		r0 = rf(ctx, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
