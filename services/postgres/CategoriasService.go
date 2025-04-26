package postgres

import (
	"context"
	"proyecto/services"

	"github.com/uptrace/bun"
)

type CategoriasService struct {
	db *bun.DB
}

func NewCategoriasService(db *bun.DB) *CategoriasService {
	return &CategoriasService{
		db: db,
	}
}

func (s *CategoriasService) Categorias(ctx context.Context, offset int, limit int) ([]*services.Categoria, error) {
	return nil, nil
}

func (s *CategoriasService) Create(ctx context.Context, cat *services.CategoriaPayload) (*services.Categoria, error) {
	return nil, nil
}
