package http

import (
	"context"
	"log"
	"proyecto/internal/config/env"
	"proyecto/internal/db"
	"proyecto/services"
	"proyecto/services/filesystem"
	"proyecto/services/postgres"

	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

type Services struct {
	AuthService       services.AuthService
	CategoriasService services.CategoriasService
	FavoritosService  services.FavoritosService
	NegociosService   services.NegociosService
	ProductosService  services.ProductosService
	ReviewsService    services.ReviewsService
	UsuariosService   services.UsuariosService
}

type Server struct {
	router *echo.Echo
	logger *log.Logger
	DB     *bun.DB
	port   string
	ctx    context.Context
	*Services
}

func NewServer() *Server {
	err := env.Init()
	if err != nil {
		panic(err)
	}

	// Here comes the server initialization
	s := &Server{
		router:   echo.New(),
		DB:       db.Init(),
		port:     env.Get("APP_PORT"),
		Services: &Services{},
	}

	is := filesystem.NewImagenesService()
	s.AuthService = postgres.NewAuthService(s.DB)
	s.CategoriasService = postgres.NewCategoriasService(s.DB)
	s.FavoritosService = postgres.NewFavoritosService(s.DB)
	s.NegociosService = postgres.NewNegociosService(s.DB, is)
	s.ProductosService = postgres.NewProductosService(s.DB, is)
	s.ReviewsService = postgres.NewReviewsService(s.DB)
	s.UsuariosService = postgres.NewUsuariosService(s.DB)

	return s
}
