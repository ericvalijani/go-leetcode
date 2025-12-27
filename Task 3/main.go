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
    subtotal := 0.0
    for _, item := range order.Items {
        subtotal += float64(item.Quantity) * item.Price
    }
    return subtotal
}

// Apply discount based on rules
func ApplyDiscountRules(order Order, rules []DiscountRule) Order {
    // Write your code here
    // Check if order total meets minimum for each rule
    // Apply the maximum possible discount
    // Update order.Discount and order.Total
    maxDiscount := func(subtotal float64, rules ...DiscountRule) float64 {
        maxD := 0.0
        for _, rule := range rules {
            if subtotal >= rule.MinAmount {
                if rule.DiscountPercent > maxD {
                    maxD = rule.DiscountPercent
                }
            }
        }
        return maxD
    }

    subtotal := CalculateSubtotal(order)
    discount := maxDiscount(subtotal, rules...)
    order.Discount = discount
    order.Total = subtotal * (1 - discount / 100)
    return order
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
    processed := make([]Order, 0, len(orders)) // Pre-allocate for efficiency
    for _, order := range orders {
        order.Status = "processing"
        order = ApplyDiscountRules(order, rules)
        if order.Total < 0 {
            order.Status = "cancelled"
        } else {
            order.Status = "completed"
        }
        processed = append(processed, order)
    }
    return processed
}

// Filter orders by status
func FilterOrdersByStatus(orders []Order, status string) []Order {
    // Write your code here
    // Return only orders with the specified status
    filtered := make([]Order, 0)
    for _, order := range orders {
        if order.Status == status {
            filtered = append(filtered, order)
        }
    }
    return filtered
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
    stats := make(map[string]interface{})
    totalOrders := len(orders)
    completedCount := 0
    var (
    totalRevenue float64
    totalDiscount float64
    )

    for _, order := range orders {
        subtotal := CalculateSubtotal(order)
        discount := subtotal * (order.Discount / 100)
        totalDiscount += discount
        switch order.Status {
        case "completed":
            completedCount++
            totalRevenue += order.Total
        default:
            // Do nothing for other statuses
        }
    }
    
    stats["total_orders"] = totalOrders
    stats["total_revenue"] = totalRevenue
    if completedCount > 0 {
        stats["average_order_value"] = totalRevenue / float64(completedCount)
    } else {
        stats["average_order_value"] = 0.0
    }
    stats["completed_orders"] = completedCount
    stats["total_discount"] = totalDiscount
    return stats
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
fmt.Println(completedOrders)

// Calculate statistics
stats := CalculateOrderStatistics(processedOrders)
fmt.Println(stats)
}
