package postgres

import (
	"context"
	"proyecto/services"

	"github.com/uptrace/bun"
)

type ProductosService struct {
	db              *bun.DB
	ImagenesService services.ImagenesService
}

func NewProductosService(db *bun.DB, is services.ImagenesService) *ProductosService {
	db.RegisterModel((*services.ProductosCategorias)(nil))
	return &ProductosService{
		db:              db,
		ImagenesService: is,
	}
}

func (s *ProductosService) Producto(ctx context.Context, id uint32) (*services.Producto, error) {
	return nil, nil
}

func (s *ProductosService) Productos(ctx context.Context, offset int, limit int) ([]*services.Producto, error) {
	return nil, nil
}

func (s *ProductosService) ProductosByNegocio(ctx context.Context, negocioId uint32, offset int, limit int) ([]*services.Producto, error) {
	return nil, nil
}

func (s *ProductosService) Create(ctx context.Context, producto *services.ProductoPayload) (*services.Producto, error) {
	return nil, nil
}

func (s *ProductosService) Update(ctx context.Context, producto *services.ProductoPayload) (*services.Producto, error) {
	return nil, nil
}

func (s *ProductosService) Delete(ctx context.Context, producto *services.ProductoPayload) error {
	return nil
}
