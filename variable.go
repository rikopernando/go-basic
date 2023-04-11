package main

import "fmt"

func main() {
	var name string

	name = "Riko"
	fmt.Println(name)

	name = "Riko Pernando"
	fmt.Println(name)

	// tanpa menggunakan kata-kata type data
	var friendName = "Eko"
	fmt.Println(friendName)

	var age = 20
	fmt.Println(age)

	// tanpa menggunakan kata kunci var
	country := "Indonesia"
	fmt.Println(country)

	// Multiple Variables
	var (
		firstName = "Iqbal"
		lastName  = "Lazuardi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
