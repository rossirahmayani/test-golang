package main

import "fmt"

func sayHello()  {
	fmt.Println("Hello World")
}

func math(a int, b int)  {
	fmt.Println("Do the math")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)
}

func compare(a int, b int) bool {
	fmt.Println("lets compare", a, ">", b)
	return a > b
}

func getCity() (string, string)  {
	return "Jakarta", "Indonesia"
}

func getProfile()(name string, gender string, age int)  {
	name = "Eko"
	gender = "Male"
	age = 30

	return
}

func getTotal(numbers ...int) int  {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	sayHello()
	fmt.Println("")

	math(20, 3)
	fmt.Println("")

	//with param
	fmt.Println("result:", compare(12, 3));
	fmt.Println("")

	city, country := getCity()
	fmt.Println("location:", city, ",", country)

	city, _ = getCity()
	fmt.Println("city:", city)

	nama, gender, umur := getProfile()
	fmt.Println("name	:", nama)
	fmt.Println("gender	:", gender)
	fmt.Println("age	:", umur)

	// variadic
	fmt.Println("")
	numbers := []int{10,10,10,10}

	total := getTotal(10, 20, 30, 40)
	fmt.Println("TOTAL:", total)

	total = getTotal(numbers...)
	fmt.Println("TOTAL:", total)

	fmt.Println("")

	//function value
	hello := sayHello
	hello()

	comp := compare
	fmt.Println("Result:", comp(2,3))
	fmt.Println("")

}
