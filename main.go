/**
* Created by GoLand.
* User: link1st
* Date: 2019-07-25
* Time: 09:59
 */

package main

import (
	"fmt"
	"github.com/link1st/gowebsocket/mod/conn"
	"github.com/link1st/gowebsocket/mod/conn/grpc_server"
	"github.com/link1st/gowebsocket/mod/system"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/link1st/gowebsocket/libs/redis_lib"
	"github.com/link1st/gowebsocket/routers"
	"github.com/spf13/viper"
)

// 入口
func main() {
	initConfig()

	initFile()

	initRedis()

	router := gin.Default()
	// 初始化路由
	routers.Init(router)
	routers.WebsocketInit()

	// 定时任务
	conn.Init()

	// 服务注册
	system.ServerInit()

	go conn.StartWebSocket()

	// grpc
	go grpc_server.Init()

	// go open()

	httpPort := viper.GetString("app.httpPort")
	_ = http.ListenAndServe(":"+httpPort, router)
}

// 初始化日志
func initFile() {
	// Disable Console Color, you don't need console color when writing the logs to file.
	gin.DisableConsoleColor()

	// Logging to a file.
	logFile := viper.GetString("app.logFile")
	f, _ := os.Create(logFile)
	gin.DefaultWriter = io.MultiWriter(f)
}

// 初始化Config
func initConfig() {
	viper.SetConfigName("config/app")
	viper.AddConfigPath(".") // 添加搜索路径

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config redis:", viper.Get("redis"))
}

// 初始化Redis
func initRedis() {
	redis_lib.NewClient()
}

// 打开页面
func open() {

	time.Sleep(1000 * time.Millisecond)

	httpUrl := viper.GetString("app.httpUrl")
	httpUrl = "http://" + httpUrl + "/home/index"

	fmt.Println("访问页面体验:", httpUrl)

	cmd := exec.Command("open", httpUrl)
	_, _ = cmd.Output()
}
