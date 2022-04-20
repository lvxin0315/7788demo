package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func HttpHandle(url, method string, data map[string]interface{}, cookies []http.Cookie, headers map[string]string) (res []byte, err error) {
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
