package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 使用方式
// response.常量.Json 或者 with 返回内容  .Json

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	Ok             = response(0, "处理成功")
	Fail           = response(1, "处理失败")
	ValidatorError = response(http.StatusUnprocessableEntity, "验证未通过")
)

// 带分页的结构
type PageData struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

func (r *Response) WithMsg(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) WithCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

func (r *Response) WithPageData(data *PageData) *Response {
	r.Data = data
	return r
}

func (r *Response) WithCodeMessage(code int, message string) *Response {
	r.Code = code
	r.Message = message
	return r
}

func (r *Response) WithMessageData(message string, data interface{}) *Response {
	r.Message = message
	r.Data = data
	return r
}

func response(code int, message string) *Response {
	return &Response{
		Code:    code,
		Message: message,
		Data:    nil,
	}
}

// 返回实例
func (r *Response) Json(c *gin.Context) {
	c.JSON(http.StatusOK, r)
	return
}
