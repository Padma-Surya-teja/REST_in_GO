package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"ecommerce.com/m/handlers"
	"ecommerce.com/m/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {
	//connect tyo database
	db_user := flag.String("user", "postgres", "database user")
	db_password := flag.String("password", "root", "database password")
	db_name := flag.String("name", "Ecommerce_microservice", "database name")
	database_link := "user=" + *db_user + " password=" + *db_password + " dbname=" + *db_name + " sslmode=disable"
	db, err := gorm.Open("postgres", database_link)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Rating{})
	db.AutoMigrate(&models.Variant{})

	var database handlers.DataBase = handlers.DataBase{Db: db}

	//Init the mux router
	router := mux.NewRouter()

	//Get all Products
	router.HandleFunc("/api/products", database.GetProducts).Methods("GET")

	//Add a Product
	router.HandleFunc("/api/products/create", database.AddProduct).Methods("POST")

	//Get a particular Product
	router.HandleFunc("/api/products/{id}", database.GetProduct).Methods("GET")

	//Add review to a particular product
	router.HandleFunc("/api/products/{id}/reviews/create", database.AddReview).Methods("POST")

	//Get a particular product review
	router.HandleFunc("/api/products/{id}/reviews", database.GetProductReviews).Methods("GET")

	//Update details of a review
	router.HandleFunc("/api/products/{id}/reviews/{rid}", database.UpdateReview).Methods("PATCH")

	//Delete a particular review
	router.HandleFunc("/api/products/{id}/reviews/{rid}", database.DeleteReview).Methods("DELETE")

	fmt.Println("Server at 8080")

	//Serving
	log.Fatal(http.ListenAndServe(":8080", router))
}
