package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrevalario/projeto-estudos-score/domain"
	ucsusuario "github.com/andrevalario/projeto-estudos-score/usecases/usuario"
	"github.com/andrevalario/projeto-estudos-score/utils"
	"github.com/gorilla/mux"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var novoUsuario domain.Usuario
	if err := json.NewDecoder(r.Body).Decode(&novoUsuario); err != nil {
		return
	}

	validacao, err := ucsusuario.CriarUsuario(r.Context(), novoUsuario)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	if len(validacao) > 0 {
		utils.ValidationJsonResponse(r.Context(), w, validacao)
		return
	}

	utils.SendJSONResponse(w, "Usuário criado com sucesso", http.StatusOK)
}

func BuscarUsuarioPorID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("ID inválido: %v", err))
		return
	}

	usuario, err := ucsusuario.FetchUsuarioById(r.Context(), id)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	utils.SendJSONResponse(w, usuario, http.StatusOK)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuarioAtualizado domain.Usuario
	if err := json.NewDecoder(r.Body).Decode(&usuarioAtualizado); err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	err := ucsusuario.Update(r.Context(), usuarioAtualizado)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	utils.SendJSONResponse(w, "Usuário atualizado com sucesso", http.StatusOK)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("ID inválido: %v", err))
		return
	}

	err = ucsusuario.Delete(r.Context(), id)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	utils.SendJSONResponse(w, "Usuário deletado com sucesso", http.StatusOK)
}
