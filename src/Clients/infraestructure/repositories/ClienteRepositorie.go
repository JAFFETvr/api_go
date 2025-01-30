package repositories

import (
    "gorm.io/gorm"
    "demo/src/Clients/domain"
	"demo/src/Clients/domain/entities"
)

type ClientRepositoryImpl struct {
    DB *gorm.DB
}

func NewClientRepository(db *gorm.DB) domain.ClientRepository {
    return &ClientRepositoryImpl{DB: db}
}

func (r *ClientRepositoryImpl) Create(client *entities.Client) error {
    return r.DB.Create(client).Error
}

func (r *ClientRepositoryImpl) GetByID(id uint) (*entities.Client, error) {
    var client entities.Client
    if err := r.DB.First(&client, id).Error; err != nil {
        return nil, err
    }
    return &client, nil
}

func (r *ClientRepositoryImpl) GetAll() ([]entities.Client, error) {
    var clients []entities.Client
    if err := r.DB.Find(&clients).Error; err != nil {
        return nil, err
    }
    return clients, nil
}

func (r *ClientRepositoryImpl) Update(client *entities.Client) error {
    return r.DB.Save(client).Error
}

func (r *ClientRepositoryImpl) Delete(id uint) error {
    return r.DB.Delete(&entities.Client{}, id).Error
}