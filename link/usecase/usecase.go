package usecase

import (
	"context"

	"github.com/goserg/links/domain"
)

type UseCase struct {
	db domain.LinkRepository
}

func New(db domain.LinkRepository) domain.LinkUseCase {
	us := UseCase{db: db}
	return &us
}

func (u UseCase) Create(ctx context.Context, link domain.Link) error {
	panic("implement me")
}

func (u UseCase) Get(ctx context.Context, h string) (*domain.Link, error) {
	panic("implement me")
}

func (u UseCase) Delete(ctx context.Context, link domain.Link) error {
	panic("implement me")
}
