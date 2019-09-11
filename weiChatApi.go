package HiWeiChatApi

import (
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

type WeiChatAPI struct {
	log          *log.Logger
	receiveFunc  WxReceiveFunc
	qrExpireTime int
}

func WeiChat(logger *log.Logger, revFunc WxReceiveFunc) *WeiChatAPI {
	if logger == nil {
		// logger = log.New(l.dest, , log.Ldate|log.Ltime|log.Lshortfile)
	}
	return &WeiChatAPI{
		log:          logger,
		receiveFunc:  revFunc,
		qrExpireTime: 2592000,
	}
}

//普通消息回调函数
func (api *WeiChatAPI) ReceiveFunc(revFunc WxReceiveFunc) {
	api.receiveFunc = revFunc
}

/**接收普通消息
func xx(w http.ResponseWriter, r *http.Request) {
 	defer r.Body.Close()
	data, _ := ioutil.ReadAll(r.Body) //获取post的数据
	weiApi := HiWeiChatApi.WeiChat(nil, nil)
	var rfun HiWeiChatApi.WxReceiveFunc = func(msg HiWeiChatApi.WxReceiveCommonMsg) error{
		fmt.Println("hellow")
                fmt.Printf("%v",msg)
		return
	}
	weiApi.ReceiveFunc(rfun)
	resp,err:=weiApi.ReceiveCommonMsg(data)
}
*/
func (api *WeiChatAPI) ReceiveCommonMsg(msgData []byte) (WxReceiveCommonMsg, error) {

	msg := WxReceiveCommonMsg{}
	err := xml.Unmarshal(msgData, &msg)
	if api.receiveFunc == nil {
		return msg, err
	}
	err = api.receiveFunc(msg)
	return msg, err
}

//获取AccessToken
func (api *WeiChatAPI) AccessToken(appId, secret string) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", appId, secret)
	body, err := Get(url, nil)
	if err != nil {
		return "", err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(body, &result)
	return fmt.Sprintf("%v", result["access_token"]), err
}

//发送模板消息
func (api *WeiChatAPI) SendTemplateMsg(accessToken string, wxt *WxTemplateMsg) (WeiChatResponse, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken)
	json, err := wxt.Marshal()
	if err != nil {
		return api.responseParse(nil), err
	}
	body, err := Post(url, json, nil)
	return api.responseParse(body), err
}

//发送客服消息
func (api *WeiChatAPI) SendCustomMsg(accessToken string, msg *CustomerMsg) (WeiChatResponse, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", accessToken)
	data, err := msg.Marshal()
	if err != nil {
		return api.responseParse(nil), err
	}

	ret, err := Post(url, data, nil)
	if err != nil {
		return api.responseParse(nil), err
	}
	return api.responseParse(ret), err
}

/**域名签名验证
 	func xx(w http.ResponseWriter, r *http.Request) {
	  	timestamp := strings.Join(r.Form["timestamp"], "")
		nonce := strings.Join(r.Form["nonce"], "")
		signatureGen := weiApi.MakeSignature("smart03", timestamp, nonce)
		signatureIn := strings.Join(r.Form["signature"], "")
		if signatureGen != signatureIn {
			return
		}
		echostr := strings.Join(r.Form["echostr"], "")
		fmt.Fprintf(w, echostr)
 	}

*/
func (api *WeiChatAPI) MakeSignature(token, timestamp, nonce string) string {
	strs := []string{token, timestamp, nonce}
	sort.Strings(strs)
	sha := sha1.New()
	io.WriteString(sha, strings.Join(strs, ""))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

//认证code获取
func (api *WeiChatAPI) Code(appId, redirectUrl string) error {

	url := fmt.Sprintf("https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=1#wechat_redirect",
		appId, redirectUrl)
	_, err := Get(url, nil)
	return err
}

//用户认知
func (api *WeiChatAPI) AuthAccessToken(appId, secret, code string) string {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appId, secret, code)
	body, _ := Get(url, nil)
	return string(body)
}

func (api *WeiChatAPI) responseParse(resp []byte) WeiChatResponse {
	wxRe := WeiChatResponse{}
	if resp == nil || len(resp) < 1 {
		return wxRe
	}
	json.Unmarshal(resp, &wxRe)
	return wxRe
}

//上传临时素材
func (api *WeiChatAPI) UploadTemporaryMedia(accessToken string, file string, mediaType MediaType) (MediaReponse, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/media/upload?access_token=%s&type=%s", accessToken, mediaType)
	meRe := MediaReponse{}
	body, err := PostFile(url, nil, "media", file)
	if err != nil {
		return meRe, err
	}
	json.Unmarshal(body, &meRe)
	return meRe, nil
}

//上传永久素材
func (api *WeiChatAPI) UploadPermanentMedia(accessToken string, file string, mediaType MediaType) (MediaReponse, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s", accessToken, mediaType)
	meRe := MediaReponse{}
	body, err := PostFile(url, nil, "media", file)
	if err != nil {
		return meRe, err
	}
	json.Unmarshal(body, &meRe)
	return meRe, nil
}

//生成参数二维码
func (api *WeiChatAPI) GenerateQRCode(accessToken string, scene interface{}, qrType QRCodeType) (QRCodeResponse, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s", accessToken)
	qrStr := fmt.Sprintf("%s", qrType)
	postData := make(map[string]interface{}, 3)
	postData["action_name"] = qrStr
	postData["expire_seconds"] = api.qRCodeExpireTime()

	keyMap := map[QRCodeType]string{
		QR_SCENE:           "scene_id",
		QR_STR_SCENE:       "scene_str",
		QR_LIMIT_SCENE:     "scene_id",
		QR_LIMIT_STR_SCENE: "scene_str",
	}

	qrResp := QRCodeResponse{}
	key := keyMap[qrType]

	if qrStr == "QR_SCENE" || qrStr == "QR_STR_SCENE" {
		postData["action_info"] = map[string]map[string]interface{}{
			"scene": map[string]interface{}{key: scene},
		}
	} else if qrStr == "QR_LIMIT_SCENE" || qrStr == "QR_LIMIT_STR_SCENE" {
		postData["action_info"] = map[string]map[string]interface{}{
			"scene": map[string]interface{}{key: scene},
		}
	} else {
		return qrResp, fmt.Errorf("undefined qrType:%s", qrStr)
	}

	postDataByte, err := json.Marshal(postData)
	if err != nil {
		return qrResp, err
	}

	body, err := Post(url, postDataByte, nil)
	if err != nil {
		return qrResp, err
	}
	json.Unmarshal(body, &qrResp)
	return qrResp, nil
}

//设置短期二维码有效时间
func (api *WeiChatAPI) SetQRCodeExpireTime(tm int) {
	api.qrExpireTime = tm
}

//获取短期二维码有效时间
func (api *WeiChatAPI) qRCodeExpireTime() int {
	return api.qrExpireTime
}

//长链接转换为短链接
func (api *WeiChatAPI) LongTransformShortUrl(accessToken, longUrl string) (ShortUrlResponse, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s", accessToken)
	postData := make(map[string]string, 2)
	postData["action"] = "long2short"
	postData["long_url"] = longUrl
	postDataByte, err := json.Marshal(postData)
	shortResp := ShortUrlResponse{}
	if err != nil {
		return shortResp, err
	}
	fmt.Println(string(postDataByte), err)
	body, err := Post(url, postDataByte, nil)
	fmt.Println(string(body), err)
	json.Unmarshal(body, &shortResp)
	return shortResp, nil
}
