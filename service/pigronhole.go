package service

import (
	"blog_gin/config"
	"blog_gin/dao"
	"blog_gin/models"
)

func FindPostPigronhole() models.PigeonholeRes {
	//	查询所有  按照月份分类
	categorys, _ := dao.GetAllCategory()
	posts, _ := dao.GetAllPost()
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	return models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        pigeonholeMap,
	}
}
