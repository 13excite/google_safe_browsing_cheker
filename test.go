package main

import "fmt"

func main() {
	x := [5]float64{
		84,
		45,
		67,
		78,
		78,
	}
	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}
