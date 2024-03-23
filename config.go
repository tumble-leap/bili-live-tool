package main

import (
	"fmt"
	"os"

	_ "embed"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

//go:embed config_example.yaml
var ExampleCfg string

func (tool *BiliDmTool) LoadConfig() error {
	result, err := os.ReadFile(tool.ConfigFile)
	if err != nil {
		return err
	}

	yaml.Unmarshal(result, tool)
	if err != nil || tool.Rooms[0].Id == 0 {
		log.Println("解析配置文件出错，请检查配置文件是否设置正确的直播间ID")
		fmt.Scanln()
		os.Exit(0)
	}
	return nil
}

func (tool *BiliDmTool) CreateBlankConfigfile() {
	os.Create(tool.ConfigFile)
	f, err := os.OpenFile(tool.ConfigFile, os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f.WriteString(ExampleCfg)
}
