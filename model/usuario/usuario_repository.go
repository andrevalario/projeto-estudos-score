package mdlusuario

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"

	"github.com/andrevalario/projeto-estudos-score/domain"
)

type UsuarioRepository struct {
	filePath string
	ultimoID uint64
	mutex    sync.Mutex
}

// Função para criar um novo repositório com caminho do arquivo JSON
func Repository() *UsuarioRepository {
	r := &UsuarioRepository{
		filePath: "./database/usuario.json",
		ultimoID: 0,
	}

	r.carregarUltimoID()

	return r
}

// Carrega o último ID utilizado ao iniciar o repositório
func (repo *UsuarioRepository) carregarUltimoID() {
	usuarios, err := repo.read()
	if err != nil {
		return
	}

	// Se houver usuários, o último ID será o maior ID existente
	for _, u := range usuarios {
		if u.Id > repo.ultimoID {
			repo.ultimoID = u.Id
		}
	}
}

// Função para gerar um novo ID incremental
func (repo *UsuarioRepository) gerarNovoID() uint64 {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.ultimoID++

	return repo.ultimoID
}

// Função para carregar todos os usuários do "banco de dados"
func (repo *UsuarioRepository) read() ([]domain.Usuario, error) {
	dados, err := ioutil.ReadFile(repo.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []domain.Usuario{}, nil // Se o arquivo não existe, retorna uma lista vazia
		}
		return nil, err
	}

	var response domain.UsuarioResponse
	if err := json.Unmarshal(dados, &response); err != nil {
		return nil, err
	}

	return response.Usuarios, nil
}

// Função para salvar a lista de usuários no "banco de dados"
func (repo *UsuarioRepository) save(usuarios []domain.Usuario) error {
	dados, err := json.MarshalIndent(usuarios, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(repo.filePath, dados, 0644)
}

// Função para adicionar um novo usuário
func (repo *UsuarioRepository) create(usuario domain.Usuario) error {
	usuarios, err := repo.read()
	if err != nil {
		return err
	}

	// Gera um novo ID incremental para o usuário
	usuario.Id = repo.gerarNovoID()

	usuarios = append(usuarios, usuario)
	return repo.save(usuarios)
}

// Função para buscar um usuário por ID
func (repo *UsuarioRepository) fetchById(idUsuario uint64) (domain.Usuario, error) {
	usuarios, err := repo.read()
	if err != nil {
		return domain.Usuario{}, err
	}

	for _, u := range usuarios {
		if u.Id == idUsuario {
			return u, nil
		}
	}
	return domain.Usuario{}, errors.New("usuário não encontrado")
}

// Função para atualizar um usuário existente
func (repo *UsuarioRepository) update(usuarioAtualizado domain.Usuario) error {
	usuarios, err := repo.read()
	if err != nil {
		return err
	}

	for i, u := range usuarios {
		if u.Id == usuarioAtualizado.Id {
			usuarios[i] = usuarioAtualizado
			return repo.save(usuarios)
		}
	}

	return errors.New("usuário informado não encontrado")
}

// Função para deletar um usuário
func (repo *UsuarioRepository) delete(idUsuario uint64) error {
	usuarios, err := repo.read()
	if err != nil {
		return err
	}

	for i, u := range usuarios {
		if u.Id == idUsuario {
			usuarios = append(usuarios[:i], usuarios[i+1:]...)
			return repo.save(usuarios)
		}
	}

	return errors.New("usuário informado não encontrado")
}
