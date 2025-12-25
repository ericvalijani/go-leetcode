package main

import (
	"errors"
	"fmt"
)

type Product struct {
    ID       int
    Name     string
    Quantity int
    Price    float64
}

// Add product to inventory
func AddProduct(inventory map[int]Product, product Product) error {
    // Write your code here
    // If product already exists, increase the quantity
    // If product is new, add it

    // First, check if the product ID is already in the inventory map.
    // This line does a safe lookup: it gets the existing product (if any) and a boolean 'ok' that tells us if it was found.
    // 'existing' will be a copy of the Product struct if found, or a zero-value Product if not.
    existing, ok := inventory[product.ID]

// Now, use a simple if-else to handle the two cases.
if ok {
    // Case 1: The product already exists in the inventory.
    // We increase its quantity by adding the new product's quantity to it.
    // (This modifies the 'existing' copy.)
    existing.Quantity += product.Quantity

    // Then, we put the updated 'existing' back into the map.
    // Why? Because maps in Go store values by copy, so we need to save the changes.
    // (Note: This step is technically needed here because we're working with a copy;
    // if we didn't update the map, the changes to 'existing' wouldn't stick.)
    inventory[product.ID] = existing
} else {
    // Case 2: The product is brand new (not found in the map).
    // Just add it directly to the inventory using its ID as the key.
    inventory[product.ID] = product
}

// No errors occurred, so return nil to indicate success.
return nil
}

// Remove product from inventory
func RemoveProduct(inventory map[int]Product, productID int) error {
    // Write your code here
    // If product doesn't exist, return an error
    if _, ok := inventory[productID]; !ok {
	// If product doesn't exist, return an error
	return errors.New("product not found")
    }
    delete(inventory, productID)
    return nil
}

// Check product stock
func CheckStock(inventory map[int]Product, productID int) (int, bool) {
    // Write your code here
    // First value: product quantity
    // Second value: whether product exists or not
    if product, ok := inventory[productID]; ok {
	// First value: product quantity
	// Second value: whether product exists or not
	return product.Quantity, true
    }
    return 0, false
}

// Calculate total inventory value
func CalculateTotalValue(inventory map[int]Product) float64 {
    // Write your code here
    // Sum of (Quantity * Price) for all products
    var total float64 = 0.0
    for _, product := range inventory {
	    total += float64(product.Quantity) * product.Price
    }
    return total
}

func main() {
	inventory := make(map[int]Product)
	
	// Add products
	product1 := Product{ID: 1, Name: "Laptop", Quantity: 10, Price: 1500.0}
	AddProduct(inventory, product1)
	
	product2 := Product{ID: 2, Name: "Mouse", Quantity: 50, Price: 25.0}
	AddProduct(inventory, product2)

	// Check stock
	quantity, exists := CheckStock(inventory, 1)
	// quantity = 10, exists = true
	fmt.Printf("Quantity: %d, Exists: %t\n", quantity, exists)

	// Calculate total value
	totalValue := CalculateTotalValue(inventory)
	// totalValue = (10 * 1500.0) + (50 * 25.0) = 16250.0
	fmt.Printf("Total Value: %.2f\n", totalValue)

	// Remove Product
	RemoveProduct(inventory, 2)
	totalValue = CalculateTotalValue(inventory)
	fmt.Printf("Total Value: %.2f\n", totalValue)
}
