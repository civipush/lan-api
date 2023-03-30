package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v9"
	"lan_api/conf"
	"lan_api/model"
	"lan_api/serializer"
	"lan_api/service"
)

// Ping 状态检查页面
func Ping(c *gin.Context) {
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "Pong",
	})
	//c.JSON(200, serializer.BuildResponse("pong pong"))
	//c.JSON(200, serializer.BuildListResponse(40000, "返回成功", "error l", nil, 0))

}

// TApi测试接口
func TApiGet(c *gin.Context) {
	var service service.TApiService
	c.JSON(200, service.Get(c.Param("id")))
}

// TApi测试接口获取列表
func TApiGetList(c *gin.Context) {
	var service service.TApiService
	c.JSON(200, service.GetList(0, 2))
}

// TApi 新增接口
func TApiPost(c *gin.Context) {
	var service service.TApiService
	if err := c.ShouldBind(&service); err == nil {
		if valid := service.Valid(); valid != nil {
			c.JSON(200, valid)
			return
		}
		// 新增
		c.JSON(200, service.Create())
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// TApi 修改接口
func TApiPut(c *gin.Context) {
	var service service.TApiService
	if err := c.ShouldBind(&service); err == nil {
		if valid := service.Valid(); valid != nil {
			c.JSON(200, valid)
			return
		}
		// 修改
		c.JSON(200, service.Update(c.Param("id")))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// TApi测试接口删除
func TApiDelete(c *gin.Context) {
	var service service.TApiService
	c.JSON(200, service.Delete(c.Param("id")))
}

// CurrentUser 获取当前用户
func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("Field.%s", e.Field()))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
			return serializer.ParamErr(
				fmt.Sprintf("%s%s", field, tag),
				err,
			)
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}
