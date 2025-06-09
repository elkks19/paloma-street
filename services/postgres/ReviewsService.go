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
	reviews := new([]*services.Review)
	err := s.db.NewSelect().
		Model(reviews).
		Where("usuario_id = ?", usuarioId).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Scan(ctx)

	return *reviews, err
}

func (s *ReviewsService) ReviewsByNegocio(ctx context.Context, negocioId uint32, offset int, limit int) ([]*services.Review, error) {
	reviews := new([]*services.Review)
	err := s.db.NewSelect().
		Model(reviews).
		JoinOn("productos.id = reviews.producto_id").
		JoinOn("productos.negocio_id = ?", negocioId).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Scan(ctx)
		
	return *reviews, err
}

func (s *ReviewsService) ReviewsByProducto(ctx context.Context, productoId uint32, offset int, limit int) ([]*services.Review, error) {
	reviews := new([]*services.Review)
	err := s.db.NewSelect().
		Model(reviews).
		Where("producto_id = ?", productoId).
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Scan(ctx)

	return *reviews, err
}

func (s *ReviewsService) Create(ctx context.Context, review *services.ReviewPayload) (*services.Review, error) {
	dbR := &services.Review{
		UsuarioID:    review.UsuarioID,
		ProductoID:   review.ProductoID,
		Contenido:    review.Contenido,
		Calificacion: review.Calificacion,
	}

	err := s.db.NewInsert().
		Model(dbR).
		Returning("*").
		Scan(ctx)


	return dbR, err
}

func (s *ReviewsService) Update(ctx context.Context, review *services.ReviewPayload) (*services.Review, error) {
	dbR := &services.Review{
		ID:           review.ID,
		UsuarioID:    review.UsuarioID,
		ProductoID:   review.ProductoID,
		Contenido:    review.Contenido,
		Calificacion: review.Calificacion,
	}

	err := s.db.NewUpdate().
		Model(dbR).
		Where("id = ?", review.ID).
		Returning("*").
		Scan(ctx)

	return dbR, err
}

func (s *ReviewsService) Delete(ctx context.Context, id uint32) error {
	_, err := s.db.NewDelete().
		Model((*services.Review)(nil)).
		Where("id = ?", id).
		Exec(ctx)

	return err
}
