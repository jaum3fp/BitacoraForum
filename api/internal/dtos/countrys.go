package dtos

type CountryDTO struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Flag      string `json:"flag"`
	Name      string `json:"name"`
}

type CountryTotalPostsDTO struct {
	CountryDTO
	TotalPosts int64 `json:"total_posts"`
}
