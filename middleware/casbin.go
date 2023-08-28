package middleware

import (
	"github.com/gin-gonic/gin"
)

var excludes = map[string]struct{}{
	`/withdraw/offline/payWallet`: {},
}

// CasBinMiddleware 认证中间件 租户
func CasBinMiddleware() gin.HandlerFunc {
	//enforcer := model.Enforcer
	return func(c *gin.Context) {
		//info, err := model.GetUserInfo(c)
		//if err != nil {
		//	model.SendResponse(c, errors.Wrap(err, t.T("未获取到token")), nil)
		//	c.Abort()
		//	return
		//}
		//
		//obj := c.Request.URL.RequestURI()
		//// 获取请求方法
		//act := c.Request.Method
		//// 获取用户的角色
		//sub := info.AdminType
		//
		//_, exist := excludes[obj]
		//if !exist {
		//	if err = model.AddUserOperaLog(info.ID, time.Now().Unix(), obj); err != nil {
		//		model.SendResponse(c, errors.Wrap(err, t.T("操作写入日志错误")), nil)
		//		c.Abort()
		//		return
		//	}
		//}
		//
		////
		//success, _ := enforcer.Enforce(sub, obj, act)
		//if success {
		//	c.Next()
		//} else {
		//	// model.SendResponse(c, errors.New(t.T("权限未通过")), nil)
		//	c.Abort()
		//	return
		//}

		c.Next()
	}
}

// AgentMiddleware 代理认证中间件
func AgentMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//info, err := model.GetUserInfo(c)
		//if err != nil {
		//	model.SendResponse(c, errors.Wrap(err, t.T("未获取到token")), nil)
		//	c.Abort()
		//	return
		//}
		//if info.RoleType != model.RoleAgent && info.RoleType != model.RoleNormal {
		//	model.SendResponse(c, errors.New(t.T("权限未通过")), nil)
		//	c.Abort()
		//	return
		//}
		c.Next()
	}
}
