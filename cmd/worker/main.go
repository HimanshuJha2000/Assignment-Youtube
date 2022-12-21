package main

import (
	"flag"
	"github.com/razorpay/MachineRound/internal/bootstrap"
	"github.com/razorpay/MachineRound/internal/youtube_worker"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
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

	youtube_worker.BeginWorkerCron()

	//ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		log.Println("Program killed !")

		// do last actions and wait for all write operations to end

		os.Exit(0)
	}()

	//c := make(chan os.Signal, 1)

	// accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// or SIGTERM. SIGKILL, SIGQUIT will not be caught.
	//signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	//
	//go func(baseCtx context.Context, baseCancel context.CancelFunc) {
	//	// Block until signal is received.
	//	<-c
	//	defer baseCancel()
	//	logrus.Error("WORKER_SHUTDOWN : Shutting down worker")
	//}(ctx, cancel)

}
