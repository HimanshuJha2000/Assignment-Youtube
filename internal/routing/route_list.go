package routing

import (
	"github.com/gin-gonic/gin"
	"github.com/razorpay/MachineRound/internal/controller"
)

func SetupRoutes(engine *gin.Engine) *gin.Engine {

	grp := engine.Group("/v1/")
	{
		grp.GET("players", controller.Hello)
		//grp.GET("players/:player_id", controller.Players.FetchPlayerByID)
		//grp.GET("players/paginate/:page_no", controller.Players.FetchPlayerByPage)
		//grp.POST("sleep", controller.Players.SleepService)
	}
	return engine
}
