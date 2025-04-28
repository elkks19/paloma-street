package postgres

import (
	"context"
	"proyecto/services"

	"github.com/uptrace/bun"
)

type UsuariosService struct {
	db *bun.DB
}

func NewUsuariosService(db *bun.DB) *UsuariosService {
	db.RegisterModel((*services.NegocioUsuarios)(nil))
	return &UsuariosService{
		db: db,
	}
}

func (s *UsuariosService) Usuario(ctx context.Context, id uint32) (*services.Usuario, error) {
	return nil, nil
}

func (s *UsuariosService) Usuarios(ctx context.Context, offset int, limit int) ([]*services.Usuario, error) {
	return nil, nil
}

func (s *UsuariosService) Create(ctx context.Context, usuario *services.UsuarioPayload) (*services.Usuario, error) {
	return nil, nil
}

func (s *UsuariosService) Update(ctx context.Context, usuario *services.UsuarioPayload) (*services.Usuario, error) {
	return nil, nil
}

func (s *UsuariosService) Delete(ctx context.Context, id uint32) error {
	return nil
}

func (s *UsuariosService) ForceDelete(ctx context.Context, id uint32) error {
	return nil
}

