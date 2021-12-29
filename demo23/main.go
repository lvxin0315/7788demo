package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var h1 = "http://www.baidu.com"
var h2 = "http://127.0.0.1"

type AppHandler struct {
}

func (h *AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	appUrl, err := url.Parse(h1)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.RequestURI)
	if r.RequestURI == "/" {
		fmt.Println(appUrl.String())
		proxy := httputil.NewSingleHostReverseProxy(appUrl)
		proxy.ServeHTTP(w, r)
	} else {
		appUrl, _ = url.Parse(h2)
		fmt.Println(appUrl.String())
		proxy := httputil.NewSingleHostReverseProxy(appUrl)
		proxy.ServeHTTP(w, r)
	}
}

func main() {
	http.ListenAndServe(":8001", new(AppHandler))
}
