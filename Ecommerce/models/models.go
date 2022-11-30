package models

import (
	"time"

	_ "github.com/lib/pq"
)

type Product struct {
	ID           int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Product_Name string
	Description  string
	Category     string
	Quantity     int
	Price        int
	Image        string `gorm:"type:varchar(255);"`
	Variant      []Variant
	Rating       []Rating
}

type Rating struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	ProductID uint
	Name      string
	Review    string
	Rating    int `gorm:"check:rating>1&rating<5"`
}

type Variant struct {
	ID        int `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ProductID uint   `json:"productid"`
	Color     string `json:"color"`
	Image     string `json:"image"`
}
