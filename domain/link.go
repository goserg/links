package domain

import (
	"context"
	"net/url"
)

type Link struct {
	ID   int
	Hash string
	URL  url.URL
}

type LinkUseCase interface {
	Create(ctx context.Context, link url.URL) (string, error)
	Get(ctx context.Context, h string) (*Link, error)
	Delete(ctx context.Context, link Link) error
}

type LinkRepository interface {
	Create(ctx context.Context, link Link) error
	Get(ctx context.Context, h string) (*Link, error)
	Delete(ctx context.Context, link Link) error
}
