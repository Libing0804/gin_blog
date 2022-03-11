package models

import (
	"blog_gin/config"
	"html/template"
	"time"
)

//这里是文章的定义结构
//数据库用的
type Post struct {
	Pid        int       `json:"pid"`        //文章id
	Title      string    `json:"title"`      //标题
	Slug       string    `json:"slug"`       // 自定义页面path
	Content    string    `json:"content"`    //文字html
	Markdown   string    `json:"markdown"`   //文章的markdown
	CategoryId int       `json:"categoryId"` //分类id
	UserId     int       `json:"userId"`     //用户id
	ViewCount  int       `json:"viewCount"`  //查看次数
	Type       int       `json:"type"`       //文章类型
	CreateAt   time.Time `json:"createAt"`   //创建时间
	UpdateAt   time.Time `json:"updateAt"`   //更新时间
}

type PostMore struct {
	Pid          int           `json:"pid"`          //文章id
	Title        string        `json:"title"`        //标题
	Slug         string        `json:"slug"`         // 自定义页面path
	Content      template.HTML `json:"content"`      //文字html
	CategoryName string        `json:"categoryName"` //费雷名
	CategoryId   int           `json:"categoryId"`   //分类id
	UserId       int           `json:"userId"`       //用户id
	UserName     string        `json:"userName"`     //用户id
	ViewCount    int           `json:"viewCount"`    //查看次数
	Type         int           `json:"type"`         //文章类型
	CreateAt     string        `json:"createAt"`     //创建时间
	UpdateAt     string        `json:"updateAt"`     //更新时间
}

//请求
type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"`
}

//搜索
type SearchResp struct {
	Pid   int    `json:"pid"`
	Title string `json:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}

type WritingRes struct {
	Title     string
	CdnURL    string
	Categorys []Category
}
type PigeonholeRes struct {
	config.Viewer
	config.SystemConfig
	Categorys []Category
	Lines     map[string][]Post
}
