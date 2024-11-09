package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/andrevalario/projeto-estudos-score/domain"
	mdlusuario "github.com/andrevalario/projeto-estudos-score/model/usuario"
	"github.com/andrevalario/projeto-estudos-score/utils"
	"github.com/dgrijalva/jwt-go"
)

func LoginUsuario(w http.ResponseWriter, r *http.Request) {
	var dadosLogin domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&dadosLogin); err != nil {
		http.Error(w, "Erro ao decodificar dados de login", http.StatusBadRequest)
		return
	}

	usuario, err := mdlusuario.FetchById(dadosLogin.ID)
	if err != nil || usuario.Id == 0 {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	if usuario.Senha != dadosLogin.Senha {
		http.Error(w, "Senha inválida", http.StatusUnauthorized)
		return
	}

	token, err := gerarTokenJWT(usuario)
	if err != nil {
		http.Error(w, "Erro ao gerar o token", http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, json.NewEncoder(w).Encode(map[string]string{"token": token}), http.StatusOK)

}

func gerarTokenJWT(usuario domain.Usuario) (string, error) {
	// Definir a chave secreta para assinar o token
	var chaveSecreta = []byte("teste")

	// Criar a estrutura do token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        usuario.Id,
		"nome":      usuario.Nome,
		"permissao": usuario.TipoUsuario,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	// Gerar o token string
	return token.SignedString(chaveSecreta)
}
