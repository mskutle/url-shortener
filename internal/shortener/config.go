package shortener

import (
	"fmt"
	"log/slog"
	"os"
)

type ServerConfig struct {
	adminUsername string
	adminPassword string
	port          string
}

func NewServerConfig(logger *slog.Logger) ServerConfig {
	adminUser := os.Getenv("ADMIN_USER")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	port := os.Getenv("PORT")

	return ServerConfig{
		adminUsername: adminUser,
		adminPassword: adminPassword,
		port:          fmt.Sprintf(":%s", port),
	}
}
