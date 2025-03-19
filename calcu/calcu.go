package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Calculator function that takes an operator and a variable number of operands
func calculator(operator string, nums ...float64) (float64, error) {
	if len(nums) == 0 {
		return 0, errors.New("at least one number is required")
	}

	// Initialize the result with the first number
	result := nums[0]

	// Switch based on the operator
	switch operator {
	case "+":
		for _, num := range nums[1:] {
			result += num
		}
	case "-":
		for _, num := range nums[1:] {
			result -= num
		}
	case "*":
		for _, num := range nums[1:] {
			result *= num
		}
	case "/":
		for _, num := range nums[1:] {
			if num == 0 {
				return 0, errors.New("cannot divide by zero")
			}
			result /= num
		}
	default:
		return 0, errors.New("invalid operator")
	}

	return result, nil
}

func main() {
	var input string

	// Read the operation and numbers from the user
	fmt.Println("Enter the operation (e.g., +, -, *, /) followed by numbers separated by space:")
	fmt.Scanln(&input)

	// Split the input into parts
	parts := strings.Fields(input)

	// The first part is the operator
	operator := parts[0]

	// Convert the remaining parts into float64 numbers
	var nums []float64
	for _, part := range parts[1:] {
		num, err := strconv.ParseFloat(part, 64)
		if err != nil {
			fmt.Println("Invalid number:", part)
			return
		}
		nums = append(nums, num)
	}

	// Call the calculator function
	result, err := calculator(operator, nums...)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
