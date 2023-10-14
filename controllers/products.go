package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/jorgsouza/webApplication/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, _ *http.Request) {
	allProducts := models.SearchAllProducts()
	temp.ExecuteTemplate(w, "Index", allProducts)
}

func New(w http.ResponseWriter, _ *http.Request) {
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
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvertedToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Error converting (Id) to int:", err)
		}

		priceConvertedToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Error converting (Price) to float:", err)
		}

		quantityConvertedToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("Error converting (Quantity) to int:", err)
		}

		models.UpdateProduct(idConvertedToInt, name, description, priceConvertedToFloat, quantityConvertedToInt)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
