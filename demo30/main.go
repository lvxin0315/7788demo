package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	ginEngine := gin.Default()
	ginEngine.Any("/u", uploadFile)

	ginEngine.Run(":8099")
}

func uploadFile(c *gin.Context) {
	fmt.Println("接收到请求：", time.Now().Format("2006-01-02 15:04:05"))
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		panic(err)
	}
	fmt.Println(header.Size, "B")
	fmt.Println(header.Size>>10, "KB")
	fmt.Println(header.Size>>10>>10, "MB")
	fmt.Println(header.Size>>10>>10>>10, "GB")

	err = c.SaveUploadedFile(header, "./大文件")
	if err != nil {
		panic(err)
	}

	fmt.Println("完成文件处理：", time.Now().Format("2006-01-02 15:04:05"))
	c.JSON(http.StatusOK, "ok")
}
