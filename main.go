package main

import (
	"fmt"
	"os"

	"github.com/XiaoMiku01/biliup-go/login"
	log "github.com/sirupsen/logrus"
)

func main() {
	myTool := NewBiliDmTool("config.yaml", "cookie.json", log.DebugLevel)

	if err := myTool.LoadConfig(); err != nil {
		log.Printf("无法加载配置: %v 已尝试创建新的配置文件", err)
		myTool.CreateBlankConfigfile()
		fmt.Print("按下回车键以退出...")
		fmt.Scanln()
		os.Exit(1)
	}

	if err := myTool.LoadCookie(); err != nil {
		log.Printf("无法加载Cookie信息: %v 已尝试创建新的Cookie文件", err)
		// myTool.CreateCookiefile()
		login.LoginBili()
	}
	// 如果roomid为0 也退出程序

	myTool.run()
}
