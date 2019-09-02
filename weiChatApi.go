package HiWeiChatApi

import (
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

type WeiChatAPI struct {
	log         *log.Logger
	receiveFunc WxReceiveFunc
}

func WeiChat(logger *log.Logger, revFunc WxReceiveFunc) *WeiChatAPI {
	if logger == nil {
		// logger = log.New(l.dest, , log.Ldate|log.Ltime|log.Lshortfile)
	}
	return &WeiChatAPI{
		log:         logger,
		receiveFunc: revFunc,
	}
}

func (api *WeiChatAPI) ReceiveFunc(revFunc WxReceiveFunc) {
	api.receiveFunc = revFunc
}

func (api *WeiChatAPI) ReceiveCommonMsg(msgData []byte) (WxReceiveCommonMsg, error) {

	msg := WxReceiveCommonMsg{}
	err := xml.Unmarshal(msgData, &msg)
	if api.receiveFunc == nil {
		return msg, err
	}
	err = api.receiveFunc(msg)
	return msg, err
}

// func (api *WeiChatAPI) GenerateWxTemplateMsgFormat(openId string, templateId string, content map[string]interface{}) ([]byte, error) {
// 	content["touser"] = openId
// 	content["template_id"] = templateId
// 	jsonData, err := json.Marshal(content)
// 	return jsonData, err
// }

func (api *WeiChatAPI) AccessToken(appId, secret string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, secret)
	body, err := Get(url, nil)
	if err != nil {
		return "", err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	return fmt.Sprintf("%v", result["access_token"]), err
}

func (api *WeiChatAPI) SendTemplateMsg(accessToken string, wxt *WxTemplateMsg) (WeiChatResponse, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken)
	json, err := wxt.Marshal()
	if err != nil {
		return api.responseParse(nil), err
	}
	body, err := Post(url, json, nil)
	return api.responseParse(body), err
}

func (api *WeiChatAPI) SendCustomMsg(accessToken string, msg *CustomerMsg) (WeiChatResponse, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)
	data, err := msg.Marshal()
	if err != nil {
		return api.responseParse(nil), err
	}

	ret, err := Post(url, data, nil)
	if err != nil {
		return api.responseParse(nil), err
	}
	return api.responseParse(ret), err
}

func (api *WeiChatAPI) MakeSignature(token, timestamp, nonce string) string {
	strs := []string{token, timestamp, nonce}
	sort.Strings(strs)
	sha := sha1.New()
	io.WriteString(sha, strings.Join(strs, ""))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func (api *WeiChatAPI) Code(appId, redirectUrl string) error {

	url := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=1#wechat_redirect",
		appId, redirectUrl)
	_, err := Get(url, nil)
	return err
}
func (api *WeiChatAPI) AuthAccessToken(appId, secret, code string) string {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appId, secret, code)
	body, _ := Get(url, nil)
	return string(body)
}
func (api *WeiChatAPI) responseParse(resp []byte) WeiChatResponse {
	wxRe := WeiChatResponse{}
	if resp == nil || len(resp) < 1 {
		return wxRe
	}
	json.Unmarshal(resp, &wxRe)
	return wxRe
}
