package main

import "fmt"

func main() {
	name := "Riko"

	switch name {
	case "Riko":
		fmt.Println("Hallo Riko")
	case "Eko":
		fmt.Println("Hallo Eko")
	default:
		fmt.Println("Kenalan dong")
	}

	// short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang")
	case false:
		fmt.Println("Nama Sudah Benar")
	}

	// tanpa kondisi
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Terlalu Pendek")
	case length > 10:
		fmt.Println("Terlalu Panjang")
	default:
		fmt.Println("Nama SUdah Benar")
	}
}
