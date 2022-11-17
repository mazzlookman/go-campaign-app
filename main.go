package main

import "go-campaign-app/dependency_injection"

func main() {
	server := dependency_injection.InitializedServerTest()
	server.Run(":2802")
}
