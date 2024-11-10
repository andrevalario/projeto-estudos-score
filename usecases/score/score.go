package ucsscore

import (
	"context"
	"fmt"

	"github.com/andrevalario/projeto-estudos-score/domain"
	mdlmiddleware "github.com/andrevalario/projeto-estudos-score/middleware"
	mdlbem "github.com/andrevalario/projeto-estudos-score/model/bem"
	mdldivida "github.com/andrevalario/projeto-estudos-score/model/divida"
	mdlusuario "github.com/andrevalario/projeto-estudos-score/model/usuario"
)

func CalcularScoreAdmin(ctx context.Context, idUsuario uint64) (score int, err error) {
	usuario, err := mdlusuario.FetchById(idUsuario)
	if err != nil {
		return
	}

	if usuario.Id == 0 {
		err = fmt.Errorf("usuário não encontrado")
		return
	}

	bens, err := mdlbem.FetchByIdProprietario(usuario.Id)
	if err != nil {
		return
	}

	dividas, err := mdldivida.FetchByIdUsuario(usuario.Id)
	if err != nil {
		return
	}

	score = CalcularScoreCredito(bens, dividas)

	return
}

func CalcularScoreCredito(bens []domain.Bem, dividas []domain.Divida) (score int) {
	score = 1000

	quantidadeDeBens := len(bens)
	quantidadeDeDividas := len(dividas)

	// Cada bem aumenta o score em 50 pontos.
	score += quantidadeDeBens * 50

	// Cada dívida diminui o score em 100 pontos.
	score -= quantidadeDeDividas * 100

	if score < 0 {
		score = 0
		return
	}

	if score > 1000 {
		score = 1000
		return
	}

	return score
}

func CalcularScoreUsuario(ctx context.Context) (score int, err error) {
	usuarioAutenticado, err := mdlmiddleware.GetUsuarioAutenticado(ctx)
	if err != nil {
		return
	}

	bens, err := mdlbem.FetchByIdProprietario(usuarioAutenticado.Id)
	if err != nil {
		return
	}

	dividas, err := mdldivida.FetchByIdUsuario(usuarioAutenticado.Id)
	if err != nil {
		return
	}

	score = CalcularScoreCredito(bens, dividas)

	return
}
