package routers

import "github.com/gin-gonic/gin"

func InitRouter(router *gin.Engine) {
	InitWeddingsRouter(router)
}
