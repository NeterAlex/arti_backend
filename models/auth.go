package models

import "gorm.io/gorm"

type Auth struct {
	gorm.Model
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=18"`
	Email    string `json:"email" validate:"required,email"`
	Avatar   string `json:"avatar"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}

func AddUser(username, password, email string) bool {
	db.Create(&Auth{Username: username, Password: password, Email: email})
	return true
}

func GetUserByUsername(username string) Auth {
	var auth Auth
	db.Select("id").Where("username = ?", username).First(&auth)
	if auth.ID > 0 {
		return auth
	}
	return Auth{}
}

func ExistUserByUsername(username string) bool {
	var auth Auth
	db.Select("id").Where("username = ?", username).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}

func DeleteUser(id int) bool {
	db.Where("id = ?", id).Delete(&Auth{})
	return true
}

func EditUser(id int, data interface{}) bool {
	db.Model(&Auth{}).Where("id = ?", id).Updates(data)
	return true
}
