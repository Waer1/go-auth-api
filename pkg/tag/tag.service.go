package tag

import (
	"api-auth/pkg/tag/dto"
	"api-auth/utils"
	"api-auth/utils/models"
	"net/http"

	"gorm.io/gorm"
)

type TagService interface {
	CreateTag(tag *dto.CreateTagDto) (*models.Tag, error)
	GetTagById(id uint) (*models.Tag, error)
	GetAllTags() ([]*models.Tag, error)
	UpdateTag(ID uint, updateTag *dto.UpdateTagDto) error
	DeleteTag(id uint) error
	GetTagsByIds(ids []uint) ([]*models.Tag, error)
	GetTagBy(tag *models.Tag) (*models.Tag, error)
}

type tagService struct {
	db *gorm.DB
}

func NewTagService(db *gorm.DB) TagService {
	return &tagService{
		db: db,
	}
}

// find tag by
func (ts *tagService) GetTagBy(tag *models.Tag) (*models.Tag, error) {
	var t models.Tag
	if err := ts.db.Where(tag).First(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

// CreateTag implements TagService.
func (t *tagService) CreateTag(tag *dto.CreateTagDto) (*models.Tag, error) {

	// check if tag already exists
	existingTag, err := t.GetTagBy(&models.Tag{Name: tag.Name})
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if existingTag != nil {
		return nil, utils.NewServiceErr(http.StatusUnprocessableEntity, map[string]string{"Name": "tag already exist"})
	}

	newTag := &models.Tag{
		Name: tag.Name,
	}

	return newTag, t.db.Create(newTag).Error
}

// DeleteTag implements TagService.
func (ts *tagService) DeleteTag(id uint) error {
	var tag models.Tag
	if err := ts.db.First(&tag, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.NewServiceErr(http.StatusNotFound, map[string]string{"message": "Tag not found"})
		}
		return err
	}
	return ts.db.Delete(&tag).Error
}

// GetAllTags implements TagService.
func (t *tagService) GetAllTags() ([]*models.Tag, error) {
	var tags []*models.Tag
	if err := t.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

// GetTagById implements TagService.
func (t *tagService) GetTagById(id uint) (*models.Tag, error) {
	var tag models.Tag

	// Query the database for the tag by ID
	if err := t.db.First(&tag, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewServiceErr(http.StatusNotFound, map[string]string{"message": "Tag not found"})
		}
		return nil, err
	}

	// Return the found tag
	return &tag, nil
}

// UpdateTag implements TagService.
func (t *tagService) UpdateTag(ID uint, updateTag *dto.UpdateTagDto) error {
	tag, err := t.GetTagById(ID)
	if err != nil {
		return err
	}

	utils.Copy(tag, updateTag)

	return t.db.Save(tag).Error
}

func (t *tagService) GetTagsByIds(ids []uint) ([]*models.Tag, error) {
	var tags []*models.Tag
	if err := t.db.Where("id IN (?)", ids).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
