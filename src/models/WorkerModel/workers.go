package workermodel

import (
	"be_hiring_app/src/config"

	"gorm.io/gorm"
)

type Worker struct {
	gorm.Model
	JobDesc     string
	JobType     string
	CompanyName string
	Skill       string
	UserId      uint
	User        User
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

func SelectAllWorker() []*Worker {
	var items []*Worker
	config.DB.Preload("User").Find(&items)
	return items
}

func SelectWorkerById(id string) *Worker {
	var item Worker
	config.DB.Preload("User").First(&item, "id = ?", id)
	return &item
}

func PostWorker(item *Worker) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateWorker(id int, newWorker *Worker) error {
	var item Worker
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newWorker)
	return result.Error
}

func DeleteWorker(id int) error {
	var item Worker
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Worker {
	var items []*Worker
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Worker {
	var items []*Worker
	config.DB.Order(sort).Limit(limit).Offset(offset).Preload("User").Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Worker{}).Count(&count)
	return count
}
