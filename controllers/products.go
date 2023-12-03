package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/pequenojoohn/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	// products := []Product{
	// 	{Name: "T-shirt", Description: "Blue shirt beautiful", Price: 30, Quantity: 4},
	// 	{Name: "Shoes", Description: "Red shoes beautiful", Price: 21, Quantity: 2},
	// }
	allProducts := models.GetAllProducts()
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

		convertFloatPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço", err)
		}

		convertIntPrice, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		models.CreateNewProduct(name, description, convertFloatPrice, convertIntPrice)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	models.DeleteProduct(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {

	productId := r.URL.Query().Get("id")
	product := models.EditProduct(productId)

	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvert, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Erro na conversão do ID para int: ", err)
		}

		priceConvert, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Println("Erro na conversão do preço para float: ", err)
		}

		quantityConvert, err := strconv.Atoi(quantity)

		if err != nil {
			log.Println("Erro na conversão do quantity para int: ", err)
		}

		models.UpdateProduct(idConvert, quantityConvert, name, description, priceConvert)

		http.Redirect(w, r, "/", 301)

	}
}
