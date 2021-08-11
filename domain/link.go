package domain

import "context"

type Link struct {
	ID   int
	Hash string
	URL  string
}

type LinkUseCase interface {
	Create(ctx context.Context, link Link) error
	Get(ctx context.Context, h string) (*Link, error)
	Delete(ctx context.Context, link Link) error
}

type LinkRepository interface {
	Create(ctx context.Context, link Link) error
	Get(ctx context.Context, h string) (*Link, error)
	Delete(ctx context.Context, link Link) error
}
