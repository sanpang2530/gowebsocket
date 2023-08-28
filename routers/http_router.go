/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 12:20
 */

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/link1st/gowebsocket/mod/home/home_controller"
	"github.com/link1st/gowebsocket/mod/system/system_controller"
	"github.com/link1st/gowebsocket/mod/user/user_controller"
)

func Init(router *gin.Engine) {
	router.LoadHTMLGlob("views/**/*")

	// 用户组
	userRouter := router.Group("/user")
	{
		userRouter.GET("/list", user_controller.List)
		userRouter.GET("/online", user_controller.Online)
		userRouter.POST("/sendMessage", user_controller.SendMessage)
		userRouter.POST("/sendMessageAll", user_controller.SendMessageAll)
	}

	// 系统
	systemRouter := router.Group("/system")
	{
		systemRouter.GET("/state", system_controller.Status)
	}

	// home
	homeRouter := router.Group("/home")
	{
		homeRouter.GET("/index", home_controller.Index)
	}

	router.POST("/user/online", user_controller.Online)
}
