package dtos

type PostDTO struct {
	ID       		uint   					`json:"id"`
	CreatedAt 		string 					`json:"created_at"`
	UpdatedAt 		string 					`json:"updated_at"`
	Title    		string 					`json:"title"`
	Description		string					`json:"description"`
	Content   		string 					`json:"content"`
	Views     		int64  					`json:"views"`
	OwnerID   		uint  					`json:"owner_id"`
	CountryAlpha 	string 					`json:"country_alpha"`
	CommentsTotal	int64					`json:"comments_total"`
	SuperID			*uint					`json:"super_id"`
}

func (PostDTO) TableName() string {
	return "posts"
}

type PostUsernameDTO struct {
	PostDTO
	OwnerUsername string `json:"owner_username"`
	OwnerAvatar   string `json:"owner_avatar"`
}
