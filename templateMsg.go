package HiWeiChatApi

import (
	"encoding/json"
)

//微信模板消息
type WxTemplateMsg struct {
	ToUser        string `json:"touser"`
	TemplateId    string `json:"template_id"`
	Url           string `json:"url"`
	*MiniPrograme `json:"miniprogram"`
	Data          map[string]WxTemplateEle `json:"data"`
}

//小程序
type MiniPrograme struct {
	AppId    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

//实例化小程序参数
func NewMiniProgram(appId, pagePath string) *MiniPrograme {
	return &MiniPrograme{
		AppId:    appId,
		PagePath: pagePath,
	}
}

//模板消息元素
type WxTemplateEle struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

//实例化模板消息元素
func NewWxTemplateEle() WxTemplateEle {
	return WxTemplateEle{
		Color: "#000000",
	}
}
func (wxt *WxTemplateMsg) Marshal() ([]byte, error) {
	data, err := json.Marshal(wxt)
	return data, err
}

//实例化模板消息
func NewWxTemplateMsg(toUser, tmpId, url string, mini *MiniPrograme, data map[string]WxTemplateEle) *WxTemplateMsg {
	return &WxTemplateMsg{
		ToUser:       toUser,
		TemplateId:   tmpId,
		Url:          url,
		MiniPrograme: mini,
		Data:         data,
	}

}
