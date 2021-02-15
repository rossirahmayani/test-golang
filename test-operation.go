package main
import "fmt"

func main() {

	//math
	var a = 10
	var b = 3

	fmt.Println("contoh math")
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)

	var i = 10
	i += 2 //i = i + 2
	fmt.Println(i)

	i++ //i = i + 1
	fmt.Println(i)

	i-- //i = i - 1
	fmt.Println(i)

	fmt.Println(-100)

	fmt.Println("")

	//compare
	fmt.Println("contoh compare")
	fmt.Println(a == b)
	fmt.Println(a > b)
	fmt.Println(b < a)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(b <= a)

	fmt.Println("")

	//boolean logic
	fmt.Println("contoh boolean logic")
	var nilai = 90
	var absensi = 70

	var lulusNilai bool = nilai > 80
	var lulusAbsen bool = absensi > 80

	fmt.Println(lulusNilai && lulusAbsen)
	fmt.Println(lulusNilai || lulusAbsen)
	fmt.Println(!lulusAbsen)

}
