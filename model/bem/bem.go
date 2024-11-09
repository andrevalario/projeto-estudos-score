package mdlbem

import "github.com/andrevalario/projeto-estudos-score/domain"

func Create(bem domain.Bem) error {
	r := Repository()
	return r.create(bem)
}

func FetchById(id uint64) (domain.Bem, error) {
	r := Repository()
	return r.fetchById(id)
}

func Update(bemAtualizado domain.Bem) error {
	r := Repository()
	return r.update(bemAtualizado)
}

func Delete(id uint64) error {
	r := Repository()
	return r.delete(id)
}
