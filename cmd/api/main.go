package main

import (
	"flag"
	"fmt"
	"github.com/razorpay/MachineRound/internal/bootstrap"
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
	fmt.Println("Starting the api layer ...")
	flag.Parse()

	bootstrap.BaseInitApi(*basePath, *env)
}
