package http

import (
	"proyecto/internal/config/env"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Register the routes, then runs the HTTP server.
func (s *Server) Serve() error {
	debug := env.Get("APP_DEBUG") == "true"
	if debug {
		s.router.Use(middleware.Logger())
	}

	s.router.Use(middleware.Recover())

	s.router.HTTPErrorHandler = s.errorHandler

	// STATIC ROUTES
	// s.registerStaticRoutes(s.router)

	// NOT AUTHENTICATED ROUTES
	noAuth := s.router.Group("")
	s.registerAuthRoutes(noAuth)

	// AUTHENTICATED ROUTES
	auth := s.router.Group("")
	s.registerNegociosRoutes(auth)
	s.registerProductosRoutes(auth)
	s.registerFavoritosRoutes(auth)


	printRoutes(s.router.Routes())
	return s.router.Start(s.port)
}
func printRoutes(routes []*echo.Route) {
	fmt.Print("Routes: ")
	for _, route := range routes {
		fmt.Println(route.Path)
	}
}
