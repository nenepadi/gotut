package main

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func multiple(x, y string) (string, string) {
	return x, y
}

func statictype() {
	var a, b float64 = 5.6, 9.5
	fmt.Println(add(a, b))
}

func dynamictype() {
	// this is dynamic ...
	// but once type is set automatically by go,
	// there kis no going back
	num1, num2 := 5.6, 9.5
	fmt.Println(add(num1, num2))
}

func main() {
	statictype()
	dynamictype()

	// throw in some strings ...
	a, b := "Hey", "there"
	fmt.Println(multiple(a, b))
}
