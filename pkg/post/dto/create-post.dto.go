package dto

type CreatePostDto struct {
	Title  string `json:"title" binding:"required,min=10,max=255"`
	Body   string `json:"body" binding:"required,min=10"`
	Tags   []uint `json:"tags" binding:"required,min=1,dive,number"`
	UserID uint   `json:"-"`
}
