package crypto

import "golang.org/x/crypto/bcrypt"

type Hasher struct {
	Cost int
}

func NewHasher(cost int) *Hasher {
	return &Hasher{Cost: cost}
}

func (h *Hasher) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.Cost)
	return string(bytes), err
}

func (h *Hasher) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
