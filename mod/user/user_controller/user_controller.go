/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:11
 */

package user_controller

import (
	"fmt"
	"github.com/link1st/gowebsocket/mod/msg_chat/chat_model"
	"github.com/link1st/gowebsocket/mod/user"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/link1st/gowebsocket/common"
	"github.com/link1st/gowebsocket/libs/cache"
)

// List 查看全部在线用户
func List(c *gin.Context) {

	appIdStr := c.Query("appId")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 查看全部在线用户", appId)

	data := make(map[string]interface{})

	userList := user.UserList(appId)
	data["userList"] = userList
	data["userCount"] = len(userList)

	common.Response(c, common.OK, "", data)
}

// Online 查看用户是否在线
func Online(c *gin.Context) {

	userId := c.Query("userId")
	appIdStr := c.Query("appId")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 查看用户是否在线", userId, appIdStr)

	data := make(map[string]interface{})

	online := user.CheckUserOnline(appId, userId)
	data["userId"] = userId
	data["online"] = online

	common.Response(c, common.OK, "", data)
}

// SendMessage 给用户发送消息
func SendMessage(c *gin.Context) {
	// 获取参数
	appIdStr := c.PostForm("appId")
	userId := c.PostForm("userId")
	msgId := c.PostForm("msgId")
	message := c.PostForm("message")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 给用户发送消息", appIdStr, userId, msgId, message)

	// TODO::进行用户权限认证，一般是客户端传入TOKEN，然后检验TOKEN是否合法，通过TOKEN解析出来用户ID
	// 本项目只是演示，所以直接过去客户端传入的用户ID(userId)

	data := make(map[string]interface{})

	if cache.SeqDuplicates(msgId) {
		fmt.Println("给用户发送消息 重复提交:", msgId)
		common.Response(c, common.OK, "", data)

		return
	}

	sendResults, err := user.SendUserMessage(appId, userId, msgId, message)
	if err != nil {
		data["sendResultsErr"] = err.Error()
	}

	data["sendResults"] = sendResults

	common.Response(c, common.OK, "", data)
}

// SendMessageAll 给全员发送消息
func SendMessageAll(c *gin.Context) {
	// 获取参数
	appIdStr := c.PostForm("appId")
	userId := c.PostForm("userId")
	msgId := c.PostForm("msgId")
	message := c.PostForm("message")
	appIdUint64, _ := strconv.ParseInt(appIdStr, 10, 32)
	appId := uint32(appIdUint64)

	fmt.Println("http_request 给全体用户发送消息", appIdStr, userId, msgId, message)

	data := make(map[string]interface{})
	if cache.SeqDuplicates(msgId) {
		fmt.Println("给用户发送消息 重复提交:", msgId)
		common.Response(c, common.OK, "", data)

		return
	}

	sendResults, err := user.SendUserMessageAll(appId, userId, msgId, chat_model.MessageCmdMsg, message)
	if err != nil {
		data["sendResultsErr"] = err.Error()

	}

	data["sendResults"] = sendResults

	common.Response(c, common.OK, "", data)
}
