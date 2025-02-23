/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:11
 */

package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JsonResult Json返回接果
type JsonResult struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type BaseController struct {
	gin.Context
}

// Response 获取全部请求解析到map
func Response(c *gin.Context, code uint32, msg string, data map[string]interface{}) {
	message := CommResponse(code, msg, data)

	// 允许跨域
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") // 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
	c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
	c.Header("Access-Control-Allow-Credentials", "true")                                                                                                                                                   //  跨域请求是否需要带cookie信息 默认设置为true
	c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json

	// message
	c.JSON(http.StatusOK, message)
	return
}

// CommResponse 响应
func CommResponse(code uint32, message string, data interface{}) *JsonResult {
	message = GetErrorMessage(code, message)
	// 按照接口格式生成原数据数组
	jsonMap := &JsonResult{
		Code: code,
		Msg:  message,
		Data: data,
	}
	return jsonMap
}
