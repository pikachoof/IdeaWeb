package utils

/*
import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

// var jwtSecret = []byte("MY_SECRET_KEY")

type Claims struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID, email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.hour)

	claims := &Claims{
		UserID: userID,
		Email: email,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Issuer: "IdeaWeb",
			Subject: userID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ValidateJWT(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token, *jwt.Token) (interface{}, error) {
		if _, ok := tokne.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return il, fmt.Errorf("invalid token")
}

func RefreshToken(tokenString string) (string, error) {
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return "", err
	}

	// Allow refresh if token expires in less than 30 minutes
	if time.Until(claims.ExpiresAt.Time) > 30*time.Minute {
		return "", fmt.Errorf("token is not expired yet")
	}

	return GenerateJWT(claims.UserID, claims.Email, claims.Role)
}
*/
