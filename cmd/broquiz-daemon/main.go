package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Cyfell/BroQuiz/internal/api"
	"github.com/Cyfell/BroQuiz/internal/config"
)

func run() int {
	config, err := config.ReadConfig("broquiz")
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to get configuration,", err)
	}

	ctx := context.Background()
	api, err := api.New(ctx, config.API)
	if err != nil {
		fmt.Fprintln(os.Stderr, "unable to init API,", err)
	}

	if err := api.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error when running history,", err)
		return 255
	}

	return 0
}

func main() {
	os.Exit(run())
}
