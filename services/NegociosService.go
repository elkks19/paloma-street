package services

import (
	"context"
	"proyecto/pkg/types"

	"github.com/uptrace/bun"
)

type NegociosService interface {
	// Negocio fetches a negocio by it's ID, if it exists.
	// otherwise, it returns an error.
	Negocio(ctx context.Context, id uint32) (*Negocio, error)

	// Negocios returns all negocios in the database
	Negocios(ctx context.Context, offset int, limit int) ([]*Negocio, error)

	// CreateNegocio creates a new negocio in the database
	Create(ctx context.Context, negocio *NegocioPayload) (*Negocio, error)

	// UpdateNegocio updates a negocio in the database
	Update(ctx context.Context, negocio *NegocioPayload) (*Negocio, error)

	// DeleteNegocio soft deletes a negocio in the database
	Delete(ctx context.Context, negocio *NegocioPayload) error
}

type NegocioPayload struct {
	ID          uint32            `json:"id"`
	Nombre      string            `json:"nombre"`
	Descripcion string            `json:"descripcion"`
	Ubicacion   []byte            `json:"ubicacion"`
	Menu        []byte            `json:"menu"`
	Tipo        string            `json:"tipo"`
	Contacto    string            `json:"contacto"`
	Usuarios    []*UsuarioPayload `json:"usuarios"`
	ImagenB64   string            `json:"imagen"`
}

type Negocio struct {
	bun.BaseModel `bun:"table:negocios"`
	ID            uint32 `bun:"id,pk,autoincrement"`

	Nombre       string        `bun:"nombre,nullzero"`
	Descripcion  string        `bun:"descripcion,nullzero"`
	Ubicacion    types.JSONMap `bun:"ubicacion,type:jsonb,json_use_number"`
	Menu         types.JSONMap `bun:"menu,type:jsonb,json_use_number"`
	Calificacion float32       `bun:"calificacion,nullzero"`
	Tipo         string        `bun:"tipo,nullzero"`
	Contacto     string        `bun:"contacto,nullzero"`
	ImagenUrl    string        `bun:"imagen_url,nullzero"`

	CreatedAt string `bun:"created_at,nullzero"`
	UpdatedAt string `bun:"updated_at,nullzero"`
	DeletedAt string `bun:"deleted_at,soft_delete,nullzero"`

	Productos []*Producto `bun:"rel:has-many,join:id=negocio_id"`
	Usuarios  []*Usuario  `bun:"m2m:negocio_usuarios,join:Negocio=Usuario"`
}

type NegocioUsuarios struct {
	bun.BaseModel `bun:"table:negocio_usuarios"`
	NegocioID     uint32   `bun:"negocio_id,pk,nullzero"`
	Negocio       *Negocio `bun:"rel:belongs-to,join:negocio_id=id"`
	UsuarioID     uint32   `bun:"usuario_id,pk,nullzero"`
	Usuario       *Usuario `bun:"rel:belongs-to,join:usuario_id=id"`
}
