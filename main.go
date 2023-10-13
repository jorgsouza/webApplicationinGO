package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"text/template"

	"github.com/jorgsouza/webApplication/infra"
	"github.com/jorgsouza/webApplication/models"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go func() {
		<-interrupt
		fmt.Println("\n Closing the application...")

		if err := infra.StopDatabaseContainer(); err != nil {
			fmt.Println("Error when stopping the database container:", err)
		}

		os.Exit(0)
	}()

	if err := infra.StartDatabaseContainer(); err != nil {
		fmt.Println("Error when starting the database container:", err)
		os.Exit(1)
	}

	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}
