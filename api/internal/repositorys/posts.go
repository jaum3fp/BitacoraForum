package repositorys

import (
	"strconv"
	"strings"

	"gorm.io/gorm"

	"github.com/jaum3fp/bitacora-forum/internal/dtos"
	"github.com/jaum3fp/bitacora-forum/internal/models"
)

type PostRepository interface {
	GetAllPosts(filters map[string]string) ([]dtos.PostUsernameDTO, error)
	GetPost(id string) (dtos.PostUsernameDTO, error)
	CreatePost(post dtos.PostDTO) error
	UpdatePost(id string, post dtos.PostUsernameDTO) error
	GetCountryPosts(flag string) ([]dtos.PostDTO, error)
	GetCountryPostsNumber(flag string) (int64, error)
	GetPostCommentsNumber(super string) (int64, error)
	DeletePost(id string) error
	IncrementPostViews(id string) error
	GetPostsByTag(id string) ([]dtos.PostDTO, error)
	GetPostComments(super string) ([]dtos.PostUsernameDTO, error)
	GetTotalPosts(filters map[string]string) (total int64, err error)
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
		Select("posts.*, users.username AS owner_username, users.profile_image AS owner_avatar").
		Joins("JOIN users ON posts.owner_id = users.id").
		Where("posts.country_alpha = ?", flag).
		Find(&posts).Error; err != nil {

		return posts, err
	}

	return posts, nil
}

func (r *postRepo) GetPostCommentsNumber(super string) (int64, error) {
	var posts int64
	if err := r.db.Model(&models.Post{}).
		Where("super_id = ?", super).
		Count(&posts).Error; err != nil {

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

	subQuery := r.db.
		Select("COUNT(*)").
		Table("posts AS comment").
		Where("comment.super_id = posts.id")

	query := r.db.Model(&models.Post{}).
		Select("posts.*, users.username AS owner_username, users.profile_image AS owner_avatar, (?) AS comments_total", subQuery).
		Joins("JOIN users ON posts.owner_id = users.id").
		Where("posts.super_id IS NULL")

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

	if startRange, ok := filters["start"]; ok {
		if endRange, ok := filters["end"]; ok {
			start, err1 := strconv.Atoi(startRange)
			end, err2 := strconv.Atoi(endRange)
			if err1 == nil && err2 == nil && end > start {
				query = query.Offset(start).Limit(end - start)
			}
		}
	}

	var posts []dtos.PostUsernameDTO
	if err := query.Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, nil
}

func (r *postRepo) GetTotalPosts(filters map[string]string) (total int64, err error) {
	query := r.db.Model(&models.Post{}).
		Joins("JOIN users ON posts.owner_id = users.id").
		Where("posts.super_id IS NULL")

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

	if err = query.Count(&total).Error; err != nil {
		return
	}
	return
}

func (r *postRepo) GetPostComments(super string) ([]dtos.PostUsernameDTO, error) {
	query := r.db.Model(&models.Post{}).
		Select("posts.*, users.username AS owner_username, users.profile_image AS owner_avatar").
		Joins("JOIN users ON posts.owner_id = users.id").
		Where("posts.super_id = ?", super)

	var posts []dtos.PostUsernameDTO
	if err := query.Find(&posts).Error; err != nil {
		return posts, err
	}

	return posts, nil
}


func (r *postRepo) GetPost(id string) (dtos.PostUsernameDTO, error) {

	var post dtos.PostUsernameDTO

	subQuery := r.db.
		Select("COUNT(*)").
		Table("posts AS comment").
		Where("comment.super_id = posts.id")

	if err := r.db.Model(&models.Post{}).
		Select("posts.*, users.username AS owner_username, users.profile_image AS owner_avatar, (?) AS comments_total", subQuery).
		Joins("JOIN users ON posts.owner_id = users.id").
		Where("posts.super_id IS NULL").
		Find(&post, id).Error; err != nil {

		return post, err
	}

	return post, nil
}

func (r *postRepo) CreatePost(post dtos.PostDTO) error {

	err := r.db.Create(&models.Post{
		Title:      	post.Title,
		Description:	post.Description,
		Content:    	post.Content,
		OwnerID:    	post.OwnerID,
		CountryAlpha:	post.CountryAlpha,
		SuperID: 		post.SuperID,
	}).Error
	return err
}

func (r *postRepo) UpdatePost(id string, post dtos.PostUsernameDTO) error {

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
