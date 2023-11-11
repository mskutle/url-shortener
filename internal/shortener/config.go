package shortener

import (
	"log/slog"
	"os"
)

type ServerConfig struct {
	AdminUsername string
	AdminPassword string
}

func NewServerConfig(logger *slog.Logger) ServerConfig {
	adminUser := os.Getenv("ADMIN_USER")
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	return ServerConfig{
		AdminUsername: adminUser,
		AdminPassword: adminPassword,
	}
}
