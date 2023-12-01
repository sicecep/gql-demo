package repository

import (
	"github.com/sicecep/gql-demo/app/models"
	"gorm.io/gorm"
)

type CountryRepository interface {
	CreateCountry(name string) (*models.Country, error)
	GetCountryByID(id int) (*models.Country, error)
}

type CountryService struct {
	Db *gorm.DB
}

var _ CountryRepository = &CountryService{}

func NewCountryService(db *gorm.DB) *CountryService {
	return &CountryService{
		Db: db,
	}
}

func (b *CountryService) CreateCountry(name string) (*models.Country, error) {
	country := &models.Country{
		Name: name,
	}
	err := b.Db.Create(&country).Error

	return country, err
}

func (b *CountryService) GetCountryByID(id int) (*models.Country, error) {
	country := &models.Country{}
	err := b.Db.Where("id = ?", id).First(country).Error
	return country, err
}
