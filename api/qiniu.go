package api

import (
	"blog_gin/config"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

func (*Api) QiniuToken(c *gin.Context) {
	//	自定义凭证有效时间
	bucket := "lbbase"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //两小时
	mac := qbox.NewMac(config.Cfg.System.QiniuAccessKey, config.Cfg.System.QiniuSecretKey)

	upToken := putPolicy.UploadToken(mac)
	c.JSON(200, gin.H{
		"code":  200,
		"error": "",
		"data":  upToken,
	})
}
