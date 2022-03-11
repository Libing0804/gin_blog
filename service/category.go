package service

import (
	"blog_gin/config"
	"blog_gin/dao"
	"blog_gin/models"
	"html/template"
	"log"
)

func GetpostByCategoryId(cid, page, pageSize int) (*models.CategoryResponse, error) {
	categorys, err := dao.GetAllCategory()

	if err != nil {
		log.Println(err)
		return nil, err

	}
	var postMores []models.PostMore
	posts, err := dao.GetPostPageByCategory(cid, page, pageSize)

	for _, post := range posts {
		categoryName, err := dao.GetCategoryById(post.CategoryId)
		if err != nil {
			log.Println("分类名查询失败")
		}
		UserName, err := dao.GetUserNameById(post.UserId)
		if err != nil {
			log.Println("用户名查询失败")
		}
		//页面展示的时候太乱了  view展示页面只显示部分内容
		contentTemp := []rune(post.Content)
		if len(contentTemp) > 30 {

			post.Content = string(contentTemp[:30])
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
		postMores = append(postMores, postMore)
	}
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//查询总文章数 和页数
	total := dao.CountGetPostByCategory(cid)
	allPages := (total-1)/10 + 1
	var pages = []int{}
	for i := 0; i < allPages; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != allPages,
	}
	categoryName, err := dao.GetCategoryById(cid)
	if err != nil {
		log.Println("categoryName 获取失败")

	}
	categoryRes := &models.CategoryResponse{
		HomeResponse: hr,
		CategoryName: categoryName,
	}
	return categoryRes, nil
}
