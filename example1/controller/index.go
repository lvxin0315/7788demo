package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lvxin0315/7788demo/example1/service"
	"net/http"
)

// Index 首页
func Index(c *gin.Context) {
	indexData := service.GetIndexDataService()
	c.HTML(http.StatusOK, "index/index.html", map[string]interface{}{
		"NavList":   service.GetNavService("/"),
		"IndexData": indexData,
	})
}

func About(c *gin.Context) {
	c.HTML(http.StatusOK, "index/about.html", map[string]interface{}{
		"NavList": service.GetNavService("/about"),
	})
}

func Courses(c *gin.Context) {
	c.HTML(http.StatusOK, "index/courses.html", map[string]interface{}{
		"NavList": service.GetNavService("/courses"),
	})
}

func Contact(c *gin.Context) {
	c.HTML(http.StatusOK, "index/contact.html", map[string]interface{}{
		"NavList": service.GetNavService("/contact"),
	})
}
