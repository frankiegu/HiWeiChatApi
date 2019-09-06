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
}

// type ATest struct {
// 	Id string
// }

// func (a *ATest) aa() {
// 	json, _ := json.Marshal(a)
// 	fmt.Println(string(json))
// }

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
func TestAccessToken(t *testing.T) {
	token, err := weiApi.AccessToken(appId, secret)
	fmt.Println(token, err)
}
func TestCode(t *testing.T) {
	err := weiApi.Code(appId, "http://smart03.com")
	fmt.Println(err)
}
func receiveDealFun(msg WxReceiveCommonMsg) error {
	fmt.Println(msg)
	return nil
}
func TesReceiveMsg(t *testing.T) {
	data := `<?xml version="1.0" encoding="UTF-8"?>
	<xml>
	  <ToUserName>231</ToUserName>
	  <FromUserName>4444</FromUserName>
	  <CreateTime>1348831860</CreateTime>
	  <MsgType>xvv</MsgType>
	  <Content>cvc</Content>
	  <MsgId>1234567890123456</MsgId>
	</xml>`
	xmlByte := []byte(data)
	weiApi.ReceiveFunc(receiveDealFun)
	weiApi.ReceiveCommonMsg(xmlByte)
}

func TestAuthByCode(t *testing.T) {}

func TestXML(t *testing.T) {
	// 	data := `<?xml version="1.0" encoding="UTF-8"?>
	// <xml>
	//   <ToUserName>231</ToUserName>
	//   <FromUserName>4444</FromUserName>
	//   <CreateTime>1348831860</CreateTime>
	//   <MsgType>xvv</MsgType>
	//   <Content>cvc</Content>
	//   <MsgId>1234567890123456</MsgId>
	// </xml>`
	// 	xmlByte := []byte(data)

	// v := WxReceiveCommonMsg{}
	// err := xml.Unmarshal(xmlByte, &v)
	// err := xml.Unmarshal(xmlByte, &v)
	// fmt.Println(err)
	// fmt.Println(v.MsgId)
	// fmt.Println("SmtpServer : ", v.SmtpServer)
	// fmt.Println("SmtpPort : ", v.SmtpPort)
	// fmt.Println("Sender : ", v.Sender)
	// fmt.Println("SenderPasswd : ", v.SenderPasswd)
	// fmt.Println("Receivers.Flag : ", v.Receivers.Flag)
	// for i, element := range v.Receivers.User {
	// 	fmt.Println(i, element)
	// }
}

// func TestPostFile(t *testing.T) {
// 	data := map[string]string{
// 		"age": "15",
// 	}
// 	_, err := PostFile("http://smart03.com/test", data, "testFile", "E:/2.png")
// 	fmt.Println(err)
// }

func TestUploadTemporaryMedia(t *testing.T) {
	token, err := weiApi.AccessToken(appId, secret)
	fmt.Println(token, err)
	resp, err := weiApi.UploadTemporaryMedia(token, "/tmp/2.png", ImageMedia)
	fmt.Println(resp, err)
}
