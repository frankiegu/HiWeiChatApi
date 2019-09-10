package HiWeiChatApi

type MediaType string

const (
	ImageMedia MediaType = MediaType("image")
	VoiceMedia MediaType = MediaType("voice")
	VideoMedia MediaType = MediaType("video")
	ThumbMedia MediaType = MediaType("thumb")
)

type QRCodeType string

const (
	QR_SCENE           QRCodeType = QRCodeType("QR_SCENE")           //短期 整型参数
	QR_STR_SCENE       QRCodeType = QRCodeType("QR_STR_SCENE")       //短期 字符串参数
	QR_LIMIT_SCENE     QRCodeType = QRCodeType("QR_LIMIT_SCENE")     //长期 整型参数
	QR_LIMIT_STR_SCENE QRCodeType = QRCodeType("QR_LIMIT_STR_SCENE") //长期 字符串参数
)
