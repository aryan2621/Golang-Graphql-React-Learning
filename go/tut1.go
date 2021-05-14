package main

import (
	"fmt"
)

// its necessary to use what u declared in Go;
// otherwise it wil through the error

func main() {
	var num1 int = 2
	var num2 = 3
	var num3 int
	num3 = 4

	var res = num1 + num2
	num1, num2 = 3, 5

	// var num6,num7 int

	result := 9 // same as var result =9

	const x = 9

	fmt.Print(2 + num1)
	fmt.Print(num2)
	fmt.Print(res)
	fmt.Print(num3)
	fmt.Print(result)

}
