package main

import "fmt"

func getName() (firstName, lastName string) {
	firstName = "Riko"
	lastName = "Pernando"

	return
}

func main() {
	firstName, lastName := getName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
