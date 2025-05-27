package dtos

import "github.com/jaum3fp/bitacora-forum/internal/models"

type UserDTO struct {
	ID        		uint  	`json:"id"`
	CreatedAt 		string	`json:"created_at"`
	UpdatedAt 		string	`json:"updated_at"`
	Username  		string	`json:"username"`
	Password  		string	`json:"password"`
	Email     		string	`json:"email"`
	Name      		string	`json:"name"`
	Surnames  		string	`json:"surnames"`
	ProfileImage	string	`json:"profile_img"`
}

type UserWithourPasswordDTO struct {
	ID        		uint  	`json:"id"`
	CreatedAt 		string	`json:"created_at"`
	UpdatedAt 		string	`json:"updated_at"`
	Username  		string	`json:"username"`
	Email     		string	`json:"email"`
	Name      		string	`json:"name"`
	Surnames  		string	`json:"surnames"`
	ProfileImage	string	`json:"profile_img"`
}

func (u *UserDTO) ParseWithoutPassword() UserWithourPasswordDTO {
	return UserWithourPasswordDTO{
		ID: u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Username: u.Username,
		Email: u.Email,
		Name: u.Name,
		Surnames: u.Surnames,
		ProfileImage: u.ProfileImage,
	}
}

func (u *UserDTO) ParseToModel() *models.User {
	return &models.User{
		Username: u.Username,
		Password: u.Password,
		Email: u.Email,
		Name: u.Name,
		Surnames: u.Surnames,
	}
}
