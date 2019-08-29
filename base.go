package HiWeiChatApi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

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

func NewCustomerNewsMsg(toUser string, article customerNewArticle) *customerNewsMsg {
	return &customerNewsMsg{
		ToUser:             toUser,
		MsgType:            "news",
		customerNewArticle: article,
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

func NewCustomerTextMsg(toUser string, content string) *customerTextMsg {
	text := customerTextContent{Content: content}
	return &customerTextMsg{
		ToUser:              toUser,
		MsgType:             "text",
		customerTextContent: text,
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

func NewCustomerImageMsg(toUser string, mediaId string) *customerImageMsg {
	Media := customerImageContent{MediaId: mediaId}
	return &customerImageMsg{
		ToUser:               toUser,
		MsgType:              "text",
		customerImageContent: Media,
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
