package main

import "fmt"

func logging() {
	message := recover() //panic error ditangkap, runApp tetap jalan
	if message != nil {
		fmt.Println("error:", message)
	}

	fmt.Println("finished")
}

func runApp(value int) {
	defer logging() //fungsi logging dipanggil kalau runApp sdh selesai (baik error atau tidak)
	fmt.Println("application is running")

	fmt.Println("value:", value)

	if value == 0 {
		panic("can't divide by zero") //dipanggil jika ada error dan stop runApp
	}

	result := 10 / value
	fmt.Println("result:", result)

}

func main() {
	runApp(2)
	runApp(0) //error
	fmt.Println("DONE")

	// ini komentar baris

	/**
	ini komentar
	banyak baris
	satu blok
	*/
}
