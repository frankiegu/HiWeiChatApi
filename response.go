package HiWeiChatApi

type WeiChatResponse struct {
	ErrNo  int    `json:"errno"`
	ErrMsg string `json:"errmsg"`
	MsgId  int64  `json:"msgid"`
}

//二维码
type QRCodeResponse struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	Url           string `json:"url"`
	WeiChatResponse
}

//长链接转换为短链接
type ShortUrlResponse struct {
	WeiChatResponse
	ShortUrl string `json:"short_url"`
}

//素材
type MediaReponse struct {
	Type     string `json:"type"`
	MediaId  string `json:"media_id"`
	CreateAt int    `json:"create_at"`
	WeiChatResponse
}
