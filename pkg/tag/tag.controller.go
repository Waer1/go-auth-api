package tag

import (
	"api-auth/pkg/tag/dto"
	"api-auth/utils"
	"api-auth/utils/parsing"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagController interface {
	CreateTag(c *gin.Context)
	GetTag(c *gin.Context)
	GetAllTags(c *gin.Context)
	UpdateTag(c *gin.Context)
	DeleteTag(c *gin.Context)
}

type tagController struct {
	tagService TagService
}

func NewTagController(tagService TagService) TagController {
	return &tagController{
		tagService: tagService,
	}
}

func (tc *tagController) CreateTag(c *gin.Context) {
	var tagDto dto.CreateTagDto
	if !utils.BindJSONAndValidate(c, &tagDto) {
		return
	}

	tag, err := tc.tagService.CreateTag(&tagDto)
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}
	c.JSON(http.StatusCreated, tag)
}

func (tc *tagController) GetTag(c *gin.Context) {
	tagId, err := parsing.GetParamUint(c, "id")
	if err != nil {
		return
	}

	tag, err := tc.tagService.GetTagById(uint(tagId))
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}
	c.JSON(http.StatusOK, tag)
}

func (tc *tagController) GetAllTags(c *gin.Context) {
	tags, err := tc.tagService.GetAllTags()
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}
	c.JSON(http.StatusOK, tags)
}

func (tc *tagController) UpdateTag(c *gin.Context) {
	tagId, err := parsing.GetParamUint(c, "id")
	if err != nil {
		return
	}

	var tagDto dto.UpdateTagDto
	if !utils.BindJSONAndValidate(c, &tagDto) {
		return
	}
	err = tc.tagService.UpdateTag(uint(tagId), &tagDto)
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tag updated successfully"})
}

func (tc *tagController) DeleteTag(c *gin.Context) {
	tagId, err := parsing.GetParamUint(c, "id")
	if err != nil {
		return
	}

	err = tc.tagService.DeleteTag(uint(tagId))
	if err != nil {
		c.Error(err) // Add error to Gin Context
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}
