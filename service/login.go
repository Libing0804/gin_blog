package service

import (
	"blog_gin/dao"
	"blog_gin/models"
	"blog_gin/utils"
	"errors"
	"log"
)

func Login(username, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "mszlu")
	//fmt.Println("********************************")
	//fmt.Println(username, passwd)
	user := dao.GetUser(username, passwd)
	if user == nil {
		log.Println("登录时用户查询失败")
		return nil, errors.New("账号密码不真切")
	}
	uid := user.Uid
	//生成token
	token, err := utils.Award(&uid)
	if err != nil {
		log.Println("生成token失败")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar

	var lr = &models.LoginRes{
		Token:    token,
		UserInfo: userInfo,
	}

	return lr, nil
}
