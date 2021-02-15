package main

import "fmt"

type CustomerData struct {
	Name, Address string
	Age           int
}

func (customer CustomerData) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var juned CustomerData
	juned.Name = "Juned"
	juned.Address = "Lombok"
	juned.Age = 16

	fmt.Println(juned.Name)
	fmt.Println(juned.Address)
	fmt.Println(juned.Age)

	//struct literal
	buret := CustomerData{"Buret", "Bandung", 20} //sesuai urutan field

	emon := CustomerData{
		Name:    "Emon",
		Age:     19,
		Address: "Bogor",
	}

	fmt.Println(buret)
	fmt.Println(emon)

	//struct method
	emon.sayHello("Bono")
}
