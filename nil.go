package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := newMap("")
	if person == nil {
		fmt.Println("data kosong")
	} else {
		fmt.Println(person)
	}
}
