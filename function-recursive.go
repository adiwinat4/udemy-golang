package main

import "fmt"

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factorialRecursive(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
}
