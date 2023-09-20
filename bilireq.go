package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Akegarasu/blivedm-go/api"
	log "github.com/sirupsen/logrus"
)

func getUname(uid int) (string, error) {
	url := fmt.Sprintf("https://api.bilibili.com/x/space/wbi/acc/info?mid=%d", uid)
	result := &GetUnameByid{}
	headers := &http.Header{}
	headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/117.0")
	err := api.GetJsonWithHeader(url, headers, result)
	if err != nil {
		return "", err
	}
	if result.Code != 0 {
		return "", errors.New(result.Message)
	}
	return result.Data.Name, nil
}

func getUnameByRoom(roomID int) (string, error) {
	res, err := api.GetRoomInfo(roomID)
	if err != nil {
		return "", err
	}
	uname, err := getUname(res.Data.Uid)
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
	tool.cookies = string(cookiesByte)
	cookieList := strings.Split(tool.cookies, ";")
	for _, v := range cookieList {
		if strings.Split(v, "=")[0] == "bili_jct" {
			tool.biliJct = strings.Split(v, "=")[1]
		} else if strings.Split(v, "=")[0] == "SESSDATA" {
			tool.sessData = strings.Split(v, "=")[1]
		}
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
