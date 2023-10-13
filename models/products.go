package models

import "github.com/jorgsouza/webApplication/db"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func SearchAllProducts() []Product {
	db := db.DatabaseConnect()

	selectAllProducts, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id int
		var name, description string
		var price float64
		var quantity int

		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}
