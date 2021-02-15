package main

import "fmt"

func main() {

	//if else
	fmt.Println("CONTOH IF ELSE")

	role := "Guest"

	fmt.Println("Role: ", role)
	if role == "Admin" {
		fmt.Println("Allowed")
	} else if role == "User"{
		fmt.Println("Need authorization")
	} else {
		fmt.Println("Need login")
	}

	fmt.Println("")

	pass := "password"
	fmt.Println("Password: ", pass)
	if length := len(pass); length < 12 {
		fmt.Println("Password kurang panjang")
	} else {
		fmt.Println("Password sudah panjang")
	}

	fmt.Println("")

	//switch
	fmt.Println("CONTOH SWITCH")
	switch role {
	case "Admin":
		fmt.Println("Hello Admin")
	case "User":
		fmt.Println("Hello User")
	default:
		fmt.Println("Who are you")
	}

	fmt.Println("")
	lth := len(pass)
	switch {
	case lth < 12:
		fmt.Println("Password kurang panjang")
	default:
		fmt.Println("Password sudah panjang")

	}



}
