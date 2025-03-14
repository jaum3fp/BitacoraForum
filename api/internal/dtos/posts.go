package dtos

type PostDTO struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Views     int64  `json:"views"`
	OwnerID   uint   `json:"owner_id"`
	CountryID uint   `json:"country_id"`
}

type PostUsernameDTO struct {
	PostDTO
	OwnerUsername string `json:"owner_username"`
}
