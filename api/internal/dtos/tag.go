package dtos

type TagDTO struct {
	Name  string    `json:"name"`
	Posts []PostDTO `json:"posts"`
}
