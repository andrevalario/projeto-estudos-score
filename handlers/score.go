package handlers

import (
	"net/http"
	"strconv"

	ucsscore "github.com/andrevalario/projeto-estudos-score/usecases/score"
	"github.com/andrevalario/projeto-estudos-score/utils"
	"github.com/julienschmidt/httprouter"
)

func CalcularScoreAdmin(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	idUsuario, err := strconv.ParseUint(p.ByName("id_usuario"), 10, 64)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	score, err := ucsscore.CalcularScoreAdmin(r.Context(), idUsuario)
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	dados := struct {
		Score int `json:"score"`
	}{
		Score: score,
	}

	utils.SendJSONResponse(w, dados, http.StatusOK)
}

func CalcularScoreUsuario(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	score, err := ucsscore.CalcularScoreUsuario(r.Context())
	if err != nil {
		utils.ErrorResponseJson(r.Context(), w, err)
		return
	}

	dados := struct {
		Score int `json:"score"`
	}{
		Score: score,
	}

	utils.SendJSONResponse(w, dados, http.StatusOK)
}
