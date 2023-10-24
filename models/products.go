package models

import (
	"go-web-app.com/infra/db"
)


type Product struct {
	Id int
	Name string
	Description string
	Price float64
	Quantity int
}


func ListProducts() []Product {
	db := db.ConnectDatabase()


	selectAllProducts, err := db.Query("SELECT * FROM products ORDER BY id ASC")

	if err != nil {
		panic(err.Error())
	}

	products := []Product{}
	p := Product{}
	
	for selectAllProducts.Next() {
		var id, quantity int 
		var name, description string
		var price float64

		err := selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		
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

func CreateProduct(p Product){
	db := db.ConnectDatabase()

	createdProduct, err := db.Prepare("insert into products(name,description,price,quantity) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	createdProduct.Exec(p.Name,p.Description,p.Price,p.Quantity)
	defer db.Close()

}

func DeleteProduct(id string) {
	db := db.ConnectDatabase()

	delete, err := db.Prepare("DELETE from products where id=$1")

	if err != nil {
		panic(err.Error())
	}

	delete.Exec(id)
	defer db.Close()

}

func EditProduct(id string) Product{
	db := db.ConnectDatabase()

	queriedProduct, err := db.Query("select * from products where id=$1", id)

	if err != nil {
		panic(err.Error())
	}

	productToUpdate := Product{}
	
	for queriedProduct.Next(){
		var id, quantity int
		var name, description string
		var price float64

		err = queriedProduct.Scan(&id, &name, &description, &quantity, &price)
		if err != nil {
			panic(err.Error())
		}

		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Quantity = quantity

	}
	

	defer db.Close()
	return productToUpdate

}

func UpdateProduct(p Product) {
	db := db.ConnectDatabase()

	updateProduct, err := db.Prepare("UPDATE products SET name=$1, description=$2, quantity=$3, price=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}

	updateProduct.Exec(p.Name,p.Description,p.Quantity,p.Price,p.Id)
	defer db.Close()
}