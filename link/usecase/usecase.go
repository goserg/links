package usecase

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/url"

	"github.com/goserg/links/domain"
)

const salt = "just for fun"

type UseCase struct {
	db domain.LinkRepository
}

func New(db domain.LinkRepository) domain.LinkUseCase {
	us := UseCase{db: db}
	return &us
}

func (u UseCase) Create(ctx context.Context, url url.URL) (string, error) {
	h := generateHash(url)
	for i := 1; i < len(h); i++ {
		currentHash := h[:i]
		link, err := u.Get(ctx, currentHash)
		if err != nil {
			link := domain.Link{
				Hash: currentHash,
				URL:  url,
			}
			err := u.db.Create(ctx, link)
			if err != nil {
				return "", err
			}
			return currentHash, nil
		}
		if link.URL == url {
			return link.Hash, nil
		}
	}
	return "", errors.New("internal server error")
}

func generateHash(s url.URL) string {
	hasher := sha256.New()
	hasher.Write([]byte(s.String() + salt))
	h := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return h
}

func (u UseCase) Get(ctx context.Context, h string) (*domain.Link, error) {
	return u.db.Get(ctx, h)
}

func (u UseCase) Delete(ctx context.Context, link domain.Link) error {
	return u.db.Delete(ctx, link)
}
