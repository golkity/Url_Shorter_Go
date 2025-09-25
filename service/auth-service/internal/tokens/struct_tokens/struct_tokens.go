package struct_tokens

import "github.com/golang-jwt/jwt/v5"

type Tokens struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshToken     string `json:"refresh_token"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	TokenType        string `json:"token_type"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
