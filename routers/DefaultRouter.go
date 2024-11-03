package routers

import (
	controllers "ApiController/controller"
	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/", controllers.DefaultController{}.Index)
		v1.GET("/status", controllers.DefaultController{}.Status)
		v1.GET("/translate", controllers.TranslateController{}.Index)
		v1.POST("/translate", controllers.TranslateController{}.Translation)
	}
}
