package HiWeiChatApi

type WeiChatResponse struct {
	ErrNo  int    `json:"errno"`
	ErrMsg string `json:"errmsg"`
	MsgId  int64  `json:"msgid"`
}
type QRCodeResponse struct {
	Ticket        string `json:"ticket"`
	ExpireSeconds int64  `json:"expire_seconds"`
	Url           string `json:"url"`
	WeiChatResponse
}

type ShortUrlResponse struct {
	WeiChatResponse
	ShortUrl string `json:"short_url"`
}
