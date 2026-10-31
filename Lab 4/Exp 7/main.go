
package main

import (
	"fmt"
	"math"
	"strconv"
)

// Division with error handling
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}

	return a / b, nil
}

// Square root with error handling
func squareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("square root of negative number is not allowed")
	}

	return math.Sqrt(num), nil
}

// Input validation
func parseNumber(input string) (int, error) {
	if input == "" {
		return 0, fmt.Errorf("input cannot be empty")
	}

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid numeric input: %s", input)
	}

	return num, nil
}

func main() {
	fmt.Println("Error Handling Demonstration")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", result)
	}

	root, err := squareRoot(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Square Root:", root)
	}

	num, err := parseNumber("123")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Number:", num)
	}

	// Demonstrating error scenarios
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = squareRoot(-25)
	if err != nil {
		fmt.Println("Error:", err)
	}

	_, err = parseNumber("")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
