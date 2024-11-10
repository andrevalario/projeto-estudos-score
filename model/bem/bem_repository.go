package mdlbem

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"

	"github.com/andrevalario/projeto-estudos-score/domain"
)

type BemRepository struct {
	filePath string
	ultimoID uint64
	mutex    sync.Mutex
}

func Repository() *BemRepository {
	repo := &BemRepository{
		filePath: "database/bens.json",
		ultimoID: 0,
	}
	repo.carregarUltimoID()
	return repo
}

func (repo *BemRepository) carregarUltimoID() {
	bens, err := repo.read()
	if err != nil {
		return
	}

	for _, d := range bens {
		if d.Id > repo.ultimoID {
			repo.ultimoID = d.Id
		}
	}
}

func (repo *BemRepository) gerarNovoID() uint64 {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.ultimoID++
	return repo.ultimoID
}

func (repo *BemRepository) read() ([]domain.Bem, error) {
	data, err := ioutil.ReadFile(repo.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []domain.Bem{}, nil
		}
		return nil, err
	}

	var bens domain.BemResponse
	if err := json.Unmarshal(data, &bens); err != nil {
		return nil, err
	}

	return bens.Bem, nil
}

func (repo *BemRepository) save(bens []domain.Bem) error {
	data, err := json.MarshalIndent(bens, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(repo.filePath, data, 0644)
}

func (repo *BemRepository) create(bem domain.Bem) error {
	bens, err := repo.read()
	if err != nil {
		return err
	}

	bem.Id = repo.gerarNovoID()
	bens = append(bens, bem)
	return repo.save(bens)
}

func (repo *BemRepository) fetchById(id uint64) (domain.Bem, error) {
	bens, err := repo.read()
	if err != nil {
		return domain.Bem{}, err
	}

	for _, bem := range bens {
		if bem.Id == id {
			return bem, nil
		}
	}
	return domain.Bem{}, errors.New("Bem não encontrado")
}

func (repo *BemRepository) update(bemAtualizado domain.Bem) error {
	bens, err := repo.read()
	if err != nil {
		return err
	}

	for i, bem := range bens {
		if bem.Id == bemAtualizado.Id {
			bens[i] = bemAtualizado
			return repo.save(bens)
		}
	}
	return errors.New("Bem não encontrado para atualização")
}

func (repo *BemRepository) delete(id uint64) error {
	bens, err := repo.read()
	if err != nil {
		return err
	}

	for i, bem := range bens {
		if bem.Id == id {
			bens = append(bens[:i], bens[i+1:]...)
			return repo.save(bens)
		}
	}
	return errors.New("Bem não encontrado para deleção")
}

func (repo *BemRepository) fetchByIdProprietario(IdProprietario uint64) (bens []domain.Bem, err error) {
	bens, err = repo.read()
	if err != nil {
		return []domain.Bem{}, err
	}

	for _, bem := range bens {
		if bem.IdProprietario == IdProprietario {
			bens = append(bens, bem)
		}
	}

	return
}
