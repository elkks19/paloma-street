package http

import (
	"proyecto/pkg/validate"
	"net/http"

	"github.com/labstack/echo/v4"
)


func (s *Server) errorHandler(err error, c echo.Context) {
 	if c.Response().Committed { 
 		return 
 	}

	code := http.StatusInternalServerError
	he, ok := err.(*echo.HTTPError)
	if ok {
		code = he.Code
	}

	switch code {
		case http.StatusNotFound:
			return

		case http.StatusInternalServerError:
			return

		case http.StatusBadRequest:
			if c.Request().Header.Get("HX-Request") != "" {
				c.NoContent(http.StatusBadRequest)
				return
			}

		case http.StatusUnprocessableEntity:
			switch e := he.Message.(type) {
				case validate.Errors:
					c.JSON(code, e)
					return

				case string:
					c.JSON(code, validate.Errors{
						"_error": []string{e},
					})
			}

		case http.StatusUnauthorized:

		case http.StatusForbidden:

	}

	c.NoContent(code)
}

