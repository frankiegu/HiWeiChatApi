package HiWeiChatApi

import (
	"fmt"
	"testing"
)

/**
 * [下列参数值需要根据自己的环境进行配置]
 * @type {String}
 */
var appId string = "wxacb25453422b57a7"
var secret string = "85a10227ade9d331f44cfc8d78904004"
var accessToken string = ""
var openId string = "omKA5uALKWNYhZk0CYN7HD43RXaw"
var tmpId string = "a2m_d3YqUBeayXo5Tr5txsIiI68yRJHO71B2SdCv1Es"
var weiApi *WeiChatAPI

func init() {
	//获取token
	weiApi = WeiChat(nil, nil)
	accessToken, err := weiApi.AccessToken(appId, secret)
}

//发送图文消息
func TestCustomerNewsMsg(t *testing.T) {
	newMsg := NewCustomerNewsMsg(openId, "测试New", "测试测试", "www.baidu.com", "")
	ret, err := weiApi.SendCustomMsg(accessToken, newMsg)
	fmt.Println(ret, err)

}

//发送客服文字消息
func TestCustomerTextMsg(t *testing.T) {
	textMsg := NewCustomerTextMsg(openId, "test TextMsg")
	ret, err := weiApi.SendCustomMsg(accessToken, textMsg)
	fmt.Println(ret, err)
}

//发送客服图片消息
func TestCustomerImageMsg(t *testing.T) {
	token, err := weiApi.AccessToken(appId, secret)
	imageMsg := NewCustomerImageMsg(openId, "5deENy942-cX15WuZsQJQSskcbpIDiOpFkriV2_a89I3Kwod9uWPtJ0_2XLVK4jZ")
	ret, err := weiApi.SendCustomMsg(token, imageMsg)
	fmt.Println(ret, err)
}

//发送模板消息
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

//获取认证code
func TestCode(t *testing.T) {
	err := weiApi.Code(appId, "http://xxx.com")
	fmt.Println(err)
}
func receiveDealFun(msg WxReceiveCommonMsg) error {
	fmt.Println(msg)
	return nil
}

func TestAuthByCode(t *testing.T) {}

//上传临时素材
func TestUploadTemporaryMedia(t *testing.T) {
	resp, err := weiApi.UploadTemporaryMedia(accessToken, "/tmp/2.png", ImageMedia)
	fmt.Println(resp, err)
}

//上传永久图片
func TestUploadPermanentMedia(t *testing.T) {
	resp, err := weiApi.UploadPermanentMedia(accessToken, "/tmp/2.png", ImageMedia)
	fmt.Println(resp, err)
}

//生成二维
func TestGenerateQRCode(t *testing.T) {
	resp, err := weiApi.GenerateQRCode(accessToken, 1223, QR_LIMIT_STR_SCENE)
	fmt.Println(resp, err)
}

//长链接转换为短链接
func TestShortUrl(t *testing.T) {
	resp, err := weiApi.LongTransformShortUrl(accessToken, "http://wwwxx")
	fmt.Println(resp, err)
}
