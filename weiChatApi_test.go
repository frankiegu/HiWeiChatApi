package HiWeiChatApi

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	_ "strings"
	"testing"
)

var appId string = ""
var secret string = ""
var accessToken string = ""
var openId string = "openxxxsssaaa"

func TestCustomerNewsMsg(t *testing.T) {

}
func TestCustomerTextMsg(t *testing.T) {
	resp := `{"errcode" : 0,"errmsg" : "ok"}`
	by := []byte(resp)

	wxRe := WeiChatResponse{}
	json.Unmarshal(by, &wxRe)
	fmt.Println(wxRe)
	// textMsg := NewCustomerTextMsg(openId, "test TextMsg")
	// data, err := textMsg.Data()
	// fmt.Println(string(data), err)
}
func TestCustomerImageMsg(t *testing.T) {

}
func TestTemplateMsg(t *testing.T) {

}
func TestAccessToken(t *testing.T) {

}
func TestCode(t *testing.T) {

}

func TestAuthByCode(t *testing.T) {}

func TestSendMsg(t *testing.T) {
	/*
		  tmpId := "a2m_d3YqUBeayXo5Tr5txsIiI68yRJHO71B2SdCv1Es"
		  openId := "omKA5uALKWNYhZk0CYN7HD43RXaw"
		  wei := WeiChat("CPS")
		  sendData := map[string]interface{} {
		 	"keyword1":"hello",
			"keyword2":"world",
			"remark":"END",

		  }
		  accessToken,err := wei.AccessToken("wxacb25453422b57a7","85a10227ade9d331f44cfc8d78904004")
		  json,err:= wei.GenerateWxTemplateMsgFormat(openId,tmpId,sendData)
		  wei.SendTemplateMsg(accessToken,json)
	*/
}

func TestWeiChat(t *testing.T) {
	//s := strings.Split("http://wwxx/ss/aa?a=1&b=2", "?")
	//fmt.Println(s)
	//params := map[string]string{
	//	"h": "2",
	//	"z": "3",
	//}
	//Get("http://wwxx/ss/aa?a=1&b=2", params)
	// content := map[string]interface{}{
	// 	"first":  "hello",
	// 	"remark": "world",
	// }
	// jsonData, _ := json.Marshal(content)
	// data := map[string]string{
	// 	"openId": "xxxxxxxxx",
	// 	"cotent": string(jsonData),
	// }
	//
	//wei := WeiChat("CPS1")
	//wei.Code("wxacb25453422b57a7","http://smart03.com/login")
	//fmt.Println(data, err)
	//fmt.Printf("%T", jsonData)
	//
}

// type SConfig struct {
// 	XMLName      xml.Name   `xml:"config"`
// 	SmtpServer   string     `xml:"smtpServer"`
// 	SmtpPort     int        `xml:"smtpPort"`
// 	Sender       string     `xml:"sender"`
// 	SenderPasswd string     `xml:"senderPasswd"`
// 	Receivers    SReceivers `xml:"receivers"`
// }
// type SReceivers struct {
// 	Flag string   `xml:"flag,attr"`
// 	User []string `xml:"user"`
// }
/*
type WxReceiveCommonMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	MsgId        int64
	PicUrl       string
	MediaId      string
	Format       string
	Recognition  string
	ThumbMediaId string
}
*/
func TestXML(t *testing.T) {
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

	v := WxReceiveCommonMsg{}
	err := xml.Unmarshal(xmlByte, &v)
	// err := xml.Unmarshal(xmlByte, &v)
	fmt.Println(err)
	fmt.Println(v.MsgId)
	// fmt.Println("SmtpServer : ", v.SmtpServer)
	// fmt.Println("SmtpPort : ", v.SmtpPort)
	// fmt.Println("Sender : ", v.Sender)
	// fmt.Println("SenderPasswd : ", v.SenderPasswd)
	// fmt.Println("Receivers.Flag : ", v.Receivers.Flag)
	// for i, element := range v.Receivers.User {
	// 	fmt.Println(i, element)
	// }
}
