package post

import (
	"api-auth/pkg/post/dto"
	"api-auth/pkg/tag"
	"api-auth/utils"
	"api-auth/utils/models"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type PostService interface {
	CreatePost(createPostDto *dto.CreatePostDto) error
	GetPostById(id uint) (*models.Post, error)
	GetAllPosts() ([]*models.Post, error)
	UpdatePost(id uint, updatePostDto *dto.UpdatePostDto) error
	DeletePost(id uint, userID uint) error
	GetPostBy(post *models.Post) (*models.Post, error)
}

type postService struct {
	db         *gorm.DB
	tagService tag.TagService
}

func NewPostService(db *gorm.DB, tagService tag.TagService) PostService {
	return &postService{
		db:         db,
		tagService: tagService,
	}
}

func (ps *postService) GetPostBy(post *models.Post) (*models.Post, error) {
	var result models.Post
	err := ps.db.Where(post).First(&result).Error

	// if the post is not found, return utils.NewServiceErr
	if err != nil {
		return nil, utils.NewServiceErr(http.StatusNotFound, map[string]string{
			"post": "not found",
		})
	}

	return &result, nil
}

func (ps *postService) CreatePost(createPostDto *dto.CreatePostDto) error {
	// check if tags exists
	tags, err := ps.tagService.GetTagsByIds(createPostDto.Tags)
	if err != nil {
		return err
	}

	// if number of returned tags is less than the number of tags in the request, return error
	if len(tags) < len(createPostDto.Tags) {
		return utils.NewServiceErr(http.StatusBadRequest, map[string]string{
			"tags": "some tags are not exist",
		})
	}

	post := &models.Post{
		Title:  createPostDto.Title,
		Body:   createPostDto.Body,
		UserID: createPostDto.UserID,
		Tags:   make([]models.Tag, len(tags)),
	}
	for i, t := range tags {
		post.Tags[i] = *t
	}
	return ps.db.Create(post).Error
}

func (ps *postService) GetPostById(id uint) (*models.Post, error) {
	var post models.Post
	post.ID = id
	var result models.Post
	err := ps.db.Preload("Tags").Where(post).First(&result).Error
	fmt.Println("Error:", err)
	// if the post is not found, return utils.NewServiceErr
	if err != nil {
		return nil, utils.NewServiceErr(http.StatusNotFound, map[string]string{
			"post": "not found",
		})
	}

	return &result, nil
}

func (ps *postService) GetAllPosts() ([]*models.Post, error) {
	var posts []*models.Post
	if err := ps.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (ps *postService) UpdatePost(id uint, updatePostDto *dto.UpdatePostDto) error {
	post, err := ps.GetPostById(id)
	if err != nil {
		return err
	}

	// validate the owner of the post is the one who try to update it
	if post.UserID != updatePostDto.UserID {
		return utils.NewServiceErr(http.StatusForbidden, map[string]string{"message": "you are not allowed to update this post"})
	}

	utils.Copy(post, updatePostDto)

	return ps.db.Save(post).Error
}

func (ps *postService) DeletePost(id uint, userID uint) error {
	post, err := ps.GetPostById(id)
	if err != nil {
		return err
	}

	// vaalidate the owner of the post is the one who try to delete
	if post.UserID != userID {
		return utils.NewServiceErr(http.StatusForbidden, map[string]string{"message": "you are not allowed to update this post"})
	}

	return ps.db.Delete(post).Error
}
