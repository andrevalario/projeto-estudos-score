package ucsscore_test

import (
	"testing"

	"github.com/andrevalario/projeto-estudos-score/domain"
	ucsscore "github.com/andrevalario/projeto-estudos-score/usecases/score"
	"github.com/stretchr/testify/assert"
)

func TestCalcularScoreComBensEDividas(t *testing.T) {
	bens := []domain.Bem{
		{Valor: 50000},
		{Valor: 30000},
	}
	dividas := []domain.Divida{
		{Valor: 20000},
		{Valor: 10000},
		{Valor: 5000},
	}

	score := ucsscore.CalcularScoreCredito(bens, dividas)
	scoreEsperado := 800

	assert.Equal(t, scoreEsperado, score, "O cálculo do score não está correto.")
}
