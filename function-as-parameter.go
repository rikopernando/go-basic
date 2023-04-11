package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello " + nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}

	return name
}

func main() {
	sayHelloWithFilter("Riko", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
