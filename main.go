package main

import (
	"fmt"
	"os"

	"github.com/XiaoMiku01/biliup-go/login"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		ForceColors:     true,
	})
}

func main() {
	myTool := NewBiliDmTool("config.yaml", "cookie.json", log.InfoLevel)

	if err := myTool.LoadConfig(); err != nil {
		log.Printf("无法加载配置: %v 已尝试创建新的配置文件", err)
		myTool.CreateBlankConfigfile()
		log.Print("按下回车键以退出...")
		fmt.Scanln()
		os.Exit(1)
	}

	if err := myTool.LoadCookie(); err != nil {
		log.Printf("无法加载Cookie信息: %v 已尝试创建新的Cookie文件", err)
		login.LoginBili()
	}

	myTool.run()
}
