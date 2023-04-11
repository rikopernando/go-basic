package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Nama anda koosng"
	}

	return "Hello " + name
}

func main() {
	fmt.Println(getHello(""))
	fmt.Println(getHello("Riko"))
}
