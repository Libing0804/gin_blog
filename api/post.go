package api

import (
	"blog_gin/dao"
	"blog_gin/models"
	"blog_gin/service"
	"blog_gin/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type SaveP struct {
	CategoryId string `json:"categoryId"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	Slug       string `json:"slug"`
	Title      string `json:"title"`
	Type       string `json:"type"`
}

func (*Api) DeletePost(c *gin.Context) {
	pidStr := c.Param("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		log.Println("获取文章pid失败")
		return
	}

	dao.DeletePost(pid)
	post, err := dao.GetAllPost()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  -999,
			"error": err,
			"data":  "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"error": "",
		"data":  post,
	})
}
func (*Api) GetPost(c *gin.Context) {
	pidStr := c.Param("pid")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		log.Println("获取文章pid失败")
		return
	}

	Post, err := dao.GetPostById(pid)
	if err != nil {
		c.HTML(http.StatusOK, "detail.html", "文章查询失败")
		return
	}

	c.JSON(200, gin.H{
		"code":  200,
		"error": "",
		"data":  Post,
	})

	//c.HTML(http.StatusOK, "detail.html", PostRes)
}
func (*Api) SavePost(c *gin.Context) {
	//获取用户id判断是否登录
	token := c.Request.Header.Get("Authorization")

	_, claim, err := utils.ParseToken(token)
	if err != nil {
		c.HTML(200, "writing.html", err)
		return
	}
	uid := claim.Uid
	var pv SaveP

	c.ShouldBind(&pv)

	//	post save

	cid, _ := strconv.Atoi(pv.CategoryId)

	typeId, _ := strconv.Atoi(pv.Type)
	post := &models.Post{
		Pid:        -1,
		Title:      pv.Title,
		Slug:       pv.Slug,
		Content:    pv.Content,
		Markdown:   pv.Markdown,
		CategoryId: cid,
		UserId:     uid,
		ViewCount:  0,
		Type:       typeId,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}
	service.SavePost(post)

	c.JSON(200, gin.H{
		"code":  200,
		"error": "这里没",
		"data":  post,
	})
}
func (*Api) UpdatePost(c *gin.Context) {
	//	post 更新
	//获取用户id判断是否登录
	token := c.Request.Header.Get("Authorization")

	_, claim, err := utils.ParseToken(token)
	if err != nil {
		c.HTML(200, "writing.html", err)
		return
	}
	uid := claim.Uid
	var pv SaveP

	c.ShouldBind(&pv)

	//	post save

	cid, _ := strconv.Atoi(pv.CategoryId)
	pidStr := c.Param("pid")
	pid, err := strconv.Atoi(pidStr)
	typeId, _ := strconv.Atoi(pv.Type)
	post := &models.Post{
		Pid:        pid,
		Title:      pv.Title,
		Slug:       pv.Slug,
		Content:    pv.Content,
		Markdown:   pv.Markdown,
		CategoryId: cid,
		UserId:     uid,
		ViewCount:  0,
		Type:       typeId,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}
	service.UpdatePost(post)
	c.JSON(200, gin.H{
		"code":  200,
		"error": "这里没错",
		"data":  post,
	})
}
