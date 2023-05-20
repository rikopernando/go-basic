package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	riko := Man{"Riko"}
	riko.Married()

	fmt.Println(riko)
}
