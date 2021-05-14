package main

import (
	"fmt"
)

func main() {
	num := 10
	if num < 5 {

		fmt.Print("hi ", num)
	} else {
		fmt.Print("byee ", num)
	}
	fmt.Print("\n")
	switch num {
	case 1:
		fmt.Print("Good Morning ", num)
		break
	case 2:
		fmt.Print("Good AfterNoon ", num)
		break
	case 3:
		fmt.Print("Good Night ", num)
		break
	default:
		fmt.Print("Not Available ", num)
	}

}
