package main

import (
	"go-campaign-app/app"
)

func main() {
	server := app.NewRouter()
	server.Run(":2802")
}
