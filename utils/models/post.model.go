package models

type Post struct {
	BaseEntity
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID uint   `json:"-"`
	User   User   `json:"-"`
	Tags   []Tag  `gorm:"many2many:post_tags" json:"tags"`
}
