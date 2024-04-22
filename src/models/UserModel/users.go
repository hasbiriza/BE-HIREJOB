package usermodel

import (
	"be_hiring_app/src/config"
	"mime/multipart"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string
	Email       string
	Password    string
	PhoneNumber string
	Address     string
	Photo       string`json:"photo,omitempty" validate:"required"`
	Role        string
	Description string
	UserToken   string
	Instagram   string
	Github      string
	Linkedin    string
}

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

func SelectAllUser() []*User {
	var items []*User
	config.DB.Find(&items)
	return items
}

func SelectUserById(id string) *User {
	var item User
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostUser(item *User) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateUser(id int, newUser *User) error {
	var item User
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newUser)
	return result.Error
}

func DeleteUser(id int) error {
	var item User
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindEmail(email string) (User, error) {
	var item User
	if err := config.DB.Where("email = ?", email).First(&item).Error; err != nil {
		return item, err
	}
	return item, nil
}

func FindRole(email string) (string, error) {
	user, err := FindEmail(email)
	if err != nil {
		return "", err
	}
	return user.Role, nil
}

func FindData(keyword string) []*User {
	var items []*User
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ? OR email LIKE ? OR phone_number LIKE ? OR address LIKE ? OR photo LIKE ? OR role LIKE ? OR description LIKE ? OR user_token LIKE ? OR instagram LIKE ? OR github LIKE ? OR linkedin LIKE ?", 
		keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword, keyword).Find(&items)
	return items
}


func FindCond(sort string, limit int, offset int) []*User {
	var items []*User
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&User{}).Count(&count)
	return count
}
