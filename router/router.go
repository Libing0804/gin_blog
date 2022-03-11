package router

import (
	"blog_gin/api"
	"blog_gin/models"
	"blog_gin/views"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	ru := models.InitRouter()

	//页面上的数据必须有定义

	ru.GET("/", views.Html.Index)
	ru.GET("/:slug", views.Html.Index)

	//http://localhost:8008/c/1 后面这个1是分类的id我们当成参数就行
	ru.GET("/c/:id", views.Html.Category)
	ru.GET("/p/:pid", views.Html.Detail)
	//登录接口
	ru.GET("/login", views.Html.Login)

	ru.POST("/api/v1/login", api.API.Login)
	//搜索
	ru.GET("/api/v1/post/search", api.API.Search)
	//写文章
	ru.GET("/writing", views.Html.Writing)
	//提交写的文章
	ru.POST("/api/v1/post", api.API.SavePost)
	//更新
	ru.PUT("/api/v1/post", api.API.UpdatePost)
	ru.GET("/api/v1/post/:pid", api.API.GetPost)
	ru.DELETE("/api/v1/post/:pid", api.API.DeletePost)
	//上传图片
	ru.GET("/api/v1/qiniu/token", api.API.QiniuToken)
	//归档
	ru.GET("/pigeonhole", views.Html.Pigeonhole)
	return ru
}
