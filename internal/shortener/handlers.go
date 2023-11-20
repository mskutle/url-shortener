package shortener

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) handleHealthz(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func (s *Server) handleAddRedirect(c echo.Context) error {

	type request struct {
		Original string `json:"original"`
		Alias    string `json:"alias"`
	}
	var req request

	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, "the request body was malformed")
	}

	redirect := Redirect{
		Original: req.Original,
		Alias:    req.Alias,
	}

	err := s.store.Save(c.Request().Context(), redirect)
	if err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			return c.String(http.StatusConflict, err.Error())
		}
		s.logger.Error("something went wrong when trying to save a new redirect", err)
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusCreated, redirect)
}

func (s *Server) handleRedirect(c echo.Context) error {
	alias := c.Param("alias")

	url, err := s.store.Get(c.Request().Context(), alias)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return c.String(http.StatusNotFound, "")
		}
		return c.String(http.StatusInternalServerError, "")
	}

	return c.Redirect(http.StatusTemporaryRedirect, url.Original)
}
