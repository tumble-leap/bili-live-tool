package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type BanResp struct {
	Code    int
	Message string
}

var (
	liveBanApi = "https://api.live.bilibili.com/xlive/web-ucenter/v1/banned/AddSilentUser"
	userAgent  = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"
)

func (tool *BiliDmTool) blockUser(userID string, roomID int) error {
	data := url.Values{}
	data.Set("room_id", fmt.Sprint(roomID))
	data.Set("tuid", fmt.Sprint(userID))
	data.Set("csrf", tool.biliJct)
	data.Set("csrf_token", tool.biliJct)

	client := &http.Client{}
	req, err := http.NewRequest("POST", liveBanApi, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Referer", "https://www.bilibili.com")
	req.Header.Set("User-Agent", userAgent)

	req.AddCookie(&http.Cookie{
		Name:  "SESSDATA",
		Value: tool.sessData,
	})

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var v BanResp
	json.Unmarshal(body, &v)
	if v.Code == 0 {
		return nil
	} else {
		return errors.New(v.Message)
	}
}
