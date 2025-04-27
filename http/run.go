package http

import (
	"proyecto/internal/config/env"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo-jwt/v4"
)

// Register the routes, then runs the HTTP server.
func (s *Server) Serve() error {
	debug := env.Get("APP_DEBUG") == "true"
	if debug {
		s.router.Use(middleware.Logger())
	}

	s.router.Use(middleware.Recover())

	// config for jwt middleware
	s.router.Use(echojwt.JWT(echojwt.Config{
		ContextKey: "token",
		SigningKey: []byte(env.Get("SESSION_SECRET")),
		SuccessHandler: s.validTokenHandler,
		ErrorHandler: s.invalidTokenHandler,
	}))

	s.router.HTTPErrorHandler = s.errorHandler

	// STATIC ROUTES
	// s.registerStaticRoutes(s.router)

	// NOT AUTHENTICATED ROUTES
	noAuth := s.router.Group("")
	noAuth.Use()
	// s.registerAuthRoutes(noAuth)

	// AUTHENTICATED ROUTES
	auth := s.router.Group("")
	auth.Use()
	// s.registerAdminRoutes(auth)
	// s.registerProductosRoutes(auth)

	if debug {
		printRoutes(s.router.Routes())
	}

	return s.router.Start(s.port)
}
func printRoutes(routes []*echo.Route) {
	fmt.Print("Routes: ")
	for _, route := range routes {
		fmt.Println(route.Path)
	}
}
