package main

import (
	linkHttpDelivery "github.com/goserg/links/link/delivery/http"
	linkUC "github.com/goserg/links/link/usecase"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	linkRepo, closeFunc, err := databaseInit()
	if err != nil {
		panic(err)
	}
	defer closeFunc()

	linkUseCase := linkUC.New(linkRepo)

	e := echo.New()
	e.HideBanner = true
	linkHttpDelivery.Register(e, linkUseCase)
	logrus.Fatal(e.Start(":8080"))
}
