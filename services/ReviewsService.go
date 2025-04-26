package services

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type ReviewsService interface {
	// Reviews by usuario returns a list of reviews for a given user.
	ReviewsByUsuario(ctx context.Context, usuarioId uint32, offset int, limit int) ([]*Review, error)

	// ReviewsByUsuario returns a list of reviews for a given user.
	ReviewsByNegocio(ctx context.Context, negocioId uint32, offset int, limit int) ([]*Review, error)

	// ReviewsByProducto returns a list of reviews for a given producto.
	ReviewsByProducto(ctx context.Context, productoId uint32, offset int, limit int) ([]*Review, error)

	// CreateReview creates a new review for a given user and negocio.
	Create(ctx context.Context, review *ReviewPayload) (*Review, error)

	// UpdateReview updates a review for a given user and negocio.
	Update(ctx context.Context, review *ReviewPayload) (*Review, error)

	// DeleteReview deletes a review for a given user and negocio.
	Delete(ctx context.Context, review *ReviewPayload) (*Review, error)
}

type ReviewPayload struct {
	ID           uint32 `json:"id"`
	Contenido    string `json:"contenido"`
	Calificacion uint32 `json:"calificacion"`
	UsuarioID    uint32 `json:"usuarioId"`
	ProductoID   uint32 `json:"productoId"`
}

type Review struct {
	bun.BaseModel `bun:"table:reviews"`

	ID        uint32   `bun:"id,pk,autoincrement"`
	UsuarioID uint32   `bun:"usuario_id,nullzero"`
	Usuario   *Usuario `bun:"rel:belongs-to,join:usuario_id=id"`
	ProductoID uint32   `bun:"producto_id,nullzero"`
	Producto  *Producto `bun:"rel:belongs-to,join:producto_id=id"`

	Contenido    string `bun:"contenido,nullzero"`
	Calificacion uint32 `bun:"calificacion,nullzero"`

	CreatedAt time.Time  `bun:"created_at,nullzero"`
	UpdatedAt time.Time  `bun:"updated_at,nullzero"`
	DeletedAt *time.Time `bun:"deleted_at,nullzero"`
}
