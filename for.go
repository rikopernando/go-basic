package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	// for with statement
	for i := 0; i <= 10; i++ {
		fmt.Println("i ke ", i)
	}

	slice := []string{"Riko", "Pernando", "Seelvia"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index ", i, "adalah", value)
		// fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Riko"
	person["title"] = "Md. BEJ"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
