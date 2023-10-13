package main

import (
	"net/http"

	"github.com/jorgsouza/webApplication/routes"
	"github.com/jorgsouza/webApplication/watchdog"
)

func main() {

	watchdog.SetupSignalHandling()

	routes.LoadsRoutes()
	http.ListenAndServe(":8080", nil)
}
