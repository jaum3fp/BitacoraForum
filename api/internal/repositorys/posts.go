package repositorys

import (
	"strings"

	"gorm.io/gorm"

	"github.com/jaum3fp/bitacora-forum/internal/dtos"
	"github.com/jaum3fp/bitacora-forum/internal/models"
)

type PostRepository interface {
	GetAllPosts(filters map[string]string) ([]dtos.PostUsernameDTO, error)
	GetPost(id string) (dtos.PostDTO, error)
	CreatePost(post models.Post) error
	UpdatePost(id string, post models.Post) error
	GetCountryPosts(flag string) ([]dtos.PostDTO, error)
	GetCountryPostsNumber(flag string) (int64, error)
	DeletePost(id string) error
	IncrementPostViews(id string) error
	GetPostsByTag(id string) ([]dtos.PostDTO, error)
}

type postRepo struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepo{db}
}

func (r *postRepo) GetCountryPosts(flag string) ([]dtos.PostDTO, error) {
	var posts []dtos.PostDTO
	if err := r.db.Model(&models.Post{}).
		Select("posts.*, users.username AS owner_username").
		Joins("JOIN users ON posts.owner_id = users.id").
		Where("posts.country_alpha = ?", flag).
		Find(&posts).Error; err != nil {

		return posts, err
	}

	return posts, nil
}

func (r *postRepo) GetCountryPostsNumber(flag string) (int64, error) {
	var posts int64
	if err := r.db.Model(&models.Post{}).
		Where("country_alpha = ?", flag).
		Count(&posts).Error; err != nil {

		return posts, err
	}

	return posts, nil
}

func (r *postRepo) GetAllPosts(filters map[string]string) ([]dtos.PostUsernameDTO, error) {

	query := r.db.Model(&models.Post{}).
		Select("posts.*, users.username AS owner_username").
		Joins("JOIN users ON posts.owner_id = users.id")

	if title, ok := filters["title"]; ok {
		query = query.Where("posts.title LIKE ?", "%"+title+"%")
	}
	if cc, ok := filters["cc"]; ok {
		countries := strings.Split(cc, ",")
		query = query.Where("posts.country_alpha IN(?)", countries)
	}
	if author, ok := filters["author"]; ok {
		query = query.Where("users.username LIKE ?", "%"+author+"%")
	}

	var posts []dtos.PostUsernameDTO
	if err := query.Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, nil
}

func (r *postRepo) GetPost(id string) (dtos.PostDTO, error) {

	var post dtos.PostDTO
	if err := r.db.Model(&models.Post{}).Find(&post, id).Error; err != nil {
		return post, err
	}

	return post, nil
}

func (r *postRepo) CreatePost(post models.Post) error {

	err := r.db.Create(&post).Error
	return err
}

func (r *postRepo) UpdatePost(id string, post models.Post) error {

	err := r.db.Model(&models.Post{}).Where("id = ?", id).Updates(post).Error
	return err
}

func (r *postRepo) DeletePost(id string) error {

	err := r.db.Delete(&models.Post{}, id).Error
	return err
}

func (r *postRepo) IncrementPostViews(id string) error {

	err := r.db.Model(&models.Post{}).Where("id = ?", id).Update("views", gorm.Expr("views + ?", 1)).Error
	return err
}

func (r *postRepo) GetPostsByTag(id string) ([]dtos.PostDTO, error) {

	var posts []dtos.PostDTO
	if err := r.db.Model(&models.Post{}).
		Select("posts.*").
		Joins("JOIN post_tags ON posts.id = post_tags.post_id").
		Where("post_tags.tag_id = ?", id).
		Find(&posts).Error; err != nil {

		return posts, err
	}

	return posts, nil
}
