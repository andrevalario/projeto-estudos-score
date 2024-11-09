package ucsdivida

import (
	"context"
	"net/http"

	"github.com/andrevalario/projeto-estudos-score/domain"
	mdldivida "github.com/andrevalario/projeto-estudos-score/model/divida"
)

func CriarDivida(ctx context.Context, divida domain.Divida) (validationErrors []domain.ApiError, err error) {
	// Validação e regras de negócios para a dívida
	if divida.Valor <= 0 {
		validationErrors = append(validationErrors, domain.ApiError{
			Detail: "O valor da dívida deve ser maior que zero.",
			Status: http.StatusBadRequest,
		})
		return
	}

	err = mdldivida.Create(divida)
	return
}

func BuscarDivida(ctx context.Context, id uint64) (domain.Divida, error) {
	return mdldivida.FetchById(id)
}

func AtualizarDivida(ctx context.Context, idDivida uint64, dividaAtualizada domain.Divida) (validationErrors []domain.ApiError, err error) {
	// Aqui você pode validar os dados antes de atualizar a dívida
	if dividaAtualizada.Valor <= 0 {
		validationErrors = append(validationErrors, domain.ApiError{
			Status: 400,
			Detail: "O valor da dívida deve ser maior que zero.",
		})
		return
	}

	err = mdldivida.Update(idDivida, dividaAtualizada)
	return
}

func DeletarDivida(ctx context.Context, idDivida uint64) error {
	return mdldivida.Delete(idDivida)
}
