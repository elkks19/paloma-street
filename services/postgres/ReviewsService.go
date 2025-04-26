package postgres

import (
	"context"
	"proyecto/services"

	"github.com/uptrace/bun"
)

type ReviewsService struct {
	db *bun.DB
}

func NewReviewsService(db *bun.DB) *ReviewsService {
	return &ReviewsService{
		db: db,
	}
}

func (s *ReviewsService) ReviewsByUsuario(ctx context.Context, usuarioId uint32, offset int, limit int) ([]*services.Review, error) {
	return nil, nil
}

func (s *ReviewsService) ReviewsByNegocio(ctx context.Context, negocioId uint32, offset int, limit int) ([]*services.Review, error) {
	return nil, nil
}

func (s *ReviewsService) ReviewsByProducto(ctx context.Context, productoId uint32, offset int, limit int) ([]*services.Review, error) {
	return nil, nil
}

func (s *ReviewsService) Create(ctx context.Context, review *services.ReviewPayload) (*services.Review, error) {
	return nil, nil
}

func (s *ReviewsService) Update(ctx context.Context, review *services.ReviewPayload) (*services.Review, error) {
	return nil, nil
}

func (s *ReviewsService) Delete(ctx context.Context, review *services.ReviewPayload) (*services.Review, error) {
	return nil, nil
}
