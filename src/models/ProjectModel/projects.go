package projectmodel

import (
	"be_hiring_app/src/config"
	"mime/multipart"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title       string
	Photo       string`json:"url,omitempty" validate:"required"`
	Repository  string
	ProjectType string
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

func SelectAllProject() []*Project {
	var items []*Project
	config.DB.Preload("User").Preload("Recipe").Find(&items)
	return items
}

func SelectProjectById(id string) *Project {
	var item Project
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostProject(item *Project) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateProject(id int, newProject *Project) error {
	var item Project
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newProject)
	return result.Error
}

func DeleteProject(id int) error {
	var item Project
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Project {
	var items []*Project
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Project {
	var items []*Project
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Project{}).Count(&count)
	return count
}
