package repositorys

import (
	"gorm.io/gorm"

	"github.com/jaum3fp/bitacora-forum/internal/dtos"
	"github.com/jaum3fp/bitacora-forum/internal/models"
)

type CountryRepository interface {
	GetCountry(flag string) (dtos.CountryTotalPostsDTO, error)
}

type countryRepo struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) CountryRepository {
	return &countryRepo{db}
}

func (r *countryRepo) GetCountry(flag string) (dtos.CountryTotalPostsDTO, error) {

	var country dtos.CountryTotalPostsDTO
	err := r.db.Model(&models.Country{}).
		Select("countries.*, COUNT(posts.id) AS total_posts").
		Joins("LEFT JOIN posts ON posts.country_id = countries.id").
		Where("countries.flag = ?", flag).
		Group("countries.id").
		First(&country).Error

	if err != nil {
		return country, err
	}

	return country, nil
}
