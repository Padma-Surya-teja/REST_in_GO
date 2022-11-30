package main

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce.com/m/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// models.Create_Tables()

	//Init the mux router
	router := mux.NewRouter()

	//Get all Products
	router.HandleFunc("/api/products", handlers.GetProducts).Methods("GET")

	//Add a Product
	router.HandleFunc("/api/products/create", handlers.AddProduct).Methods("POST")

	//Get a particular Product
	router.HandleFunc("/api/products/{id}", handlers.GetProduct).Methods("GET")

	//Add review to a particular product
	router.HandleFunc("/api/products/{id}/reviews/create", handlers.AddReview).Methods("PATCH")

	//Get a particular product review
	router.HandleFunc("/api/products/{id}/reviews", handlers.GetProductReviews).Methods("GET")

	//Update details of a review
	router.HandleFunc("/api/products/{id}/reviews/{rid}", handlers.UpdateReview).Methods("PATCH")

	//Delete a particular review
	router.HandleFunc("/api/products/{id}/reviews/{rid}", handlers.DeleteReview).Methods("DELETE")

	fmt.Println("Server at 8080")

	//Serving
	log.Fatal(http.ListenAndServe(":8080", router))
}
