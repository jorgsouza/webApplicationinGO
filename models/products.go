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
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(name string, description string, price float64, quantity int) {
	db := db.DatabaseConnect()

	sqlInsertInto, err := db.Prepare("INSERT INTO products (name,description,price,quantity) values($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	sqlInsertInto.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.DatabaseConnect()

	deleteProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}
