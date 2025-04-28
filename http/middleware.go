package http

import (
	"proyecto/internal/config/env"

	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func (s *Server) validTokenHandler(c echo.Context) {

}

func (s *Server) invalidTokenHandler(c echo.Context, err error) error {
	return nil
}


func (s *Server) User(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {


		return nil
	}
}

func (s *Server) Guest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		
		return nil
	}
}

func (s *Server) Authorized(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		echojwt.JWT(echojwt.Config{
			ContextKey: "token",
			SigningKey: []byte(env.Get("SESSION_SECRET")),
			SuccessHandler: s.validTokenHandler,
			ErrorHandler: s.invalidTokenHandler,
		})

		err := next(c)
		
		return err
	}
}
