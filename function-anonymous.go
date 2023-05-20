package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You're Blocked ", name)
	} else {
		fmt.Println("Welcome ", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("riko", blacklist)

	registerUser("root", func(name string) bool {
		return name == "admin"
	})

	registerUser("admin", func(name string) bool {
		return name == "admin"
	})
}
