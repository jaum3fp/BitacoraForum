package dtos

type UserDTO struct {
	ID        uint      `json:"id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Surnames  string    `json:"surnames"`
}
