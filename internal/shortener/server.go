package shortener

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	port   string
	store  Store
	logger *slog.Logger
	config ServerConfig
}

func NewServer(port string, store Store) Server {
	logger := slog.Default()
	return Server{
		port:   port,
		store:  store,
		logger: logger,
		config: NewServerConfig(logger),
	}
}

func (s *Server) Start() error {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/r/:alias", s.redirect)
	app.POST("/redirects", s.handleAddRedirect, middleware.BasicAuth(s.validateUser))

	return app.Start(s.port)
}