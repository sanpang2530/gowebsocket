/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-27
 * Time: 13:12
 */

package conn

import (
	"encoding/json"
	"fmt"
	"github.com/link1st/gowebsocket/mod/user/user_model"
	"time"

	"github.com/link1st/gowebsocket/common"
	"github.com/link1st/gowebsocket/libs/cache"
	"github.com/link1st/gowebsocket/models"

	"github.com/redis/go-redis/v9"
)

// PingController ping
func PingController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	fmt.Println("webSocket_request ping接口", client.Addr, seq, message)

	data = "pong"

	return
}

// LoginController 用户登录
func LoginController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	currentTime := uint64(time.Now().Unix())

	request := &models.Login{}
	if err := json.Unmarshal(message, request); err != nil {
		code = common.ParameterIllegal
		fmt.Println("用户登录 解析数据失败", seq, err)

		return
	}

	fmt.Println("webSocket_request 用户登录", seq, "Token", request.Token)

	// TODO::进行用户权限认证，一般是客户端传入TOKEN，然后检验TOKEN是否合法，通过TOKEN解析出来用户ID
	// 本项目只是演示，所以直接过去客户端传入的用户ID
	if request.UserId == "" || len(request.UserId) >= 20 {
		code = common.UnauthorizedUserId
		fmt.Println("用户登录 非法的用户", seq, request.UserId)

		return
	}

	if !InAppIds(request.AppId) {
		code = common.Unauthorized
		fmt.Println("用户登录 不支持的平台", seq, request.AppId)
		return
	}

	if client.IsLogin() {
		fmt.Println("用户登录 用户已经登录", client.AppId, client.UserId, seq)
		code = common.OperationFailure
		return
	}

	client.Login(request.AppId, request.UserId, currentTime)

	// 存储数据
	userOnline := user_model.UserLogin(serverIp, serverPort, request.AppId, request.UserId, client.Addr, currentTime)
	err := cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		code = common.ServerError
		fmt.Println("用户登录 SetUserOnlineInfo", seq, err)

		return
	}

	// 用户登录
	login := &Login{
		AppId:  request.AppId,
		UserId: request.UserId,
		Client: client,
	}
	connManager.Login <- login

	fmt.Println("用户登录 成功", seq, client.Addr, request.UserId)

	return
}

// HeartbeatController 心跳接口
func HeartbeatController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = common.OK
	currentTime := uint64(time.Now().Unix())

	request := &models.HeartBeat{}
	if err := json.Unmarshal(message, request); err != nil {
		code = common.ParameterIllegal
		fmt.Println("心跳接口 解析数据失败", seq, err)

		return
	}

	fmt.Println("webSocket_request 心跳接口", client.AppId, client.UserId)

	if !client.IsLogin() {
		fmt.Println("心跳接口 用户未登录", client.AppId, client.UserId, seq)
		code = common.NotLoggedIn

		return
	}

	userOnline, err := cache.GetUserOnlineInfo(client.GetKey())
	if err != nil {
		if err == redis.Nil {
			code = common.NotLoggedIn
			fmt.Println("心跳接口 用户未登录", seq, client.AppId, client.UserId)

			return
		} else {
			code = common.ServerError
			fmt.Println("心跳接口 GetUserOnlineInfo", seq, client.AppId, client.UserId, err)

			return
		}
	}

	client.Heartbeat(currentTime)
	userOnline.Heartbeat(currentTime)
	err = cache.SetUserOnlineInfo(client.GetKey(), userOnline)
	if err != nil {
		code = common.ServerError
		fmt.Println("心跳接口 SetUserOnlineInfo", seq, client.AppId, client.UserId, err)

		return
	}

	return
}
