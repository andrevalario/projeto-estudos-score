package ucsbem

import (
	"context"

	"github.com/andrevalario/projeto-estudos-score/domain"
	mdlbem "github.com/andrevalario/projeto-estudos-score/model/bem"
)

func CriarBem(ctx context.Context, dadosBem domain.Bem) (err error) {
	err = mdlbem.Create(dadosBem)

	return
}

func BuscarBemPorID(ctx context.Context, idBem uint64) (bem domain.Bem, err error) {
	bem, err = mdlbem.FetchById(idBem)

	return
}

func Update(ctx context.Context, dadosBem domain.Bem) (err error) {
	err = mdlbem.Update(dadosBem)

	return
}

func Delete(ctx context.Context, idBem uint64) (err error) {
	err = mdlbem.Delete(idBem)

	return
}
