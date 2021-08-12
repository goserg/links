package mocks

import (
	"context"
	"net/url"

	"github.com/goserg/links/domain"
	"github.com/stretchr/testify/mock"
)

type UseCase struct {
	mock.Mock
}

func (m *UseCase) Create(ctx context.Context, u url.URL) (string, error) {
	ret := m.Called(ctx, u)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, url.URL) string); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, url.URL) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *UseCase) Get(ctx context.Context, h string) (*domain.Link, error) {
	ret := m.Called(ctx, h)

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

func (m *UseCase) Delete(ctx context.Context, link domain.Link) error {
	ret := m.Called(ctx, link)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Link) error); ok {
		r0 = rf(ctx, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
