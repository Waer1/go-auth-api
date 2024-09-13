package models

type Tag struct {
	BaseEntity
	Name  string `json:"name" validate:"required,min=3,max=255"`
	Posts []Post `gorm:"many2many:post_tags" json:"posts"`
}
