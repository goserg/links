package postgres

import "database/sql"

func migrate(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS link (
    id SERIAL PRIMARY KEY,
    hash TEXT UNIQUE NOT NULL,
    url TEXT NOT NULL
)`)
	return err
}
