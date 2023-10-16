package main

import (
	"net/http"

	"github.com/jorgsouza/webApplication/internal/watchdog"
	"github.com/jorgsouza/webApplication/routes"
)

func main() {

	watchdog.SetupSignalHandling()

	routes.LoadsRoutes()
	http.ListenAndServe(":8080", nil)
}
