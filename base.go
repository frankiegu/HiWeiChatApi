package HiWeiChatApi

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//Post请求
func Post(url string, paramBody []byte, header map[string]string) ([]byte, error) {
	client := &http.Client{}
	paramsData := bytes.NewBuffer(paramBody)
	req, err := http.NewRequest("POST", url, paramsData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if header != nil {
		for hkey, hval := range header {
			req.Header.Set(hkey, hval)
		}
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}

/**Post上传文件
 * [PostFile description]
 * @param {[type]} url           string            [description]
 * @param {[type]} params        map[string]string [description]
 * @param {[type]} fileFieldName [description]
 * @param {[type]} path          string)           ([]byte,      error [description]
 //结束函数
 func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	reader,err := r.MultipartReader()
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	for {
		part,err := reader.NextPart()
		if err == io.EOF {
			break;
		}
		fmt.Printf("FileName=[%s],FormName[%s]\n",part.FileName(),part.FormName())
		if part.FileName() == "" {  //非文件属性值
		  data,_ := ioutil.ReadAll(part)
		  fmt.Printf("FormData=[%s]\n",string(data))
		} else { //文件
			dst,_ := os.Create("/tmp/"+part.FileName()+".upload")
			defer dst.Close()
			io.Copy(dst,part)
		}
	}
	fmt.Fprintf(w, "OK")
}
*/
func PostFile(url string, params map[string]string, fileFieldName, path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileFieldName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", url, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	return respBody, err
}

//Get请求
func Get(url string, params map[string]string) ([]byte, error) {
	paramsStr := ""
	if params != nil && len(params) > 0 {
		for k, v := range params {
			if paramsStr != "" {
				paramsStr = fmt.Sprintf("%s&%s=%s", paramsStr, k, v)
			} else {
				paramsStr = fmt.Sprintf("%s=%s", k, v)
			}
		}
	}
	urls := strings.Split(url, "?")
	targetUrl := url
	if len(urls) > 1 && paramsStr != "" {
		targetUrl = fmt.Sprintf("%s&%s", url, paramsStr)
	} else if paramsStr != "" {
		targetUrl = fmt.Sprintf("%s?%s", url, paramsStr)
	}

	resp, err := http.Get(targetUrl)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return body, nil
}
