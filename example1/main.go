package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/7788demo/example1/controller"
)

func main() {
	// 初始化gin框架
	engine := gin.New()
	// 加载静态资源
	static(engine)
	// 加载路由
	route(engine)
	// 启动服务，端口 8090
	engine.Run(":8090")
}

// 静态资源处理
func static(engine *gin.Engine) {
	// css、js、图片等文件的静态目录
	engine.Static("/assets", "assets")
	// html 模板目录
	engine.LoadHTMLGlob("view/**/*")
}

// 路由
func route(engine *gin.Engine) {
	// html
	engine.GET("/", controller.Index)
	engine.GET("/index", controller.Index)
	engine.GET("/about", controller.About)
	engine.GET("/courses", controller.Courses)
	engine.GET("/contact", controller.Contact)
}
