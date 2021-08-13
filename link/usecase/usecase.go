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

/*
Create
Creating hash from url, then trying its first letter as current hash.
If url stored with current hash is equal to request url - return current hash, otherwise try first two letters and so on.
When an unused hash is - save it and return the current hash.

My goal was to create algorithm that was relatively fast and generated as short a hash as possible.
The drawback is the multiple storage calls per url save.

Possible improvement if storage calls will slow things down:
  make one storage request with all the hashes, and then process the result here.
*/
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
		if link.URL.String() == url.String() {
			return link.Hash, nil
		}
	}
	return "", errors.New("internal server error")
}

/*
sha256 chosen here since it's twice as long as the md5 hash, although it's twice as slow.
See benchmark test in /bench folder.
*/
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
