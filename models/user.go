package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:varchar(36)"`
	Username string    `json:"username" gorm:"type:varchar(255);unique"`
	Email    string    `json:"email" gorm:"type:varchar(255);unique"`
	Password string    `json:"password" gorm:"text"`
}
type UserLogin struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
