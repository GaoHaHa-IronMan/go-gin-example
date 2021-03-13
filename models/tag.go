package models

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (tag *Tag) BeforeCreate(tx *gorm.DB) error {

	tag.CreatedOn = time.Now().Unix()
	return nil
}

func (tag *Tag) BeforeUpdate(tx *gorm.DB) error {
	tag.ModifiedOn = time.Now().Unix()

	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {

	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return

}

func GetTagTotal(maps interface{}) int {
	var res int64
	db.Model(&Tag{}).Where(maps).Count(&res)
	return int(res)
}

func ExistTagByName(name string) bool {

	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if err != nil {
		return false
	}
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {

	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if err != nil {
		return false
	}
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createdBy string) bool {
	err := db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error
	if err != nil {
		return false
	}

	return true
}

func EditTag(id int, data map[string]interface{}) bool {
	err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return false
	}
	return true
}

func DeleteTag(id int) bool {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		return false
	}

	return true
}
