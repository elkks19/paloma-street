package http

import (
	"net/http"
	v "proyecto/pkg/validate"
	"proyecto/services"

	"github.com/labstack/echo/v4"
)

func (s *Server) registerAuthRoutes(r *echo.Group) {
	r.POST("/register", s.handleRegister)
	r.POST("/login", s.handleLogin)
	r.POST("/logout", s.handleLogout)
}

func (s *Server) handleRegister(c echo.Context) error {
	// Parse the request body into a struct
	up := new(services.UsuarioPayload)
	err := c.Bind(up)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request body")
	}

	schema := v.Schema {
		"nombre": v.Rules(v.Required, v.Min(2), v.Max(255)),
		"email": v.Rules(v.Required, v.Email, v.Min(2), v.Max(255)),
		"password": v.Rules(v.Required, v.Min(8), v.Max(255)),
		"passwordConfirmation": v.Rules(v.Required, v.Min(8), v.Max(255)),
	}

	errors, valid := v.Validate(up, schema)
	if !valid {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors)
	}

	// Validate the password and password confirmation
	if up.Password != up.PasswordConfirmation {
		errors.Add("password", "Las contraseñas no coinciden")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors)
	}

	ses, err := s.AuthService.Register(c.Request().Context(), up)
	if err != nil {
		if err == services.ErrEmailExists || err == services.ErrWrongPassword {
			errors.Add("email", "Credenciales invalida")
			return echo.NewHTTPError(http.StatusUnauthorized, errors)
		}
	}

	return c.JSON(http.StatusOK, map[string]any {
		"usuario": up,
		"token": ses.Token,
	})
}

func (s *Server) handleLogin(c echo.Context) error {
	// Parse the request body into a struct
	up := new(services.UsuarioPayload)
	err := c.Bind(up)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, "Invalid request body")
	}

	schema := v.Schema {
		"email": v.Rules(v.Required, v.Email, v.Min(2), v.Max(255)),
		"password": v.Rules(v.Required, v.Min(8), v.Max(255)),
	}

	errors, valid := v.Validate(up, schema)
	if !valid {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors)
	}

	ses, err := s.AuthService.Attempt(c.Request().Context(), up)
	if err != nil {
		if err == services.ErrWrongPassword {
			errors.Add("password", "Contraseña incorrecta")
			return echo.NewHTTPError(http.StatusUnprocessableEntity, errors)
		}

		errors.Add("email", "Credenciales invalida")
		return echo.NewHTTPError(http.StatusUnprocessableEntity, errors)
	}

	return c.JSON(http.StatusOK, map[string]any {
		"usuario": up,
		"token": ses.Token,
	})
}

func (s *Server) handleLogout(c echo.Context) error {
	// token := c.Response().Header().Get("Authorization")
	// if token == "" {
	// 	return echo.NewHTTPError(http.StatusUnauthorized, "No token provided")
	// }
	//
	// err := s.AuthService.Logout(c.Request().Context(), token)

	return nil
}



