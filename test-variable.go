package main
import "fmt"

func main(){

	//variable
	var name string = "Hello World"
	var boolean bool = true
	var age = 20

	fmt.Println(age)
	fmt.Println(boolean)

	fmt.Println(name)
	fmt.Println(len(name))
	fmt.Println(name[2])

	country := "Indonesia"
	fmt.Println("Negara : ", country)

	country = "USA"
	fmt.Println("Negara : ", country)

	var (
		firstName = "Rossi"
		lastName = "Rahmayani"
	)
	fmt.Println(firstName, lastName)

	//constant
	const food string = "Nasi"
	const price float32 = 3.2
	fmt.Println(food)
	fmt.Println(price)

	const (
		drink string = "Aqua"
		quantity int = 1
	)
	fmt.Println("Minum  : ", drink)
	fmt.Println("Jumlah : ", quantity)

	//type conversion
	var nilai1 int32 = 327689
	var nilai2 int64 = int64(nilai1)
	var nilai3 int8 = int8(nilai1)

	fmt.Println(nilai1)
	fmt.Println(nilai2)
	fmt.Println(nilai3)

	var e = drink[2]
	var eString = string(e)
	fmt.Println(eString)

	//type declaration
	type NoHP string
	var hp NoHP = "081234567890"
	fmt.Println(hp)

}
