package routing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/razorpay/MachineRound/internal/config"
	"github.com/razorpay/MachineRound/internal/constants"
	"github.com/tylerb/graceful"
	"time"
)

// Initialize all the web routes of reminders service
func InitializeRoutes() {
	router := gin.Default()
	SetupRoutes(router)

	LaunchServer(router, constants.API)
}

func LaunchServer(router *gin.Engine, service string) {
	err := graceful.RunWithErr(GetListenAddress(service), constants.GracefulTimeoutDuration*time.Second, router)
	if err != nil {
		fmt.Errorf(err.Error())
		panic("Error: Error occurred while starting the server")
	}
}

// GetListenAddress will give the address in string to listen to
func GetListenAddress(service string) string {
	if service == constants.API {
		application := config.GetConfig().Application
		return fmt.Sprintf("%s:%d", application.ListenIP, application.ListenPort)
	}
	worker := config.GetConfig().Worker
	return fmt.Sprintf("%s:%d", worker.ListenIP, worker.ListenPort)
}
