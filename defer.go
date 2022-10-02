package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}
func runApplication(value int) {
	defer logging() //jalan di akhir function induk meskipun terjadi error
	fmt.Println("Run Application")
	result := 12 / value
	fmt.Println("hasil bagi: ", result)
	// logging()		//tidak jalan ketika error
}
func main() {
	runApplication(0)
}
