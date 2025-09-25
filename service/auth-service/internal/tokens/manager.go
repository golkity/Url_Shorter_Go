package tokens

import (
	"auth-service/internal/tokens/struct_tokens"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type Manager struct {
	secret            []byte
	accessTTLSeconds  int64
	refreshTTLSeconds int64
}

func NewManager(secret []byte, accessTTL, refreshTTL int64) *Manager {
	return &Manager{secret, accessTTL, refreshTTL}
}

func (m *Manager) Parse(tokenStr string) (*struct_tokens.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &struct_tokens.Claims{},
		func(t *jwt.Token) (interface{}, error) {
			return m.secret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	return token.Claims.(*struct_tokens.Claims), nil
}

func (m *Manager) SignedToken(UserID string, exp time.Time) (string, error) {
	jti := uuid.New().String()

	cls := struct_tokens.Claims{
		UserID: UserID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jti,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, cls)
	return tkn.SignedString(m.secret)
}

func (m *Manager) AccessTTLSeconds() int64  { return m.accessTTLSeconds }
func (m *Manager) RefreshTTLSeconds() int64 { return m.refreshTTLSeconds }
