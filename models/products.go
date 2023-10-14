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

	selectAllProducts, err := db.Query("SELECT * FROM products ORDER BY id ASC")
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

func EditProduct(id string) Product {
	db := db.DatabaseConnect()

	productDbReturned, err := db.Query("SELECT * FROM products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	productUpdate := Product{}

	for productDbReturned.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productDbReturned.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		productUpdate.Id = id
		productUpdate.Name = name
		productUpdate.Description = description
		productUpdate.Price = price
		productUpdate.Quantity = quantity
	}
	defer db.Close()
	return productUpdate
}

func UpdateProduct(id int, name string, description string, price float64, quantity int) {
	db := db.DatabaseConnect()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
