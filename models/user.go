package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID          int    `json:"id"`
	FullName    string `json:"fullName"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Picture     string `json:"picture"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
