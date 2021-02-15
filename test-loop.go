package main

import "fmt"

func main() {

	//for loop
	fmt.Println("CONTOH FOR LOOP")
	count := 1
	for count <= 5 {
		fmt.Println("ulang ke -", count)
		count++
	}

	fmt.Println("")

	slice := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	for i := 0; i < len(slice); i++{
		fmt.Println("Day :", slice[i])
	}

	fmt.Println("")

	for _, val := range slice{
		fmt.Println(val)
	}

	fmt.Println("")

	//break and continue
	fmt.Println("CONTOH BREAK CONTINUE")
	for i := 0; i < 10; i ++ {
		if i % 2 == 1{
			continue //skip to next loop
		}
		if i == 6 {
			break //stop looping
		}
		fmt.Println("Nilai ", i)
	}


}
