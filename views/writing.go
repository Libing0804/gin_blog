package views

import (
	"blog_gin/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (HTMLApi) Writing(c *gin.Context) {

	wr := service.Writing()
	c.HTML(http.StatusOK, "writing.html", wr)

}
