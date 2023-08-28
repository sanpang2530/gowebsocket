/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-30
* Time: 12:27
 */

package user

import (
	"errors"
	"fmt"
	"github.com/link1st/gowebsocket/mod/conn"
	"github.com/link1st/gowebsocket/mod/conn/grpc_client"
	"github.com/link1st/gowebsocket/mod/msg_chat/chat_model"
	"github.com/link1st/gowebsocket/mod/system/system_model"
	"time"

	"github.com/link1st/gowebsocket/libs/cache"
	"github.com/redis/go-redis/v9"
)

// UserList 查询所有用户
func UserList(appId uint32) (userList []string) {

	userList = make([]string, 0)
	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("给全体用户发消息", err)

		return
	}

	for _, server := range servers {
		var (
			list []string
		)
		if conn.IsLocal(server) {
			list = conn.GetUserList(appId)
		} else {
			list, _ = grpc_client.GetUserList(server, appId)
		}
		userList = append(userList, list...)
	}

	return
}

// CheckUserOnline 查询用户是否在线
func CheckUserOnline(appId uint32, userId string) (online bool) {
	// 全平台查询
	if appId == 0 {
		for _, appId := range conn.GetAppIds() {
			online, _ = checkUserOnline(appId, userId)
			if online == true {
				break
			}
		}
	} else {
		online, _ = checkUserOnline(appId, userId)
	}

	return
}

// 查询用户 是否在线
func checkUserOnline(appId uint32, userId string) (online bool, err error) {
	key := conn.GetUserKey(appId, userId)
	userOnline, err := cache.GetUserOnlineInfo(key)
	if err != nil {
		if err == redis.Nil {
			fmt.Println("GetUserOnlineInfo", appId, userId, err)

			return false, nil
		}

		fmt.Println("GetUserOnlineInfo", appId, userId, err)

		return
	}

	online = userOnline.IsOnline()

	return
}

// SendUserMessage 给用户发送消息
func SendUserMessage(appId uint32, userId string, msgId, message string) (sendResults bool, err error) {

	data := chat_model.GetTextMsgData(userId, msgId, message)

	client := conn.GetUserConn(appId, userId)

	if client != nil {
		// 在本机发送
		sendResults, err = SendUserMessageLocal(appId, userId, data)
		if err != nil {
			fmt.Println("给用户发送消息", appId, userId, err)
		}

		return
	}

	key := conn.GetUserKey(appId, userId)
	info, err := cache.GetUserOnlineInfo(key)
	if err != nil {
		fmt.Println("给用户发送消息失败", key, err)

		return false, nil
	}
	if !info.IsOnline() {
		fmt.Println("用户不在线", key)
		return false, nil
	}
	server := system_model.NewServer(info.AccIp, info.AccPort)
	msg, err := grpc_client.SendMsg(server, msgId, appId, userId, chat_model.MessageCmdMsg, chat_model.MessageCmdMsg, message)
	if err != nil {
		fmt.Println("给用户发送消息失败", key, err)

		return false, err
	}
	fmt.Println("给用户发送消息成功-rpc", msg)
	sendResults = true

	return
}

// SendUserMessageLocal 给本机用户发送消息
func SendUserMessageLocal(appId uint32, userId string, data string) (sendResults bool, err error) {

	client := conn.GetUserConn(appId, userId)
	if client == nil {
		err = errors.New("用户不在线")

		return
	}

	// 发送消息
	client.SendMsg([]byte(data))
	sendResults = true

	return
}

// SendUserMessageAll 给全体用户发消息
func SendUserMessageAll(appId uint32, userId string, msgId, cmd, message string) (sendResults bool, err error) {
	sendResults = true

	currentTime := uint64(time.Now().Unix())
	servers, err := cache.GetServerAll(currentTime)
	if err != nil {
		fmt.Println("给全体用户发消息", err)

		return
	}

	for _, server := range servers {
		if conn.IsLocal(server) {
			data := chat_model.GetMsgData(userId, msgId, cmd, message)
			conn.AllSendMessages(appId, userId, data)
		} else {
			_, _ = grpc_client.SendMsgAll(server, msgId, appId, userId, cmd, message)
		}
	}

	return
}
