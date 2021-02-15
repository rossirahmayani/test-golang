package main

import "fmt"

func main() {

	//array
	var names[3] string
	names[0] = "apa"
	names[1] = "siapa"
	names[2] = "dimana"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3] int {1,2,3}
	fmt.Println(values)
	fmt.Println(len(values))

	//slice
	var days = [...] string {
		"Senin", "Selasa", "Rabu","Kamis", "Jumat", "Sabtu", "Minggu",
	}

	slice1 := days[2:6] //rabu - sabtu (index 2 - index 5)
	slice2 := days[4:] //jumat - minggu (index 4 - ...)
	slice3 := days[:5] //senin - jumat (... - index 4)

	fmt.Println(slice1)
	fmt.Println(slice2[0]) //data 0 dari slice
	fmt.Println(cap(slice3)) //kapasitas dari slice awal sampai akhir array utama
	fmt.Println(len(slice3)) //panjang slice array

	slice1[2] = "Jumat lagi"
	fmt.Println(slice1) //berubah
	fmt.Println(days)

	slice4 := append(slice2, "Libur")
	fmt.Println(slice4) //nambah
	fmt.Println(days) //ga nambah karena di ujung, kapasitas ga muat

	slice5 := append(slice1, "Kerja")
	fmt.Println(slice5) //nambah
	fmt.Println(days) //ketimpa, nilai nya ganti

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Kerja"
	newSlice[1] = "Libur"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	copySlice2 := make([]string, 1, cap(newSlice))
	copy(copySlice2, newSlice)
	fmt.Println(copySlice2)

	iniSlice := []int{1,2,3} //ini slice bukan array
	fmt.Println(iniSlice)

	//map
	person := map[string]string{
		"name": "Budi",
		"city": "Solo",
	}

	person["gender"] = "Male"

	fmt.Println(person["name"])
	fmt.Println(person["city"])
	fmt.Println(person["gender"])
	fmt.Println(person)
	fmt.Println(len(person))


	movie := make(map[string]string)
	movie["title"] = "Batman"
	movie["genre"] = "Action"
	movie["rate"] = "8"

	delete(movie, "rate")
	fmt.Println(movie)
}
