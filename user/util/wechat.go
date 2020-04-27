package util

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"teamup/config"
)

type WechatSession struct {
	Openid string
	SessionKey string
	Unionid string
	Errcode int32
	Errmsg string
}

func Code2Session(code string) (*WechatSession, error){
	wechatSession := &WechatSession{}
	appid := config.Cfg.Get("wechat", "app_id").String("")
	secret := config.Cfg.Get("wechat", "app_secret").String("")
	if appid == "" || secret == "" {
		return wechatSession, errors.New("config not valid")
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session" +
		"?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appid, secret, code)
	res, err := http.Get(url)
	if err != nil {
		return wechatSession, err
	}
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, wechatSession)
	return wechatSession, nil
}

func Hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}