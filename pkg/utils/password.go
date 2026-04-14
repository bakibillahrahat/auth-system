package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	// 10 is the cost of hashing, which determines the comutational complexity. Higer cost means more security but also more time to hash.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with a plain text password and returns true if they match.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}