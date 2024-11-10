package handlers

import (
	"net/http"

	"github.com/andrevalario/projeto-estudos-score/utils"
	"github.com/julienschmidt/httprouter"
)

// Função que trata a rota /alive
func Alive(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	utils.SendJSONResponse(w, "API de estudos de score está online", http.StatusOK)
}
