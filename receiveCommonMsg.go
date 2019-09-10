package HiWeiChatApi

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
