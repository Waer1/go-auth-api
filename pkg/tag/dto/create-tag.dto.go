package dto

type CreateTagDto struct {
	Name string `json:"name" binding:"required,min=3,max=255"`
}
