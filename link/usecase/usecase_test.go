package usecase

import (
	"context"
	"errors"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/goserg/links/domain"

	"github.com/stretchr/testify/mock"

	"github.com/goserg/links/link/mocks"
)

func TestUseCase_Create(t *testing.T) {
	repoMock := new(mocks.Repository)
	uc := New(repoMock)

	u := url.URL{Path: "link"}
	h := generateHash(u)
	repoMock.On("Get", mock.Anything, h[:1]).Return(&domain.Link{}, errors.New("err"))
	repoMock.On("Create", mock.Anything, domain.Link{Hash: h[:1], URL: u}).Return(nil)
	h, err := uc.Create(context.Background(), u)
	assert.NoError(t, err)
	assert.Equal(t, h[:1], h)
}

func TestUseCase_Get(t *testing.T) {
	repoMock := new(mocks.Repository)
	uc := New(repoMock)

	link := &domain.Link{}
	repoMock.On("Get", mock.Anything, "hash").Return(link, nil)
	l, err := uc.Get(context.Background(), "hash")
	assert.NoError(t, err)
	assert.Equal(t, l, link)
}

func TestUseCase_Delete(t *testing.T) {
	repoMock := new(mocks.Repository)
	uc := New(repoMock)

	link := domain.Link{}
	repoMock.On("Delete", mock.Anything, link).Return(nil)
	err := uc.Delete(context.Background(), link)
	assert.NoError(t, err)
}

func TestUseCase_Delete_Error(t *testing.T) {
	repoMock := new(mocks.Repository)
	uc := New(repoMock)

	link := domain.Link{}
	expectedError := errors.New("error")
	repoMock.On("Delete", mock.Anything, link).Return(expectedError)
	err := uc.Delete(context.Background(), link)
	if assert.Error(t, err) {
		assert.Equal(t, err, expectedError)
	}
}
