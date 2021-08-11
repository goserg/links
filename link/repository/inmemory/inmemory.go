package inmemory

import (
	"context"

	"github.com/goserg/links/domain"
)

type Repository struct {
	storage map[string]string
}

func New() domain.LinkRepository {
	r := Repository{
		storage: make(map[string]string),
	}
	return &r
}

func (r Repository) Create(ctx context.Context, link domain.Link) error {
	panic("implement me")
}

func (r Repository) Get(ctx context.Context, h string) (*domain.Link, error) {
	panic("implement me")
}

func (r Repository) Delete(ctx context.Context, link domain.Link) error {
	panic("implement me")
}
