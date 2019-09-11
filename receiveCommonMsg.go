package HiWeiChatApi

//接收普通消息
type WxReceiveCommonMsg struct {
	ToUserName   string //接收者
	FromUserName string //发送者
	Content      string //文本内容
	CreateTime   int64  //创建时间
	MsgType      string //消息类型
	MsgId        int64  //消息id
	PicUrl       string //图片url
	MediaId      string //媒体id
	Format       string
	Recognition  string
	ThumbMediaId string //缩略图媒体ID
}

//callback(接收到消息之后，会将消息交于这个函数处理)
type WxReceiveFunc func(msg WxReceiveCommonMsg) error
