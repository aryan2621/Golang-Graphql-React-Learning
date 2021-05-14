package main

import "fmt"

func add(x int, y int) int {
	var res = x + y
	return res
}
func calc(x int, y int) (int, int) {
	var res = x + y
	var res2 = x - y
	return res, res2
}


func main() {
	// fmt.Print("hello world")
	num1 := 60
	num2 := 5
	var result = add(num1, num2)
	a, b := calc(num1, num2)
	fmt.Print(result)
	fmt.Print(a)
	fmt.Print(b)
}
