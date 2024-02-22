package models

import "github.com/google/uuid"

type Post struct {
	ID       uint      `json:"id" gorm:"primarykey"`
	UserUUID uuid.UUID `json:"userUUID" gorm:"type:varchar(36)"`
	UserName string    `json:"username" gorm:"type:varchar(255)"`
	Title    string    `json:"title" gorm:"type:varchar(255)"`
	Text     string    `json:"text" gorm:"text"`
}
type NewPostData struct {
	Title string `json:"title" gorm:"type:varchar(255)"`
	Text  string `json:"text" gorm:"text"`
}
