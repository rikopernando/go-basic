package helper

import "fmt"

var version = 1
var Application = "Basic Golang"

func SayHello(name string) {
	fmt.Println("Hello " + name)
}

func sayGoodBy(name string) {
	fmt.Println("Good by " + name)
}
