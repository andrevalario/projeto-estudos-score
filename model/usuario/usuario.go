package mdlusuario

import "github.com/andrevalario/projeto-estudos-score/domain"

func Create(usuario domain.Usuario) error {
	r := Repository()
	return r.create(usuario)
}

// Função para buscar um usuário por ID
func FetchById(idUsuario uint64) (domain.Usuario, error) {
	r := Repository()
	return r.fetchById(idUsuario)
}

// Função para atualizar um usuário
func Update(usuarioAtualizado domain.Usuario) error {
	r := Repository()
	return r.update(usuarioAtualizado)
}

// Função para deletar um usuário
func Delete(idUsuario uint64) error {
	r := Repository()
	return r.delete(idUsuario)
}

// Função para buscar um usuário por um Email
func FetchByEmail(emailUsuario string) (domain.Usuario, error) {
	r := Repository()
	return r.fetchByEmail(emailUsuario)
}
