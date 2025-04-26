package services

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type CategoriasService interface {
	// Categorias returns a list of all categorias.
	Categorias(ctx context.Context, offset int, limit int) ([]*Categoria, error)

	// CreateCategoria creates a new categoria in the database.
	Create(ctx context.Context, categoria *CategoriaPayload) (*Categoria, error)
}

type CategoriaPayload struct {
	ID     uint32 `json:"id"`
	Nombre string `json:"nombre"`
}

type Categoria struct {
	bun.BaseModel `bun:"table:categorias"`
	ID            uint32    `bun:"id,pk,autoincrement"`
	Nombre        string    `bun:"nombre,nullzero"`
	CreatedAt     time.Time `bun:"created_at,nullzero"`

	Productos []*Producto `bun:"m2m:categoria_has_producto,join:Categoria=Producto"`
}
