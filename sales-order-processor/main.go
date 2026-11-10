package main

import (
	"fmt"
	"strings"
)

type itemPrices map[string]float64

var products = itemPrices{
	"T-Shirt": 20.00,
	"Mug":     12.50,
	"Hat":     18,
	"Book":    25.99,
}

func calculateItemPrice(itemCode string) (float64, bool) {
	basePrice, found := products[itemCode]
	if !found {
		suffix := "-SALE"
		if strings.HasSuffix(itemCode, suffix) {
			originalItemCode := strings.TrimSuffix(itemCode, suffix)
			basePrice, found = products[originalItemCode]
			if found {
				salePrice := basePrice * 0.90
				fmt.Printf(" - Item %s (Sale!) Price: $%.2f, Sale: $%.2f\n",
					originalItemCode, basePrice, salePrice)
				return salePrice, true
			}
		}

		fmt.Printf(" - Item %s (Product not found!)\n", itemCode)

		return 0.0, false
	}
	fmt.Printf(" - Item %s Price: $%.2f\n", itemCode, basePrice)

	return basePrice, true
}

func main() {
	fmt.Println("--------------------Simple Sales Order Processor------------------------")

	orderItems := []string{
		"T-Shirt", "Mug-SALE", "Hat", "Book",
	}

	var subTotal float64
	fmt.Println("Processing Order Items:")

	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subTotal += price
		}
	}

	fmt.Printf("==== Subtotal Price: $%.2f\n", subTotal)

}
