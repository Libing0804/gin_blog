package api

import (
	"blog_gin/models"
	"blog_gin/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func (*Api) Login(c *gin.Context) {
	//接收用户名和密码
	//返回数据

	//必须使用shouldBind获取参数  postform 不行
	userMima := models.Mima{}
	err := c.ShouldBind(&userMima)
	if err != nil {
		log.Println("从页面获取的用户参数失败")
		return
	}
	username := userMima.UserName
	passwd := userMima.Passwd

	loginRes, err := service.Login(username, passwd)
	//fmt.Println(loginRes.Usf.Uid, loginRes.Usf.UserName)
	if err != nil {
		fmt.Println("我到了这")
		var res models.Result
		res.Code = -999
		res.Error = err.Error()
		//resultJson, err := json.Marshal(res)
		//c.Set("Content-Type", "application/json")
		//if err != nil {
		//	log.Println("json 序列化失败")
		//	return
		//}
		c.JSON(200, gin.H{
			"code":  res.Code,
			"data":  "",
			"error": res.Error,
		})
		return
	}
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = loginRes
	//_, err = json.Marshal(result)
	//if err != nil {
	//	log.Println("json 序列化失败")
	//	return
	//}
	c.JSON(200, gin.H{
		"code":     result.Code,
		"data":     result.Data,
		"error":    result.Error,
		"userName": loginRes.UserInfo.UserName,
	})
}
