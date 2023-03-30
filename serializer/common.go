package serializer

import "github.com/gin-gonic/gin"

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`            // 请求接口的结果状态
	Data  interface{} `json:"data,omitempty"`  // 主要的返回内容，多种格式
	Msg   string      `json:"msg"`             // 直接显示出来请求接口的情况
	Error string      `json:"error,omitempty"` // 调试用的，内部检查问题
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 五开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 四开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	CodeSuccess       = 20000 // 查询成功
	CodeCreateSuccess = 20001 // 新增成功
	CodeUpdateSuccess = 20002 // 更新成功
	CodeDeleteSuccess = 20003 // 删除成功
	CodeCheckError    = 40000 // 查询错误
	CodeCreateError   = 40001 // 新增错误
	CodeUpdateError   = 40002 // 更新错误
	CodeDeleteError   = 40003 // 删除错误

	// CodeCheckLogin 未登录
	CodeCheckLogin = 401
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403
	// CodeDBError 数据库操作失败
	CodeDBError = 50001
	// CodeEncryptError 加密失败
	CodeEncryptError = 50002
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40001
)

// DataList 基础列表结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// BuildResponse 一般请求构建器
func BuildResponse(data interface{}) *Response {
	return &Response{
		Data:  data,
		Code:  CodeSuccess,
		Msg:   "请求成功",
		Error: "",
	}
}

// BuildResponseGet 查询单行数据构建器
func BuildResponseGet(data interface{}) *Response {
	return &Response{
		Data:  data,
		Code:  CodeSuccess,
		Msg:   "查询成功",
		Error: "",
	}
}

// BuildListResponse 查询列表构建器
func BuildListResponse(err string, items interface{}, total uint) *Response {
	return &Response{
		Data: DataList{
			Items: items,
			Total: total,
		},
		Code:  CodeSuccess,
		Msg:   "查询成功",
		Error: err,
	}
}

// CheckLogin 检查登录
func CheckLogin() Response {
	return Response{
		Code: CodeCheckLogin,
		Msg:  "未登录",
	}
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}
