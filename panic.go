package main

import "fmt"

func endApp() {
	fmt.Println("Sudah berjalan functionnya")
}

func runApp(err bool) {
	defer endApp()
	if err {
		panic("ERROR")
	}
	fmt.Println("Aplikasi berjalan normal")
}

func main() {
	runApp(true)
}
