package api

import (
	"blog_gin/dao"
	"fmt"
	"github.com/gin-gonic/gin"
)

type SearchRes struct {
	Pid   int    `json:"pid"`
	Title string `json:"title"`
}

func (*Api) Search(c *gin.Context) {
	search := c.Query("val")
	//获取表单

	//每页显示的数量
	posts, err := dao.GetPostPageBySearch(search)
	if len(posts) == 0 {
		fmt.Println("不对")
	}

	if err != nil {
		fmt.Println("********************")
		c.JSON(200, gin.H{
			"code": 501,
			"err":  err,
			"data": "",
		})
		return
	}
	var searchRes []SearchRes
	for _, v := range posts {
		searchRes = append(searchRes, SearchRes{v.Pid, v.Title})
	}
	c.JSON(200, gin.H{
		"code":  200,
		"error": "",
		"data":  searchRes,
	})
}
