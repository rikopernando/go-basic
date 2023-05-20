package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "my name is "+customer.Name)
}

func main() {
	var riko Customer
	riko.Name = "Riko"
	riko.Address = "Lampung"
	riko.Age = 28
	riko.sayHi("Joko")

	fmt.Println(riko)

	nando := Customer{
		Name:    "Nando",
		Address: "Lampung",
		Age:     27,
	}

	fmt.Println(nando)

	budi := Customer{"Budi", "Lampung", 20}
	fmt.Println(budi)

}
