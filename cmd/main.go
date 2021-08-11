package main

import (
	inmemoryLinkRepo "github.com/goserg/links/link/repository/inmemory"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("starting app")

	e := echo.New()

	inmemoryLinkRepo.New()

	logrus.Fatal(e.Start(":8080"))
}
