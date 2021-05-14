package main

import (
	"fmt"
	"math"
)

func main() {
	num := 8
	fmt.Printf("This is the output %.2g Thanks !!!!1\n",math.Sqrt(float64(num)))
	var second=math.Round(math.Sqrt(float64(num)))
	// var second=math.Floor(math.Sqrt(float64(num)))
	fmt.Print(second)
//	fmt.Printf("This is the output %g Thanks !!!!1",math.Sqrt(float64(num)))

}
