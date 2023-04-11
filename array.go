package main

import "fmt"

func main() {
	var names [2]string

	names[0] = "Riko"
	names[1] = "Pernando"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [...]int{
		90, 80, 92,
	}

	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi)) // 10 not 0
}
