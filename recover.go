package main

import "fmt"

func endApp() {
	message := recover()
	fmt.Println("Terjadi kesalahan:", message)
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
	fmt.Println("Memastikan aplikasi masih berjalan")
}
