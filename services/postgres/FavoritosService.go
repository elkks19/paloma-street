package postgres

import (
	"context"
	"proyecto/services"

	"github.com/uptrace/bun"
)

type FavoritosService struct {
	db *bun.DB
}

func NewFavoritosService(db *bun.DB) *FavoritosService {
	return &FavoritosService{
		db: db,
	}
}

func (s *FavoritosService) Favoritos(ctx context.Context, u *services.Usuario, offset int, limit int) ([]*services.Favorito, error) {
	return nil, nil
}

func (s *FavoritosService) RegisterFavorito(ctx context.Context, favorito *services.FavoritoPayload) error {
	return nil
}

func (s *FavoritosService) UnregisterFavorito(ctx context.Context, favorito *services.FavoritoPayload) error {
	return nil
}


