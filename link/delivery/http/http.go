package http

import (
	"github.com/goserg/links/domain"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	useCase domain.LinkUseCase
}

func Register(e *echo.Echo, useCase domain.LinkUseCase) {
	handler := Handler{useCase: useCase}

	e.POST("/create", handler.Get)
	e.GET("/:h", handler.Create)
}

func (h *Handler) Get(c echo.Context) error {
	return nil
}

func (h *Handler) Create(c echo.Context) error {
	return nil
}
