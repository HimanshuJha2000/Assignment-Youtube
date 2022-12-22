package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/razorpay/MachineRound/internal/controller"
)

func SetupRoutes(engine *gin.Engine) *gin.Engine {

	grp := engine.Group("/v1/youtube")
	{
		grp.POST("/add/api_key", controller.ApiPod.CreateAPIKey)
		//grp.GET("/search/:title/:description", controller.Hello)
		//grp.GET("/videos", controller.Hello)
	}
	return engine
}
