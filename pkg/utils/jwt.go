package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key used to sign the JWT. In a real application, this should be stored securely and not hardcoded.
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// GenerateToekn: if login is successful, generate a JWT token for the user.
func GenerateToken(email string) (string, error) {
	// Set token claims
	claims := jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	// Create a new token object, specifying sigining method and the claims (payload).
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// Sign the token with the secret key and return it as a string.
	return token.SignedString(jwtSecret)
}