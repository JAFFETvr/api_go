package applications


import (
    "demo/src/Clients/domain/entities"
	"demo/src/Clients/domain"
)

type CreateClientUseCase struct {
    ClientRepo domain.ClientRepository
}

func NewCreateClientUseCase(repo domain.ClientRepository) *CreateClientUseCase {
    return &CreateClientUseCase{ClientRepo: repo}
}

func (uc *CreateClientUseCase) Execute(client *entities.Client) error {
    return uc.ClientRepo.Create(client)
}