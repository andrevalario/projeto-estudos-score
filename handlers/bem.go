package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/andrevalario/projeto-estudos-score/domain"
	ucsbem "github.com/andrevalario/projeto-estudos-score/usecases/bem"
	"github.com/andrevalario/projeto-estudos-score/utils"
	"github.com/julienschmidt/httprouter"
)

func CriarBem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var novoBem domain.Bem
	if err := json.NewDecoder(r.Body).Decode(&novoBem); err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	err := ucsbem.CriarBem(r.Context(), novoBem)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	utils.SendJSONResponse(w, "Bem criado com sucesso", http.StatusOK)
}

func BuscarBemPorID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.ParseUint(p.ByName("id"), 10, 64)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	bem, err := ucsbem.BuscarBemPorID(r.Context(), id)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	utils.SendJSONResponse(w, bem, http.StatusOK)
}

func AtualizarBem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var bemAtualizado domain.Bem
	if err := json.NewDecoder(r.Body).Decode(&bemAtualizado); err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	err := ucsbem.Update(r.Context(), bemAtualizado)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	utils.SendJSONResponse(w, "Bem atualizado com sucesso", http.StatusOK)
}

func DeletarBem(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, err := strconv.ParseUint(p.ByName("id"), 10, 64)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	err = ucsbem.Delete(r.Context(), id)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	utils.SendJSONResponse(w, "Bem deletado com sucesso", http.StatusOK)
}
