package skillmodel

import (
	"be_hiring_app/src/config"

	"gorm.io/gorm"
)

type Skill struct {
	gorm.Model
	Name   string
	UserId uint
	User   User
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

func SelectAllSkill() []*Skill {
	var items []*Skill
	config.DB.Find(&items)
	return items
}

func SelectSkillById(id string) *Skill {
	var item Skill
	config.DB.First(&item, "id = ?", id)
	return &item
}

func PostSkill(item *Skill) error {
	result := config.DB.Create(&item)
	return result.Error
}

func UpdateSkill(id int, newSkill *Skill) error {
	var item Skill
	result := config.DB.Model(&item).Where("id = ?", id).Updates(newSkill)
	return result.Error
}

func DeleteSkill(id int) error {
	var item Skill
	result := config.DB.Delete(&item, "id = ?", id)
	return result.Error
}

func FindData(keyword string) []*Skill {
	var items []*Skill
	keyword = "%" + keyword + "%"
	config.DB.Where("CAST(id AS TEXT) LIKE ? OR name LIKE ? OR CAST(day AS TEXT) LIKE ?", keyword, keyword, keyword).Find(&items)
	return items
}

func FindCond(sort string, limit int, offset int) []*Skill {
	var items []*Skill
	config.DB.Order(sort).Limit(limit).Offset(offset).Find(&items)
	return items
}

func CountData() int64 {
	var count int64
	config.DB.Model(&Skill{}).Count(&count)
	return count
}
