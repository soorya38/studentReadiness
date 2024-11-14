package main

import "manageProfiles/handlers"

func main() {
	handlers.RegisterHandlers()

	handlers.StartServer(8080)
}