package inmemory

import (
	"context"
	"errors"

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

func (r Repository) Create(_ context.Context, link domain.Link) error {
	_, exist := r.storage[link.Hash]
	if exist {
		return errors.New("link hash is not unique")
	}
	r.storage[link.Hash] = link.URL
	return nil
}

func (r Repository) Get(_ context.Context, h string) (*domain.Link, error) {
	url, exist := r.storage[h]
	if !exist {
		return nil, errors.New("can not find link with hash = " + h)
	}
	return &domain.Link{
		Hash: h,
		URL:  url,
	}, nil
}

func (r Repository) Delete(_ context.Context, link domain.Link) error {
	_, exist := r.storage[link.Hash]
	if exist {
		return errors.New("can not find link with hash = " + link.Hash)
	}
	delete(r.storage, link.Hash)
	return nil
}
