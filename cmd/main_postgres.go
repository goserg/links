//+build postgres

package main

import (
	"database/sql"

	linkHttpDelivery "github.com/goserg/links/link/delivery/http"
	postgresLinkRepo "github.com/goserg/links/link/repository/postgres"
	linkUC "github.com/goserg/links/link/usecase"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("starting app with postgres storage")

	e := echo.New()
	e.HideBanner = true

	db, err := sql.Open("postgres", "host=postgres port=5432 user=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	linkRepo, err := postgresLinkRepo.New(db)
	if err != nil {
		panic(err)
	}
	linkUseCase := linkUC.New(linkRepo)

	linkHttpDelivery.Register(e, linkUseCase)

	logrus.Fatal(e.Start(":8080"))
}
