package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) registerProductosRoutes(r *echo.Group) {
	r.GET("/productos", s.handleIndexProductos)
	r.POST("/productos", s.handleStoreProducto)
	r.GET("/productos/:id", s.handleShowProducto)
	r.PUT("/productos/:id", s.handleUpdateProducto)
	r.DELETE("/productos/:id", s.handleDeleteProducto)
}

func (s *Server) handleIndexProductos(c echo.Context) error {
	// Parse the query parameters
	var limit, offset int

	errors := echo.QueryParamsBinder(c).
		Int("offset", &offset).
		Int("limit", &limit).
		BindErrors()

	if len(errors) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, errors)
	}

	return nil
}

func (s *Server) handleShowProducto(c echo.Context) error {
	return nil
}

func (s *Server) handleStoreProducto(c echo.Context) error {
	return nil
}

func (s *Server) handleUpdateProducto(c echo.Context) error {
	return nil
}

func (s *Server) handleDeleteProducto(c echo.Context) error {
	return nil
}
