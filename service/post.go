package service

import (
	"blog_gin/config"
	"blog_gin/dao"
	"blog_gin/models"
	"html/template"
	"log"
)

func SavePost(post *models.Post) {
	dao.SavePost(post)
}
func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}
func GetDetailPost(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pid)

	if err != nil {
		log.Println("查询数据库失败")
		return nil, err
	}
	categoryName, err := dao.GetCategoryById(post.CategoryId)
	if err != nil {
		log.Println("显示文章详情时分类名查询失败")
	}
	UserName, err := dao.GetUserNameById(post.UserId)
	if err != nil {
		log.Println("显示文章详情时用户名查询失败")
	}
	var postMore models.PostMore
	postMore.Pid = post.Pid
	postMore.Title = post.Title

	postMore.Slug = post.Slug
	postMore.Content = template.HTML(post.Content)
	postMore.CategoryName = categoryName
	postMore.CategoryId = post.CategoryId
	postMore.UserId = post.UserId
	postMore.UserName = UserName
	postMore.ViewCount = post.ViewCount
	postMore.Type = post.Type
	postMore.CreateAt = post.CreateAt.Format("2006-01-02 15:04:05")
	postMore.UpdateAt = post.UpdateAt.Format("2006-01-02 15:04:05")
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}

	return postRes, nil
}

func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}
	wr.Categorys = category
	return
}
