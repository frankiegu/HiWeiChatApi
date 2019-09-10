package HiWeiChatApi

import (
	"encoding/json"
)

type WxTemplateMsg struct {
	ToUser        string `json:"touser"`
	TemplateId    string `json:"template_id"`
	Url           string `json:"url"`
	*MiniPrograme `json:"miniprogram"`
	Data          map[string]WxTemplateEle `json:"data"`
}
type MiniPrograme struct {
	AppId    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

func NewMiniProgram(appId, pagePath string) *MiniPrograme {
	return &MiniPrograme{
		AppId:    appId,
		PagePath: pagePath,
	}
}

type WxTemplateEle struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

func NewWxTemplateEle() WxTemplateEle {
	return WxTemplateEle{
		Color: "#000000",
	}
}
func (wxt *WxTemplateMsg) Marshal() ([]byte, error) {
	data, err := json.Marshal(wxt)
	return data, err
}
func NewWxTemplateMsg(toUser, tmpId, url string, mini *MiniPrograme, data map[string]WxTemplateEle) *WxTemplateMsg {
	return &WxTemplateMsg{
		ToUser:       toUser,
		TemplateId:   tmpId,
		Url:          url,
		MiniPrograme: mini,
		Data:         data,
	}

}
