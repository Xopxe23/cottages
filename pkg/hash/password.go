package hash

import (
	"crypto/sha1"
	"errors"
	"fmt"
)

type SHA1Hasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *SHA1Hasher {
	return &SHA1Hasher{salt: salt}
}

func (h *SHA1Hasher) Hash(password string) (string, error) {
	hash := sha1.New()
	if _, err := hash.Write([]byte(password)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum([]byte(h.salt))), nil
}

func (h *SHA1Hasher) VerifyPassword(password, hashedPassword string) error {
	passwordHash, err := h.Hash(password)
	if err != nil {
		return err
	}
	if passwordHash != hashedPassword {
		return errors.New("password incorrect")
	}
	return nil
}
