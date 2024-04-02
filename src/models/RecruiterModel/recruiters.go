package recruitermodel

import (
	"be_hiring_app/src/config"

	"gorm.io/gorm"
)

type Recruiter struct {
	gorm.Model
	CompanyName string
	Position    string
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
	Linkedin    string
}

func SelectAllRecruiter() []*Recruiter {
	var items []*Recruiter
	config.DB.Preload("User").Preload("Recipe").Find(&items)
	return items
}

func SelectRecruiterById(id string) *Recruiter {
	var item Recruiter
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostRecruiter(item *Recruiter) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateRecruiter(id int, newRecruiter *Recruiter) error {
	var item Recruiter
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newRecruiter)
	return result.Error
}

func DeleteRecruiter(id int) error {
	var item Recruiter
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Recruiter {
	var items []*Recruiter
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR company_name LIKE ? OR CAST(user_id AS TEXT) LIKE ? ", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Recruiter {
	var items []*Recruiter
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Recruiter{}).Count(&count)
	return count
}
