package models

import (
	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	ID           int `gorm:"auto_increment"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Product_Name string `sql:"unique_index:id_product"`
	Description  string
	Category     string
	Quantity     int
	Price        int
	Image        string `gorm:"type:varchar(255);"`
	Variant      []Variant
	Rating       []Rating
}

type Rating struct {
	ID        int `gorm:"auto_increment"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ProductID uint   `sql:"unique_index:id_name"`
	Name      string `sql:"unique_index:id_name"`
	Review    string
	Rating    int `gorm:"check:rating>1&rating<5"`
}

type Variant struct {
	ID        int `gorm:"auto_increment"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ProductID uint   `json:"productid" sql:"unique_index:product_and_color"`
	Color     string `json:"color" sql:"unique_index:product_and_color"`
	Image     string `json:"image"`
}
