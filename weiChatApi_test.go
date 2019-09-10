package HiWeiChatApi

import (
	"fmt"
	"testing"
)

var appId string = "wxacb25453422b57a7"
var secret string = "85a10227ade9d331f44cfc8d78904004"
var accessToken string = ""
var openId string = "omKA5uALKWNYhZk0CYN7HD43RXaw"
var tmpId string = "a2m_d3YqUBeayXo5Tr5txsIiI68yRJHO71B2SdCv1Es"
var weiApi *WeiChatAPI

func init() {
	weiApi = WeiChat(nil, nil)
	accessToken, err := weiApi.AccessToken(appId, secret)
}

func TestCustomerNewsMsg(t *testing.T) {
	newMsg := NewCustomerNewsMsg(openId, "测试New", "测试测试", "www.baidu.com", "")
	ret, err := weiApi.SendCustomMsg(accessToken, newMsg)
	fmt.Println(ret, err)

}
func TestCustomerTextMsg(t *testing.T) {
	textMsg := NewCustomerTextMsg(openId, "test TextMsg")
	ret, err := weiApi.SendCustomMsg(accessToken, textMsg)
	fmt.Println(ret, err)
}
func TestCustomerImageMsg(t *testing.T) {
	token, err := weiApi.AccessToken(appId, secret)
	imageMsg := NewCustomerImageMsg(openId, "5deENy942-cX15WuZsQJQSskcbpIDiOpFkriV2_a89I3Kwod9uWPtJ0_2XLVK4jZ")
	ret, err := weiApi.SendCustomMsg(token, imageMsg)
	fmt.Println(ret, err)
}
func TestTemplateMsg(t *testing.T) {
	keyword1 := NewWxTemplateEle()
	keyword1.Value = "hello"
	keyword2 := NewWxTemplateEle()
	keyword2.Value = "world"
	remark := NewWxTemplateEle()
	remark.Value = "备注一下"

	data := make(map[string]WxTemplateEle, 3)
	data["keyword1"] = keyword1
	data["keyword2"] = keyword2
	data["remark"] = remark
	tmpMsg := NewWxTemplateMsg(openId, tmpId, "www.baidu.com", nil, data)
	ret, err := weiApi.SendTemplateMsg(accessToken, tmpMsg)
	fmt.Println(ret, err)
}

// func TestAccessToken(t *testing.T) {
// 	token, err := weiApi.AccessToken(appId, secret)
// 	fmt.Println(token, err)
// }
func TestCode(t *testing.T) {
	err := weiApi.Code(appId, "http://smart03.com")
	fmt.Println(err)
}
func receiveDealFun(msg WxReceiveCommonMsg) error {
	fmt.Println(msg)
	return nil
}

func TestAuthByCode(t *testing.T) {}

func TestUploadTemporaryMedia(t *testing.T) {
	resp, err := weiApi.UploadTemporaryMedia(accessToken, "/tmp/2.png", ImageMedia)
	fmt.Println(resp, err)
}

func TestUploadPermanentMedia(t *testing.T) {

	resp, err := weiApi.UploadPermanentMedia(accessToken, "/tmp/2.png", ImageMedia)
	fmt.Println(resp, err)
}

func TestGenerateQRCode(t *testing.T) {
	resp, err := weiApi.GenerateQRCode(accessToken, 1223, QR_LIMIT_STR_SCENE)
	fmt.Println(resp, err)
}
func TestShortUrl(t *testing.T) {
	resp, err := weiApi.LongTransformShortUrl(accessToken, "http://wwwxx")
	fmt.Println(resp, err)
}
