package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 20; i++ {
		frac1, frac2 := i%3, i%5
		switch {
		case frac1 == 0 && frac2 == 0:
			fmt.Println("fizz buzz")
		case frac1 == 0:
			fmt.Println("fizz")
		case frac2 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}
