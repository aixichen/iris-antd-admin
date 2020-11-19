package main

import (
	"car-tms/libs"
	"car-tms/seeder"
	web_server "car-tms/web_service"
	"fmt"
	logger "github.com/sirupsen/logrus"
	"os"
)

var Version = "master"
var f *os.File

func init() {
	f, _ := os.OpenFile(fmt.Sprintf("%s/go_iris.log", libs.LogDir()), os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	logger.SetFormatter(&logger.JSONFormatter{})
	logger.SetOutput(f)
	logger.SetLevel(logger.DebugLevel)
}

func main() {
	irisServer := web_server.NewServer()
	if irisServer == nil {
		panic("Http 初始化失败")
	}
	irisServer.NewApp()

	if len(os.Args) == 2 {
		if os.Args[1] == "version" {
			fmt.Println(fmt.Sprintf("版本号：%s", Version))
			return
		} else if os.Args[1] == "seeder" {
			seeder.Seeds.Perms = irisServer.GetRoutes()
			seeder.Run()
			return
		} else if os.Args[1] == "perms" {
			fmt.Println("系统权限：")
			fmt.Println()
			routes := irisServer.GetRoutes()
			for _, route := range routes {
				fmt.Println("+++++++++++++++")
				fmt.Println(fmt.Sprintf("名称 ：%s ", route.DisplayName))
				fmt.Println(fmt.Sprintf("路由地址 ：%s ", route.Name))
				fmt.Println(fmt.Sprintf("请求方式 ：%s", route.Act))
				fmt.Println()
			}
			return
		} else {
			if libs.IsPortInUse(libs.Config.Port) {
				panic(fmt.Sprintf("端口 %d 已被使用", libs.Config.Port))
				return
			}
		}
	}
	seeder.Seeds.Perms = irisServer.GetRoutes()
	seeder.Run()

	err := irisServer.Serve()
	if err != nil {
		panic(err)
	}

}
