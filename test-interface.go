package main

import (
	"errors"
	"fmt"
)

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type HasName interface {
	GetName() string
}

func sayHi(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func bagi(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("error divide by 0")
	} else {
		return a / b, nil
	}
}

func main() {

	//interface
	var eko Person
	eko.Name = "Eko"
	sayHi(eko)

	doge := Animal{"Doge"}
	sayHi(doge)
	fmt.Println("")

	//interface kosong
	var data interface{} = Ups(1)

	fmt.Println(data)
	fmt.Println(Ups(2))
	fmt.Println(Ups(3))
	fmt.Println("")

	//error interface
	hasil, err := bagi(10, 0)
	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err)
	}
	fmt.Println("")

	//type assertion
	result := Ups(3)

	switch value := result.(type) {
	case string:
		fmt.Println(value, "is string")
	case int:
		fmt.Println(value, "is int")
	default:
		fmt.Println("unknown type")
	}

}
