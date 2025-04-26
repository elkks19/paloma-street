package services

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type ProductosService interface {
	// Producto returns a producto by ID
	Producto(ctx context.Context, id uint32) (*Producto, error)

	// Productos returns a list of all productos in the database
	Productos(ctx context.Context, offset int, limit int) ([]*Producto, error)

	// ProductosByNegocio returns a list of all productos by negocio ID
	ProductosByNegocio(ctx context.Context, negocioID uint32, offset int, limit int) ([]*Producto, error)

	// CreateProducto creates a new producto in the database.
	Create(ctx context.Context, producto *ProductoPayload) (*Producto, error)

	// UpdateProducto updates a producto in the database.
	Update(ctx context.Context, producto *ProductoPayload) (*Producto, error)

	// DeleteProducto soft deletes a producto in the database.
	Delete(ctx context.Context, producto *ProductoPayload) error
}

type ProductoPayload struct {
	ID          uint32   `json:"id"`
	NegocioID   uint32   `json:"negocioId"`
	Nombre      string   `json:"nombre"`
	Descripcion string   `json:"descripcion"`
	Precio      float32  `json:"precio"`
	Categorias  []string `json:"categorias"`
}

type Producto struct {
	bun.BaseModel `bun:"table:productos"`
	ID            uint32   `bun:"id,pk,autoincrement"`
	NegocioID     uint32   `bun:"negocio_id,nullzero"`
	Negocio       *Negocio `bun:"rel:belongs-to,join:negocio_id=id"`

	Nombre      string  `bun:"nombre,nullzero"`
	Descripcion string  `bun:"descripcion,nullzero"`
	Precio      float32 `bun:"precio,nullzero"`

	CreatedAt time.Time  `bun:"created_at,nullzero"`
	UpdatedAt time.Time  `bun:"updated_at,nullzero"`
	DeletedAt *time.Time `bun:"deleted_at,nullzero"`

	Categorias []*Categoria `bun:"m2m:categoria_has_producto,join:Producto=Categoria"`
	Reviews    []*Review    `bun:"rel:has-many,join:id=producto_id"`
}

type ProductosCategorias struct {
	bun.BaseModel `bun:"table:productos_categorias"`
	ProductoID    uint32     `bun:"producto_id,pk,nullzero"`
	Producto      *Producto  `bun:"rel:belongs-to,join:producto_id=id"`
	CategoriaID   uint32     `bun:"categoria_id,pk,nullzero"`
	Categoria     *Categoria `bun:"rel:belongs-to,join:categoria_id=id"`
}
