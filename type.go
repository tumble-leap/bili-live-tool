package main

import log "github.com/sirupsen/logrus"

type GenQrcodeRes struct {
	Code    int
	Message string
	Data    struct {
		Url       string
		QrcodeKey string `json:"qrcode_key"`
	}
}

type CookieType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		IsNew        bool   `json:"is_new"`
		Mid          int    `json:"mid"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		TokenInfo    struct {
			Mid          int    `json:"mid"`
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			ExpiresIn    int    `json:"expires_in"`
		} `json:"token_info"`
		CookieInfo struct {
			Cookies []struct {
				Name     string `json:"name"`
				Value    string `json:"value"`
				HTTPOnly int    `json:"http_only"`
				Expires  int    `json:"expires"`
				Secure   int    `json:"secure"`
			} `json:"cookies"`
			Domains []string `json:"domains"`
		} `json:"cookie_info"`
		Sso  []string `json:"sso"`
		Hint string   `json:"hint"`
	} `json:"data"`
}

type BiliDmTool struct {
	ConfigFile string `yaml:"-"`
	CookieFile string `yaml:"-"`
	LogLevel   log.Level
	Nick       string `yaml:"bot"`
	Admin      int
	Rooms      []struct {
		Id                int
		Enable            bool
		AutoSend          bool     `yaml:"auto_send"`
		ThankGift         bool     `yaml:"thank_gift"`
		ThankGuard        bool     `yaml:"thank_guard"`
		AutoWelcome       bool     `yaml:"auto_welcome"`
		WelcomeMessage    string   `yaml:"welcome_message"`
		Messages          []string `yaml:"auto_send_message"`
		ThankGiftMessage  []string `yaml:"thank_gift_message"`
		ThankGuardMessage []string `yaml:"thank_guard_message"`
		EnterMessage      string   `yaml:"enter_message"`
		AutoBan           bool     `yaml:"auto_ban"`
		BanWords          []string `yaml:"ban_words"`
		LimitTime         int      `yaml:"limit_time"`
		LimitNum          int      `yaml:"limit_num"`
	}

	MinDura   int        `yaml:"min_dura"`
	MaxDura   int        `yaml:"max_dura"`
	biliJct   string     `yaml:"-"`
	sessData  string     `yaml:"-"`
	cookieStr string     `yaml:"-"`
	allCookie CookieType `yaml:"-"`
}

type GetUnameByid struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    struct {
		Mid       int    `json:"mid"`
		Name      string `json:"name"`
		Sex       string `json:"sex"`
		Face      string `json:"face"`
		Sign      string `json:"sign"`
		Rank      int    `json:"rank"`
		Level     int    `json:"level"`
		Jointime  int    `json:"jointime"`
		Moral     int    `json:"moral"`
		Silence   int    `json:"silence"`
		Coins     int    `json:"coins"`
		FansBadge bool   `json:"fans_badge"`
		FansMedal struct {
			Show  bool `json:"show"`
			Wear  bool `json:"wear"`
			Medal struct {
				UID           int    `json:"uid"`
				TargetID      int    `json:"target_id"`
				MedalID       int    `json:"medal_id"`
				Level         int    `json:"level"`
				MedalName     string `json:"medal_name"`
				Intimacy      int    `json:"intimacy"`
				NextIntimacy  int    `json:"next_intimacy"`
				DayLimit      int    `json:"day_limit"`
				IsLighted     int    `json:"is_lighted"`
				LightStatus   int    `json:"light_status"`
				WearingStatus int    `json:"wearing_status"`
				Score         int    `json:"score"`
			} `json:"medal"`
		} `json:"fans_medal"`
		Official struct {
			Role  int    `json:"role"`
			Title string `json:"title"`
			Desc  string `json:"desc"`
			Type  int    `json:"type"`
		} `json:"official"`
		UserHonourInfo struct {
			Mid    int           `json:"mid"`
			Colour interface{}   `json:"colour"`
			Tags   []interface{} `json:"tags"`
		} `json:"user_honour_info"`
		IsFollowed bool   `json:"is_followed"`
		TopPhoto   string `json:"top_photo"`
		LiveRoom   struct {
			RoomStatus    int    `json:"roomStatus"`
			LiveStatus    int    `json:"liveStatus"`
			URL           string `json:"url"`
			Title         string `json:"title"`
			Cover         string `json:"cover"`
			Roomid        int    `json:"roomid"`
			RoundStatus   int    `json:"roundStatus"`
			BroadcastType int    `json:"broadcast_type"`
		} `json:"live_room"`
		Birthday string `json:"birthday"`
		School   struct {
			Name string `json:"name"`
		} `json:"school"`
		Profession struct {
			Name       string `json:"name"`
			Department string `json:"department"`
			Title      string `json:"title"`
			IsShow     int    `json:"is_show"`
		} `json:"profession"`
		Tags   interface{} `json:"tags"`
		IsRisk bool        `json:"is_risk"`
		Elec   struct {
			ShowInfo struct {
				Show    bool   `json:"show"`
				State   int    `json:"state"`
				Title   string `json:"title"`
				Icon    string `json:"icon"`
				JumpURL string `json:"jump_url"`
			} `json:"show_info"`
		} `json:"elec"`
	} `json:"data"`
}
