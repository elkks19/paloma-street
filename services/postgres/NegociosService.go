package postgres

import (
	"context"
	"proyecto/services"

	"github.com/uptrace/bun"
)

type NegociosService struct {
	db              *bun.DB
	ImagenesService services.ImagenesService
}

func NewNegociosService(db *bun.DB, is services.ImagenesService) *NegociosService {
	db.RegisterModel(
		(*services.NegocioUsuarios)(nil),
		(*services.ProductosCategorias)(nil),
	)

	return &NegociosService{
		db: db,
		ImagenesService: is,
	}
}

func (s *NegociosService) Negocio(ctx context.Context, id uint32) (*services.Negocio, error) {
	negocio := new(services.Negocio)
	err := s.db.NewSelect().
		Model(negocio).
		Where("id = ?", id).
		Relation("Productos").
		Relation("Usuarios").
		Scan(ctx)

	return negocio, err
}

func (s *NegociosService) Negocios(ctx context.Context, offset int, limit int) ([]*services.Negocio, error) {
	negocios := new([]*services.Negocio)
	err := s.db.NewSelect().
		Model(negocios).
		Offset(offset).
		Limit(limit).
		Scan(ctx)

	return *negocios, err
}

func (s *NegociosService) Create(ctx context.Context, negocio *services.NegocioPayload) (*services.Negocio, error) {
	nDB := &services.Negocio{
		Nombre:      negocio.Nombre,
		Descripcion: negocio.Descripcion,
		Ubicacion:   negocio.Ubicacion,
		Menu:        negocio.Menu,
		Tipo:        negocio.Tipo,
		Contacto:    negocio.Contacto,
		// TODO: GET THIS TO THE FILESYSTEM WITH THE IMAGENSERVICE
		// ImagenUrl:   negocio.,
	}

	err := s.db.NewInsert().
		Model(nDB).
		Returning("*").
		Scan(ctx)

	negocio.ID = nDB.ID
	return nDB, err
}

func (s *NegociosService) Update(ctx context.Context, negocio *services.NegocioPayload) (*services.Negocio, error) {
	nDB := &services.Negocio{
		ID:          negocio.ID,
		Nombre:      negocio.Nombre,
		Descripcion: negocio.Descripcion,
		Ubicacion:   negocio.Ubicacion,
		Menu:        negocio.Menu,
		Tipo:        negocio.Tipo,
		Contacto:    negocio.Contacto,
		// TODO: GET THIS TO THE FILESYSTEM WITH THE IMAGENSERVICE
		// ImagenUrl:   negocio.ImagenUrl,
	}

	err := s.db.NewUpdate().
		Model(nDB).
		WherePK().
		Returning("*").
		Scan(ctx)

	negocio.ID = nDB.ID
	return nDB, err
}

func (s *NegociosService) Delete(ctx context.Context, negocio *services.NegocioPayload) error {
	_, err := s.db.NewDelete().
		Model((*services.Negocio)(nil)).
		Where("id = ?", negocio.ID).
		Where("deleted_at IS NULL").
		Exec(ctx)

	return err
}
