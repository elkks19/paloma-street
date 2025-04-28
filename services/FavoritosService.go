package services

import (
	"context"

	"github.com/uptrace/bun"
)

type FavoritosService interface {
	// Favoritos returns a list of Favoritos for a given user.
	Favoritos(ctx context.Context, u *Usuario, offset int, limit int) ([]*Favorito, error)

	// RegisterFavorito registers a new Favorito for a given user and negocio.
	RegisterFavorito(ctx context.Context, favorito *FavoritoPayload) error

	// UnregisterFavorito unregisters a Favorito for a given user and negocio.
	UnregisterFavorito(ctx context.Context, favorito *FavoritoPayload) error
}

type FavoritoPayload struct {
	UsuarioID uint32 `json:"usuarioId"`
	NegocioID uint32 `json:"negocioId"`
}

type Favorito struct {
	bun.BaseModel `bun:"table:favoritos"`
	UsuarioID     uint32   `bun:"usuario_id,pk,nullzero"`
	Usuario       *Usuario `bun:"rel:belongs-to,join:usuario_id=id"`
	NegocioID     uint32   `bun:"negocio_id,pk,nullzero"`
	Negocio       *Negocio `bun:"rel:belongs-to,join:negocio_id=id"`

	CreatedAt string `bun:"created_at,nullzero"`
	UpdatedAt string `bun:"updated_at,nullzero"`
	DeletedAt string `bun:"deleted_at,soft_delete,nullzero"`
}
