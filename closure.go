package main

import "fmt"

func main() {
	count := 0
	name := "anwar"

	increment := func() {
		name := "joko"
		count++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(name)
	fmt.Println(count)
}
