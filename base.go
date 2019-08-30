package HiWeiChatApi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type WeiChatConfig struct {
}

type WeiChatResponse struct {
	ErrNo  int    `json:"errno"`
	ErrMsg string `json:"errmsg"`
	MsgId  int64  `json:"msgid"`
}

type CustomerMsg struct {
	Msg interface{}
}

func (c *CustomerMsg) Data() ([]byte, error) {
	if c.Msg == nil {
		return nil, errors.New("msg is empty")
	}
	data, err := json.Marshal(c.Msg)
	return data, err
}

func ResponseParse(resp []byte) WeiChatResponse {
	wxRe := WeiChatResponse{}
	json.Unmarshal(resp, &wxRe)
	return wxRe
}

type customerNewsMsg struct {
	ToUser             string `json:"touser"`
	MsgType            string `json:"msgtype"`
	customerNewArticle `json:"news"`
}
type customerNewArticle struct {
	Articles customeNewArticleEle `json:"articles"`
}
type customeNewArticleEle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:'picurl'`
}

func NewCustomerNewsMsg(toUser string, article customerNewArticle) *CustomerMsg {

	msg := customerNewsMsg{
		ToUser:             toUser,
		MsgType:            "news",
		customerNewArticle: article,
	}
	return &CustomerMsg{
		Msg: msg,
	}
}

func NewCustomeNewArticleEle(title string, desc string, url string, picUrl string) customerNewArticle {
	ele := customeNewArticleEle{
		Title:       title,
		Description: desc,
		Url:         url,
		PicUrl:      picUrl,
	}
	return customerNewArticle{Articles: ele}
}

type customerTextMsg struct {
	ToUser              string `json:"touser"`
	MsgType             string `json:"msgtype"`
	customerTextContent `json:"text"`
}
type customerTextContent struct {
	Content string `json:"content"`
}

func NewCustomerTextMsg(toUser string, content string) *CustomerMsg {
	text := customerTextContent{Content: content}
	msg := customerTextMsg{
		ToUser:              toUser,
		MsgType:             "text",
		customerTextContent: text,
	}
	return &CustomerMsg{
		Msg: msg,
	}
}

type customerImageMsg struct {
	ToUser               string `json:"touser"`
	MsgType              string `json:"msgtype"`
	customerImageContent `json:"image"`
}
type customerImageContent struct {
	MediaId string `json:"media_id"`
}

func NewCustomerImageMsg(toUser string, mediaId string) *CustomerMsg {
	Media := customerImageContent{MediaId: mediaId}
	msg := customerImageMsg{
		ToUser:               toUser,
		MsgType:              "text",
		customerImageContent: Media,
	}
	return &CustomerMsg{
		Msg: msg,
	}
}

//接收普通消息
type WxReceiveCommonMsg struct {
	ToUserName   string
	FromUserName string
	Content      string
	CreateTime   int64
	MsgType      string
	MsgId        int64
	PicUrl       string
	MediaId      string
	Format       string
	Recognition  string
	ThumbMediaId string
}

type WxReceiveFunc func(msg WxReceiveCommonMsg) error

type WxTemplateMsg struct {
	ToUser       string `json:"touser"`
	TemplateId   string `json:"template_id"`
	Url          string `json:"url"`
	MiniPrograme `json:"miniprogram"`
	Data         map[string]WxTemplateEle `json:"data"`
}
type MiniPrograme struct {
	AppId    string `json:"appid"`
	PagePath string `json:"pagepath"`
}
type WxTemplateEle struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

func (wxt *WxTemplateMsg) Data() {
	//会不会序列话函数？
	data, err := json.Marshal(c.Msg)
}
func NewWxTemplateMsg(toUser, tmpId, url string, mini MiniPrograme, data map[string]WxTemplateEle) {
	tMsg := WxTemplateMsg{
		ToUser:       toUser,
		TemplateId:   tmpId,
		Url:          url,
		MiniPrograme: mini,
		Data:         data,
	}

}

func Post(url string, paramBody []byte, header map[string]string) ([]byte, error) {
	client := &http.Client{}
	//	fmt.Println("******", url, paramBody)
	paramsData := bytes.NewBuffer(paramBody)
	req, err := http.NewRequest("POST", url, paramsData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if header != nil {
		for hkey, hval := range header {
			req.Header.Set(hkey, hval)
		}
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}

func Get(url string, params map[string]string) ([]byte, error) {
	paramsStr := ""
	if params != nil && len(params) > 0 {
		for k, v := range params {
			if paramsStr != "" {
				paramsStr = fmt.Sprintf("%s&%s=%s", paramsStr, k, v)
			} else {
				paramsStr = fmt.Sprintf("%s=%s", k, v)
			}
		}
	}

	urls := strings.Split(url, "?")
	targetUrl := url
	if len(urls) > 1 && paramsStr != "" {
		targetUrl = fmt.Sprintf("%s&%s", url, paramsStr)
	} else if paramsStr != "" {
		targetUrl = fmt.Sprintf("%s?%s", url, paramsStr)
	}

	resp, err := http.Get(targetUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}
