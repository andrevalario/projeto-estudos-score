package mdldivida

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"sync"

	"github.com/andrevalario/projeto-estudos-score/domain"
)

type DividaRepository struct {
	filePath string
	ultimoID uint64
	mutex    sync.Mutex
}

func Repository() *DividaRepository {
	repo := &DividaRepository{
		filePath: "database/dividas.json",
		ultimoID: 0,
	}
	repo.carregarUltimoID()
	return repo
}

func (repo *DividaRepository) carregarUltimoID() {
	dividas, err := repo.read()
	if err != nil {
		return
	}

	for _, divida := range dividas {
		if divida.Id > repo.ultimoID {
			repo.ultimoID = divida.Id
		}
	}
}

func (repo *DividaRepository) gerarNovoID() uint64 {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	repo.ultimoID++
	return repo.ultimoID
}

func (repo *DividaRepository) read() ([]domain.Divida, error) {
	data, err := ioutil.ReadFile(repo.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []domain.Divida{}, nil
		}
		return nil, err
	}

	var dividas domain.DividaResponse
	if err := json.Unmarshal(data, &dividas); err != nil {
		return nil, err
	}

	return dividas.Dividas, nil
}

func (repo *DividaRepository) save(dividas []domain.Divida) error {
	dividaResponse := domain.DividaResponse{
		Dividas: dividas,
	}

	data, err := json.MarshalIndent(dividaResponse, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(repo.filePath, data, 0644)
}

func (repo *DividaRepository) create(divida domain.Divida) error {
	dividas, err := repo.read()
	if err != nil {
		return err
	}

	divida.Id = repo.gerarNovoID()
	dividas = append(dividas, divida)
	return repo.save(dividas)
}

func (repo *DividaRepository) fetchById(id uint64) (domain.Divida, error) {
	dividas, err := repo.read()
	if err != nil {
		return domain.Divida{}, err
	}

	for _, divida := range dividas {
		if divida.Id == id {
			return divida, nil
		}
	}

	return domain.Divida{}, errors.New("dívida não encontrada")
}

func (repo *DividaRepository) update(id uint64, dividaAtualizada domain.Divida) error {
	dividas, err := repo.read()
	if err != nil {
		return err
	}

	for i, divida := range dividas {
		if divida.Id == id {
			dividas[i] = dividaAtualizada
			return repo.save(dividas)
		}
	}
	return errors.New("dívida não encontrada")
}

func (repo *DividaRepository) delete(id uint64) error {
	dividas, err := repo.read()
	if err != nil {
		return err
	}

	for i, divida := range dividas {
		if divida.Id == id {
			dividas = append(dividas[:i], dividas[i+1:]...)
			return repo.save(dividas)
		}
	}
	return errors.New("dívida não encontrada")
}

func (repo *DividaRepository) fetchByIdUsuario(idUsuario uint64) (dividas []domain.Divida, err error) {
	dividas, err = repo.read()
	if err != nil {
		return []domain.Divida{}, err
	}

	for _, divida := range dividas {
		if divida.IdUsuario == idUsuario {
			dividas = append(dividas, divida)
		}
	}

	return
}
