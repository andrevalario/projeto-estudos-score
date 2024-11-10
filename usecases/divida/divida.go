package ucsdivida

import (
	"context"
	"net/http"

	"github.com/andrevalario/projeto-estudos-score/domain"
	mdldivida "github.com/andrevalario/projeto-estudos-score/model/divida"
)

func CriarDivida(ctx context.Context, divida domain.Divida) (validacao []domain.ApiError, err error) {
	validacao = validateDivida(divida)
	if len(validacao) > 0 {
		return
	}

	err = mdldivida.Create(divida)
	return
}

func BuscarDivida(ctx context.Context, id uint64) (domain.Divida, error) {
	return mdldivida.FetchById(id)
}

func AtualizarDivida(ctx context.Context, idDivida uint64, dividaAtualizada domain.Divida) (validacao []domain.ApiError, err error) {
	validacao = validateDivida(dividaAtualizada)
	if len(validacao) > 0 {
		return
	}

	err = mdldivida.Update(idDivida, dividaAtualizada)
	return
}

func DeletarDivida(ctx context.Context, idDivida uint64) error {
	return mdldivida.Delete(idDivida)
}

func validateDivida(divida domain.Divida) (validacao []domain.ApiError) {
	if divida.Valor <= 0 {
		validacao = append(validacao, domain.ApiError{
			Detail: "O valor da dívida deve ser maior que zero.",
			Status: http.StatusBadRequest,
		})
		return
	}

	return
}
