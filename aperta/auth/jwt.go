package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

// The structure of Claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// The private key to generate the JWT token.
// TODO: move it to config file and keep it secure
var jwtKey = []byte("my_secret_key")

// go-kit : encapsulating our methods into the interface
type JwtService interface {
	GetToken(username string) (string, error)
	ValidateToken(token string) (bool, int)
}

type jwtService struct{}

// Service constructor
func NewJwtService() *jwtService {
	return &jwtService{}
}

// GetToken: This function generate the JWT token, it also add the token into the claims so that it can be validated.
// It signs the token with the private key.
func (j jwtService) GetToken(username string) (string, error) {

	// Here we set the expiration for the token, it is better if token is short lived.
	// setting token expiration to be 5 minutes.
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, it includes expiry date and issuer information.
	// Later we could validate our JWT tokens
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(expirationTime),
		Issuer:    username,
	}

	// Using SigningMethodHS256 to sign the token, any other algorithm can be used instead.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("internal server error")
	}

	// Retrun the generated token
	return tokenString, nil
}

// ValidateToken: This function takes a token as input and tries to validate it against the claims and also checks if it is not expired.
func (j jwtService) ValidateToken(token string) (bool, int) {

	// Initialize a new instance of `Claims`
	claims := &Claims{}

	// Please note that here we are parsing the incoming token based on the claims.
	// If the token is not valid it generates different kind of errors
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return false, http.StatusUnauthorized
		}
		return false, http.StatusBadRequest
	}
	if !tkn.Valid {
		return false, http.StatusUnauthorized
	}
	return true, http.StatusOK
}
