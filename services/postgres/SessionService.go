package postgres

import (
	"proyecto/services"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

type SessionService struct {
	db      *bun.DB
}

func NewSessionService(db *bun.DB) *SessionService {
	return &SessionService{
		db: db,
	}
}

func (s *SessionService) Create(ctx context.Context, ses *services.Session) error {
	err := s.db.NewInsert().
		Model(ses).
		Column(
			"usuario_id",
			"key",
			"data",
			"token",
			"created_at",
			"last_access",
			"expires_at",
		).
		Returning("*").
		Scan(ctx)

	return err
}

func (s *SessionService) Update(ctx context.Context, ses *services.Session) error {
	err := s.db.NewUpdate().
		Model(ses).
		Set("usuario_id = ?", ses.UsuarioID).
		Set("token = ?", ses.Token).
		Set("last_access = ?", time.Now()).
		Returning("*").
		Scan(ctx)

	return err
}

func (s *SessionService) Validate(ctx context.Context, key []byte) (*services.Usuario, error) {
	dbSession := new(services.Session)
	err := s.db.NewSelect().
		Model(dbSession).
		Where("key = ?", key).
		Relation("Usuario").
		Where("expires_at > ?", time.Now()).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
	}

	return dbSession.Usuario, err
}

func (s *SessionService) Exists(ctx context.Context, key []byte) (*services.Session, error) {
	dbSession := new(services.Session)
	err := s.db.NewSelect().
		Model(dbSession).
		Where("key = ?", key).
		Where("expires_at > ?", time.Now()).
		Scan(ctx)

	return dbSession, err
}

func (s *SessionService) Delete(ctx context.Context, key []byte) error {
	_, err := s.db.NewDelete().
		Model((*services.Session)(nil)).
		Where("key = ?", key).
		Exec(ctx)

	return err
}
