package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	myTool := NewBiliDmTool("config.yaml", "cookies.txt", log.DebugLevel)
	log.SetLevel(log.InfoLevel)

	if err := myTool.LoadConfig(); err != nil {
		log.Printf("无法加载配置: %v 已尝试创建新的配置文件", err)
		myTool.CreateBlankConfigfile()
		fmt.Print("按下回车键以退出...")
		fmt.Scanln()
		os.Exit(1)
	}

	if err := myTool.LoadCookie(); err != nil {
		log.Printf("无法加载Cookie信息: %v 已尝试创建新的Cookie文件", err)
		myTool.CreateCookiefile()
		myTool.LoginBilibili()
	}

	myTool.run()
}
