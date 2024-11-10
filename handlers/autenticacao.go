package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/andrevalario/projeto-estudos-score/domain"
	mdlusuario "github.com/andrevalario/projeto-estudos-score/model/usuario"
	"github.com/andrevalario/projeto-estudos-score/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func LoginUsuario(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var dadosLogin domain.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&dadosLogin); err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	usuario, err := mdlusuario.FetchByEmail(dadosLogin.Email)
	if err != nil || usuario.Id == 0 {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	//Validação simples de senha
	if usuario.Senha != dadosLogin.Senha {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	token, err := gerarTokenJWT(usuario)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	dados := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	utils.SendJSONResponse(w, dados, http.StatusOK)
}

func gerarTokenJWT(usuario domain.Usuario) (string, error) {
	var chaveSecreta = []byte(os.Getenv("JWT_TOKEN"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        usuario.Id,
		"nome":      usuario.Nome,
		"permissao": usuario.TipoUsuario,
		"exp":       time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(chaveSecreta)
}
