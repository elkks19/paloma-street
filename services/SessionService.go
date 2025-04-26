package services

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/uptrace/bun"
)

// Auxiliar functions for the auth service that separate db operations
// from the store implementation of gorilla/sessions
type SessionService interface {
	// Create stores a new [Session] for the provided [Usuario]
	// in the database and returns an error if the operation fails.
	Create(ctx context.Context, ses *Session) error

	// Update updates the provided [Session] in the database
	// and returns an error if the operation fails.
	Update(ctx context.Context, ses *Session) error

	// Validate checks if the provided [Session] is valid.
	// It returns the user associated with the session and an error
	// if the validation fails.
	Validate(ctx context.Context, key []byte) (*Usuario, error)

	// Select retrieves a [Session] from the database
	// using the provided key. It returns an error if the
	// retrieval fails.
	Exists(ctx context.Context, key []byte) (*Session, error)

	// Delete removes the provided [Session] from the database
	// and returns an error if the operation fails.
	Delete(ctx context.Context, key []byte) error
}

// See https://github.com/golang-jwt/jwt for more examples
type CustomClaims struct {
	jwt.RegisteredClaims
	*Usuario
}

type SessionPayload struct {
	ID        uint64 `json:"id"`
	UsuarioID uint32 `json:"usuario_id"`
}

type Session struct {
	bun.BaseModel `bun:"table:sessions"`
	ID            uint32   `bun:",pk,autoincrement"`
	UsuarioID     *uint32  `bun:"usuario_id,nullzero"`
	Usuario       *Usuario `bun:"rel:belongs-to,join:usuario_id=id"`

	Token string `bun:"token,nullzero"`

	CreatedAt  time.Time `bun:"created_at,nullzero"`
	LastAccess time.Time `bun:"last_access,nullzero"`
	ExpiresAt  time.Time `bun:"expires_at,nullzero"`
}
