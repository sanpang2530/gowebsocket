/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-27
 * Time: 14:41
 */

package models

/************************  请求数据  **************************/

// Request 通用请求数据格式
type Request struct {
	Seq  string      `json:"seq"`            // 消息的唯一Id
	Cmd  string      `json:"cmd"`            // 请求命令字
	Data interface{} `json:"data,omitempty"` // 数据 json
}

// Login 登录请求数据
type Login struct {
	Token  string `json:"token"` // 验证用户是否登录
	AppId  uint32 `json:"app_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

// HeartBeat 心跳请求数据
type HeartBeat struct {
	UserId string `json:"user_id,omitempty"`
}
