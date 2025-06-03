package repositorys

import (
	"fmt"

	token "github.com/jaum3fp/bitacora-forum/internal/auth"
	"github.com/jaum3fp/bitacora-forum/internal/dtos"
	"github.com/jaum3fp/bitacora-forum/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(username string, password string) (string, error)
	Register(user dtos.UserDTO) (string, error)
}

type authRepo struct {
	db   *gorm.DB
	user UserRepository
}

func NewAuthRepository(db *gorm.DB, user UserRepository) AuthRepository {
	return &authRepo{db, user}
}

func (a *authRepo) Login(username string, password string) (string, error) {

	user, err := a.user.GetUser(username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	tkn, err := token.GenerateUserToken(getUserModel(user))
	if err != nil {
		return "", err
	}

	return tkn, nil
}

func (a *authRepo) Register(user dtos.UserDTO) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user.Password = string(hashedPassword)

	res, err := a.user.CreateUser(user)
	if err != nil {
		return "", err
	}

	fmt.Println(res)

	tkn, err := token.GenerateUserToken(*user.ParseToModel())
	if err != nil {
		return "", err
	}

	return tkn, nil
}

func getUserModel(user dtos.UserDTO) (model models.User) {
	model = models.User{
		Username: user.Username,
		Email:    user.Email,
		Name:     user.Name,
		Surnames: user.Surnames,
	}
	return
}

func getUserToken(user dtos.UserDTO) (tkn string, err error) {
	userModel := getUserModel(user)
	tkn, err = token.GenerateUserToken(userModel)
	return
}
