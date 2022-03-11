package dao

import (
	"blog_gin/models"
	"errors"
	"log"
)

func GetUserNameById(uid int) (string, error) {
	row := DB.QueryRow("select user_name from blog_user where uid = ?", uid)
	var userName string
	if row.Err() != nil {
		return userName, errors.New("查询用户名失败")
	}
	err := row.Scan(&userName)
	if err != nil {
		return userName, err
	}
	return userName, nil
}

func GetUser(userName, passwd string) *models.User {
	//fmt.Println(userName, passwd)
	row := DB.QueryRow("select * from blog_user where user_name=? and passwd=? limit 1",
		userName,
		passwd,
	)
	if row.Err() != nil {
		//
		log.Println("在getuser中获取失败")
		log.Println(row.Err())
		return nil
	}

	var user = &models.User{}
	err := row.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	//fmt.Println("**************************************")
	//fmt.Println("打印uid", user.Uid)
	if err != nil {

		log.Println(err)
		return nil
	}
	return user
}
