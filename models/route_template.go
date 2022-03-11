package models

import (
	"blog_gin/common"
	"github.com/gin-gonic/gin"
	"html/template"
)

func InitRouter()*gin.Engine{
	ru:=gin.Default()
	//解析文件之前要加载静态文件‘
	ru.Static("/resource","./public/resource")
	//应用模板函数  在解析之前加上
	ru.SetFuncMap(template.FuncMap{
		"isODD": common.IsODD,
		"getNextName": common.GetNextName,
		"date": common.Date,
		"dateDay": common.DateDay,
	})
	ru.LoadHTMLGlob("./template/*")
	return  ru
}

