package shortener

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	store  Store
	logger *slog.Logger
	config ServerConfig
}

func NewServer(store Store) Server {
	logger := slog.Default()
	return Server{
		store:  store,
		logger: logger,
		config: NewServerConfig(logger),
	}
}

func (s *Server) Start() error {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/healthz", s.healthCheck)
	app.GET("/r/:alias", s.redirect)
	app.POST("/redirects", s.handleAddRedirect, middleware.BasicAuth(s.validateUser))

	return app.Start(s.config.port)
}
