package main

import (
	"flag"
	"github.com/razorpay/MachineRound/internal/bootstrap"
	"github.com/razorpay/MachineRound/internal/youtube_video"
	"github.com/sirupsen/logrus"
)

var (
	env      *string
	basePath *string
)

// fetch all the cli inputs options provided.
func init() {
	basePath = flag.String("base_path", ".", "Reminders base path")
	env = flag.String("env", "dev", "Application env : prod/dev")
}

func main() {
	flag.Parse()
	logrus.Println("Starting the Youtube fetch worker layer ...")

	bootstrap.BaseInitWorker(*basePath, *env)
	youtube_video.BeginWorkerCron()
}
