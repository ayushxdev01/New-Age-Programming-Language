package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

// Read and validate integer input
func readInt(message string) int64 {
	for {
		fmt.Print(message)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.ParseInt(input, 10, 64)

		if err != nil || num < -1000000 || num > 1000000 {
			fmt.Println("Invalid input! Enter an integer between -1000000 and 1000000.")
			continue
		}

		return num
	}
}

// Read and validate floating-point input
func readFloat(message string) float64 {
	for {
		fmt.Print(message)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		num, err := strconv.ParseFloat(input, 64)

		if err != nil || math.IsNaN(num) || math.IsInf(num, 0) ||
			math.Abs(num) > 1000000000 {
			fmt.Println("Invalid input! Enter a valid number within the allowed range.")
			continue
		}

		return num
	}
}

func main() {
	fmt.Println("================================")
	fmt.Println("   NAP LAB 1 - GO FUNDAMENTALS")
	fmt.Println("================================")

	// Task 1: Go variables and arithmetic
	fmt.Println("\n--- Task 1: Integer Arithmetic ---")

	var a int = 20
	var b int = 10

	fmt.Println("Integer A:", a)
	fmt.Println("Integer B:", b)
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)

	fmt.Println("\n--- Floating Point Arithmetic ---")

	var x float64 = 15.5
	var y float64 = 2.5

	fmt.Println("Float X:", x)
	fmt.Println("Float Y:", y)
	fmt.Println("Addition:", x+y)
	fmt.Println("Subtraction:", x-y)
	fmt.Println("Multiplication:", x*y)

	// Task 2: Integer Calculator
	fmt.Println("\n--- Task 2: Integer Calculator ---")

	num1 := readInt("Enter first integer: ")
	num2 := readInt("Enter second integer: ")

	fmt.Println("Integer Addition:", num1+num2)
	fmt.Println("Integer Subtraction:", num1-num2)
	fmt.Println("Integer Multiplication:", num1*num2)

	// Floating Point Calculator
	fmt.Println("\n--- Floating Point Calculator ---")

	float1 := readFloat("Enter first decimal number: ")
	float2 := readFloat("Enter second decimal number: ")

	fmt.Printf("Float Addition: %.2f\n", float1+float2)
	fmt.Printf("Float Subtraction: %.2f\n", float1-float2)
	fmt.Printf("Float Multiplication: %.2f\n", float1*float2)

	fmt.Println("\n================================")
	fmt.Println("       PROGRAM COMPLETED")
	fmt.Println("================================")
}
