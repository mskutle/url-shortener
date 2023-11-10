package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) handleAddUrl(c echo.Context) error {
	var url URL
	if err := c.Bind(&url); err != nil {
		return c.String(http.StatusBadRequest, "the request body was malformed")
	}

	err := s.store.Save(url)
	if err != nil {
		if errors.Is(err, ErrAlreadyExists) {
			return c.String(http.StatusConflict, err.Error())
		}
		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusCreated, url)
}

func (s *Server) handleGetByAlias(c echo.Context) error {
	alias := c.Param("alias")
	if alias == "" {
		return c.String(http.StatusBadRequest, "you must provide an alias")
	}

	url, err := s.store.Get(alias)

	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return c.String(http.StatusNotFound, err.Error())
		}

		return c.String(http.StatusInternalServerError, "something went wrong")
	}

	return c.JSON(http.StatusOK, url)
}

func (s *Server) handleGetAllUrls(c echo.Context) error {
	urls, err := s.store.GetAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, urls)
}
