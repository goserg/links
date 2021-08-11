package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("starting app")

	e := echo.New()

	logrus.Fatal(e.Start(":8080"))
}
