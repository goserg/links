package main

import (
	inmemoryLinkRepo "github.com/goserg/links/link/repository/inmemory"
	linkUC "github.com/goserg/links/link/usecase"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("starting app")

	e := echo.New()

	linkRepo := inmemoryLinkRepo.New()
	linkUC.New(linkRepo)

	logrus.Fatal(e.Start(":8080"))
}
