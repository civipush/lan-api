package model

import (
	"github.com/jinzhu/gorm"
)

// TApi 测试用模型
type TApi struct {
	gorm.Model
	A int
	B string `gorm:"size:100"`
}

func GetTApi(tApi *TApi, id string) error {
	return DB.Find(tApi, id).Error
}

func GetAllTApi(offset, limit int64) (tApis []TApi, err error) {
	err = DB.Offset(offset).Limit(limit).Find(&tApis).Error
	return tApis, err
}

func CreateTApi(tApi *TApi) error {
	return DB.Create(tApi).Error
}

func UpdateTApi(tApi *TApi, attr map[string]interface{}) error {
	if attr == nil {
		return DB.Save(tApi).Error
	}
	return DB.Model(tApi).Update(attr).Error
}

func DeleteTApi(tApi *TApi) error {
	return DB.Delete(tApi).Error
}
