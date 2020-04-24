package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	appid := "appid"
	js_code := "jscode"
	secret := "secret"
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session" +
		"?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appid, secret, js_code)
	res, err := http.Get(url)
	if err != nil {
		return wechatSession, err
	}
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, wechatSession)
	return wechatSession, nil
}