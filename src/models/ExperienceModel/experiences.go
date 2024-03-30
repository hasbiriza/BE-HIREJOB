package experiencemodel

import (
	"be_hiring_app/src/config"
	"mime/multipart"

	"gorm.io/gorm"
)

type Experience struct {
	gorm.Model
	Position    string
	Photo       string`json:"url,omitempty" validate:"required"`
	StartDate   string
	EndDate     string
	Description string
	UserId      uint
	User        User
}

type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type User struct {
	gorm.Model
	Name        string
	Email       string
	PhoneNumber string
	Address     string
	Photo       string
	Role        string
	Description string
	Instagram   string
	Github      string
	Linkedin    string
}

func SelectAllExperience() []*Experience {
	var items []*Experience
	config.DB.Preload("User").Find(&items)
	return items
}

func SelectExperienceById(id string) *Experience {
	var item Experience
	config.DB.Preload("User").First(&item, "id = ?", id)
	return &item
}

func PostExperience(item *Experience) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateExperience(id int, newExperience *Experience) error {
	var item Experience
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newExperience)
	return result.Error
}

func DeleteExperience(id int) error {
	var item Experience
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Experience {
	var items []*Experience
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Experience {
	var items []*Experience
	config.DB.Order(sort).Limit(limit).Offset(offset).Preload("User").Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Experience{}).Count(&count)
	return count
}
