package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"ecommerce.com/m/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type DataBase struct {
	Db *gorm.DB
}

// Function for checking errors
func (s DataBase) checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Function for handling messages
func (s DataBase) printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// returns all the products
func (s DataBase) GetProducts(w http.ResponseWriter, r *http.Request) {

	s.printMessage("Getting Products...")

	var products []models.Product
	s.Db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Find(&products)

	s.printMessage("Received the Products...")
	var response = JsonResponse{Type: "success", Data: products}

	json.NewEncoder(w).Encode(response)
}

// Add products
func (s DataBase) AddProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t models.Product
	err := decoder.Decode(&t)

	s.checkErr(err)

	var response = JsonResponse{}

	fmt.Println(t)
	if t.ID == 0 {
		response = JsonResponse{Type: "error", Message: "You are missing some parameters."}
	} else {

		s.printMessage("Inserting Product into Db")

		s.Db.Create(&t)

		response = JsonResponse{Type: "success", Message: "The product has been inserted successfully!"}

	}

	json.NewEncoder(w).Encode(response)
}
func (s DataBase) GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	s.printMessage("Getting Products...")

	var product []models.Product
	s.Db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Where("id=?", id).Find(&product)

	json.NewEncoder(w).Encode(product)
}

func (s DataBase) AddReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	decoder := json.NewDecoder(r.Body)
	var t models.Rating
	err := decoder.Decode(&t)

	s.checkErr(err)

	var response = JsonResponse{}

	fmt.Println(t)

	s.printMessage("Inserting Review into Db")

	var product models.Product
	s.Db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Where("id=?", id).Find(&product)

	product.Rating = append(product.Rating, t)
	s.Db.Save(&product)

	response = JsonResponse{Type: "success", Message: "The product review has been inserted successfully!"}
	s.printMessage("Inserted Review into Db")

	json.NewEncoder(w).Encode(response)
}

func (s DataBase) GetProductReviews(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	s.printMessage("Getting Product Reviews...")

	var product models.Product
	s.Db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Where("id=?", id).Find(&product)

	json.NewEncoder(w).Encode(product.Rating)
}

func (s DataBase) UpdateReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	s.checkErr(err)

	rid, err := strconv.Atoi(params["rid"])
	s.checkErr(err)

	decoder := json.NewDecoder(r.Body)
	var t models.Rating
	err = decoder.Decode(&t)

	s.checkErr(err)

	var response = JsonResponse{}

	if t.ID == 0 {
		response = JsonResponse{Type: "error", Message: "You are missing some parameters."}
	} else {

		s.printMessage("Inserting Review into Db")

		var rating models.Rating

		s.Db.Model(&models.Rating{}).Where("product_id=? and id=?", id, rid).Find(&rating)

		s.Db.First(&rating)
		rating.Rating = t.Rating
		rating.Review = t.Review
		s.Db.Save(&rating)
		// Db.Save(&t);

		response = JsonResponse{Type: "success", Message: "The product review has been inserted successfully!"}

	}

	json.NewEncoder(w).Encode(response)
}

func (s DataBase) DeleteReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	pid, err := strconv.Atoi(params["id"])
	s.checkErr(err)
	rid, err := strconv.Atoi(params["rid"])

	var response = JsonResponse{}
	if pid == 0 || rid == 0 {
		response = JsonResponse{Type: "error", Message: "You are missing some parameters."}
	} else {

		s.printMessage("Deleting a Review in Db")

		s.Db.Model(models.Rating{}).Where("product_id=? and id=?", pid, rid).Delete(&models.Rating{})

		response = JsonResponse{Type: "success", Message: "The product review has been inserted successfully!"}

	}

	json.NewEncoder(w).Encode(response)

}

type JsonResponse struct {
	Type    string           `json:"type"`
	Data    []models.Product `json:"data"`
	Message string           `json:"message"`
}
