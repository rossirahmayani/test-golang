package main

import (
	"fmt"
	"./helper"
	"./database"
	_"errors" // blank identifier, kalo ga kepake ga error
)

func main() {
	//test package
	result := helper.SayHello("Bujang")
	fmt.Println(result)

	//test package initialization
	db := database.GetDatabase()
	fmt.Println("Database connection:", db)
}
