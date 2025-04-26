package postgres

import (
	"context"
	"proyecto/services"

	"github.com/uptrace/bun"
)

type NegociosService struct {
	db *bun.DB
}

func NewNegociosService(db *bun.DB) *NegociosService {
	return &NegociosService{
		db: db,
	}
}

func (s *NegociosService) Negocio(ctx context.Context, id uint32) (*services.Negocio, error) {
	return nil, nil
}

func (s *NegociosService) Negocios(ctx context.Context, offset int, limit int) ([]*services.Negocio, error) {
	return nil, nil
}

func (s *NegociosService) Create(ctx context.Context, negocio *services.NegocioPayload) (*services.Negocio, error) {
	return nil, nil
}

func (s *NegociosService) Update(ctx context.Context, negocio *services.NegocioPayload) (*services.Negocio, error) {
	return nil, nil
}

func (s *NegociosService) Delete(ctx context.Context, negocio *services.NegocioPayload) error {
	return nil
}
