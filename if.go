package main

import "fmt"

func main() {
	name := "Rikoo"

	if name == "Riko" {
		fmt.Println("Hello", name)
	} else if name == "Joko" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Halo, boleh kenalan?", name)
	}

	// short statement IF
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
