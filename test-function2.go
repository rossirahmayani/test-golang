package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter)  {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string  {
	if name == "Anjing" {
		return "Cuk"
	} else {
		return name
	}
}

func main() {
	//funct as param
	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	
}
