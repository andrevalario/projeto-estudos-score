package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andrevalario/projeto-estudos-score/domain"
	"github.com/andrevalario/projeto-estudos-score/middleware"
	ucsdivida "github.com/andrevalario/projeto-estudos-score/usecases/divida"
	"github.com/andrevalario/projeto-estudos-score/utils"
	"github.com/gorilla/mux"
)

func CriarDivida(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.ValidarTokenEPermissoes(r, 1) // 1 para admin
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var novaDivida domain.Divida
	if err := json.NewDecoder(r.Body).Decode(&novaDivida); err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	validacao, err := ucsdivida.CriarDivida(r.Context(), novaDivida)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	if len(validacao) > 0 {
		utils.ValidationJsonResponse(r.Context(), w, validacao)
		return
	}

	utils.SendJSONResponse(w, "Dívida criada com sucesso", http.StatusOK)
}

func BuscarDivida(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.ValidarTokenEPermissoes(r, 1) // 1 para admin
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("ID inválido: %v", err))
		return
	}

	divida, err := ucsdivida.BuscarDivida(r.Context(), id)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	utils.SendJSONResponse(w, divida, http.StatusOK)
}

func AtualizarDivida(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.ValidarTokenEPermissoes(r, 1) // 1 para admin
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	var dividaAtualizada domain.Divida
	if err := json.NewDecoder(r.Body).Decode(&dividaAtualizada); err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("ID inválido: %v", err))
		return
	}

	validacao, err := ucsdivida.AtualizarDivida(r.Context(), id, dividaAtualizada)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	if len(validacao) > 0 {
		utils.ValidationJsonResponse(r.Context(), w, validacao)
		return
	}

	utils.SendJSONResponse(w, "Dívida atualizada com sucesso", http.StatusOK)
}

func DeletarDivida(w http.ResponseWriter, r *http.Request) {
	_, err := middleware.ValidarTokenEPermissoes(r, 1) // 1 para admin
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, fmt.Errorf("ID inválido: %v", err))
		return
	}

	err = ucsdivida.DeletarDivida(r.Context(), id)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	utils.SendJSONResponse(w, "Dívida deletada com sucesso", http.StatusOK)
}
