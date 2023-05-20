package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "root", "Put your database password")
	number := flag.Int("number", 123, "Put your number")

	flag.Parse()

	fmt.Println("HOST: ", *host)
	fmt.Println("USERNAME : ", *username)
	fmt.Println("PASSWORD", *password)
	fmt.Println("NUMBER: ", *number)
}
