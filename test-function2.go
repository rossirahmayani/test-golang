package main

import "fmt"

type Filter func(string) string
type Blacklist func(string) bool

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "Cuk"
	} else {
		return name
	}
}

func register(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func factorialLoop(value int) int {
	res := 1
	for i := value; i > 0; i-- {
		res *= i
	}
	return res
}

func factorialRecursive(value int) int {
	if value <= 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	//function as param
	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	fmt.Println("")

	//anonymous function
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	register("Eko", blacklist)

	register("Anjing", func(name string) bool {
		return name == "Anjing"
	})

	//recursive function
	fmt.Println("FACTORIAL LOOP RESULT:", factorialLoop(5))      //loop biasa
	fmt.Println("RECURSIVE FUNC RESULT:", factorialRecursive(5)) //pakai recursive function
}
