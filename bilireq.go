package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/AceXiamo/blivedm-go/api"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func getUname(uid int, biliJct string, sessData string) (string, error) {
	url := fmt.Sprintf("https://api.bilibili.com/x/space/wbi/acc/info?mid=%d", uid)
	b, err := WbiRequst(url, biliJct, sessData)
	if err != nil {
		log.Println(err)
	}
	uname := gjson.Get(string(b), "data.name").String()
	return uname, nil
}

func getUnameByRoom(roomID int, biliJct string, sessData string) (string, error) {
	res, err := api.GetRoomInfo(roomID)
	if err != nil {
		return "", err
	}
	uname, err := getUname(res.Data.Uid, biliJct, sessData)
	if err != nil {
		return "", err
	}
	return uname, nil
}

func (tool *BiliDmTool) LoadCookie() error {
	cookiesByte, err := os.ReadFile(tool.CookieFile)
	if err != nil {
		return err
	}
	// tool.cookies = string(cookiesByte)
	json.Unmarshal(cookiesByte, &tool.allCookie)

	for _, v := range tool.allCookie.Data.CookieInfo.Cookies {
		if v.Name == "bili_jct" {
			tool.biliJct = v.Value
		} else if v.Name == "SESSDATA" {
			tool.sessData = v.Value
		}
	}
	if tool.biliJct == "" || tool.sessData == "" {
		return errors.New("cookie lost")
	}

	for _, v := range tool.allCookie.Data.CookieInfo.Cookies {
		tool.cookieStr += v.Name + "=" + v.Value + ";"
	}

	return nil
}

func (tool *BiliDmTool) CreateCookiefile() {
	_, err := os.Stat(tool.CookieFile)
	if os.IsNotExist(err) {
		log.Println("检测到未登录状态，请扫码登录～")
		os.Create(tool.CookieFile)
	}
}
