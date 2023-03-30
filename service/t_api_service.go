package service

import (
	"lan_api/model"
	"lan_api/serializer"
)

// TApiService 测试服务
type TApiService struct {
	Id uint
	A  int    `binding:"required,min=2,max=30"`
	B  string `binding:"required,min=2,max=30"`
}

func (c *TApiService) Valid() *serializer.Response {
	if c.A == 10 {
		return &serializer.Response{
			Code:  41000,
			Msg:   "不能等于10",
			Error: "== 10!",
		}
	}
	return nil
}

func (c *TApiService) Get(id string) *serializer.Response {
	var tApi model.TApi
	if err := model.GetTApi(&tApi, id); err != nil {
		return &serializer.Response{
			Code:  serializer.CodeCreateError,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}
	return serializer.BuildResponseGet(serializer.BuildTApi(&tApi))
}

func (c *TApiService) GetList(offset, limit int64) *serializer.Response {
	tApis, err := model.GetAllTApi(offset, limit)
	if err != nil {
		return &serializer.Response{
			Code:  serializer.CodeCreateError,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}
	items, total := serializer.BuildTApis(tApis)
	return serializer.BuildListResponse("", items, total)
}

func (c *TApiService) Create() *serializer.Response {
	tApi := model.TApi{
		A: c.A,
		B: c.B,
	}

	if err := model.CreateTApi(&tApi); err != nil {
		return &serializer.Response{
			Code:  serializer.CodeCreateError,
			Msg:   "新增失败",
			Error: err.Error(),
		}
	}

	return &serializer.Response{
		Code: serializer.CodeCreateSuccess,
		Msg:  "新增成功",
		Data: tApi,
	}
}

func (c *TApiService) Update(id string) *serializer.Response {
	var tApi model.TApi
	if err := model.GetTApi(&tApi, id); err != nil {
		return &serializer.Response{
			Code:  serializer.CodeUpdateError,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}
	tApi.A = c.A
	tApi.B = c.B

	// 更新全部
	if err := model.UpdateTApi(&tApi, nil); err != nil {
		return &serializer.Response{
			Code:  serializer.CodeUpdateError,
			Msg:   "更新失败",
			Error: err.Error(),
		}
	}

	// 更新部分
	//attr := map[string]interface{}{
	//	"a": c.A,
	//}
	//if err := model.UpdateTApi(&tApi, attr); err != nil {
	//	return &serializer.Response{
	//		Code:  serializer.CodeUpdateError,
	//		Msg:   "更新失败",
	//		Error: err.Error(),
	//	}
	//}
	return &serializer.Response{
		Code: serializer.CodeUpdateSuccess,
		Msg:  "更新成功",
		Data: tApi,
	}
}

func (c *TApiService) Delete(id string) *serializer.Response {
	var tApi model.TApi
	if err := model.GetTApi(&tApi, id); err != nil {
		return &serializer.Response{
			Code:  serializer.CodeDeleteError,
			Msg:   "查询失败",
			Error: err.Error(),
		}
	}

	if err := model.DeleteTApi(&tApi); err != nil {
		return &serializer.Response{
			Code:  serializer.CodeDeleteError,
			Msg:   "删除失败",
			Error: err.Error(),
		}
	}
	return &serializer.Response{
		Code: serializer.CodeDeleteSuccess,
		Msg:  "删除成功",
	}
}
