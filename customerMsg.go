package HiWeiChatApi

import (
	"encoding/json"
	"errors"
)

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
