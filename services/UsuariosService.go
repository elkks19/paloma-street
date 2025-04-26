package services

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type UsuariosService interface {
	// Usuario returns the [Usuario] that matches the provided id
	Usuario(ctx context.Context, id uint32) (*Usuario, error)

	// Usuarios returns all [Usuario] in the database
	Usuarios(ctx context.Context, offset int, limit int) ([]*Usuario, error)

	// Create creates a new [Usuario] in the database
	Create(ctx context.Context, u *UsuarioPayload) (*Usuario, error)

	// Update updates the provided [Usuario] in the database
	Update(ctx context.Context, u *UsuarioPayload) (*Usuario, error)

	// Delete soft deletes the [Usuario] with the provided id from the database
	Delete(ctx context.Context, id uint32) error

	// ForceDelete removes the [Usuario] with the provided id from the database
	ForceDelete(ctx context.Context, id uint32) error
}

type UsuarioPayload struct {
	ID              uint32     `json:"id"`
	Rol             string     `json:"rol"`
	Nombre          string     `json:"nombre"`
	Email           string     `json:"email"`
	EmailVerificado *time.Time `json:"email_verificado"`
	Password        string     `json:"password"`
	Permisos        []string   `json:"permisos"`
}

type Usuario struct {
	bun.BaseModel `bun:"table:usuarios"`
	ID            uint32 `bun:"id,pk,autoincrement"`
	RolID         uint32 `bun:"rol_id,nullzero"`
	Rol           *Rol   `bun:"rel:belongs-to,join:rol_id=id"`

	Nombre          string     `bun:"nombre,nullzero"`
	Email           string     `bun:"email,nullzero,unique"`
	EmailVerificado *time.Time `bun:"email_verificado,nullzero,unique"`
	Password        string     `bun:"password,nullzero"`

	CreatedAt time.Time  `bun:"created_at,nullzero"`
	UpdatedAt time.Time  `bun:"updated_at,nullzero"`
	DeletedAt *time.Time `bun:"deleted_at,soft_delete,nullzero"`

	Permisos []*Permiso `bun:"m2m:usuario_has_permiso,join:Usuario=Permiso"`
	Negocios []*Negocio `bun:"m2m:negocio_has_usuario,join:Usuario=Negocio"`
}
