package main

import (
	"context"
	"github.com/joho/godotenv"
	"pwd/pkg/resources"
)

const envFile = ".env-test"

func main() {
	// The code implements deployment using serverless function and webhooks (pls see api/index.go)
	// In the main function I left the method using the long polling (all that is needed is to add a loop and change the Update type to UpdateChannels in the SendMsg function (sender.go))

	if err := godotenv.Load(envFile); err != nil {
		panic("can't load env")
	}

	_, telegramService := resources.InitBot()

	telegramService.SendMsgWithoutWebhook(context.Background(), telegramService.ReceiveRequest())
}
