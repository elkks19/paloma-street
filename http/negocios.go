package http

import (
	"database/sql"
	"net/http"
	"proyecto/services"

	"github.com/labstack/echo/v4"
)

func (s *Server) registerNegociosRoutes(r *echo.Group) {
	r.GET("/negocios", s.handleIndexNegocios)
	r.POST("/negocios", s.handleStoreNegocio)
	r.GET("/negocios/:id", s.handleShowNegocio)
	r.PUT("/negocios/:id", s.handleUpdateNegocio)
	r.DELETE("/negocios/:id", s.handleDeleteNegocio)
}

func (s *Server) handleIndexNegocios(c echo.Context) error {
	// Parse the query parameters
	var limit, offset int
	errors := echo.QueryParamsBinder(c).
		Int("offset", &offset).
		Int("limit", &limit).
		BindErrors()

	if len(errors) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, errors)
	}

	negocios, err := s.NegociosService.Negocios(c.Request().Context(), offset, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(negocios) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	// Return the list of negocios
	return c.JSON(http.StatusOK, map[string]any{
		"negocios": negocios,
	})
}

func (s *Server) handleShowNegocio(c echo.Context) error {
	id := new(uint32)
	// Parse the path parameter
	err := echo.PathParamsBinder(c).
		MustUint32("id", id).
		BindError()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// Get the negocio by ID
	negocio, err := s.NegociosService.Negocio(c.Request().Context(), *id)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Negocio not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid request")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"negocio": negocio,
	})
}

func (s *Server) handleStoreNegocio(c echo.Context) error {
	np := new(services.NegocioPayload)

	err := echo.PathParamsBinder(c).
		MustUint32("id", &np.ID).
		BindError()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = echo.QueryParamsBinder(c).
		MustString("nombre", &np.Nombre).
		MustString("descripcion", &np.Descripcion).
		MustJSONUnmarshaler("ubicacion", &np.Ubicacion).
		JSONUnmarshaler("menu", &np.Menu).
		MustString("tipo", &np.Tipo).
		String("contacto", &np.Contacto).
		String("imagen", &np.ImagenB64).
		BindError()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	nDB, err := s.NegociosService.Create(c.Request().Context(), np)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(http.StatusNotFound, "Negocio not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Invalid request")
	}

	// Return the created negocio
	return c.JSON(http.StatusCreated, map[string]any{
		"negocio": nDB,
	})
}

func (s *Server) handleUpdateNegocio(c echo.Context) error {
	return nil
}

func (s *Server) handleDeleteNegocio(c echo.Context) error {
	return nil
}
