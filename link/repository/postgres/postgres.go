package postgres

import (
	"context"
	"database/sql"
	"net/url"

	"github.com/goserg/links/domain"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) (domain.LinkRepository, error) {
	err := migrate(db)
	if err != nil {
		return nil, err
	}
	r := Repository{
		db: db,
	}
	return &r, nil
}

func (r Repository) Create(ctx context.Context, link domain.Link) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO link (hash, url) VALUES ($1, $2);", link.Hash, link.URL.String())
	return err
}

func (r Repository) Get(ctx context.Context, h string) (*domain.Link, error) {
	var link domain.Link
	var rawUrl string
	err := r.db.QueryRowContext(ctx, "SELECT * FROM link WHERE hash=$1;", h).Scan(&link.ID, &link.Hash, &rawUrl)
	if err != nil {
		return nil, err
	}
	u, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	link.URL = *u
	return &link, nil
}

func (r Repository) Delete(ctx context.Context, link domain.Link) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM link WHERE hash=$1;", link.Hash)
	return err
}
