package http

import "github.com/labstack/echo/v4"

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
		err := next(c)
		
		return err
	}
}
