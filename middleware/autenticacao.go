package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/andrevalario/projeto-estudos-score/domain"
	"github.com/dgrijalva/jwt-go"
)

// ValidarTokenEPermissoes - Função que valida o token e verifica permissões do usuário
func ValidarTokenEPermissoes(r *http.Request, requiredPermission uint64) (domain.Usuario, error) {
	// Obter o token do header Authorization
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return domain.Usuario{}, fmt.Errorf("Token não encontrado")
	}

	// Remover o prefixo "Bearer "
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	// Validar e decodificar o token JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validar o método de assinatura
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Método de assinatura inválido")
		}
		// Retornar a chave secreta para validação
		return []byte("minha-chave-secreta"), nil
	})

	// Se o token for inválido ou ocorrer algum erro
	if err != nil || !token.Valid {
		return domain.Usuario{}, fmt.Errorf("Token inválido")
	}

	// Obter as claims do token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return domain.Usuario{}, fmt.Errorf("Erro ao ler as claims do token")
	}

	// Criar o usuário a partir das claims
	user := domain.Usuario{
		Id:          uint64(claims["id"].(float64)),
		Nome:        claims["nome"].(string),
		TipoUsuario: uint64(claims["tipoUsuario"].(float64)),
	}

	// Verificar se o usuário tem a permissão necessária
	if user.TipoUsuario != requiredPermission {
		return domain.Usuario{}, fmt.Errorf("Permissão insuficiente")
	}

	return user, nil
}
