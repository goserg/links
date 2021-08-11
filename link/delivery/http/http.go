package http

import (
	"net/http"

	"github.com/goserg/links/domain"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	useCase domain.LinkUseCase
}

func Register(e *echo.Echo, useCase domain.LinkUseCase) {
	handler := Handler{useCase: useCase}

	e.POST("/create", handler.Create)
	e.GET("/:h", handler.Get)
}

func (h *Handler) Get(c echo.Context) error {
	ctx := c.Request().Context()

	hash := c.Param("h")

	link, err := h.useCase.Get(ctx, hash)
	if err != nil {
		return c.JSON(http.StatusOK, struct {
			Url string `json:"url"`
			Err string `json:"err"`
		}{
			Url: "",
			Err: err.Error(),
		})
	}
	return c.Redirect(http.StatusSeeOther, "http://"+link.URL)
}

type Link struct {
	Url string `json:"url"`
}

func (h *Handler) Create(c echo.Context) error {
	var link Link
	if err := c.Bind(&link); err != nil {
		// todo log
		return err
	}
	ctx := c.Request().Context()

	hash, err := h.useCase.Create(ctx, link.Url)
	if err != nil {
		// todo log
		return err
	}
	return c.String(http.StatusOK, hash)
}
