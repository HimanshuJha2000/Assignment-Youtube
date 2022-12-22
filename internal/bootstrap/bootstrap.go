package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/razorpay/MachineRound/internal/config"
	"github.com/razorpay/MachineRound/internal/constants"
	"github.com/razorpay/MachineRound/internal/providers/database"
	"github.com/razorpay/MachineRound/internal/routing"
	"github.com/razorpay/MachineRound/internal/video"
)

var (
	YoutubeClient video.YoutubeService
)

// InitializeRouter will initialize the web server for the application
func initializeRouter() {
	routing.InitializeRoutes()
}

// BaseInitApi Function will be used to load config for api layer
func BaseInitApi(basePath string, env string) {
	config.LoadConfig(basePath, env)

	initProviders()

	initializeRouter()
}

// BaseInitWorker Function will be used to load config for worker layer
func BaseInitWorker(basePath string, env string) {
	config.LoadConfig(basePath, env)

	initProviders()

	router := gin.Default()

	YoutubeClient.Initialize()

	go routing.LaunchServer(router, constants.WORKER)
}

// initProviders : Provider initialization is done here
// There initiated providers will be available across the application
func initProviders() {
	database.Initialize()
}

//func WaitForShutDown(ctx context.Context) {
//	select {
//	case <-ctx.Done():
//		fmt.Fprint(os.Stderr, "request cancelled\n")
//		return
//	}
//}
