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

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "root"
	DB_NAME     = "Ecommerce_microservice"
)

// DB set up
func SetupDB() *gorm.DB {
	database_link := "user=" + DB_USER + " password=" + DB_PASSWORD + " dbname=" + DB_NAME + " sslmode=disable"
	db, err := gorm.Open("postgres", database_link)

	checkErr(err)
	// db.Debug().Model(&models.Rating{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	// db.Debug().Model(&models.Variant{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Rating{})
	db.AutoMigrate(&models.Variant{})

	return db
}

// Function for checking errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// returns all the products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	db := SetupDB()

	printMessage("Getting Products...")

	var products []models.Product
	db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Find(&products)

	printMessage("Received the Products...")
	var response = JsonResponse{Type: "success", Data: products}

	json.NewEncoder(w).Encode(response)
}

// Add products
func AddProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t models.Product
	err := decoder.Decode(&t)

	checkErr(err)

	var response = JsonResponse{}

	fmt.Println(t)
	if t.ID == 0 {
		response = JsonResponse{Type: "error", Message: "You are missing some parameters."}
	} else {
		db := SetupDB()

		printMessage("Inserting Product into DB")

		db.Create(&t)

		response = JsonResponse{Type: "success", Message: "The product has been inserted successfully!"}

	}

	json.NewEncoder(w).Encode(response)
}
func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	db := SetupDB()

	printMessage("Getting Products...")

	var product []models.Product
	db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Where("id=?", id).Find(&product)

	json.NewEncoder(w).Encode(product)
}

func AddReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]

	decoder := json.NewDecoder(r.Body)
	var t models.Rating
	err := decoder.Decode(&t)

	checkErr(err)

	var response = JsonResponse{}

	fmt.Println(t)
	if t.ID == 0 {
		response = JsonResponse{Type: "error", Message: "You are missing some parameters."}
	} else {
		db := SetupDB()

		printMessage("Inserting Review into DB")

		var product models.Product
		db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Where("id=?", id).Find(&product)

		product.Rating = append(product.Rating, t)
		db.Save(&product)

		response = JsonResponse{Type: "success", Message: "The product review has been inserted successfully!"}

	}

	json.NewEncoder(w).Encode(response)
}

func GetProductReviews(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	db := SetupDB()

	printMessage("Getting Product Reviews...")

	var product models.Product
	db.Model(&models.Product{}).Preload("Rating").Preload("Variant").Where("id=?", id).Find(&product)

	json.NewEncoder(w).Encode(product.Rating)
}

func UpdateReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	checkErr(err)

	rid, err := strconv.Atoi(params["rid"])
	checkErr(err)

	decoder := json.NewDecoder(r.Body)
	var t models.Rating
	err = decoder.Decode(&t)

	checkErr(err)

	var response = JsonResponse{}

	if t.ID == 0 {
		response = JsonResponse{Type: "error", Message: "You are missing some parameters."}
	} else {
		db := SetupDB()

		printMessage("Inserting Review into DB")

		var rating models.Rating

		db.Model(&models.Rating{}).Where("product_id=? and id=?", id, rid).Find(&rating)

		db.First(&rating)
		rating.Rating = t.Rating
		rating.Review = t.Review
		db.Save(&rating)
		// db.Save(&t);

		response = JsonResponse{Type: "success", Message: "The product review has been inserted successfully!"}

	}

	json.NewEncoder(w).Encode(response)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	pid, err := strconv.Atoi(params["id"])
	checkErr(err)
	rid, err := strconv.Atoi(params["rid"])

	var response = JsonResponse{}
	if pid == 0 || rid == 0 {
		response = JsonResponse{Type: "error", Message: "You are missing some parameters."}
	} else {
		db := SetupDB()

		printMessage("Deleting a Review in DB")

		db.Model(models.Rating{}).Where("product_id=? and id=?", pid, rid).Delete(&models.Rating{})

		response = JsonResponse{Type: "success", Message: "The product review has been inserted successfully!"}

	}

	json.NewEncoder(w).Encode(response)

}

type JsonResponse struct {
	Type    string           `json:"type"`
	Data    []models.Product `json:"data"`
	Message string           `json:"message"`
}
