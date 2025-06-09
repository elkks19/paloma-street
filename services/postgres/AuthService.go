package postgres

import (
	"context"
	"database/sql"
	"proyecto/internal/config/env"
	"proyecto/services"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/uptrace/bun"
)

type AuthService struct {
	db *bun.DB
}

func NewAuthService(db *bun.DB) *AuthService {
	db.RegisterModel((*services.UsuarioHasPermiso)(nil))
	return &AuthService{
		db: db,
	}
}

func (s *AuthService) Register(ctx context.Context, up *services.UsuarioPayload) (*services.Session, error) {
	exists, err := s.db.NewSelect().
		Model((*services.Usuario)(nil)).
		Where("email = ?", up.Email).
		Exists(ctx)

	if exists {
		return nil, services.ErrEmailExists
	}

	if err != nil {
		return nil, err
	}

	// if there's no user with that email, we can create a new one
	b := env.Get("BCRYPT_ROUNDS")
	bcryptRounds, err := strconv.Atoi(b)
	if err != nil {
		panic("ENV VARIABLE BCRYPT_ROUNDS IS NOT AN INTEGER")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(up.Password), bcryptRounds)
	if err != nil {
		return nil, err
	}

	up.Password = string(hash)

	dbU := &services.Usuario{
		RolID:    1,
		Nombre:   up.Nombre,
		Email:    up.Email,
		Password: up.Password,
	}

	err = s.db.NewInsert().
		Model(dbU).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	// create a session for the user
	ses, err := s.Attempt(ctx, up)

	return ses, err
}

func (s *AuthService) Attempt(ctx context.Context, u *services.UsuarioPayload) (*services.Session, error) {
	uDB := new(services.Usuario)
	err := s.db.NewSelect().
		Model(uDB).
		Where("email = ?", u.Email).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, services.ErrUserNotFound
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(uDB.Password), []byte(u.Password))
	if err != nil {
		return nil, services.ErrWrongPassword
	}

	exp := env.Get("SESSION_EXPIRATION_HRS")
	expirationHrs, err := strconv.Atoi(exp)
	if err != nil {
		panic("ENV VARIABLE SESSION_EXPIRATION_HRS IS NOT AN INTEGER")
	}

	claims := &services.CustomClaims{
		jwt.RegisteredClaims{
			Issuer:    env.Get("APP_NAME"),
			Subject:   u.Nombre,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expirationHrs) * time.Hour)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		uDB,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ses := &services.Session{
		UsuarioID: &u.ID,
		Token:     token.Raw,
	}

	return ses, err
}

func (s *AuthService) Logout(ctx context.Context, ses *services.Session) error {
	return nil
}

func (s *AuthService) UserCan(ctx context.Context, u *services.UsuarioPayload, permisos ...string) (bool, error) {

	return false, nil
}

func (s *AuthService) UserHasRole(ctx context.Context, u *services.UsuarioPayload, rol string) (bool, error) {
	return false, nil
}
