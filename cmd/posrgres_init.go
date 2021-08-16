//+build postgres

package main

import (
	"database/sql"

	"github.com/goserg/links/domain"

	postgresLinkRepo "github.com/goserg/links/link/repository/postgres"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func databaseInit() (domain.LinkRepository, func() error, error) {
	logrus.Println("starting app with postgres storage")

	db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	repo, err := postgresLinkRepo.New(db)
	return repo, db.Close, err

}
