package model

import "github.com/jinzhu/gorm"

// Friend 朋友的模型
type Friend struct {
	gorm.Model
	UserId uint
}

func GetFriend(friend *Friend, id string) error {
	return DB.Find(friend, id).Error
}

func CreateFriend(friend *Friend) error {
	return DB.Create(friend).Error
}

func UpdateFriend(friend *TApi, attr map[string]interface{}) error {
	if attr == nil {
		return DB.Save(friend).Error
	}
	return DB.Model(friend).Update(attr).Error
}

func DeleteFriend(friend *Friend) error {
	return DB.Delete(friend).Error
}
