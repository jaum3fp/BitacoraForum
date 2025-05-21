package repositorys

import (
	"github.com/jaum3fp/bitacora-forum/internal/dtos"
	"github.com/jaum3fp/bitacora-forum/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]dtos.UserDTO, error)
	GetUser(username string) (dtos.UserDTO, error)
	CreateUser(user dtos.UserDTO) (dtos.UserDTO, error)
	UpdateUser(user dtos.UserDTO) (dtos.UserDTO, error)
	DeleteUser(username string) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (u *userRepo) GetAllUsers() ([]dtos.UserDTO, error) {
	var users []dtos.UserDTO
	if err := u.db.Model(&models.User{}).Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u *userRepo) GetUser(username string) (dtos.UserDTO, error) {
	var user dtos.UserDTO

	if err := u.db.Model(&models.User{}).
		Where("username = ?", username).
		First(&user).Error; err != nil {

		return user, err
	}

	return user, nil
}

func (u *userRepo) CreateUser(user dtos.UserDTO) (dtos.UserDTO, error) {
	panic("unimplemented")
}

func (u *userRepo) DeleteUser(username string) error {
	panic("unimplemented")
}

func (u *userRepo) UpdateUser(user dtos.UserDTO) (dtos.UserDTO, error) {
	panic("unimplemented")
}
