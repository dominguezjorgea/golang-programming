package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

func main() {
	product := &Product{
		ProductID:      123,
		Manufacturer:   "Big Box Company",
		Sku:            "45687kjhlj",
		Upc:            "6546542",
		PricePerUnit:   "12.99",
		QuantityOnHand: 28,
		ProductName:    "Gizmo",
	}

	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(productJSON))

}
