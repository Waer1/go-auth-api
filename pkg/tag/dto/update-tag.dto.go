package dto

type UpdateTagDto struct {
	Name string `json:"name" binding:"omitempty,min=3,max=255"`
}
