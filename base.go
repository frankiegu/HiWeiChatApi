package HiWeiChatApi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
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

func (c *CustomerMsg) Marshal() ([]byte, error) {
	if c.Msg == nil {
		return nil, errors.New("msg is empty")
	}
	data, err := json.Marshal(c.Msg)
	return data, err
}

// func (c *CustomerMsg) ResponseParse(resp []byte) WeiChatResponse {
// 	wxRe := WeiChatResponse{}
// 	if resp == nil || len(resp) < 1 {
// 		return wxRe
// 	}
// 	json.Unmarshal(resp, &wxRe)
// 	return wxRe
// }

type customerNewsMsg struct {
	ToUser             string `json:"touser"`
	MsgType            string `json:"msgtype"`
	customerNewArticle `json:"news"`
}
type customerNewArticle struct {
	Articles []customeNewArticleEle `json:"articles"`
}
type customeNewArticleEle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PicUrl      string `json:'picurl'`
}

func NewCustomerNewsMsg(toUser string, title string, desc string, url string, picUrl string) *CustomerMsg {

	ele := customeNewArticleEle{
		Title:       title,
		Description: desc,
		Url:         url,
		PicUrl:      picUrl,
	}
	eles := make([]customeNewArticleEle, 1)
	eles[0] = ele
	article := customerNewArticle{Articles: eles}
	msg := customerNewsMsg{
		ToUser:             toUser,
		MsgType:            "news",
		customerNewArticle: article,
	}

	return &CustomerMsg{
		Msg: msg,
	}
}

// func NewCustomeNewArticleEle(title string, desc string, url string, picUrl string) customerNewArticle {
// 	ele := customeNewArticleEle{
// 		Title:       title,
// 		Description: desc,
// 		Url:         url,
// 		PicUrl:      picUrl,
// 	}
// 	return customerNewArticle{Articles: ele}
// }

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
		MsgType:              "image",
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

type MediaType string

const (
	ImageMedia MediaType = MediaType("image")
	VoiceMedia MediaType = MediaType("voice")
	VideoMedia MediaType = MediaType("video")
	ThumbMedia MediaType = MediaType("thumb")
)

func Post(url string, paramBody []byte, header map[string]string) ([]byte, error) {
	client := &http.Client{}
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

func PostFile(url string, params map[string]string, fileFieldName, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileFieldName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", url, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	return respBody, err
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
