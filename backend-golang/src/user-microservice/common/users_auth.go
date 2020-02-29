package common

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("my_secret_key")

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"user_name" bson:"user_name"`
	Role     string `json:"role" bson:"role"`
	Email    string `json:"email" bson:"email"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT token
func GenerateJWT(username, role, email string) (string, error) {
	// Declare the expiration time of the token here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(45 * time.Minute)

	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		Role:     role,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // In JWT, the expiry time is expressed as unix milliseconds
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return "", err
	}
	return tokenString, nil
}

func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// Initialize a new instance of `Claims`
	claims := &Claims{}
	// Get token from request
	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			DisplayError(w, err, http.StatusUnauthorized,
				"Unauthorized",
			)
			return
		}
		DisplayError(w, err, http.StatusBadRequest,
			"Bad Request",
		)
		return
	}
	if !tkn.Valid {
		DisplayError(w, err, http.StatusUnauthorized,
			"Unauthorized",
		)
		return
	} else if tkn.Valid {
		next(w, r)
	}
}
