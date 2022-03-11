package views

import (
	"blog_gin/models"
	"blog_gin/service"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Index(c *gin.Context) {

	//获取表单
	pageStr := c.Param("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10
	//去数据库获取需要的数据
	pathurl := c.Request.RequestURI
	slug := strings.TrimPrefix(pathurl, "/")

	var err error
	var hr *models.HomeResponse

	hr, err = service.GetAllIndexInfo(slug, page, pageSize)

	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "index.html", errors.New("系统错误"))
		return
	}
	c.HTML(http.StatusOK, "index.html", hr)

}
