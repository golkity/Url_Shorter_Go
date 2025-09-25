package tokens

import (
	"auth-service/internal/tokens/struct_tokens"
	"log"
	"time"
)

func (m *Manager) GenerateToken(UserID string) (*struct_tokens.Tokens, error) {
	now := time.Now()

	access, err := m.SignedToken(UserID, now.Add(time.Duration(m.accessTTLSeconds)*time.Second))
	if err != nil {
		return nil, err
	}
	refresh, err := m.SignedToken(UserID, now.Add(time.Duration(m.refreshTTLSeconds)*time.Second))
	if err != nil {
		return nil, err
	}

	exp := now.Add(time.Duration(m.refreshTTLSeconds) * time.Second)
	log.Printf("New access exp=%s (ttl=%d)", exp.Format(time.RFC3339), m.accessTTLSeconds)

	return &struct_tokens.Tokens{
		AccessToken:      access,
		ExpiresIn:        m.accessTTLSeconds,
		RefreshToken:     refresh,
		RefreshExpiresIn: m.refreshTTLSeconds,
		TokenType:        "bearer",
	}, nil
}
