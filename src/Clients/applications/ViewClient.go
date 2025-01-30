package applications

import (
    "demo/src/Clients/domain/entities"
	"demo/src/Clients/domain"
)

type GetClientsUseCase struct {
    ClientRepo domain.ClientRepository
}

func NewGetClientsUseCase(repo domain.ClientRepository) *GetClientsUseCase {
    return &GetClientsUseCase{ClientRepo: repo}
}

func (uc *GetClientsUseCase) Execute() ([]entities.Client, error) {
    return uc.ClientRepo.GetAll()
}
