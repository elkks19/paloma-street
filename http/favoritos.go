package http

import "github.com/labstack/echo/v4"

func (s *Server) registerFavoritosRoutes(r *echo.Group) {
	r.GET("/favoritos", s.handleIndexFavoritos)
	r.POST("/favoritos/:id", s.handleToggleFavorito)
}

func (s *Server) handleIndexFavoritos(c echo.Context) error {
	return nil
}

func (s *Server) handleToggleFavorito(c echo.Context) error {
	return nil
}
