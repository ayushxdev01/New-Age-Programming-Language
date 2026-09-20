package main

import (
	"fmt"
	"lab2/mathutil"
	"lab2/strop"
)

func main() {
	text := "Hello World"
	fmt.Println("Original String :", text)
	fmt.Println("Reversed String :", strop.Reverse(text))
	fmt.Println("Number of Vowels:", strop.CountVowels(text))
	fmt.Println("Factorial of 5  :", mathutil.Factorial(5))
	fmt.Println("2 raised to 5   :", mathutil.Power(2, 5))

	sliceOperations()
	mapOperations()
}
