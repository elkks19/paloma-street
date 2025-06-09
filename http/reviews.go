package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) registerReviewsRoutes(r *echo.Group) {
	r.GET("/reviews/:usuario", s.handleReviewsByUsuario)
	r.GET("/reviews/:negocio", s.handleReviewsByNegocio)
	r.GET("/reviews/:producto", s.handleReviewsByProducto)
	r.POST("/reviews", s.handleStoreReview)
	r.PUT("/reviews/:id", s.handleUpdateReview)
	r.DELETE("/reviews/:id", s.handleDeleteReview)
}

func (s *Server) handleReviewsByNegocio(c echo.Context) error {
	var limit, offset int
	err := echo.QueryParamsBinder(c).
		MustInt("offset", &offset).
		MustInt("limit", &limit).
		BindError()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var negocioID uint32
	err = echo.PathParamsBinder(c).MustUint32("negocio", &negocioID).BindError()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	reviews, err := s.ReviewsService.ReviewsByNegocio(c.Request().Context(), negocioID, offset, limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]any{
		"reviews": reviews,
	})
}

func (s *Server) handleReviewsByProducto(c echo.Context) error {
	return nil
}

func (s *Server) handleReviewsByUsuario(c echo.Context) error {
	return nil
}

func (s *Server) handleStoreReview(c echo.Context) error {
	return nil
}

func (s *Server) handleUpdateReview(c echo.Context) error {
	return nil
}

func (s *Server) handleDeleteReview(c echo.Context) error {
	return nil
}
