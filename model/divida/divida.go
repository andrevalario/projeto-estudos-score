package mdldivida

import (
	"github.com/andrevalario/projeto-estudos-score/domain"
)

// Função para criar uma dívida
func Create(divida domain.Divida) error {
	r := Repository()
	return r.create(divida)
}

// Função para buscar uma dívida por ID
func FetchById(idDivida uint64) (domain.Divida, error) {
	r := Repository()
	return r.fetchById(idDivida)
}

// Função para atualizar uma dívida
func Update(idDivida uint64, dividaAtualizada domain.Divida) error {
	r := Repository()
	return r.update(idDivida, dividaAtualizada)
}

// Função para deletar uma dívida
func Delete(idDivida uint64) error {
	r := Repository()
	return r.delete(idDivida)
}
