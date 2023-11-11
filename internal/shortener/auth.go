package shortener

import (
	"github.com/labstack/echo/v4"
)

func (s *Server) validateUser(username, password string, c echo.Context) (bool, error) {

	if username == s.config.AdminUsername && password == s.config.AdminPassword {
		return true, nil
	}

	return false, nil
}
