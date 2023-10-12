package main

import (
	"net/http"
	"text/template"
)

type Product struct {
	Name, Description string
	Price             float64
	Quantity          int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	products := []Product{
		{
			Name:        "T-Shirt",
			Description: "A coll t-shirt",
			Price:       29.99,
			Quantity:    9,
		},
		{"Hood", "A colorful hood", 199.99, 2},
		{"Fone", "Hyper tech Fone", 1999.89, 3},
		{"New secret product", "a new secret for a great person", 91999.89, 6},
	}

	temp.ExecuteTemplate(w, "Index", products)
}
