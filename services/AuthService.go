package services

import (
	"context"

	"github.com/uptrace/bun"
)

type AuthService interface {
	// Register stores the information of the new user in the database
	// and attempts to login, if successful, it returns a session.
	// If unsuccessful, it returns an error.
	Register(ctx context.Context, u *UsuarioPayload) (*Session, error)

	// Attempt tries to authenticate a user with the provided user password
	// and email. If successful, it returns the user and a token.
	// If unsuccessful, it returns an error.
	Attempt(ctx context.Context, u *UsuarioPayload) (*Session, error)

	// Logout invalidates the provided [Session], effectively
	// logging out the user it returns an error if the logout fails.
	Logout(ctx context.Context, s *Session) error

	// UserCan verifies if provided user currently has the
	// permission to perform a specific action.
	// It returns true if the user has the permission, false otherwise.
	// It returns an error if the verification fails.
	UserCan(ctx context.Context, u *UsuarioPayload, permisos ...string) (bool, error)

	// UserHasRole verifies if the provided user has a specific role.
	// It returns true if the user has the role, false otherwise.
	// It returns an error if the verification fails.
	UserHasRole(ctx context.Context, u *UsuarioPayload, rol string) (bool, error)
}

type Rol struct {
	bun.BaseModel `bun:"table:roles"`
	ID            uint32 `bun:"id,pk,autoincrement"`
	Nombre        string `bun:"nombre,nullzero"`

	Usuarios []*Usuario `bun:"rel:has-many,join:id=rol_id"`
}

type Permiso struct {
	bun.BaseModel `bun:"table:permisos"`
	ID            uint16 `bun:"id,pk,autoincrement"`
	Nombre        string `bun:"nombre,nullzero"`

	Usuario []*Usuario `bun:"m2m:usuario_has_permiso,join:Permiso=Usuario"`
}

type UsuarioHasPermiso struct {
	bun.BaseModel `bun:"table:usuario_has_permiso"`
	UsuarioID     uint64   `bun:"usuario_id,pk"`
	Usuario       *Usuario `bun:"rel:belongs-to,join:usuario_id=id"`
	PermisoID     uint16   `bun:"permiso_id,pk"`
	Permiso       *Permiso `bun:"rel:belongs-to,join:permiso_id=id"`
}
