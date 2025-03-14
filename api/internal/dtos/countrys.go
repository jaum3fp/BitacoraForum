package dtos

type CountryDTO struct {
	ID        uint      `json:"id"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	Flag      string    `json:"flag"`
	Name      string    `json:"name"`
	Posts     []PostDTO `json:"posts"`
}

type CountryTotalPostsDTO struct {
	CountryDTO
	TotalPosts int64 `json:"total_posts"`
}
