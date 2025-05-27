package repositorys

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jaum3fp/bitacora-forum/internal/dtos"
	"github.com/jaum3fp/bitacora-forum/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() ([]dtos.UserDTO, error)
	GetUser(username string) (dtos.UserDTO, error)
	CreateUser(user dtos.UserDTO) (dtos.UserDTO, error)
	UpdateUser(username string, user dtos.UserWithourPasswordDTO) (error)
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
	if err := u.db.Create(user.ParseToModel()).Error; err != nil {
		return dtos.UserDTO{}, err
	}
	return user, nil
}

func (u *userRepo) DeleteUser(username string) error {
	panic("unimplemented")
}

func (u *userRepo) UpdateUser(username string, user dtos.UserWithourPasswordDTO) error {
	updates := map[string]any{}

	if user.Email != "" {
		updates["email"] = user.Email
	}
	if user.Name != "" {
		updates["name"] = user.Name
	}
	if user.Surnames != "" {
		updates["surnames"] = user.Surnames
	}
	if user.ProfileImage != "" {
		println("Writing image...")
		fileName, err := storeBase64Image(user.Username, user.ProfileImage)
		if err != nil {
			println(err)
			return err
		}
		updates["profile_image"] = fileName
	}

	if len(updates) == 0 {
		return errors.New("no hay campos para actualizar")
	}

	if err := u.db.Model(&models.User{}).
		Where("username = ?", username).
		Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func storeBase64Image(owner string, encodedImage string) (string, error) {

	cwd, err := os.Getwd()
	if err != nil {
		return "", errors.New("Cwd not found")
	}
	profileImagesPath := filepath.Join(cwd, "internal", "static", "images", "avatars")
	splitedEncodedImage := strings.SplitN(encodedImage, ",", 2)
	if len(splitedEncodedImage) != 2 {
		return "", errors.New("Incorrect format")
	}

	data, err := base64.StdEncoding.DecodeString(splitedEncodedImage[1])
	if err != nil {
		return "", errors.New("Image decode failed")
	}

	if err := os.MkdirAll(profileImagesPath, 0755); err != nil {
		return "", errors.New("Avatar folder not created")
	}

	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s_%s.png", timestamp, owner)
	filePath := filepath.Join(profileImagesPath, filename)

	if err = os.WriteFile(filePath, data, 0744); err != nil {
		return "", errors.New("Image writing failed")
	}

	return filename, nil
}
