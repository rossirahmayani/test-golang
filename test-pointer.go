package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountry(address Address, country string)  {
	address.Country = country
}

func changeCountry2(address *Address, country string)  {
	address.Country = country
}

type Woman struct {
	Name string
}

func (woman *Woman) Married()  {
	woman.Name = "Mrs. " + woman.Name
}

func (woman Woman) Single(){
	woman.Name = "Miss " + woman.Name
}

func main() {

	//pass by value
	var address1 Address = Address{"Sleman", "DI Yogyakarta", "Indonesia"}
	address2 := address1
	address2.City = "Bantul"

	fmt.Println(address1) //city tidak berubah
	fmt.Println(address2)
	fmt.Println("")

	//pass by reference
	address3 := &address1
	address3.City = "Yogyakarta"
	fmt.Println(address1) //city berubah
	fmt.Println(address3)
	fmt.Println("")

	var address4 *Address = &address1
	address4 = &Address{"Surabaya", "Jawa Timur", "Indonesia"}
	fmt.Println(address1) //semua tdk berubah
	fmt.Println(address4)
	fmt.Println("")

	address5 := &address1
	*address5 = Address{"Bandung", "Jawa Barat", "Indonesia"}
	fmt.Println(address1) //semua berubah
	fmt.Println(address5)
	fmt.Println("")

	//function new
	var address6 *Address = new(Address)
	address6.City = "Padang"
	fmt.Println(address6)
	fmt.Println("")

	//pointer di function
	address7 := Address{"Medan", "Sumatera Utara", ""}
	changeCountry(address7, "Indonesia")
	fmt.Println(address7) //country tidak berubah
	changeCountry2(&address7, "Indonesia")
	fmt.Println(address7) //country berubah
	fmt.Println("")

	//pointer di method
	puff := Woman{"Puff"}
	puff.Single()
	fmt.Println(puff) //tidak berubah
	puff.Married()
	fmt.Println(puff) //berubah

}
