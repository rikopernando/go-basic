package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello " + hasName.GetName())
}

// type Animal struct {
// 	Name string
// }

func (person Person) GetName() string {
	return person.Name
}

// func (animal Animal) GetName() string {
// 	return animal.Name
// }

func main() {
	var riko Person
	riko.Name = "riko"
	SayHello(riko)

	// cat := Animal{Name: "push"}
	// SayHello(cat)
}
