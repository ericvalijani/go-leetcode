package main

import "fmt"

func sumNumbers(numbers []int) int {
	var total int = 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func main() {
	example := []int{1, 2, 3, 4, 5}
	sum := sumNumbers(example)
	fmt.Println("Sum:", sum)
}
