package serializer

import (
	"lan_api/model"
	"lan_api/util"
)

// TApi 测试接口列化器
type TApi struct {
	ID        uint   `json:"id"`
	A         int    `json:"a"`
	B         string `json:"b"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
	DeletedAt int64  `json:"deleted_at"`
}

// BuildTApi 序列化测试接口
func BuildTApi(item *model.TApi) (tApi *TApi) {
	tApi = &TApi{
		ID:        item.ID,
		A:         item.A,
		B:         item.B,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	if item.DeletedAt != nil {
		tApi.DeletedAt = item.DeletedAt.Unix()
	}
	return
}

// BuildTApis 列表序列化测试接口
func BuildTApis(items []model.TApi) ([]TApi, uint) {
	util.LogD("items", items)
	var tApis []TApi
	for _, i := range items {
		tApi := BuildTApi(&i)
		tApis = append(tApis, *tApi)
	}
	return tApis, uint(len(tApis))
	//return DataList{
	//	Items: tApis,
	//	Total: len(tApis),
	//}
}
