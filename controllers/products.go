package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/jorgsouza/webApplication/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvertedToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error in price conversion:", err)
		}

		quantityConvertedToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error in quantity conversion:", err)
		}

		models.CreateNewProduct(name, description, priceConvertedToFloat, quantityConvertedToInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}
