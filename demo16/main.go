package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const maitixURL = "https://cloud.maitix.com"

const (
	getToken = "/getToken"
	projects = "/tic/sale/projects"
)

var userAgent = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36`
var headers = map[string]string{
	"User-Agent":       userAgent,
	"x-requested-with": "XMLHttpRequest",
	"Content-Type":     "application/json;charset=UTF-8",
}

// {"headerName":"X-XSRF-TOKEN","parameterName":"_csrf","token":"3691462b-3154-474e-8692-d92458d69262"}
type csrfToken struct {
	HeaderName    string `json:"headerName"`
	ParameterName string `json:"parameterName"`
	Token         string `json:"token"`
}

func main() {
	cookieData, err := _readCookie("")
	if err != nil {
		panic(err)
	}

	// token
	tokenRes, err := httpHandle(maitixURL+getToken, http.MethodGet, nil, cookieData, headers)
	if err != nil {
		panic(err)
	}
	var ct csrfToken
	err = json.Unmarshal(tokenRes, &ct)
	if err != nil {
		panic(err)
	}
	// 拉取项目
	ctHeaders := headers
	ctHeaders["X-XSRF-TOKEN"] = ct.Token
	projectsRes, err := httpHandle(maitixURL+projects, http.MethodPost, map[string]interface{}{
		"pageNo":   1,
		"pageSize": 20,
	}, cookieData, ctHeaders)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(projectsRes))
}

// 读cookie文件
func _readCookie(cookieFileName string) ([]http.Cookie, error) {
	b, err := ioutil.ReadFile("/Users/lvxin/PycharmProjects/piao_lock/data/cookies/xuezhi001_cookies.txt")
	if err != nil {
		return nil, err
	}
	//fmt.Println(string(b))
	var data []http.Cookie
	err = json.Unmarshal(b, &data)
	//fmt.Println(data)
	return data, err
}

// request
func httpHandle(url, method string, data map[string]interface{}, cookies []http.Cookie, headers map[string]string) (res []byte, err error) {
	client := &http.Client{}
	var reqBody io.Reader
	if len(data) > 0 {
		b, _ := json.Marshal(data)
		reqBody = bytes.NewBuffer(b)
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		fmt.Println("http.NewRequest:", err)
		return
	}
	//可以添加多个cookie
	for _, cookie := range cookies {
		req.AddCookie(&cookie)
	}

	//添加header，key为X-Xsrftoken，value为b6d695bbdcd111e8b681002324e63af81
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client.Do:", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	res = b
	return
}
