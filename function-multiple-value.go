package main

import "fmt"

func getFullName() (string, string) {
	return "Riko", "Pernando"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	name, _ := getFullName()
	fmt.Println(name)
}
