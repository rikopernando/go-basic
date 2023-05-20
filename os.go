package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument : ")
	fmt.Println(args)

	name, err := os.Hostname()
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println("Hostname : ", name)
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username, password)
}
