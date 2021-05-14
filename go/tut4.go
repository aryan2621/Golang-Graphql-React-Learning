package main

import "fmt"

var a = 9

// global variable

func demo() {
	a := 8
	// local variable
	fmt.Print(a)
}
func main() {
	demo()
	fmt.Print(a)

}
