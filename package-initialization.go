package main

import (
	"fmt"
	"golang-basic/database"
)

// import (
// 	_ "golang-basic/database" // black identifier
// )

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
