package shortener

import (
	"github.com/labstack/echo/v4"
)

func (s *Server) validateUser(username, password string, c echo.Context) (bool, error) {

	if username == s.config.adminUsername && password == s.config.adminPassword {
		return true, nil
	}

	return false, nil
}
