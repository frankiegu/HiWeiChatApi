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

//图文消息
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

/**实例化图片消息
 * toUser:openId
 * title:
 * desc:
 * url:跳转链接
 * picUrl:图片链接
 */
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

type customerTextMsg struct {
	ToUser              string `json:"touser"`
	MsgType             string `json:"msgtype"`
	customerTextContent `json:"text"`
}
type customerTextContent struct {
	Content string `json:"content"`
}

/**
 * 实例化文字消息
 * toUser:openId
 * content:文本内容
 */
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

/**
 * 实例化图片消息
 * toUser:openId
 * mediaId:微信媒体ID
 */
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
