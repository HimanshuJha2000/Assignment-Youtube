package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/razorpay/MachineRound/internal/controller"
)

func SetupRoutes(engine *gin.Engine) *gin.Engine {

	grp := engine.Group("/v1/youtube")
	{
		grp.POST("/add/api_key", controller.ApiPod.CreateAPIKey)
		grp.POST("/search", controller.YoutubePod.SearchVideo)
		grp.GET("/videos/paginate/:page_no", controller.YoutubePod.FetchVideoDetails)
	}
	return engine
}
