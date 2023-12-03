package models

import (
	"github.com/pequenojoohn/db"
	_ "github.com/pequenojoohn/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func GetAllProducts() []Product {
	db := db.ConnectionDB()

	selectAllProducts, err := db.Query("select * from products order by id asc")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64
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

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.ConnectionDB()

	insertData, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(name, description, price, quantity)
	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.ConnectionDB()

	deleteProduct, err := db.Prepare("delete from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deleteProduct.Exec(id)
	defer db.Close()
}

func EditProduct(id string) Product {
	db := db.ConnectionDB()

	productDb, err := db.Query("select * from products where  id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productEdit := Product{}

	for productDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productDb.Scan(&id, &name, &description, &quantity, &price)

		if err != nil {
			panic(err.Error())
		}

		productEdit.Id = id
		productEdit.Name = name
		productEdit.Description = description
		productEdit.Quantity = quantity
		productEdit.Price = price

	}

	defer db.Close()
	return productEdit
}

func UpdateProduct(id int, quantity int, name string, description string, price float64) {
	db := db.ConnectionDB()

	updateProduct, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(name, description, price, quantity, id)

	defer db.Close()
}
