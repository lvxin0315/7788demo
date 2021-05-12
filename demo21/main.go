package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"math/rand"
	"time"
)

// StartChrome 启动谷歌浏览器headless模式

func StartChrome() {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		//"profile.managed_default_content_settings.images": 2,
	}
	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			//"--headless", // 设置Chrome无头模式
			//"--no-sandbox",
			"--disable-gpu",
			"disable-infobars",
			//"--window-size=1280x1024'",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/604.4.7 (KHTML, like Gecko) Version/11.0.2 Safari/604.4.7", // 模拟user-agent，防反爬
		},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	service, err := selenium.NewChromeDriverService("./chromedriver", 9515, opts...)
	if err != nil {
		panic(err)
	}
	defer service.Stop() // 停止chromedriver
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		panic(err)
	}
	defer webDriver.Quit() // 关闭浏览器

	//目标网站
	targetUrl := "https://cloud.maitix.com"
	// 导航到目标网站
	err = webDriver.Get(targetUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}

	time.Sleep(5 * time.Second)
	log.Println(webDriver.GetCookies())
	//webDriver.MaximizeWindow("")
	// 操作
	loginFrame, err := webDriver.FindElement(selenium.ByID, "alibaba-login-box")
	if err != nil {
		panic(err)
	}
	webDriver.SwitchFrame(loginFrame)
	loginIdEle, err := webDriver.FindElement(selenium.ByCSSSelector, "#fm-login-id")
	if err != nil {
		fmt.Println(loginIdEle)
		panic(err)
	}

	passwordEle, err := webDriver.FindElement(selenium.ByCSSSelector, "#fm-login-password")
	if err != nil {
		panic(err)
	}
	time.Sleep(1 * time.Second)
	// 输入账号密码
	loginIdEle.SendKeys("asdasdasd")
	passwordEle.SendKeys("123456789")
	// 点击登录
	loginButtonEle, err := webDriver.FindElement(selenium.ByCSSSelector, "#fm-login-submit")
	if err != nil {
		panic(err)
	}
	loginButtonEle.Click()
	// 移动滑块
	time.Sleep(time.Second)
	huaKuaiEle, err := webDriver.FindElement(selenium.ByXPATH, "//*[@id=\"nc_1_n1z\"]")
	if err != nil {
		panic(err)
	}

	moveX := 0
	for moveX < 250 {
		webDriver.ButtonDown()
		moveX += 1 + rand.Intn(5)
		huaKuaiEle.MoveTo(moveX, 0)
		fmt.Println("MoveTo:", moveX)
		time.Sleep(10 * time.Millisecond)
	}
	webDriver.ButtonUp()

	time.Sleep(time.Minute)

}

func main() {
	StartChrome()
}
