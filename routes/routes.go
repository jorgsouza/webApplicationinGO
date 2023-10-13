package routes

import (
	"net/http"

	"github.com/jorgsouza/webApplication/controllers"
)

func LoadsRoutes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
}
