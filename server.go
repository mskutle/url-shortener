package main

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port   string
	store  URLStore
	logger slog.Logger
}

func NewServer(port string, store URLStore) Server {
	return Server{
		port:   port,
		store:  store,
		logger: *slog.Default(),
	}
}

func (s *Server) Start() error {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/", s.handleAddUrl)
	e.GET("/", s.handleGetAllUrls)
	e.GET("/:alias", s.handleGetByAlias)

	return e.Start(s.port)
}
