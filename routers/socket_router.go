/**
 * Created by GoLand.
 * User: link1st
 * Date: 2019-07-25
 * Time: 16:02
 */

package routers

import (
	"github.com/link1st/gowebsocket/mod/conn"
)

// WebsocketInit 路由
func WebsocketInit() {
	conn.Register("login", conn.LoginController)
	conn.Register("heartbeat", conn.HeartbeatController)
	conn.Register("ping", conn.PingController)
}
