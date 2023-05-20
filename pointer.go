package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"Liwa", "Lampung", "Indonesia"}
	address2 := &address1
	address3 := &address1
	// jika ingin membuat pointer baru dengan reference berbeda
	address4 := &Address{"Kobum", "Lampung", "Indonesia"}

	address2.City = "Balam"

	// Jika ingin merubah semua pointer tambahkan * didepan variable
	*address2 = Address{"Palembang", "Sumsel", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)

	address5 := new(Address)
	address5.City = "Bandung"
	fmt.Println(address5)

	newAlamat := Address{
		City:     "Lampung",
		Province: "Lampung",
		Country:  "",
	}

	changeCountryToIndonesia(&newAlamat)
	fmt.Println(newAlamat)
}
