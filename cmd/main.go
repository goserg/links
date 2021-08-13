//+build inmemory

package main

import (
	linkHttpDelivery "github.com/goserg/links/link/delivery/http"
	inmemoryLinkRepo "github.com/goserg/links/link/repository/inmemory"
	linkUC "github.com/goserg/links/link/usecase"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("starting app with inmemory storage")

	e := echo.New()

	linkRepo := inmemoryLinkRepo.New()
	linkUseCase := linkUC.New(linkRepo)

	linkHttpDelivery.Register(e, linkUseCase)

	logrus.Fatal(e.Start(":8080"))
}
