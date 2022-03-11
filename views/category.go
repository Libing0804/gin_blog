package views

import (
	"blog_gin/service"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (*HTMLApi) Category(c *gin.Context) {
	cidStr := c.Param("id")
	cid, err := strconv.Atoi(cidStr)
	if err != nil {
		log.Println("获取传输的id出错")
		c.HTML(200, "category.html", errors.New("分类id查询错误"))
		return
	}
	//获取表单
	pageStr := c.Param("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	//每页显示的数量
	pageSize := 10
	//去数据库获取需要的数据
	categoryResponse, err := service.GetpostByCategoryId(cid, page, pageSize)
	if err != nil {
		c.HTML(http.StatusOK, "category.html", err)

		return
	}
	c.HTML(http.StatusOK, "category.html", categoryResponse)
}
