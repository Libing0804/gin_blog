package views

import (
	"blog_gin/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Detail(c *gin.Context) {
	pidstr := c.Param("pid")
	pidstr = strings.TrimSuffix(pidstr, ".html")
	pid, err := strconv.Atoi(pidstr)
	if err != nil {
		log.Println("文章id截取失败")
		return
	}

	PostRes, err := service.GetDetailPost(pid)
	if err != nil {
		c.HTML(http.StatusOK, "detail.html", "文章查询失败")
		return
	}
	c.HTML(http.StatusOK, "detail.html", PostRes)
}
