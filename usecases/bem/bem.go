package ucsbem

import (
	"context"
	"fmt"

	"github.com/andrevalario/projeto-estudos-score/domain"
	mdlmiddleware "github.com/andrevalario/projeto-estudos-score/middleware"
	mdlbem "github.com/andrevalario/projeto-estudos-score/model/bem"
)

func CriarBem(ctx context.Context, dadosBem domain.Bem) (err error) {
	usuarioAutenticado, err := mdlmiddleware.GetUsuarioAutenticado(ctx)
	if err != nil {
		return
	}

	if usuarioAutenticado.TipoUsuario != domain.Admin && usuarioAutenticado.Id != dadosBem.IdProprietario {
		return fmt.Errorf("cadastro de bem não permitido para outro usuário")
	}

	err = mdlbem.Create(dadosBem)

	return
}

func BuscarBemPorID(ctx context.Context, idBem uint64) (bem domain.Bem, err error) {
	usuarioAutenticado, err := mdlmiddleware.GetUsuarioAutenticado(ctx)
	if err != nil {
		return
	}

	bem, err = mdlbem.FetchById(idBem)
	if err != nil {
		return
	}

	if usuarioAutenticado.TipoUsuario != domain.Admin && usuarioAutenticado.Id != bem.IdProprietario {
		return domain.Bem{}, fmt.Errorf("não é permitido a visualização de bens de outro usuário")
	}

	return
}

func Update(ctx context.Context, dadosBem domain.Bem) (err error) {
	usuarioAutenticado, err := mdlmiddleware.GetUsuarioAutenticado(ctx)
	if err != nil {
		return
	}

	if usuarioAutenticado.TipoUsuario != domain.Admin && usuarioAutenticado.Id != dadosBem.IdProprietario {
		return fmt.Errorf("não é permitido atualizar o bem de outro usuário")
	}

	err = mdlbem.Update(dadosBem)

	return
}

func Delete(ctx context.Context, idBem uint64) (err error) {
	usuarioAutenticado, err := mdlmiddleware.GetUsuarioAutenticado(ctx)
	if err != nil {
		return
	}

	bem, err := mdlbem.FetchById(idBem)
	if err != nil {
		return
	}

	if usuarioAutenticado.TipoUsuario != domain.Admin && usuarioAutenticado.Id != bem.IdProprietario {
		return fmt.Errorf("não é permitido a deleção do bem de outro usuário")
	}

	err = mdlbem.Delete(idBem)

	return
}
