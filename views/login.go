package views

import (
	"blog_gin/config"
	"github.com/gin-gonic/gin"
)

func (*HTMLApi) Login(c *gin.Context) {
	c.HTML(200, "login.html", config.Cfg.Viewer)
}
