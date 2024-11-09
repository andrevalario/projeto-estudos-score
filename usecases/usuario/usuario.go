package ucsusuario

import (
	"context"

	"github.com/andrevalario/projeto-estudos-score/domain"
	mdlusuario "github.com/andrevalario/projeto-estudos-score/model/usuario"
)

func CriarUsuario(ctx context.Context, dadosUsuario domain.Usuario) (validacao []domain.ApiError, err error) {
	err = mdlusuario.Create(domain.Usuario{
		Nome:        dadosUsuario.Nome,
		Email:       dadosUsuario.Email,
		Senha:       dadosUsuario.Senha,
		TipoUsuario: dadosUsuario.TipoUsuario,
	})

	return
}

func FetchUsuarioById(ctx context.Context, idUsuario uint64) (usuario domain.Usuario, err error) {
	usuario, err = mdlusuario.FetchById(idUsuario)

	return
}

func Update(ctx context.Context, dadosUsuario domain.Usuario) (err error) {
	err = mdlusuario.Update(domain.Usuario{
		Id:          dadosUsuario.Id,
		Nome:        dadosUsuario.Nome,
		Email:       dadosUsuario.Email,
		Senha:       dadosUsuario.Senha,
		TipoUsuario: dadosUsuario.TipoUsuario,
	})

	return
}

func Delete(ctx context.Context, idUsuario uint64) (err error) {
	err = mdlusuario.Delete(idUsuario)

	return
}
