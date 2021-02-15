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

func main() {
	sayHello()
	fmt.Println("")

	math(20, 3)
	fmt.Println("")

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

}
