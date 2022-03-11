package views

import (
	"blog_gin/service"
	"github.com/gin-gonic/gin"
)

func (*HTMLApi) Pigeonhole(c *gin.Context) {
	pigeonholeRes := service.FindPostPigronhole()
	c.HTML(200, "pigeonhole.html", pigeonholeRes)
}
