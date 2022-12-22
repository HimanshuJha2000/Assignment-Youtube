package main

import (
	"flag"
	"github.com/razorpay/MachineRound/internal/bootstrap"
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
	logrus.Println("Starting the api layer ...")
	flag.Parse()

	bootstrap.BaseInitApi(*basePath, *env)
}
