package domain

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(os.Getenv("JWT_TOKEN"))

// Estrutura do token
type Claims struct {
	ID         uint64   `json:"id"`
	Nome       string   `json:"nome"`
	Permissoes []string `json:"permissoes"`
	jwt.StandardClaims
}

// Função para gerar o JWT
func GerarJWT(id uint64, nome string, permissoes []string) (string, error) {
	claims := Claims{
		ID:         id,
		Nome:       nome,
		Permissoes: permissoes,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "sistema",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
