package repositorys

import (
	token "github.com/jaum3fp/bitacora-forum/internal/auth"
	"github.com/jaum3fp/bitacora-forum/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Login(username string, password string) (string, error)
	Register(username string, password string) (string, error)
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

	tkn, err := token.GenerateUserToken(models.User{
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Name:     user.Name,
		Surnames: user.Surnames,
	})
	if err != nil {
		return "", err
	}

	return tkn, nil
}

func (a *authRepo) Register(username string, password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	res := a.db.Create(&user)
	if res.Error != nil {
		return "", res.Error
	}

	return user.Username, nil
}
