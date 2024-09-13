package dto

type UpdatePostDto struct {
	Title  string `json:"title" validate:"omitempty,min=10,max=255"`
	Body   string `json:"body" validate:"omitempty,min=10"`
	Tags   []uint `json:"tags" validate:"omitempty,dive,required,gt=0,min=1"`
	UserID uint   `json:"-"`
}
