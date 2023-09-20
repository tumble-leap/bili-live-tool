package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	log "github.com/sirupsen/logrus"
)

var (
	genQrcodeUrl  = "https://passport.bilibili.com/x/passport-login/web/qrcode/generate?source=main_web"
	valiQrcodeUrl = "https://passport.bilibili.com/x/passport-login/web/qrcode/poll?qrcode_key=%s&source=main_web"
	userAgent     = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36"
)

func genValiUrl(key string) string {
	return fmt.Sprintf(valiQrcodeUrl, key)
}

func myRequest(method string, url string, cookies map[string]string, postData io.Reader) []byte {
	c := http.Client{}
	req, _ := http.NewRequest(method, url, postData)
	for k, v := range cookies {
		req.AddCookie(&http.Cookie{Name: k, Value: v})
	}
	req.Header.Set("User-Agent", userAgent)
	resp, err := c.Do(req)
	if err != nil {
		log.Println("请求发送失败", err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("内容读取失败", err)
	}
	return body
}

func createQrcode() string {
	res := myRequest(http.MethodGet, genQrcodeUrl, nil, nil)
	var v GenQrcodeRes
	_ = json.Unmarshal(res, &v)
	if v.Code != 0 {
		log.Println("登陆链接获取失败")
	}
	url := v.Data.Url
	obj := qrcodeTerminal.New()
	obj.Get(url).Print()
	return v.Data.QrcodeKey
}

func (tool *BiliDmTool) LoginBilibili() {
	key := createQrcode()
	time.Sleep(time.Second * 5)
	url := genValiUrl(key)
	log.Println("打开哔哩哔哩app扫码登陆\n若二维码显示不完整则按下ctrl后波动滚轮调整大小")
	log.Println("正在准备获取扫码状态...")
	f, err := os.OpenFile(tool.CookieFile, os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	for {
		res := myRequest(http.MethodGet, url, nil, nil)
		var v valiQrcodeRes
		_ = json.Unmarshal(res, &v)
		if v.Message != "0" {
			log.Println("网络异常")
			break
		} else if v.Data.Message == "" {
			cookies := strings.Split(v.Data.Url, "?")[1] + ";"
			cookies = strings.ReplaceAll(cookies, "&", ";")
			f.WriteString(cookies)
			log.Println("登陆成功！")
			break
		} else if v.Data.Message == "二维码已失效" {
			log.Println("二维码已失效")
			break
		}
		time.Sleep(time.Second * 2)
	}
}
