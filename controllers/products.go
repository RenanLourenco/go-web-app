package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"go-web-app.com/models"
)


var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.ListProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w,"New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceString := r.FormValue("price")
		quantityString := r.FormValue("quantity")

		price,err := strconv.ParseFloat(priceString,64)
		
		if err != nil {
			log.Println("Error converting price to float64")
		}

		quantity, err := strconv.Atoi(quantityString)

		if err != nil {
			log.Println("Error converting quantity to integer")
		}

		p := models.Product{
			Name: name,
			Description: description,
			Price: price,
			Quantity: quantity,
		}

		models.CreateProduct(p)
	}

	http.Redirect(w,r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request){
	productId := r.URL.Query().Get("id")
	models.DeleteProduct(productId)

	http.Redirect(w,r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request){
	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)

	temp.ExecuteTemplate(w,"Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		idString := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceString := r.FormValue("price")
		quantityString := r.FormValue("quantity")

		id, err := strconv.Atoi(idString)


		price,err := strconv.ParseFloat(priceString,64)
		
		if err != nil {
			log.Println("Error converting price to float64")
		}

		quantity, err := strconv.Atoi(quantityString)

		if err != nil {
			log.Println("Error converting quantity to integer")
		}

		product := models.Product{
			Id: id,
			Name: name,
			Description: description,
			Price: price,
			Quantity: quantity,
		}

		models.UpdateProduct(product)

	}
	http.Redirect(w,r,"/",301)
}