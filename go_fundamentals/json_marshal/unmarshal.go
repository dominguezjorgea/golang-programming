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
	productJSON := `{
		"productID": 123,
		"manufacturer":"Big Box Company",
		"sku":"45687kjhlj",
		"upc":"6546542",
		"pricePerUnit":"12.99",
		"quantityOnHand":28,
		"productName":"Gizmo"
	}`

	product := Product{}

	err := json.Unmarshal([]byte(productJSON), &product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("json unmarshal poduct: %s", product.ProductName)

}
