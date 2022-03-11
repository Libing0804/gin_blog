package main

import (
	"blog_gin/router"
)

func main() {
	ru := router.Router()
	ru.Run("127.0.0.1:8000")
}
