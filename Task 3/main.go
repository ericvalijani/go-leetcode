package main

import "fmt"

type OrderItem struct {
    ProductID int
    Quantity  int
    Price     float64
}

type Order struct {
    ID        int
    Items     []OrderItem
    Discount  float64 // Discount percentage (0-100)
    Total     float64
    Status    string // "pending", "processing", "completed", "cancelled"
}

type DiscountRule struct {
    MinAmount      float64
    DiscountPercent float64
    Description    string
}

// Calculate order subtotal without discount
func CalculateSubtotal(order Order) float64 {
    // Write your code here
    // Sum of (Quantity * Price) for all items
}

// Apply discount based on rules
func ApplyDiscountRules(order Order, rules []DiscountRule) Order {
    // Write your code here
    // Check if order total meets minimum for each rule
    // Apply the maximum possible discount
    // Update order.Discount and order.Total
}

// Process multiple orders simultaneously
func ProcessOrders(orders []Order, rules []DiscountRule) []Order {
    // Write your code here
    // For each order:
    //   1. Change status to "processing"
    //   2. Apply discount
    //   3. Calculate final total
    //   4. Change status to "completed"
    // If order total is less than 0, set status to "cancelled"
}

// Filter orders by status
func FilterOrdersByStatus(orders []Order, status string) []Order {
    // Write your code here
    // Return only orders with the specified status
}

// Calculate order statistics
func CalculateOrderStatistics(orders []Order) map[string]interface{} {
    // Write your code here
    // Return a map containing:
    //   - "total_orders": total number of orders
    //   - "total_revenue": total revenue
    //   - "average_order_value": average order value
    //   - "completed_orders": number of completed orders
    //   - "total_discount": total discounts applied
}

func main(){
// Define discount rules
rules := []DiscountRule{
    {MinAmount: 100.0, DiscountPercent: 5.0, Description: "5% off for orders over $100"},
    {MinAmount: 500.0, DiscountPercent: 10.0, Description: "10% off for orders over $500"},
    {MinAmount: 1000.0, DiscountPercent: 15.0, Description: "15% off for orders over $1000"},
}

// Create orders
orders := []Order{
    {
        ID: 1,
        Items: []OrderItem{
            {ProductID: 1, Quantity: 2, Price: 50.0},
            {ProductID: 2, Quantity: 1, Price: 30.0},
        },
        Status: "pending",
    },
    {
        ID: 2,
        Items: []OrderItem{
            {ProductID: 3, Quantity: 10, Price: 100.0},
        },
        Status: "pending",
    },
}

// Process orders
processedOrders := ProcessOrders(orders, rules)

// Filter completed orders
completedOrders := FilterOrdersByStatus(processedOrders, "completed")

// Calculate statistics
stats := CalculateOrderStatistics(processedOrders)
}
