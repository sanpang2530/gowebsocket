/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 16:24
 */

package conn

import (
	"fmt"
	"time"
)

// GetUserKey 获取用户key
func GetUserKey(appId uint32, userId string) (key string) {
	key = fmt.Sprintf("%d_%s", appId, userId)

	return
}

// GetManagerInfo 获取管理者信息
func GetManagerInfo(isDebug string) (managerInfo map[string]interface{}) {
	managerInfo = make(map[string]interface{})

	managerInfo["clientsLen"] = connManager.GetClientsLen()        // 客户端连接数
	managerInfo["usersLen"] = connManager.GetUsersLen()            // 登录用户数
	managerInfo["chanRegisterLen"] = len(connManager.Register)     // 未处理连接事件数
	managerInfo["chanLoginLen"] = len(connManager.Login)           // 未处理登录事件数
	managerInfo["chanUnregisterLen"] = len(connManager.Unregister) // 未处理退出登录事件数
	managerInfo["chanBroadcastLen"] = len(connManager.Broadcast)   // 未处理广播事件数

	if isDebug == "true" {
		addrList := make([]string, 0)
		connManager.ClientsRange(func(client *Client, value bool) (result bool) {
			addrList = append(addrList, client.Addr)

			return true
		})

		users := connManager.GetUserKeys()

		managerInfo["clients"] = addrList // 客户端列表
		managerInfo["users"] = users      // 登录用户列表
	}

	return
}

// GetUserConn 获取用户所在的连接
func GetUserConn(appId uint32, userId string) (client *Client) {
	client = connManager.GetUserClient(appId, userId)

	return
}

// ClearTimeoutConnections 定时清理超时连接
func ClearTimeoutConnections() {
	currentTime := uint64(time.Now().Unix())

	clients := connManager.GetClients()
	for client := range clients {
		if client.IsHeartbeatTimeout(currentTime) {
			fmt.Println("心跳时间超时 关闭连接", client.Addr, client.UserId, client.LoginTime, client.HeartbeatTime)

			_ = client.Socket.Close()
		}
	}
}

// GetUserList 获取全部用户
func GetUserList(appId uint32) (userList []string) {
	fmt.Println("获取全部用户", appId)

	userList = connManager.GetUserList(appId)

	return
}

// AllSendMessages 全员广播
func AllSendMessages(appId uint32, userId string, data string) {
	fmt.Println("全员广播", appId, userId, data)

	ignoreClient := connManager.GetUserClient(appId, userId)
	connManager.sendAppIdAll([]byte(data), appId, ignoreClient)
}
