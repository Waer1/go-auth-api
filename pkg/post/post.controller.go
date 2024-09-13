package post

import (
	"api-auth/pkg/post/dto"
	"api-auth/utils"
	"api-auth/utils/helpers"
	"api-auth/utils/parsing"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController interface {
	CreatePost(c *gin.Context)
	GetPostById(c *gin.Context)
	// GetPostsByTag(c *gin.Context)
	GetAllPosts(c *gin.Context)
	UpdatePost(c *gin.Context)
	DeletePost(c *gin.Context)
}

type postController struct {
	postService PostService
}

func NewPostController(postService PostService) PostController {
	return &postController{
		postService: postService,
	}
}

func (pc *postController) CreatePost(c *gin.Context) {
	var postDto dto.CreatePostDto
	if !utils.BindJSONAndValidate(c, &postDto) {
		return
	}

	// get the current user
	user, err := helpers.GetCurrentUser(c)
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}
	postDto.UserID = user.UserId

	err = pc.postService.CreatePost(&postDto)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

func (pc *postController) GetPostById(c *gin.Context) {
	id, err := parsing.GetParamUint(c, "id")
	if err != nil {
		return
	}

	post, err := pc.postService.GetPostById(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, post)
}

func (pc *postController) GetAllPosts(c *gin.Context) {
	posts, err := pc.postService.GetAllPosts()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, posts)
}

func (pc *postController) UpdatePost(c *gin.Context) {
	var postDto dto.UpdatePostDto
	if !utils.BindJSONAndValidate(c, &postDto) {
		return
	}

	postId, err := parsing.GetParamUint(c, "id")
	if err != nil {
		return
	}

	// get the current user
	user, err := helpers.GetCurrentUser(c)
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}
	postDto.UserID = user.UserId

	err = pc.postService.UpdatePost(postId, &postDto)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{"message": "Post updated successfully"})
}

func (pc *postController) DeletePost(c *gin.Context) {
	postId, err := parsing.GetParamUint(c, "id")
	if err != nil {
		return
	}

	// get the current user
	user, err := helpers.GetCurrentUser(c)
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}

	err = pc.postService.DeletePost(postId, user.UserId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}
