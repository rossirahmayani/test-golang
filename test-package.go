package main

import (
	"fmt"
	"os"
)

func main() {
	//package os
	args := os.Args
	fmt.Println(args)

	hostname, err := os.Hostname()
	if err == nil{
		fmt.Println("Hostname:", hostname)
	} else {
		fmt.Println("Error", err.Error())
	}

	fmt.Println(os.Getenv("USERNAME"))
	fmt.Println(os.Getenv("PASSWORD"))

}
