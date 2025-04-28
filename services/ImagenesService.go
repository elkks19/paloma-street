package services

import (
	"time"

	"github.com/uptrace/bun"
)

type ImagenesService interface {
	// Save takes a b64 string, stores it in the filesystem, and returns the full url to the image.
	// The model parameter is used to determine the path where the image will be stored
	// it can be a [Producto] or a [Negocio]; any other model will cause the app to panic.
	Save(b64 string, model any) (string, error)

	// Delete takes a url and deletes the image from the filesystem.
	Delete(url string) error

	// Get takes a url and returns the image as a byte array.
	Get(url string) ([]byte, error)
}

type ProductoImagen struct {
	bun.BaseModel `bun:"table:producto_imagenes"`

	ID         uint32    `bun:"id,pk,autoincrement"`
	ProductoID uint32    `bun:"producto_id,pk,nullzero"`
	Producto   *Producto `bun:"rel:belongs-to,join:producto_id=id"`
	ImagenUrl  string    `bun:"imagen_url,nullzero"`

	CreatedAt time.Time  `bun:"created_at,nullzero"`
	UpdatedAt time.Time  `bun:"updated_at,nullzero"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete,nullzero"`
}
