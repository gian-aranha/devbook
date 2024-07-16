package security

import "golang.org/x/crypto/bcrypt"

// Hash receives a string and turns it into a hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compares a hash and a password and returns if they're equal
func VerifyPassword(hashPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}