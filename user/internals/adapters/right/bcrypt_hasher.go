package adapters

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter struct {
	cost int
}

func NewBcryptAdapter(cost int) *BcryptAdapter {
	return &BcryptAdapter{
		cost,
	}
}

func (bc *BcryptAdapter) Hash(password string) (string, error) {
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(password), bc.cost)

	if err != nil {
		return "", fmt.Errorf("password hashing failed: %w", err)
	}

	return string(hashedPwd), nil

}

func (bc *BcryptAdapter) Compare(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
